package repository

import (
	"container/list"
	"context"
	profileStorage "profile/storage"
	authmodel "backend/model"
	"time"

	//authRepository "backend/repository"
)

func SaveTweet(tweetUser string,tweetContent string,ctx context.Context)(error){
	time.Sleep(10*time.Millisecond)
	resultChan := make(chan bool)
	deleteChan := make(chan bool)
	go profileStorage.SaveTweetDB(tweetUser,tweetContent,resultChan,deleteChan,ctx)
	select {
	case <-resultChan:
		return nil
	case <-deleteChan:
		return ctx.Err()
	}
}

func GetTweetList(followUsername string,ctx context.Context)(*list.List,error) {
	resultChan := make(chan *list.List)
	deleteChan := make(chan bool)
	dummyList := list.New()
	go profileStorage.GetTweetListDB(followUsername,resultChan,deleteChan,ctx)
	select {
	case res := <-resultChan:
		return res,nil
	case <-deleteChan:
		return dummyList, ctx.Err()
	}
}

func InitialiseTweets(user authmodel.User,ctx context.Context)(error){
	resultChan := make(chan bool)
	deleteChan := make(chan bool)
	go profileStorage.InitialiseTweetsDB(user, resultChan,deleteChan,ctx)
	select {
	case <-resultChan:
		return nil
	case <-deleteChan:
		//authRepository.DeleteUser(user)
		return ctx.Err()
	}
}

func GetUsers(ctx context.Context)(string,error)  {
	resultChan := make(chan string)
	deleteChan := make(chan bool)
	go profileStorage.GetUsersDB(resultChan,deleteChan,ctx)
	select {
	case res := <-resultChan:
		return res,nil
	case <-deleteChan:
		return "",ctx.Err()
	}
}