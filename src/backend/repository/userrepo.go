package repository

import (
	model "backend/model"
	"context"
	"time"
)

var Users = make(map[string]model.User)

func ReturnUser(username string, ctx context.Context)(model.User, bool, error){
	resultChan := make(chan model.User)
	errChan := make(chan bool)
	deleteChan := make(chan bool)
	dummy := new(model.User)
	dummyUser := *dummy
	go ReturnUserDB(username,resultChan,errChan,deleteChan,ctx)

	select {
	case res := <-resultChan :
		return res, <-errChan, nil
	case <-deleteChan:
		return dummyUser,false,ctx.Err()
	}

}

func ReturnUserDB(username string, resultChan chan model.User, errChan chan bool,deleteChan chan bool, ctx context.Context)  {
	model.UsersMux.Lock()
	user, exists := Users[username]
	select {
	case <-ctx.Done():
		model.UsersMux.Unlock()
		deleteChan <- true
	default:
		model.UsersMux.Unlock()
		resultChan <- user
		errChan <- exists
	}
}

func SaveUserRegister(user model.User, ctx context.Context)(error){
	time.Sleep(10*time.Millisecond)
	resultChan := make(chan bool)
	deleteChan := make(chan bool)
	go SaveUserRegisterDB(user,resultChan,deleteChan,ctx)

	select {
	case <-resultChan:
		return nil
	case <-deleteChan:
		return ctx.Err()
	}

}

func SaveUserRegisterDB(user model.User,resultChan chan bool,deleteChan chan bool,ctx context.Context)  {
	model.UsersMux.Lock()
	Users[user.Username] = user

	select {
	case <-ctx.Done():
		model.UsersMux.Unlock()
		channel := make(chan bool)
		go DeleteUserDB(user,channel)
		<-channel
		deleteChan <- true
	default:
		model.UsersMux.Unlock()
		resultChan <- true
	}
}

func DeleteUserDB(user model.User,resultChan chan bool)  {
	model.UsersMux.Lock()
	delete(Users,user.Username)
	model.UsersMux.Unlock()
	resultChan <- true
	return
}

func SaveUser(user model.User, ctx context.Context, bkpUser model.User)(error){
	time.Sleep(10*time.Millisecond)
	resultChan := make(chan bool)
	deleteChan := make(chan bool)
	go SaveUserDB(user,bkpUser,resultChan,deleteChan,ctx)

	select {
	case <-resultChan:
		return nil
	case <-deleteChan:
		return ctx.Err()
	}

}

func SaveUserDB(user model.User,bkpUser model.User,resultChan chan bool,deleteChan chan bool,ctx context.Context)  {
	model.UsersMux.Lock()
	Users[user.Username] = user

	select {
	case <-ctx.Done():
		model.UsersMux.Unlock()
		channel := make(chan bool)
		go ModifyUserDB(bkpUser,channel)
		<-channel
		deleteChan <- true
	default:
		model.UsersMux.Unlock()
		resultChan <- true
	}
}

func ModifyUserDB(user model.User, channel chan bool) {
	model.UsersMux.Lock()
	Users[user.Username] = user
	model.UsersMux.Unlock()
	channel <- true
}

func GetUsers(ctx context.Context)(string,error){
	resultChan := make(chan string)
	deleteChan := make(chan bool)
	go GetUsersDB(resultChan,deleteChan,ctx)
	select {
	case res := <-resultChan:
		return res,nil
	case <-deleteChan:
		return "",ctx.Err()
	}
}

func GetUsersDB(resultChan chan string,deleteChan chan bool, ctx context.Context)  {
	model.UsersMux.Lock()
	keys := ""
	for k := range Users {
		keys += k + ","
	}
	keys = keys[:len(keys)-1]
	select {
	case <-ctx.Done():
		model.UsersMux.Unlock()
		deleteChan <- true
	default:
		model.UsersMux.Unlock()
		resultChan <- keys
	}
}
