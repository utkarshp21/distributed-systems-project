package service

import (
	authmodel "auth/model"
	repository "auth/repository"
	"container/list"
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

func LoginService(w http.ResponseWriter ,r *http.Request)(string)  {


	r.ParseForm()

	user, usernameExists :=  repository.ReturnUser(r.Form["username"][0])

	if usernameExists {

		passowrdErr := repository.CheckLoginPassword(user.Password, r.Form["password"][0])

		if  passowrdErr != nil {
			return "Invalid password"
		}

		tokenString, loginErr := repository.GenerateToken(user)

		if loginErr != nil {
			return "Error while generating token,Try again"
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
		})

		return ""

	}
	return "Invalid username"
}