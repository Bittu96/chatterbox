package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	auth "projects/chatterbox/server/pkgs/auth"
	"projects/chatterbox/server/pkgs/dao"
	"projects/chatterbox/server/pkgs/database"
	"projects/chatterbox/server/pkgs/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

const (
	appName = "chatterbox"
	appDesc = "Chat chat chat!"
	appPort = 8080
)

type Envs struct {
	DB struct {
		Host     string
		Port     int
		Password string
		User     string
		Database string
	}
	Redis struct {
		Host     string
		Port     string
		Password string
		User     string
		Database string
	}
}

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appDesc
	envs := Envs{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// fmt.Println("envs", os.Environ())

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "DB_NAME",
			Value:       "postgres",
			Usage:       "database name",
			Destination: &envs.DB.Database,
		},
		cli.StringFlag{
			Name:        "DB_HOST",
			Value:       "localhost",
			Destination: &envs.DB.Host,
		},
		cli.StringFlag{
			Name:        "DB_USER",
			Value:       "postgres",
			Usage:       "database name",
			Destination: &envs.DB.User,
		},
		cli.IntFlag{
			Name:        "DB_PORT",
			Value:       5432,
			Destination: &envs.DB.Port,
		},
		cli.StringFlag{
			Name:        "DB_PASSWORD",
			Value:       "postgres",
			Destination: &envs.DB.Password,
		},
	}

	app.Action = func(ctx *cli.Context) {
		gin.SetMode(gin.ReleaseMode)
		routes := gin.Default()
		routes.Use(CORSMiddleware())

		//connect postgres db
		pgClient, err := database.PostgresConnect()
		if err != nil {
			return
		}

		//close database
		defer pgClient.Close()

		//connect mongo db
		// mgClient, err := database.MongoConnect()
		// if err != nil {
		// 	return
		// }

		//connect redis
		rdClient, err := database.RedisConnect()
		if err != nil {
			return
		}

		//close redis
		defer rdClient.Close()

		//init data access service
		daoService := dao.DAO{
			PgClient: pgClient,
			// MgClient: mgClient,
			RdClient: rdClient,
		}

		//create handles
		handles := handlers.Handles{
			Dao: daoService,
		}

		//test server connection
		routes.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		routes.LoadHTMLFiles("./welcome.html")
		routes.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "welcome.html", gin.H{
				"content": "This is an welcome page...",
			})
		})

		//public auth routes
		v1 := routes.Group("/v1")
		{
			v1.POST("/register", handles.Register)
			v1.POST("/login", handles.Login)
			v1.GET("/logout", handles.Logout)
		}

		//private user routes
		v2 := routes.Group("/v2")
		{
			v2.Use(BearerAuthMiddleware())
			v2.GET("/home", handles.Home)
			v2.GET("/users", handles.GetUsers)
			v2.GET("/followers", handles.GetFollowers)
			v2.GET("/following", handles.GetFollowing)
			v2.GET("/follow", handles.FollowUser)
			v2.GET("/unfollow", handles.UnfollowUser)
		}

		//private admin routes
		v3 := routes.Group("/v3")
		{
			v3.Use(BearerAuthMiddleware())
			v3.POST("/admin", handles.Premium)
		}

		//private socket routes
		v4 := routes.Group("/v4")
		{
			// v2.Use(BearerAuthMiddleware())
			// v4.GET("/chat/:room_id", handles.WebSocGin)
			v4.GET("/chat/:user_id", handles.WebSocPrivate)
			v4.GET("/todo_list.html", func(c *gin.Context) {
				handles.WebSocPage(c.Writer, c.Request)
			})
		}

		// for _, r := range routes.Routes() {
		// 	fmt.Println(r.Method, r.Path)
		// }

		//run server
		fmt.Println("server running on port.. :", appPort)
		routes.Run(fmt.Sprintf(":%v", appPort))
	}
	app.Run(os.Args)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Writer.Header().Set("Content-Type", "*; application/json; charset=utf-8; application/x-www-form-urlencoded;")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Cookie")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}

		c.Next()
	}
}

func CookieAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		_, err = auth.ParseToken(cookie)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// if claims.Role != "user" && claims.Role != "admin" {
		// 	c.JSON(401, gin.H{"error": "unauthorized"})
		// 	return
		// }

		// if claims.Role != "admin" {
		// 	c.JSON(401, gin.H{"error": "unauthorized"})
		// 	return
		// }

		c.Next()
	}
}

func BearerAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if bearerToken == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		splitToken := strings.Split(bearerToken, " ")
		if len(splitToken) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := auth.ParseToken(splitToken[1])
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// if claims.Role != "user" && claims.Role != "admin" {
		// 	c.JSON(401, gin.H{"error": "unauthorized"})
		// 	return
		// }

		// if claims.Role != "admin" {
		// 	c.JSON(401, gin.H{"error": "unauthorized"})
		// 	return
		// }

		c.AddParam("auth_user_id", fmt.Sprintf("%d", claims.UserId))
		c.Next()
	}
}
