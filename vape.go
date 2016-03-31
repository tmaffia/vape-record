// Package main provides ...
package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)
import _ "github.com/lib/pq"

var db *gorm.DB

func main() {
	dbConString := "postgres://" + AppConf.DBUsername + ":" +
		AppConf.DBPassword + "@" + AppConf.DBID
	vapedb, err := gorm.Open("postgres", dbConString)
	if err != nil {
		panic(err.Error())
	}
	defer vapedb.Close()

	db = vapedb
	// db.CreateTable(&Vape{})
	// db.Create(&Vape{})

	r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/api/vapes", VapeHandler)
	r.HandleFunc("/api/vape", VapeCreateHandler)
	r.HandleFunc("/api/delete-vape", VapeDeleteHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public/")))
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}

type Vape struct {
	gorm.Model
}

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// v := struct {
// Vapes int
// }{
// Vapes: GetTotalVapes(),
// }
// t, err := template.ParseFiles("public/index.html")
// if err != nil {
// panic(err.Error())
// }
// t.Execute(w, v)
// }

type VapeReponseAll struct {
	Count int
	Vapes []Vape
}

func VapeHandler(w http.ResponseWriter, r *http.Request) {
	vapes := []Vape{}
	db.Find(&vapes)

	jsn, err := json.Marshal(VapeReponseAll{
		len(vapes),
		vapes,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsn)
}

func VapeCreateHandler(w http.ResponseWriter, r *http.Request) {
	v := &Vape{}
	db.Create(v)
}

func VapeDeleteHandler(w http.ResponseWriter, r *http.Request) {
	v := &Vape{}
	db.Last(v)
	db.Delete(v)
}

func GetTotalVapes() int {
	vapes := []Vape{}
	db.Find(&vapes)
	return len(vapes)
}
