package controller

import (
	"auth/model"
	"container/list"
	"time"

	// "context"
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	// "io/ioutil"
	// "log"
	"net/http"
)

var Users = make(map[string]model.User)
var Tweets = make(map[string]*list.List)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	m := map[string]interface{}{}

	t, _ := template.ParseFiles("register.gtpl")

	if r.Method == "GET" {
		t.Execute(w, m)
		return 
    }else{
		
		r.ParseForm()

		userPresent := Users[r.Form["username"][0]]
		
		if userPresent.Username != "" {
			m["Error"] = "Email already in use!"
			t.Execute(w, m)
			return 
		}
	
		user := model.User{
			Username:r.Form["username"][0], 
			Password: r.Form["password"][0],
			FirstName: r.Form["firstname"][0],
			LastName: r.Form["lastname"][0],
			Followers: list.New(),
		}

		//fmt.Println(user)

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

		if err != nil {
			m["Error"] = "Error While Hashing Password, Try Again"
			t.Execute(w, m)
			return
		}

		user.Password = string(hash)

		Users[user.Username] = user
		Tweets[user.Username] = list.New()

		http.Redirect(w, r, "/login", http.StatusFound)
	}
	
	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	
	t, _ := template.ParseFiles("login.gtpl")

	m := map[string]interface{}{}

	if r.Method == "GET" {
		t.Execute(w, nil)
		return 
	}

	r.ParseForm()
		
	userPresent := Users[r.Form["username"][0]]

	if userPresent.Username != "" {
		var err = bcrypt.CompareHashAndPassword([]byte(userPresent.Password), []byte(r.Form["password"][0]))
		fmt.Println("login error", err)
		if err != nil {
			m["Error"] = "Invalid password"
			t.Execute(w, m)
			return			
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username":  userPresent.Username,
			"firstname": userPresent.FirstName,
			"lastname":  userPresent.LastName,
		})

		tokenString, err := token.SignedString([]byte("secret"))

		if err != nil {
			m["Error"] = "Error while generating token,Try again"
			t.Execute(w, m)
			return	
		}

		userPresent.Token = tokenString

		Users[userPresent.Username] = userPresent

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
		})

		http.Redirect(w, r, "/dashboard", http.StatusFound)
		return

	}

	m["Error"] = "Invalid username"
	t.Execute(w, m)
	return	
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	
	t, _ := template.ParseFiles("profile.gtpl")
	c, err := r.Cookie("token")

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		// if err == http.ErrNoCookie {
		// 	// If the cookie is not set, return an unauthorized status
		// 	//w.WriteHeader(http.StatusUnauthorized)
		// 	http.Redirect(w, r, "/login", http.StatusFound)
		// 	return	
		// }
		// // For any other type of error, return a bad request status
		// http.Redirect(w, r, "/login", http.StatusFound)
		// return
	}
	tokenString := c.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

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

	c, err := r.Cookie("token")
	if err != nil {
		// if err == http.ErrNoCookie {
		// 	// If the cookie is not set, return an unauthorized status
		// 	//w.WriteHeader(http.StatusUnauthorized)
		// 	res.Error = "No cookie"
		// 	json.NewEncoder(w).Encode(res)
		// 	return
		// }
		// // For any other type of error, return a bad request status
		// //w.WriteHeader(http.StatusBadRequest)
		// res.Error = "Bad request"
		// json.NewEncoder(w).Encode(res)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	tokenString := c.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if !token.Valid{
		// res.Error = "Invalid token"
		// json.NewEncoder(w).Encode(res)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	r.ParseForm()
	//fmt.Println(r.Form["username"][0])
	userPresent := Users[r.Form["username"][0]]
	claims, _ := token.Claims.(jwt.MapClaims)
	followUser := Users[claims["username"].(string)]

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
		Users[followUser.Username] = followUser
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
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			//w.WriteHeader(http.StatusUnauthorized)
			// res.Error = "No cookie"
			// json.NewEncoder(w).Encode(res)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		// For any other type of error, return a bad request status
		//w.WriteHeader(http.StatusBadRequest)
		// res.Error = "Bad request"
		// json.NewEncoder(w).Encode(res)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	tokenString := c.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if !token.Valid{
		// res.Error = "Invalid token"
		// json.NewEncoder(w).Encode(res)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	r.ParseForm()
	//fmt.Println(r.Form["username"][0])
	tweetContent := r.Form["tweet"][0]

	if tweetContent != "" {
		claims, _ := token.Claims.(jwt.MapClaims)
		tweetUser := claims["username"].(string)
		for e := Tweets[tweetUser].Front(); e != nil; e = e.Next() {
			fmt.Println("before",e.Value)
		}
		Tweets[tweetUser].PushBack(tweetContent)
		for e := Tweets[tweetUser].Front(); e != nil; e = e.Next() {
			fmt.Println("after",e.Value)
		}
		//Users[followUser.Username] = followUser

		m["Error"] = nil
		m["Success"] = "Succesfully tweeted!"
		t.Execute(w, m)
		return

	} else {

		m["Error"] = "Username doesnt exist"
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
		// if err == http.ErrNoCookie {
		// 	// If the cookie is not set, return an unauthorized status
		// 	//w.WriteHeader(http.StatusUnauthorized)
			
		// 	http.Redirect(w, r, "/login", http.StatusFound)
		// 	return
		// }
		// // For any other type of error, return a bad request status
		// //w.WriteHeader(http.StatusBadRequest)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	tokenString := c.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if !token.Valid{
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	feedUserName := claims["username"].(string)
	feedUser := Users[feedUserName]
	feed := ""
	for e:= feedUser.Followers.Front(); e != nil; e = e.Next(){
		followUser := e.Value.(model.User)
		tweetList := Tweets[followUser.Username]
		numOfTweets := 5
		feed = feed + " Top 5 tweets from "+ followUser.Username + " : \n"
		for k := tweetList.Back(); k != nil && numOfTweets > 0; k = k.Prev() {
			numOfTweets = numOfTweets - 1
			feed = feed + k.Value.(string) + "\n"
		}
	}
	if feed != "" {
		fmt.Println("feed succesfull")
		res.Result = feed
		json.NewEncoder(w).Encode(res)
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
		// if err == http.ErrNoCookie {
		// 	// If the cookie is not set, return an unauthorized status
		// 	//w.WriteHeader(http.StatusUnauthorized)
		// 	res.Error = "No cookie"
		// 	json.NewEncoder(w).Encode(res)
		// 	return
		// }
		// // For any other type of error, return a bad request status
		// //w.WriteHeader(http.StatusBadRequest)
		// res.Error = "Bad request"
		// json.NewEncoder(w).Encode(res)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	tokenString := c.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if !token.Valid || err != nil{
		// res.Error = "Invalid token"
		// json.NewEncoder(w).Encode(res)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	signoutUserName := claims["username"].(string)
	signoutUser := Users[signoutUserName]

	if signoutUser.Username != "" {

		signoutUser.Token = ""
		Users[signoutUserName] = signoutUser
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   "",
			Expires: time.Unix(0, 0),
		})
		fmt.Println(Users)
		fmt.Println("Logout succesfull")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	} else {
		res.Error = "Couldnt logout"
		json.NewEncoder(w).Encode(res)
		return
	}
}