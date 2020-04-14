package repository

import (
	authStorage "auth/storage"
	authmodel "auth/model"
	"context"
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

func SaveUser(user authmodel.User, ctx context.Context)(error){
	resultChan := make(chan bool)
	go authStorage.SaveUserDB(user,resultChan)

	select {
	case <-resultChan:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}

}


