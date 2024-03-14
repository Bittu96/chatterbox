package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// local db
// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "password"
// 	dbname   = "chatterbox"
// )

// ec2 db
const (
	host     = "ec2-52-55-96-26.compute-1.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func PostgresConnect() (*sql.DB, error) {
	_, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// connection string
	psqlConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println("psqlConnectionString:", psqlConnectionString)

	// open database
	db, err := sql.Open("postgres", psqlConnectionString)
	fmt.Println("db, err:", db, err)

	// CheckError(err)
	if err != nil {
		fmt.Println("postgres connection failed!")
		return db, err
	}

	// check db
	err = db.Ping()
	// CheckError(err)
	if err != nil {
		fmt.Println("postgres connection failed!", err)
		return db, err
	}

	fmt.Println("postgres connection success!")
	return db, err
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
