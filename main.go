package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type event struct {
	ID        int    `json:"id"`
	Messaggio string `json:"messaggio"`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := os.Getenv("SQL_USER")
	dbPass := os.Getenv("SQL_PASSWORD")
	dbUrl := os.Getenv("SQL_URL")
	dbName := os.Getenv("SQL_NAME")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbUrl+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var messaggi []event

func getData() {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM messaggi")
	if err != nil {
		panic(err.Error())
	}

	messaggi = nil
	for selDB.Next() {
		var id int
		var messaggio string
		err = selDB.Scan(&id, &messaggio)
		if err != nil {
			panic(err.Error())
		}
		a := event{ID: id, Messaggio: messaggio}
		messaggi = append(messaggi, a)

	}
	defer db.Close()
	fmt.Println("Caricato")
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Lavagna GO API\n")
}

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}

		messaggio := r.Form.Get("messaggio")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Dati ottenuti")

		db := dbConn()
		insForm, err := db.Prepare("INSERT INTO messaggi(messaggio) VALUE(?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(messaggio)
		log.Println("INSERT: messaggio: " + messaggio)

		defer db.Close()

		getData()

	}
}

func getAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(messaggi)
}

func main() {
	getData()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/add", add).Methods("POST")
	router.HandleFunc("/all", getAll).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
