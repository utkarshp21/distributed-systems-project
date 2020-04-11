package repository

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	authmodel "auth/model"
	authStorage "auth/storage"
	profilemodel "profile/model"
	profileStorage "profile/storage"
)

func GetToken(c *http.Cookie) (*jwt.Token,error) {
	tokenString := c.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})
	return token, err
}

func CheckIfUserAlreadyFollowed(followUser authmodel.User, userPresent authmodel.User)(bool){
	for e := followUser.Followers.Front() ; e != nil ; e = e.Next(){
		k := e.Value.(authmodel.User)
		if userPresent == k{
			return true
		}
	}
	return false
}

func FollowUser(userPresent authmodel.User,followUser authmodel.User)  {
	followUser.Followers.PushBack(userPresent)
	authmodel.UsersMux.Lock()
	authStorage.Users[followUser.Username] = followUser
	authmodel.UsersMux.Unlock()
}

func SaveTweet(tweetUser string,tweetContent string){
	profilemodel.TweetsMux.Lock()
	profileStorage.Tweets[tweetUser].PushBack(tweetContent)
	profilemodel.TweetsMux.Unlock()
}

func GetTopFiveTweets(followUsername string)(string){
	profilemodel.TweetsMux.Lock()
	tweetList := profileStorage.Tweets[followUsername]
	profilemodel.TweetsMux.Unlock()
	
	numOfTweets := 5
	feed := ""
	for k := tweetList.Back(); k != nil && numOfTweets > 0; k = k.Prev() {
		numOfTweets = numOfTweets - 1
		feed = feed + k.Value.(string) + "\n"
	}
	if feed != ""{
		feed = "Top 5 tweets from "+ followUsername + " : \n" + feed
	}
	return feed
	
}

func FeedGenerator(feedUser authmodel.User)(string){
	feed := ""
	for e:= feedUser.Followers.Front(); e != nil; e = e.Next(){
		followUser := e.Value.(authmodel.User)
		followUsername := followUser.Username
		feed = feed + GetTopFiveTweets(followUsername)
	}
	return feed
}
