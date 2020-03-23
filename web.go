package main

import (
    "net/http"
    // "strings"
    "log"
    "auth/controller"
    // "auth/model"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/register", controller.RegisterHandler)
    r.HandleFunc("/login", controller.LoginHandler)  
	log.Fatal(http.ListenAndServe(":8080", r))
}