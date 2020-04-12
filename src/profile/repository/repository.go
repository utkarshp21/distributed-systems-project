package repository

import (
	"container/list"
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

func GetTweetList(followUsername string)(*list.List) {
	profilemodel.TweetsMux.Lock()
	tweetList := profileStorage.Tweets[followUsername]
	profilemodel.TweetsMux.Unlock()
	return tweetList
}

