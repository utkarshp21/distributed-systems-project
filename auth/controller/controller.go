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
	// jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var Users []model.User

func UserExsits(email string) bool{
	if len(Users) != 0 {
		for _,v := range Users {
			//Check if user is already in the database
			if v.Username == email {
				return true
			}
		}
	}
	return false
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
		
		if userPresent {
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

		return
	}
	
	
	

	// var result model.User
	// err = collection.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&result)

	// if err != nil {
	// 	if err.Error() == "mongo: no documents in result" {
	// 		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

	// 		if err != nil {
	// 			res.Error = "Error While Hashing Password, Try Again"
	// 			json.NewEncoder(w).Encode(res)
	// 			return
	// 		}
	// 		user.Password = string(hash)

	// 		_, err = collection.InsertOne(context.TODO(), user)
	// 		if err != nil {
	// 			res.Error = "Error While Creating User, Try Again"
	// 			json.NewEncoder(w).Encode(res)
	// 			return
	// 		}
	// 		res.Result = "Registration Successful"
	// 		json.NewEncoder(w).Encode(res)
	// 		return
	// 	}

	// 	res.Error = err.Error()
	// 	json.NewEncoder(w).Encode(res)
	// 	return
	// }

	// res.Result = "Username already Exists!!"
	// json.NewEncoder(w).Encode(res)
	return
}

// func RegisterHandler(w http.ResponseWriter, r *http.Request) {
// 	return "test"
// }
