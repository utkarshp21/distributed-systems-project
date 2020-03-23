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
    r.HandleFunc("/dashboard",controller.DashboardHandler)
    r.HandleFunc("/follow", controller.FollowHandler)
    r.HandleFunc("/tweet", controller.TweetHandler)
    r.HandleFunc("/feed", controller.FeedHandler)
    r.HandleFunc("/signout", controller.SignoutHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}