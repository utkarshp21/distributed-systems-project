package controller

import (
	authmodel "auth/model"
	authStorage "auth/storage"
	repository "auth/repository"
	"container/list"
	"fmt"
	"log"
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

		usernameExists :=  repository.CheckUserExists(r.Form["username"][0])

		if usernameExists == true {
			m["Error"] = "Email already in use!"
			log.Println("Email already in use!")
			t.Execute(w, m)
			return 
		}

		r.ParseForm()
		
		registerFromInput := authmodel.User{
			Username: r.Form["username"][0],
			Password: r.Form["password"][0],
			FirstName: r.Form["firstname"][0],
			LastName: r.Form["lastname"][0],
			Followers: list.New(),
		}

		err := repository.SaveUser(registerFromInput)

		if err != nil {
			m["Error"] = "Error While Hashing Password, Try Again"
			log.Println("Error While Hashing Password, Try Again")
			t.Execute(w, m)
			return
		}

		log.Println("User Registered succesfully")
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	
	return
}

func LoginUser(userPresent authmodel.User) (string,error){

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  userPresent.Username,
		"firstname": userPresent.FirstName,
		"lastname":  userPresent.LastName,
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		return "",err
	}

	userPresent.Token = tokenString

	authmodel.UsersMux.Lock()
	authStorage.Users[userPresent.Username] = userPresent
	authmodel.UsersMux.Unlock()

	return tokenString,nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	
	t, _ := template.ParseFiles("login.gtpl")

	m := map[string]interface{}{}

	if r.Method == "GET" {
		t.Execute(w, nil)
		return 
	}

	r.ParseForm()

	authmodel.UsersMux.Lock()
	userPresent := authStorage.Users[r.Form["username"][0]]
	authmodel.UsersMux.Unlock()

	if userPresent.Username != "" {
		var err = bcrypt.CompareHashAndPassword([]byte(userPresent.Password), []byte(r.Form["password"][0]))
		if err != nil {
			m["Error"] = "Invalid password"
			fmt.Println("Invalid password")
			t.Execute(w, m)
			return			
		}

		tokenString, loginerr := LoginUser(userPresent)

		if loginerr != nil {
			m["Error"] = "Error while generating token,Try again"
			fmt.Println("Error while generating token,Try again")
			t.Execute(w, m)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
		})
		fmt.Println("Login successful")
		http.Redirect(w, r, "/profile", http.StatusFound)
		return

	}

	m["Error"] = "Invalid username"
	fmt.Println("Invalid username")
	t.Execute(w, m)
	return	
}