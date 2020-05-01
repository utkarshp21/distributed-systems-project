package repository

import (
	authmodel "backend/model"
	"context"
	"time"
)

var Users = make(map[string]authmodel.User)

func ReturnUser(username string, ctx context.Context)(authmodel.User, bool, error){
	resultChan := make(chan authmodel.User)
	errChan := make(chan bool)
	deleteChan := make(chan bool)
	dummy := new(authmodel.User)
	dummyUser := *dummy
	go ReturnUserDB(username,resultChan,errChan,deleteChan,ctx)

	select {
	case res := <-resultChan :
		return res, <-errChan, nil
	case <-deleteChan:
		return dummyUser,false,ctx.Err()
	}

}

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

func SaveUserRegister(user authmodel.User, ctx context.Context)(error){
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

func SaveUser(user authmodel.User, ctx context.Context, bkpUser authmodel.User)(error){
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


