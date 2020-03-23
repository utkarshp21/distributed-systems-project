package controller

import (
	// "context"
	"encoding/json"
	"html/template"
	"fmt"
	"auth/model"
	// "io/ioutil"
	// "log"
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var Users []model.User

func UserExsits(email string) model.User{
	if len(Users) != 0 {
		for k,v := range Users {
			//Check if user is already in the database
			if v.Username == email {
				return Users[k]
			}
		}
	}
	
	return model.User{}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
        t, _ := template.ParseFiles("register.gtpl")
		t.Execute(w, nil)
		return 
    }else{
		var res model.ResponseResult

		w.Header().Set("Content-Type", "application/json")

		r.ParseForm()

		userPresent := UserExsits(r.Form["username"][0])
		
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
		}


		fmt.Println(user)

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

		if err != nil {
			res.Error = "Error While Hashing Password, Try Again"
			json.NewEncoder(w).Encode(res)
			return
		}

		user.Password = string(hash)

		Users = append(Users, user)

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

	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()

	var res model.ResponseResult
		
	userPresent := UserExsits(r.Form["username"][0])

	if userPresent.Username != "" {
		var err = bcrypt.CompareHashAndPassword([]byte(userPresent.Password), []byte(r.Form["password"][0]))
		fmt.Println("login error", err)
		if err != nil {
			res.Error = "Invalid password"
			json.NewEncoder(w).Encode(res)
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
		userPresent.Password = ""

		json.NewEncoder(w).Encode(userPresent)

	}else{
		res.Error = "Invalid username"
		json.NewEncoder(w).Encode(res)
		return
	}

}