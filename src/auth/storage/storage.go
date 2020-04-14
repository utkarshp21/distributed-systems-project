package storage

import (
	authmodel "auth/model"
	//authStorage "auth/storage"
)

var Users = make(map[string]authmodel.User)

func ReturnUserDB(username string, resultChan chan authmodel.User, errChan chan bool)  {
	authmodel.UsersMux.Lock()
	user, exists := Users[username]
	authmodel.UsersMux.Unlock()
	resultChan <- user
	errChan <- exists
}

func SaveUserDB(user authmodel.User,resultChan chan bool)  {
	authmodel.UsersMux.Lock()
	Users[user.Username] = user
	authmodel.UsersMux.Unlock()
	resultChan <- true
}
