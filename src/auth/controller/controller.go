package controller

import (
	//authmodel "auth/model"
	repository "auth/repository"
	service "auth/service"
	//"container/list"
	"log"
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

    	errMsg := service.RegisterService(r)

    	if errMsg != "" {
			m["Error"] = errMsg
			log.Println(errMsg)
			t.Execute(w, m)
			return
		}

		log.Println("User Registered succesfully")
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

	user, usernameExists :=  repository.ReturnUser(r.Form["username"][0])

	if usernameExists {

		passowrdErr := repository.CheckLoginPassword(user.Password, r.Form["password"][0])
		
		if  passowrdErr != nil {
			m["Error"] = "Invalid password"
			log.Println("Invalid password")
			t.Execute(w, m)
			return			
		}

		tokenString, loginErr := repository.GenerateToken(user)

		if loginErr != nil {
			m["Error"] = "Error while generating token,Try again"
			log.Println("Error while generating token,Try again")
			t.Execute(w, m)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
		})

		log.Println("Login successful")
		http.Redirect(w, r, "/profile", http.StatusFound)

		return

	}

	m["Error"] = "Invalid username"
	log.Println("Invalid username")
	t.Execute(w, m)
	return	
}