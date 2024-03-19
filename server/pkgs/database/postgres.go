package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// local db
const (
	host     = "localhost"
	port     = 5432
	user     = "bittu"
	password = "bittu"
	dbname   = "postgres"
)

// ec2 db
// const (
// 	host     = "ec2-52-55-96-26.compute-1.amazonaws.com"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "postgres"
// 	dbname   = "postgres"
// )

func PostgresConnect() (*sql.DB, error) {
	_, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// connection string
	psqlConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// fmt.Println("psqlConnectionString:", psqlConnectionString)

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
