package service

import (
	authmodel "auth/model"
	repository "auth/repository"
	"container/list"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	jwt "github.com/dgrijalva/jwt-go"
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

	hash, err := bcrypt.GenerateFromPassword([]byte(registerFromInput.Password), 5)
	if err != nil {
		return "Error While Hashing Password, Try Again"
	}
	registerFromInput.Password = string(hash)

	repository.SaveUser(registerFromInput)
	return ""
}

func LoginService(w http.ResponseWriter ,r *http.Request)(string)  {


	r.ParseForm()

	user, usernameExists :=  repository.ReturnUser(r.Form["username"][0])

	if usernameExists {
		passowrdErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Form["password"][0]))

		if  passowrdErr != nil {
			return "Invalid password"
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username":  user.Username,
			"firstname": user.FirstName,
			"lastname":  user.LastName,
		})

		tokenString, loginErr := token.SignedString([]byte("secret"))

		if loginErr != nil {
			return "Error while generating token,Try again"
		}

		user.Token = tokenString
		repository.SetCurrentUser(user.Username, user)

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
		})

		return ""

	}
	return "Invalid username"
}