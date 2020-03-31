package controller

import (
	authmodel "auth/model"
	profilemodel "profile/model"
	"time"
	authStorage "auth/storage"
	profileStorage "profile/storage"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"html/template"
	"net/http"
)


func getToken(c *http.Cookie) (*jwt.Token,error) {

	tokenString := c.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	return token, err
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("profile.gtpl")
	c, err := r.Cookie("token")

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	token, tokenerr := getToken(c)

	if token.Valid && tokenerr == nil{
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
	c, err := r.Cookie("token")

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	token, tokenerr := getToken(c)

	if !token.Valid || tokenerr != nil{
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	r.ParseForm()

	authmodel.UsersMux.Lock()
	userPresent := authStorage.Users[r.Form["username"][0]]
	authmodel.UsersMux.Unlock()

	claims, _ := token.Claims.(jwt.MapClaims)

	authmodel.UsersMux.Lock()
	followUser := authStorage.Users[claims["username"].(string)]
	authmodel.UsersMux.Unlock()

	if userPresent == followUser{
		m["Error"] = "Cant follow yourself"
		m["Success"] = nil
		t.Execute(w, m)
		return
	}

	for e := followUser.Followers.Front() ; e != nil ; e.Next(){
		k := e.Value.(authmodel.User)
		if userPresent == k{
			m["Error"] = "User already followed!"
			m["Success"] = nil
			t.Execute(w, m)
			return
		}
	}
	if userPresent.Username != "" {
		followUser.Followers.PushBack(userPresent)

		authmodel.UsersMux.Lock()
		authStorage.Users[followUser.Username] = followUser
		authmodel.UsersMux.Unlock()

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

	c, err := r.Cookie("token")

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	token, tokenerr := getToken(c)

	if !token.Valid || tokenerr != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	r.ParseForm()
	tweetContent := r.Form["tweet"][0]

	if tweetContent != "" {
		claims, _ := token.Claims.(jwt.MapClaims)
		tweetUser := claims["username"].(string)

		profilemodel.TweetsMux.Lock()
		profileStorage.Tweets[tweetUser].PushBack(tweetContent)
		profilemodel.TweetsMux.Unlock()

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

	c, err := r.Cookie("token")

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	token, tokenerr := getToken(c)

	if !token.Valid || tokenerr != nil{
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	feedUserName := claims["username"].(string)

	authmodel.UsersMux.Lock()
	feedUser := authStorage.Users[feedUserName]
	authmodel.UsersMux.Unlock()

	feed := ""
	for e:= feedUser.Followers.Front(); e != nil; e = e.Next(){
		followUser := e.Value.(authmodel.User)

		profilemodel.TweetsMux.Lock()
		tweetList := profileStorage.Tweets[followUser.Username]
		profilemodel.TweetsMux.Unlock()

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
		return
	} else {
		m["Error"] = "No feed"
		m["Success"] = nil
		t.Execute(w, m)
		return
	}
}

func SignoutHandler(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("token")

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	token, tokenerr := getToken(c)

	if !token.Valid || tokenerr != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	signoutUserName := claims["username"].(string)

	authmodel.UsersMux.Lock()
	signoutUser := authStorage.Users[signoutUserName]
	authmodel.UsersMux.Unlock()

	if signoutUser.Username != "" {

		signoutUser.Token = ""

		authmodel.UsersMux.Lock()
		authStorage.Users[signoutUserName] = signoutUser
		authmodel.UsersMux.Unlock()

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   "",
			Expires: time.Unix(0, 0),
		})
		fmt.Println("Logout succesfull")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	} else {
		fmt.Println("Logout error")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}