package repository

import (
	authStorage "auth/storage"
	authmodel "auth/model"
	"context"
	"time"
)

func ReturnUser(username string, ctx context.Context)(authmodel.User, bool, error){
	resultChan := make(chan authmodel.User)
	errChan := make(chan bool)
	dummy := new(authmodel.User)
	dummyUser := *dummy
	go authStorage.ReturnUserDB(username,resultChan,errChan)

	select {
	case res := <-resultChan :
		return res, <-errChan, nil
	case <-ctx.Done():
		return dummyUser,false,ctx.Err()
	}

}

func SaveUserRegister(user authmodel.User, ctx context.Context)(error){
	time.Sleep(10*time.Millisecond)
	resultChan := make(chan bool)
	deleteChan := make(chan bool)
	go authStorage.SaveUserRegisterDB(user,resultChan,deleteChan,ctx)

	select {
	case <-resultChan:
		return nil
	case <-deleteChan:
		return ctx.Err()
	}

}


func SaveUser(user authmodel.User, ctx context.Context, bkpUser authmodel.User)(error){
	time.Sleep(10*time.Millisecond)
	resultChan := make(chan bool)
	deleteChan := make(chan bool)
	go authStorage.SaveUserDB(user,bkpUser,resultChan,deleteChan,ctx)

	select {
	case <-resultChan:
		return nil
	case <-deleteChan:
		return ctx.Err()
	}

}



