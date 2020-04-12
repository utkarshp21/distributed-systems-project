package controller

import (
	//profileRepository "profile/repository"
	//authRepository "auth/repository"
	//"time"
	"log"
	//jwt "github.com/dgrijalva/jwt-go"
	"html/template"
	"net/http"
	service "profile/service"
)

func redirectToLogin(w http.ResponseWriter){
	t, _ := template.ParseFiles("login.gtpl")
	m := map[string]interface{}{}
	m["Error"] = "Please login to continue!"
	m["Success"] = nil
	log.Println("Please login to continue")
	t.Execute(w, m)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	
	t, _ := template.ParseFiles("profile.gtpl")

	err := service.ProfileService(r)

	if err != nil {
		redirectToLogin(w)
		return
	}else{
		log.Println("Profile loaded succesfully")
		t.Execute(w, nil)
		return
	}
}

func FollowHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return 
	}
	m := map[string]interface{}{}
	t, _ := template.ParseFiles("profile.gtpl")

	loginerr, err := service.FollowService(r)

	if loginerr != nil {
		redirectToLogin(w)
		return
	}else if err != "" {
		m["Error"] = err
		m["Success"] = nil
		log.Println(err)
		t.Execute(w, m)
		return
	}else {
		m["Error"] = nil
		m["Success"] = "Succesfully followed"
		log.Println("Succesfully followed")
		t.Execute(w, m)
		return
	}
}

func UnfollowHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}
	m := map[string]interface{}{}
	t, _ := template.ParseFiles("profile.gtpl")

	loginerr, err := service.UnfollowService(r)

	if loginerr != nil {
		redirectToLogin(w)
		return
	}else if err != "" {
		m["Error"] = err
		m["Success"] = nil
		log.Println(err)
		t.Execute(w, m)
		return
	}else {
		m["Error"] = nil
		m["Success"] = "Succesfully unfollowed"
		log.Println("Succesfully unfollowed")
		t.Execute(w, m)
		return
	}
}

func TweetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return 
	}

	t, _ := template.ParseFiles("profile.gtpl")
	m := map[string]interface{}{}

	loginerr, err := service.TweetService(r)

	if loginerr != nil {
		redirectToLogin(w)
		return
	}else if err != "" {
		m["Error"] = err
		m["Success"] = nil
		log.Println(err)
		t.Execute(w, m)
		return
	}else {
		m["Error"] = nil
		m["Success"] = "Succesfully tweeted"
		log.Println("Succesfully tweeted")
		t.Execute(w, m)
		return
	}
}

func FeedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	t, _ := template.ParseFiles("profile.gtpl")
	m := map[string]interface{}{}

	loginerr, err, feed := service.FeedService(r)

	if loginerr != nil {
		redirectToLogin(w)
		return
	}else if err != "" {
		m["Error"] = err
		m["Success"] = nil
		log.Println(err)
		t.Execute(w, m)
		return
	}else {
		m["Error"] = nil
		m["Success"] = nil
		m["Feed"] = feed
		log.Println("Feed Succesfull")
		t.Execute(w, m)
		return
	}
}
