package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbConnect() *sql.DB {
	dbtype, dbname, host, port := ImportConn()

	db, err := sql.Open(dbtype, "dbname="+dbname+" host="+host+" port="+port+" sslmode=disable")
	if err != nil {
		log.Fatal("FIRST: ", err)
	}
	return db
}
