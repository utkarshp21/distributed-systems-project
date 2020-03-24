package main

import (
    "net/http"
    // "strings"
    "log"
    authController "auth/controller"
    profileController "profile/controller"
    // "auth/model"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/register", authController.RegisterHandler)
    r.HandleFunc("/login", authController.LoginHandler)
    r.HandleFunc("/profile",profileController.ProfileHandler)
    r.HandleFunc("/follow", profileController.FollowHandler)
    r.HandleFunc("/tweet", profileController.TweetHandler)
    r.HandleFunc("/feed", profileController.FeedHandler)
    r.HandleFunc("/signout", profileController.SignoutHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}