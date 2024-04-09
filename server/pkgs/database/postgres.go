package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"projects/chatterbox/server/pkgs/configs"
	"time"

	_ "github.com/lib/pq"
)

var (
	host     = configs.ServerConfig.DB_Host
	port     = configs.ServerConfig.DB_Port
	user     = configs.ServerConfig.DB_User
	password = configs.ServerConfig.DB_Pass
	dbname   = configs.ServerConfig.DB_Name
)

func PostgresConnect() (*sql.DB, error) {
	_, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	fmt.Println("postgres creds", host, port, user, password, dbname)
	// connection string
	psqlConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configs.ServerConfig.DB_Host,
		configs.ServerConfig.DB_Port,
		configs.ServerConfig.DB_User,
		configs.ServerConfig.DB_User,
		configs.ServerConfig.DB_Name)
	fmt.Println("psqlConnectionString:", psqlConnectionString)

	// open database
	db, err := sql.Open("postgres", psqlConnectionString)
	if err != nil {
		fmt.Println("postgres connection failed!", err)
		log.Fatal("postgres connection failed!", err)
		return db, err
	}

	// check db
	if err = db.Ping(); err != nil {
		fmt.Println("postgres connection failed!", err)
		return db, err
	}

	fmt.Println("postgres connection success!")
	return db, err
}
