package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	username = "user"
	password = "xyz13juli"
	hostname = "perpus.crossnet.co.id"
	port     = "3306"
	dbname   = "taman_baca"
)

func DbConnection() (*sql.DB, error) {
	connectionString := username + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + dbname

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Printf("Error #(err) when opening DB\n")
		return nil, err
	}
	err = db.Ping()

	if err != nil {
		fmt.Println("Error : " + err.Error())
		return nil, err
	}
	//defer db.Close()

	return db, nil
}
