package main

import (
	"log"
	"net/http"

	"github.com/clarch/handler/rest"
	"github.com/jmoiron/sqlx"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()
	//all technology init here
	db, err := sqlx.Connect("postgres", "user=postgres dbname=election sslmode=disable password=postgres host=localhost")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	controller := rest.RestController{
		DB: db,
		//any dependency for controller injected here
	}

	router.HandleFunc("/join_election", controller.RESTJoinElection).Methods("GET")
	log.Println("Serving REST service...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
