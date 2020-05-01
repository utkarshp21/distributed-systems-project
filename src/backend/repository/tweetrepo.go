package repository

import (
	model "backend/model"
	//authRepository "backend/repository"
	"container/list"
	"context"
	//profileStorage "profile/storage"
	"time"
)

var Tweets = make(map[string]*list.List)

func SaveTweet(tweetUser string,tweetContent string,ctx context.Context)(error){
	time.Sleep(10*time.Millisecond)
	resultChan := make(chan bool)
	deleteChan := make(chan bool)
	go SaveTweetDB(tweetUser,tweetContent,resultChan,deleteChan,ctx)
	select {
	case <-resultChan:
		return nil
	case <-deleteChan:
		return ctx.Err()
	}
}

func SaveTweetDB(tweetUser string,tweetContent string,resultChan chan bool, deleteChan chan bool, ctx context.Context){
	tweetContent += "*"+time.Now().Format("2006-01-02 15:04:05")
	model.TweetsMux.Lock()
	Tweets[tweetUser].PushBack(tweetContent)

	select {
	case <-ctx.Done():
		model.TweetsMux.Unlock()
		channel := make(chan bool)
		go DeleteTweetDB(tweetUser,tweetContent,channel)
		<-channel
		deleteChan <- true
	default:
		model.TweetsMux.Unlock()
		resultChan <- true
	}
}

func DeleteTweetDB(tweetUser string,tweetContent string,resultChan chan bool) {
	model.TweetsMux.Lock()
	for e := Tweets[tweetUser].Front(); e != nil ; e = e.Next(){
		if tweetContent == e.Value{
			Tweets[tweetUser].Remove(e)
		}
	}
	model.TweetsMux.Unlock()
	resultChan <- true
}

func GetTweetList(followUsername string,ctx context.Context)(*list.List,error) {
	resultChan := make(chan *list.List)
	deleteChan := make(chan bool)
	dummyList := list.New()
	go GetTweetListDB(followUsername,resultChan,deleteChan,ctx)
	select {
	case res := <-resultChan:
		return res,nil
	case <-deleteChan:
		return dummyList, ctx.Err()
	}
}

func GetTweetListDB(followUsername string, resultChan chan *list.List,deleteChan chan bool, ctx context.Context){
	model.TweetsMux.Lock()
	tweetList := Tweets[followUsername]
	select {
	case <-ctx.Done():
		model.TweetsMux.Unlock()
		deleteChan <- true
	default:
		model.TweetsMux.Unlock()
		resultChan <- tweetList
	}
}

func InitialiseTweets(user model.User,ctx context.Context)(error){
	resultChan := make(chan bool)
	deleteChan := make(chan bool)
	go InitialiseTweetsDB(user, resultChan,deleteChan,ctx)
	select {
	case <-resultChan:
		return nil
	case <-deleteChan:
		return ctx.Err()
	}
}

func InitialiseTweetsDB(user model.User, resultChan chan bool,deleteChan chan bool, ctx context.Context) {
	model.TweetsMux.Lock()
	Tweets[user.Username] = list.New()
	select {
	case <-ctx.Done():
		model.TweetsMux.Unlock()
		channel := make(chan bool)
		go DeleteUserDB(user,channel)
		<-channel
		deleteChan <- true
	default:
		model.TweetsMux.Unlock()
		resultChan <- true
	}
}
