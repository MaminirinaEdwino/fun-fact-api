package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)
var DB *sql.DB

func ConnectDB() {
	// var db sql.DB
	var err error
	DB, err = sql.Open("sqlite3", "database/db.sqlite")
	if err != nil {
		fmt.Println("Erreur pendant l'ouverture de la base de donnée")
		return 
	}
}

func CloseDB(){
	fmt.Println("Close DB")
	if  DB != nil {
		DB.Close()
	}

}