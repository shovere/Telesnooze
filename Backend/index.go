package main

import (
	"database/sql"
	//"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)


func setupDB() *sql.DB {
	db, err := sql.Open("sqlite3","telesnooze.db");
	fmt.Print(err);
	return db;
}

func sayHello(writer http.ResponseWriter, request *http.Request){
	fmt.Println("hello new user")
	writer.Header().Set("Content-Type", "application/json")
}

func main(){
	setupDB()
	router := mux.NewRouter();
	router.HandleFunc("/api/v1/", sayHello).Methods("GET");
	fmt.Println("Server at 8123")
    log.Fatal(http.ListenAndServe(":8123", router))
}