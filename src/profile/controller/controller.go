package controller

import (
	profileRepository "profile/repository"
	authRepository "auth/repository"
	//"time"
	"log"
	jwt "github.com/dgrijalva/jwt-go"
	"html/template"
	"net/http"
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
	c, err := r.Cookie("token")

	if err != nil {
		redirectToLogin(w)
		return
	}

	token, tokenerr := profileRepository.GetToken(c)

	if token.Valid && tokenerr == nil{
		log.Println("Profile loaded succesfully")
		t.Execute(w, nil)
		return
	}else{
		redirectToLogin(w)
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
	c, err := r.Cookie("token")

	if err != nil {
		redirectToLogin(w)
		return
	}

	token, tokenerr := profileRepository.GetToken(c)

	if !token.Valid || tokenerr != nil{
		redirectToLogin(w)
		return
	}

	r.ParseForm()

	userPresent, _ := authRepository.ReturnUser(r.Form["username"][0])

	claims, _ := token.Claims.(jwt.MapClaims)

	followUser, _ := authRepository.ReturnUser(claims["username"].(string))

	if userPresent == followUser{
		m["Error"] = "Cant follow yourself"
		m["Success"] = nil
		log.Println("Cant follow yourself")
		t.Execute(w, m)
		return
	}

	alreadyFollowed := profileRepository.CheckIfUserAlreadyFollowed(followUser, userPresent)

	if alreadyFollowed {
		m["Error"] = "User already followed!"
		m["Success"] = nil
		log.Println("User already followed!")
		t.Execute(w, m)
		return
	}
	
	if userPresent.Username != "" {
		profileRepository.FollowUser(userPresent,followUser)
		m["Error"] = nil
		m["Success"] = "Succesfully followed!"
		log.Println("Succesfully followed!")
		t.Execute(w, m)
		return
	} else {
		m["Error"] = "Username doesnt exist!"
		m["Success"] = nil
		log.Println("Username doesnt exist!")
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

	c, err := r.Cookie("token")

	if err != nil {
		redirectToLogin(w)
		return
	}

	token, tokenerr := profileRepository.GetToken(c)

	if !token.Valid || tokenerr != nil {
		redirectToLogin(w)
		return
	}
	r.ParseForm()

	tweetContent := r.Form["tweet"][0]

	if tweetContent != "" {
		claims, _ := token.Claims.(jwt.MapClaims)
		tweetUser := claims["username"].(string)
		profileRepository.SaveTweet(tweetUser,tweetContent)
		m["Error"] = nil
		m["Success"] = "Succesfully tweeted!"
		log.Println("Succesfully tweeted!")
		t.Execute(w, m)
		return

	} else {

		m["Error"] = "Enter tweet content"
		m["Success"] = nil
		log.Println("Enter tweet content")
		t.Execute(w, m)
		return
	}
}

func FeedHandler(w http.ResponseWriter, r *http.Request) {
	
	t, _ := template.ParseFiles("profile.gtpl")
	m := map[string]interface{}{}

	c, err := r.Cookie("token")

	if err != nil {
		redirectToLogin(w)
		return
	}

	token, tokenerr := profileRepository.GetToken(c)

	if !token.Valid || tokenerr != nil{
		redirectToLogin(w)
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	feedUser, _ := authRepository.ReturnUser(claims["username"].(string))

	feed := profileRepository.FeedGenerator(feedUser)

	if feed != "" {
		log.Println("Feed succesfull")
		m["Error"] = nil
		m["Success"] = nil
		m["Feed"] = feed
		t.Execute(w, m)
		return
	} else {
		m["Error"] = "No feed"
		m["Success"] = nil
		log.Println("No feed")
		t.Execute(w, m)
		return
	}
}
