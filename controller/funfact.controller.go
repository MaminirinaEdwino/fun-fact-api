package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MaminirinaEdwino/fun-fact-api/database"
	"github.com/MaminirinaEdwino/fun-fact-api/model"
	_ "github.com/mattn/go-sqlite3"
)



func CreateTable(w http.ResponseWriter, r *http.Request){
	database.ConnectDB()
	_, err := database.DB.Exec("CREATE TABLE IF NOT EXISTS funfactlist (id INTEGER PRIMARY KEY AUTOINCREMENT, funfact TEXT)")
	if err != nil {
		fmt.Printf("Erreur lors de la creation de la table %s\n", err)
	}
	database.CloseDB()
}

func GetAll(w http.ResponseWriter, r *http.Request){
	var FfList model.FunFactList
	database.ConnectDB()
	rows, err := database.DB.Query("SELECT * FROM funfactlist")

	if err != nil {
		fmt.Println("Erreur lors du select", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var funfact model.Funfact
		err = rows.Scan(&funfact.Id, &funfact.FunFact)
		
		if err != nil {
			fmt.Println("Erreur lors du scan", err)
			continue
		}
		FfList = append(FfList, funfact)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(FfList)
	if err != nil {
		fmt.Println("Erreur d'encodage du json", err)
		return
	}
}

func GetById(w http.ResponseWriter, r *http.Request){
	// id := r.URL.Query().Get("id")
	// fmt.Println("teste de sqlite avec golang by id")
	var funfact model.Funfact
	id := r.PathValue("id")
	database.ConnectDB()
	result, err := database.DB.Query("SELECT * FROM funfactlist where id = ? ", id)
	if err!= nil {
		fmt.Println("Erreur lors de la recuperation des données")
	}
	
	result.Next()
	result.Scan(&funfact.Id, &funfact.FunFact)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(funfact)


}

func Post(w http.ResponseWriter, r *http.Request){
	database.ConnectDB()
	// var ff model.FunFactPost
	// _, err := database.DB.Exec("INSERT INTO funfactlist (funfact) VALUES (?)", "rose are red")
	// if err != nil {
	// 	fmt.Printf("Erreur lord de l'insertion %s\n", err)
	// 	return
	// }
	// body, err := io.ReadAll(r.Body)
	// if err!= nil {
	// 	fmt.Println("Bpody error")
	// }
	var ff model.FunFactPost

	decodeur := json.NewDecoder(r.Body)
	err := decodeur.Decode(&ff)
	if err != nil {
		fmt.Println("Erreur de decodage", err)
	}
	_, err = database.DB.Exec("INSERT INTO funfactlist (funfact) VALUES (?)", ff.Funfact)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ff)
	
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