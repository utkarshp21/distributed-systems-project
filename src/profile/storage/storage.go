package storage

import (
	"container/list"
	"context"
	profilemodel "profile/model"
	authmodel "auth/model"
	authStorage "auth/storage"
)

var Tweets = make(map[string]*list.List)

func SaveTweetDB(tweetUser string,tweetContent string,resultChan chan bool, deleteChan chan bool, ctx context.Context){
	profilemodel.TweetsMux.Lock()
	Tweets[tweetUser].PushBack(tweetContent)

	select {
	case <-ctx.Done():
		profilemodel.TweetsMux.Unlock()
		channel := make(chan bool)
		go DeleteTweetDB(tweetUser,tweetContent,channel)
		<-channel
		deleteChan <- true
	default:
		profilemodel.TweetsMux.Unlock()
		resultChan <- true
	}
}

func DeleteTweetDB(tweetUser string,tweetContent string,resultChan chan bool) {
	profilemodel.TweetsMux.Lock()
	for e := Tweets[tweetUser].Front(); e != nil ; e = e.Next(){
		if tweetContent == e.Value{
			Tweets[tweetUser].Remove(e)
		}
	}
	profilemodel.TweetsMux.Unlock()
	resultChan <- true
}

func GetTweetListDB(followUsername string, resultChan chan *list.List,deleteChan chan bool, ctx context.Context){
	profilemodel.TweetsMux.Lock()
	tweetList := Tweets[followUsername]
	select {
	case <-ctx.Done():
		profilemodel.TweetsMux.Unlock()
		deleteChan <- true
	default:
		profilemodel.TweetsMux.Unlock()
		resultChan <- tweetList
	}
}

func InitialiseTweetsDB(user authmodel.User, resultChan chan bool,deleteChan chan bool, ctx context.Context) {
	profilemodel.TweetsMux.Lock()
	Tweets[user.Username] = list.New()
	select {
	case <-ctx.Done():
		profilemodel.TweetsMux.Unlock()
		channel := make(chan bool)
		go authStorage.DeleteUserDB(user,channel)
		<-channel
		deleteChan <- true
	default:
		profilemodel.TweetsMux.Unlock()
		resultChan <- true
	}
}

func GetUsersDB(resultChan chan string,deleteChan chan bool, ctx context.Context)  {
	authmodel.UsersMux.Lock()
	keys := ""
	for k := range authStorage.Users {
		keys += k + ","
	}
	keys = keys[:len(keys)-1]
	select {
	case <-ctx.Done():
		authmodel.UsersMux.Unlock()
		deleteChan <- true
	default:
		authmodel.UsersMux.Unlock()
		resultChan <- keys
	}
}