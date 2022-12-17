package main

//Main file of our go project

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/y0gesh02/go-bookstore/pkg/routes"
)

func main(){
	r := mux.NewRouter()   //Initializing route
	routes.RegisterBookStoreRoutes(r)  //passing r to routes folder 
	http.Handle("/", r)  
	log.Fatal(http.ListenAndServe("localhost:8080", r)) //creating server
}