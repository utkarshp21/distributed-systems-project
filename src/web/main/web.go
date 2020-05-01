package main

import (
	//authController "backend/controller"
	"log"
	"net/http"
	//profileController "profile/controller"
	"github.com/gorilla/mux"
	controller "web/controller"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.LoginHandler)
	r.HandleFunc("/register", controller.RegisterHandler)
	r.HandleFunc("/login", controller.LoginHandler)
	r.HandleFunc("/unfollow", controller.UnfollowHandler)
	r.HandleFunc("/follow", controller.FollowHandler)
	r.HandleFunc("/tweet", controller.TweetHandler)
	r.HandleFunc("/feed", controller.FeedHandler)
	r.HandleFunc("/signout", controller.SignoutHandler)
	r.HandleFunc("/userlist", controller.UserListHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
