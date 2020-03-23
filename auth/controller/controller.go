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

//func UserExsits(email string) model.User {
//	if len(Users) != 0 {
//		for k,v := range Users {
//			//Check if user is already in the database
//			if v.Username == email {
//				return Users[k]
//			}
//		}
//	}
//
//	return model.User{}
//}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
        t, _ := template.ParseFiles("register.gtpl")
		t.Execute(w, nil)
		return 
    }else{
		var res model.ResponseResult

		w.Header().Set("Content-Type", "application/json")

		r.ParseForm()

		userPresent := Users[r.Form["username"][0]]
		
		if userPresent.Username != "" {
			res.Result = "Email already Exists!!"
			json.NewEncoder(w).Encode(res)
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
			res.Error = "Error While Hashing Password, Try Again"
			json.NewEncoder(w).Encode(res)
			return
		}

		user.Password = string(hash)

		Users[user.Username] = user
		Tweets[user.Username] = list.New()

		res.Result = "Registration Successful"
		json.NewEncoder(w).Encode(res)

		fmt.Println(Users)

		// t, _ := template.ParseFiles("login.gtpl")
		// t.Execute(w, nil)
		return 
	}
	
	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
		return 
	}

	w.Header().Set("Content-Type", "text/html")

	r.ParseForm()

	var res model.ResponseResult
		
	userPresent := Users[r.Form["username"][0]]

	if userPresent.Username != "" {
		var err = bcrypt.CompareHashAndPassword([]byte(userPresent.Password), []byte(r.Form["password"][0]))
		fmt.Println("login error", err)
		if err != nil {
			res.Error = "Invalid password"
			json.NewEncoder(w).Encode(res)
			//http.Error(w,"Invalid password",302)
			//http.Redirect(w,r,"/login",302)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username":  userPresent.Username,
			"firstname": userPresent.FirstName,
			"lastname":  userPresent.LastName,
		})

		tokenString, err := token.SignedString([]byte("secret"))

		if err != nil {
			res.Error = "Error while generating token,Try again"
			json.NewEncoder(w).Encode(res)
			return
		}

		userPresent.Token = tokenString
		//userPresent.Password = ""

		Users[userPresent.Username] = userPresent

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
		})
		fmt.Println(Users)
		json.NewEncoder(w).Encode(userPresent)
		//http.Redirect(w,r,"/dashboard",302)

		return

	}else{
		res.Error = "Invalid username"
		json.NewEncoder(w).Encode(res)
		return
	}

}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var res model.ResponseResult
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			//w.WriteHeader(http.StatusUnauthorized)
			res.Error = "No cookie"
			json.NewEncoder(w).Encode(res)
			return
		}
		// For any other type of error, return a bad request status
		//w.WriteHeader(http.StatusBadRequest)
		res.Error = "Bad request"
		json.NewEncoder(w).Encode(res)
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

	if token.Valid{
		t, _ := template.ParseFiles("profile.gtpl")
		t.Execute(w, nil)
		return
	}else{
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
}


func FollowHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res model.ResponseResult
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			//w.WriteHeader(http.StatusUnauthorized)
			res.Error = "No cookie"
			json.NewEncoder(w).Encode(res)
			return
		}
		// For any other type of error, return a bad request status
		//w.WriteHeader(http.StatusBadRequest)
		res.Error = "Bad request"
		json.NewEncoder(w).Encode(res)
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
		res.Error = "Invalid token"
		json.NewEncoder(w).Encode(res)
		return
	}
	r.ParseForm()
	//fmt.Println(r.Form["username"][0])
	userPresent := Users[r.Form["username"][0]]
	claims, _ := token.Claims.(jwt.MapClaims)
	followUser := Users[claims["username"].(string)]

	if userPresent == followUser{
		res.Error = "Cant follow yourself"
		json.NewEncoder(w).Encode(res)
		return
	}

	for e := followUser.Followers.Front() ; e != nil ; e.Next(){
		k := e.Value.(model.User)
		if userPresent == k{
			res.Error = "User already followed"
			json.NewEncoder(w).Encode(res)
			return
		}
	}
	if userPresent.Username != "" {
		followUser.Followers.PushBack(userPresent)
		Users[followUser.Username] = followUser
		fmt.Println("succesfully followed")
		res.Result = "Succesfully followed"
		json.NewEncoder(w).Encode(res)
		return
	} else {
		res.Error = "Username doesnt exist"
		json.NewEncoder(w).Encode(res)
		return
	}
}

func TweetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res model.ResponseResult
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			//w.WriteHeader(http.StatusUnauthorized)
			res.Error = "No cookie"
			json.NewEncoder(w).Encode(res)
			return
		}
		// For any other type of error, return a bad request status
		//w.WriteHeader(http.StatusBadRequest)
		res.Error = "Bad request"
		json.NewEncoder(w).Encode(res)
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
		res.Error = "Invalid token"
		json.NewEncoder(w).Encode(res)
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
		fmt.Println("succesfully tweeted")
		res.Result = "Succesfully tweeted"
		json.NewEncoder(w).Encode(res)
		return
	} else {
		res.Error = "Username doesnt exist"
		json.NewEncoder(w).Encode(res)
		return
	}
}

func FeedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res model.ResponseResult
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			//w.WriteHeader(http.StatusUnauthorized)
			res.Error = "No cookie"
			json.NewEncoder(w).Encode(res)
			return
		}
		// For any other type of error, return a bad request status
		//w.WriteHeader(http.StatusBadRequest)
		res.Error = "Bad request"
		json.NewEncoder(w).Encode(res)
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
		res.Error = "Invalid token"
		json.NewEncoder(w).Encode(res)
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
		res.Error = "No feed"
		json.NewEncoder(w).Encode(res)
		return
	}
}

func SignoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res model.ResponseResult
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			//w.WriteHeader(http.StatusUnauthorized)
			res.Error = "No cookie"
			json.NewEncoder(w).Encode(res)
			return
		}
		// For any other type of error, return a bad request status
		//w.WriteHeader(http.StatusBadRequest)
		res.Error = "Bad request"
		json.NewEncoder(w).Encode(res)
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
		res.Error = "Invalid token"
		json.NewEncoder(w).Encode(res)
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
		res.Result = "Logout successful"
		json.NewEncoder(w).Encode(res)
		return
	} else {
		res.Error = "Couldnt logout"
		json.NewEncoder(w).Encode(res)
		return
	}
}