package database

import (
	"os"
)

func ImportConn() (string, string, string, string) {
	database := os.Getenv("DATABASE")
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	dbname := os.Getenv("DBNAME")

	return database, dbname, host, port
}

func ImportSingle() string {
	single := os.Getenv("SINGLE")
	return single
}

func ImportDouble() string {
	double := os.Getenv("DOUBLE")
	return double
}

func ImportContest() string {
	contest := os.Getenv("CONTEST")
	return contest
}
