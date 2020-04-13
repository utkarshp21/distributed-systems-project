package service

import (
	"container/list"
	//"errors"
	//"fmt"
	profileRepository "profile/repository"
	authRepository "auth/repository"
	authmodel "auth/model"
	//jwt "github.com/dgrijalva/jwt-go"
	//"net/http"
	//authStorage "auth/storage"
)

//func GetToken(c *http.Cookie) (*jwt.Token,error) {
//	tokenString := c.Value
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("Unexpected signing method")
//		}
//		return []byte("secret"), nil
//	})
//	return token, err
//}
//
//func ProfileService(r *http.Request)error{
//
//	c, err := r.Cookie("token")
//
//	if err != nil {
//		return err
//	}
//	token, tokenerr := GetToken(c)
//
//	if token.Valid && tokenerr == nil{
//		return nil
//	}else{
//		return errors.New("Token error")
//	}
//}

func FollowService(userPresentUsername string,followuserUsername string) (string) {

	userPresent, _ := authRepository.ReturnUser(userPresentUsername)
	followUser, _ := authRepository.ReturnUser(followuserUsername)

	if userPresent == followUser{
		return "Cant follow yourself"
	}

	for e := followUser.Followers.Front() ; e != nil ; e = e.Next(){
		k := e.Value.(authmodel.User)
		if userPresent == k{
			return "User already followed"
		}
	}

	if userPresent.Username != "" {
		followUser.Followers.PushBack(userPresent)
		authRepository.SaveUser(followUser)
		return ""
	} else {
		return "Username doesnt exist"
	}
}

func UnfollowService(userPresentUsername string,unfollowuserUsername string) (string) {

	userPresent, _ := authRepository.ReturnUser(userPresentUsername)
	unfollowUser, _ := authRepository.ReturnUser(unfollowuserUsername)

	if userPresent == unfollowUser{
		return "Cant unfollow yourself"
	}

	if userPresent.Username == "" {
		return "Username doesnt exist"
	}

	for e := unfollowUser.Followers.Front() ; e != nil ; e = e.Next(){
		k := e.Value.(authmodel.User)
		if userPresent == k{
			unfollowUser.Followers.Remove(e)
			authRepository.SaveUser(unfollowUser)
			return ""
		}
	}

	return "Follow user first"
}

func TweetService(tweetContent string, tweetUser string) (string) {

	if tweetContent != "" {
		profileRepository.SaveTweet(tweetUser,tweetContent)
		return ""
	} else {
		return "Enter tweet content"
	}
}

func FeedService(feedUserUsename string)(string,string){

	feedUser, _ := authRepository.ReturnUser(feedUserUsename)
	feed := ""

	for e:= feedUser.Followers.Front(); e != nil; e = e.Next(){
		followUser := e.Value.(authmodel.User)
		followUsername := followUser.Username
		feed = feed + GetTopFiveTweets(profileRepository.GetTweetList(followUsername),followUsername)
	}

	if feed != "" {
		return "",feed
	} else {
		return "No feed",""
	}
}

func GetTopFiveTweets(tweetList *list.List,followUsername string)(string){
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
