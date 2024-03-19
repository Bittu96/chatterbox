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

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appDesc

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	} else {
		// fmt.Println("envs", os.Environ())
		fmt.Println("loaded", len(os.Environ()), "envs")
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
		// rdClient, err := database.RedisConnect()
		// if err != nil {
		// 	return
		// }

		//close redis
		// defer rdClient.Close()

		//init data access service
		daoService := dao.DAO{
			PgClient: pgClient,
			// MgClient: mgClient,
			// RdClient: rdClient,
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
		auth := routes.Group("/auth")
		{
			auth.POST("/register", handles.Register)
			auth.POST("/login", handles.Login)
			auth.GET("/users", handles.GetUsers)
		}

		//private user routes
		svc := routes.Group("/svc")
		svc.Use(userAuthMiddleware())
		{
			svc.GET("/logout", handles.Logout)
			svc.GET("/home", handles.Home)
			svc.GET("/followers", handles.GetFollowers)
			svc.GET("/following", handles.GetFollowing)
			svc.GET("/follow", handles.FollowUser)
			svc.GET("/unfollow", handles.UnfollowUser)
		}

		//private admin routes
		admin := routes.Group("/admin")
		admin.Use(adminAuthMiddleware())
		{
			admin.POST("/admin", handles.Premium)
		}

		//private socket routes
		webs := routes.Group("/webs")
		webs.Use(userAuthMiddleware())
		{
			// v4.GET("/chat/:room_id", handles.WebSocGin)
			webs.GET("/chat/:user_id", handles.WebSocPrivate)
			webs.GET("/todo_list.html", func(c *gin.Context) {
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

func userAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionToken, err := c.Cookie("session_token")
		if err != nil {
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

			sessionToken = splitToken[1]
		}

		claims, err := auth.ParseToken(sessionToken)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims.Role != "user" && claims.Role != "admin" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AddParam("auth_user_id", fmt.Sprintf("%d", claims.UserId))
		c.Next()
	}
}

func adminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionToken, err := c.Cookie("session_token")
		if err != nil {
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

			sessionToken = splitToken[1]
		}

		claims, err := auth.ParseToken(sessionToken)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims.Role != "admin" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AddParam("auth_user_id", fmt.Sprintf("%d", claims.UserId))
		c.Next()
	}
}
