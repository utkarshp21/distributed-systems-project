package repository

import (
	"container/list"
	"context"
	profileStorage "profile/storage"
	authmodel "auth/model"
	"time"

	//authRepository "auth/repository"
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
	dummyList := list.New()
	go profileStorage.GetTweetListDB(followUsername,resultChan)
	select {
	case res := <-resultChan:
		return res,nil
	case <-ctx.Done():
		return dummyList, ctx.Err()
	}
}

func InitialiseTweets(user authmodel.User,ctx context.Context)(error){
	resultChan := make(chan bool)
	go profileStorage.InitialiseTweetsDB(user, resultChan)
	select {
	case <-resultChan:
		return nil
	case <-ctx.Done():
		//authRepository.DeleteUser(user)
		return ctx.Err()
	}
}