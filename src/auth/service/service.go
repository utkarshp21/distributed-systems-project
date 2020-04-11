package service

import (
	authmodel "auth/model"
	repository "auth/repository"
	"container/list"
	//"log"
	//"html/template"
	"net/http"
)

func RegisterService (r *http.Request)(string){

	r.ParseForm()

	_, usernameExists :=  repository.ReturnUser(r.Form["username"][0])

	if usernameExists {
		return "Email already in use!"
	}

	registerFromInput := authmodel.User{
		Username: r.Form["username"][0],
		Password: r.Form["password"][0],
		FirstName: r.Form["firstname"][0],
		LastName: r.Form["lastname"][0],
		Followers: list.New(),
	}

	err := repository.SaveUser(registerFromInput)

	if err != nil {
		return "Error While Hashing Password, Try Again"
	}

	return ""
}
