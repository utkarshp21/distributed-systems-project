package repository

import (
	"container/list"
	profilemodel "profile/model"
	profileStorage "profile/storage"
)

func SaveTweet(tweetUser string,tweetContent string){
	profilemodel.TweetsMux.Lock()
	profileStorage.Tweets[tweetUser].PushBack(tweetContent)
	profilemodel.TweetsMux.Unlock()
}

func GetTweetList(followUsername string)(*list.List) {
	profilemodel.TweetsMux.Lock()
	tweetList := profileStorage.Tweets[followUsername]
	profilemodel.TweetsMux.Unlock()
	return tweetList
}

