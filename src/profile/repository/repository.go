package repository

import (
	"container/list"
	"context"
	profileStorage "profile/storage"
	authmodel "auth/model"
)

func SaveTweet(tweetUser string,tweetContent string,ctx context.Context)(error){
	resultChan := make(chan bool)
	go profileStorage.SaveTweetDB(tweetUser,tweetContent,resultChan)
	select {
	case <-resultChan:
		return nil
	case <-ctx.Done():
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
		return ctx.Err()
	}
}