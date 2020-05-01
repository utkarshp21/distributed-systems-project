package storage

import (
	authmodel "backend/model"
	"context"
)

var Users = make(map[string]authmodel.User)

func ReturnUserDB(username string, resultChan chan authmodel.User, errChan chan bool,deleteChan chan bool, ctx context.Context)  {
	authmodel.UsersMux.Lock()
	user, exists := Users[username]
	select {
	case <-ctx.Done():
		authmodel.UsersMux.Unlock()
		deleteChan <- true
	default:
		authmodel.UsersMux.Unlock()
		resultChan <- user
		errChan <- exists
	}
}

func SaveUserRegisterDB(user authmodel.User,resultChan chan bool,deleteChan chan bool,ctx context.Context)  {
	authmodel.UsersMux.Lock()
	Users[user.Username] = user

	select {
	case <-ctx.Done():
		authmodel.UsersMux.Unlock()
		channel := make(chan bool)
		go DeleteUserDB(user,channel)
		<-channel
		deleteChan <- true
	default:
		authmodel.UsersMux.Unlock()
		resultChan <- true
	}
}

func DeleteUserDB(user authmodel.User,resultChan chan bool)  {
	authmodel.UsersMux.Lock()
	delete(Users,user.Username)
	authmodel.UsersMux.Unlock()
	resultChan <- true
	return
}

func SaveUserDB(user authmodel.User,bkpUser authmodel.User,resultChan chan bool,deleteChan chan bool,ctx context.Context)  {
	authmodel.UsersMux.Lock()
	Users[user.Username] = user

	select {
	case <-ctx.Done():
		authmodel.UsersMux.Unlock()
		channel := make(chan bool)
		go ModifyUserDB(bkpUser,channel)
		<-channel
		deleteChan <- true
	default:
		authmodel.UsersMux.Unlock()
		resultChan <- true
	}
}

func ModifyUserDB(user authmodel.User, channel chan bool) {
	authmodel.UsersMux.Lock()
	Users[user.Username] = user
	authmodel.UsersMux.Unlock()
	channel <- true
}