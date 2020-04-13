package storage

import (
	"container/list"
	profilemodel "profile/model"
	authmodel "auth/model"
)

var Tweets = make(map[string]*list.List)

func SaveTweetDB(tweetUser string,tweetContent string,resultChan chan bool){
	profilemodel.TweetsMux.Lock()
	Tweets[tweetUser].PushBack(tweetContent)
	profilemodel.TweetsMux.Unlock()
	resultChan <- true
}

func GetTweetListDB(followUsername string, resultChan chan *list.List){
	profilemodel.TweetsMux.Lock()
	tweetList := Tweets[followUsername]
	profilemodel.TweetsMux.Unlock()
	resultChan <- tweetList
}

func InitialiseTweetsDB(user authmodel.User, resultChan chan bool) {
	profilemodel.TweetsMux.Lock()
	Tweets[user.Username] = list.New()
	profilemodel.TweetsMux.Unlock()
	resultChan <- true
}