package repository

import (
	authStorage "auth/storage"
	authmodel "auth/model"
	//"container/list"
	//profilemodel "profile/model"
	//profileStorage "profile/storage"
	//jwt "github.com/dgrijalva/jwt-go"
	//"golang.org/x/crypto/bcrypt"
)

func ReturnUser(username string)(authmodel.User, bool){
	resultChan := make(chan authmodel.User)
	errChan := make(chan bool)
	go authStorage.ReturnUserDB(username,resultChan,errChan)
	return <-resultChan, <-errChan
}

func SaveUser(user authmodel.User){
	resultChan := make(chan bool)
	go authStorage.SaveUserDB(user,resultChan)
	<-resultChan
}

func SetCurrentUser(username string, user authmodel.User) {
	resultChan := make(chan bool)
	go authStorage.SetCurrentUserDB(username,user,resultChan)
	<-resultChan
}

