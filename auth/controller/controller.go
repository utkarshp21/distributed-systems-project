package controller

import (
	"auth/model"
	"container/list"
	authStorage "auth/storage"
	profileStorage "profile/storage"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)


func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	m := map[string]interface{}{}

	t, _ := template.ParseFiles("register.gtpl")

	if r.Method == "GET" {
		t.Execute(w, m)
		return 
    }else{
		
		r.ParseForm()

		userPresent := authStorage.Users[r.Form["username"][0]]
		
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

		authStorage.Users[user.Username] = user
		profileStorage.Tweets[user.Username] = list.New()

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
		
	userPresent := authStorage.Users[r.Form["username"][0]]

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

		authStorage.Users[userPresent.Username] = userPresent

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
		})

		http.Redirect(w, r, "/profile", http.StatusFound)
		return

	}

	m["Error"] = "Invalid username"
	t.Execute(w, m)
	return	
}