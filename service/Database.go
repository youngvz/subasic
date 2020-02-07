package service

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"subasic/utilities"
)

func get() *sql.DB {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=disable", config.User, config.Password, config.Server, config.Port)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return db
}