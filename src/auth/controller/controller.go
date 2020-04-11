package controller

import (
	//authmodel "auth/model"
	//repository "auth/repository"
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
		}else{
			log.Println("User Registered succesfully")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	
	t, _ := template.ParseFiles("login.gtpl")

	m := map[string]interface{}{}

	if r.Method == "GET" {
		t.Execute(w, nil)
		return 
	}else{

		errMsg := service.LoginService(w,r)

		if errMsg != "" {
			m["Error"] = errMsg
			log.Println(errMsg)
			t.Execute(w, m)
			return
		}else{
			log.Println("Login successful")
			http.Redirect(w, r, "/profile", http.StatusFound)
			return
		}
	}
}