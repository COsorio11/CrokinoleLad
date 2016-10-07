package main

import (
	"CrokinoleLad/server/database"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := database.DbConnect()
	single := database.ImportSingle()
	rows, err := db.Query("SELECT id, name, password from " + single)
	if err != nil {
		log.Fatal("SECOND: ", err)
	}

	fmt.Println("SUCCESS: ", rows)
	defer rows.Close()
	var id int
	var name string
	// var wins int
	// var losses int
	var password string

	for rows.Next() {
		_ = rows.Scan(&id, &name, &password)
		fmt.Println("ID: ", id)
		fmt.Println("NAME: ", name)
		// fmt.Println("WINS: ", wins)
		// fmt.Println("LOSSES: ", losses)
		fmt.Println("PASSWORD: ", password)
	}
}
