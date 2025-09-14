package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MaminirinaEdwino/fun-fact-api/database"
	"github.com/MaminirinaEdwino/fun-fact-api/model"
	_ "github.com/mattn/go-sqlite3"
)

func MainController(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, ReadTemplate("Template/mainController.html"))
}



func CreateTable(w http.ResponseWriter, r *http.Request) {
	database.ConnectDB()
	_, err := database.DB.Exec("CREATE TABLE IF NOT EXISTS funfactlist (id INTEGER PRIMARY KEY AUTOINCREMENT, funfact TEXT)")
	if err != nil {
		fmt.Printf("Erreur lors de la creation de la table %s\n", err)
	}
	database.CloseDB()
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	var FfList model.FunFactList
	database.ConnectDB()
	_, err := database.DB.Exec("CREATE TABLE IF NOT EXISTS funfactlist (id INTEGER PRIMARY KEY AUTOINCREMENT, funfact TEXT)")
	if err != nil {
		fmt.Printf("Erreur lors de la creation de la table %s\n", err)
	}
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

func GetById(w http.ResponseWriter, r *http.Request) {

	var funfact model.Funfact
	id := r.PathValue("id")
	database.ConnectDB()

	_, err := database.DB.Exec("CREATE TABLE IF NOT EXISTS funfactlist (id INTEGER PRIMARY KEY AUTOINCREMENT, funfact TEXT)")
	if err != nil {
		fmt.Printf("Erreur lors de la creation de la table %s\n", err)
	}

	result, err := database.DB.Query("SELECT * FROM funfactlist where id = ? ", id)
	if err != nil {
		fmt.Println("Erreur lors de la recuperation des données")
	}

	result.Next()
	defer result.Close()
	err = result.Scan(&funfact.Id, &funfact.FunFact)
	
	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(model.BasicResponse{
			Action: "get funfact by id",
			Message: "funfact not found",
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(funfact)

}

func Post(w http.ResponseWriter, r *http.Request) {
	database.ConnectDB()


	_, err := database.DB.Exec("CREATE TABLE IF NOT EXISTS funfactlist (id INTEGER PRIMARY KEY AUTOINCREMENT, funfact TEXT)")
	if err != nil {
		fmt.Printf("Erreur lors de la creation de la table %s\n", err)
	}
	
	var ff model.FunFactPost

	decodeur := json.NewDecoder(r.Body)
	err = decodeur.Decode(&ff)
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

func Put(w http.ResponseWriter, r *http.Request) {
	database.ConnectDB()

	_, err := database.DB.Exec("CREATE TABLE IF NOT EXISTS funfactlist (id INTEGER PRIMARY KEY AUTOINCREMENT, funfact TEXT)")
	if err != nil {
		fmt.Printf("Erreur lors de la creation de la table %s\n", err)
	}
	defer database.DB.Close()
	id := r.PathValue("id")
	var ff model.FunFactPost

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&ff)
	if err != nil {
		fmt.Println("Erreur lors du decodage du contenu de la requette", err)
	}

	_, err = database.DB.Exec("UPDATE funfactlist SET funfact = ? WHERE id = ?", ff.Funfact, id)
	if err != nil {
		fmt.Println("Erreur lors de la modification", err)
		return
	}
	database.CloseDB()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	nbr, err := strconv.Atoi(id)

	response := model.Funfact{
		Id:      nbr,
		FunFact: ff.Funfact,
	}
	json.NewEncoder(w).Encode(response)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	database.ConnectDB()
	defer database.CloseDB()
	_, err := database.DB.Exec("CREATE TABLE IF NOT EXISTS funfactlist (id INTEGER PRIMARY KEY AUTOINCREMENT, funfact TEXT)")
	if err != nil {
		fmt.Printf("Erreur lors de la creation de la table %s\n", err)
	}
	_, err = database.DB.Exec("DELETE FROM funfactlist WHERE id = ?", r.PathValue("id"))

	if err != nil {
		fmt.Println("erreur de suppression", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := model.BasicResponse{
		Action:  "Delete funfact",
		Message: "Funfact deleted",
	}
	json.NewEncoder(w).Encode(response)

}
