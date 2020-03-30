package controller

import (
	"auth/model"
	"time"
	authStorage "auth/storage"
	profileStorage "profile/storage"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"html/template"
	"net/http"
)


//Checks if session exists otherwise takes to login page
func getToken(w http.ResponseWriter, r *http.Request) *jwt.Token {
	c, err := r.Cookie("token")

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	tokenString := c.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	return token
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("profile.gtpl")

	var token = getToken(w, r)

	if token.Valid{
		t.Execute(w, nil)
		return
	}else{
		http.Redirect(w, r, "/login", http.StatusFound)
		return 
	}
}

func FollowHandler(w http.ResponseWriter, r *http.Request) {
	
	m := map[string]interface{}{}
	
	t, _ := template.ParseFiles("profile.gtpl")
	
	var token = getToken(w, r)

	if !token.Valid{
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	r.ParseForm()
	//fmt.Println(r.Form["username"][0])
	userPresent := authStorage.Users[r.Form["username"][0]]
	claims, _ := token.Claims.(jwt.MapClaims)
	followUser := authStorage.Users[claims["username"].(string)]

	if userPresent == followUser{
		m["Error"] = "Cant follow yourself"
		m["Success"] = nil
		t.Execute(w, m)
		return
	}

	for e := followUser.Followers.Front() ; e != nil ; e.Next(){
		k := e.Value.(model.User)
		if userPresent == k{
			m["Error"] = "User already followed!"
			m["Success"] = nil
			t.Execute(w, m)
			return	
		}
	}
	if userPresent.Username != "" {
		followUser.Followers.PushBack(userPresent)
		authStorage.Users[followUser.Username] = followUser
		m["Error"] = nil
		m["Success"] = "Succesfully followed!"
		t.Execute(w, m)
		return
	} else {
		m["Error"] = "Username doesnt exist!"
		m["Success"] = nil
		t.Execute(w, m)
		return
	}
}

func TweetHandler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("profile.gtpl")
	m := map[string]interface{}{}

	var token = getToken(w, r)

	if !token.Valid{
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	r.ParseForm()
	//fmt.Println(r.Form["username"][0])
	tweetContent := r.Form["tweet"][0]

	if tweetContent != "" {
		claims, _ := token.Claims.(jwt.MapClaims)
		tweetUser := claims["username"].(string)
		for e := profileStorage.Tweets[tweetUser].Front(); e != nil; e = e.Next() {
			fmt.Println("before",e.Value)
		}
		profileStorage.Tweets[tweetUser].PushBack(tweetContent)
		for e := profileStorage.Tweets[tweetUser].Front(); e != nil; e = e.Next() {
			fmt.Println("after",e.Value)
		}

		m["Error"] = nil
		m["Success"] = "Succesfully tweeted!"
		t.Execute(w, m)
		return

	} else {

		m["Error"] = "Enter tweet content"
		m["Success"] = nil
		t.Execute(w, m)
		return
	}
}

func FeedHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("profile.gtpl")
	m := map[string]interface{}{}

	var token = getToken(w, r)

	if !token.Valid{
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	feedUserName := claims["username"].(string)
	feedUser := authStorage.Users[feedUserName]
	feed := ""
	for e:= feedUser.Followers.Front(); e != nil; e = e.Next(){
		followUser := e.Value.(model.User)
		tweetList := profileStorage.Tweets[followUser.Username]
		numOfTweets := 5
		feed = feed + " Top 5 tweets from "+ followUser.Username + " : \n"
		for k := tweetList.Back(); k != nil && numOfTweets > 0; k = k.Prev() {
			numOfTweets = numOfTweets - 1
			feed = feed + k.Value.(string) + "\n"
		}
	}
	if feed != "" {
		fmt.Println("feed succesfull")
		fmt.Println(feed)
		m["Error"] = nil
		m["Success"] = nil
		m["Feed"] = feed
		t.Execute(w, m)
		// res.Result = feed
		// json.NewEncoder(w).Encode(res)
		return
	} else {		
		m["Error"] = "No feed"
		m["Success"] = nil
		t.Execute(w, m)
		return
	}
}

func SignoutHandler(w http.ResponseWriter, r *http.Request) {
	
	var token = getToken(w, r)

	if !token.Valid {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	signoutUserName := claims["username"].(string)
	signoutUser := authStorage.Users[signoutUserName]

	if signoutUser.Username != "" {

		signoutUser.Token = ""
		authStorage.Users[signoutUserName] = signoutUser
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   "",
			Expires: time.Unix(0, 0),
		})
		fmt.Println(authStorage.Users)
		fmt.Println("Logout succesfull")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	} else {
		// res.Error = "Couldnt logout"
		// json.NewEncoder(w).Encode(res)
		fmt.Println("Logout error")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}