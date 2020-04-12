package main

import (
    "net/http"
    "log"
    authController "auth/controller"
    profileController "profile/controller"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    // r.HandleFunc("/", authController.LoginHandler)
    r.HandleFunc("/register", authController.RegisterHandler)
    // r.HandleFunc("/login", authController.LoginHandler)
    // r.HandleFunc("/profile",profileController.ProfileHandler)
    r.HandleFunc("/unfollow",profileController.UnfollowHandler)
    r.HandleFunc("/follow", profileController.FollowHandler)
    r.HandleFunc("/tweet", profileController.TweetHandler)
    r.HandleFunc("/feed", profileController.FeedHandler)
    // r.HandleFunc("/signout", authController.SignoutHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}