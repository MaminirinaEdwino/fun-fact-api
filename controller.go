package main

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

func CreateTable(w http.ResponseWriter, r *http.Request){
	fmt.Println("teste de sqlite avec golang")
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		fmt.Println("Erreur pendant l'ouverture de la base de donnée")
		return
	}
	defer db.Close()


	_, err = db.Exec("CREATE TABLE IF NOT EXISTS funfactlist (id INTEGER PRIMARY KEY AUTOINCREMENT, funfact TEXT)")
	if err != nil {
		fmt.Printf("Erreur lors de la creation de la table %s\n", err)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request){
	fmt.Println("teste de sqlite avec golang")
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		fmt.Println("Erreur pendant l'ouverture de la base de donnée")
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM funfactlist")

	if err != nil {
		fmt.Println("Erreur lors du select", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var funfact string
		err = rows.Scan(&id, &funfact)
		if err != nil {
			fmt.Println("Erreur lors du scan", err)
			continue
		}
		fmt.Fprintf(w, "id: %d funfact:%s \n", id, funfact)
	}
}

func GetById(w http.ResponseWriter, r *http.Request){
	// id := r.URL.Query().Get("id")
	fmt.Println("teste de sqlite avec golang by id")
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		fmt.Println("Erreur pendant l'ouverture de la base de donnée")
		return
	}
	defer db.Close()
	fmt.Fprintf(w, "id: %s", r.PathValue("id"))

}

func Post(w http.ResponseWriter, r *http.Request){
	fmt.Println("teste de sqlite avec golang")
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		fmt.Println("Erreur pendant l'ouverture de la base de donnée")
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO funfactlist (funfact) VALUES (?)", "rose are red")
	if err != nil {
		fmt.Printf("Erreur lord de l'insertion %s\n", err)
		return
	}
}

func Put(w http.ResponseWriter, r *http.Request){
	fmt.Println("teste de sqlite avec golang")
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		fmt.Println("Erreur pendant l'ouverture de la base de donnée")
		return
	}
	defer db.Close()

}

func Delete(w http.ResponseWriter, r *http.Request){
	fmt.Println("teste de sqlite avec golang")
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		fmt.Println("Erreur pendant l'ouverture de la base de donnée")
		return
	}
	defer db.Close()

}