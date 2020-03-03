package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Lavagna struct {
	Id        int
	Messaggio string
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

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM messaggi ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	lav := Lavagna{}
	res := []Lavagna{}
	for selDB.Next() {
		var id int
		var messaggio string
		err = selDB.Scan(&id, &messaggio)
		if err != nil {
			panic(err.Error())
		}
		lav.Id = id
		lav.Messaggio = messaggio
		res = append(res, lav)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM messaggi WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	lav := Lavagna{}
	for selDB.Next() {
		var id int
		var messaggio string
		err = selDB.Scan(&id, &messaggio)
		if err != nil {
			panic(err.Error())
		}
		lav.Id = id
		lav.Messaggio = messaggio
	}
	tmpl.ExecuteTemplate(w, "Show", lav)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM messaggi WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	lav := Lavagna{}
	for selDB.Next() {
		var id int
		var messaggio string
		err = selDB.Scan(&id, &messaggio)
		if err != nil {
			panic(err.Error())
		}
		lav.Id = id
		lav.Messaggio = messaggio
	}
	tmpl.ExecuteTemplate(w, "Edit", lav)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		messaggio := r.FormValue("messaggio")
		insForm, err := db.Prepare("INSERT INTO messaggi(messaggio) VALUE(?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(messaggio)
		log.Println("INSERT: messaggio: " + messaggio)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		messaggio := r.FormValue("messaggio")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE messaggi SET messaggio=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(messaggio, id)
		log.Println("UPDATE: Messaggio: " + messaggio)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM messaggi WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
