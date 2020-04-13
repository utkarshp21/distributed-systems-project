package repository

import (
	"container/list"
	profileStorage "profile/storage"
	authmodel "auth/model"
)

func SaveTweet(tweetUser string,tweetContent string){
	resultChan := make(chan bool)
	go profileStorage.SaveTweetDB(tweetUser,tweetContent,resultChan)
	<-resultChan
}

func GetTweetList(followUsername string)(*list.List) {
	resultChan := make(chan *list.List)
	go profileStorage.GetTweetListDB(followUsername,resultChan)
	return <-resultChan
}

func InitialiseTweets(user authmodel.User){
	resultChan := make(chan bool)
	go profileStorage.InitialiseTweetsDB(user, resultChan)
	<-resultChan
}