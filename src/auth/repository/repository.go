package repository

import (
	authStorage "auth/storage"
	authmodel "auth/model"
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


