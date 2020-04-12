package service

import (
	"container/list"
	"errors"
	"fmt"
	profileRepository "profile/repository"
	authRepository "auth/repository"
	authmodel "auth/model"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
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

func ProfileService(r *http.Request)error{

	c, err := r.Cookie("token")

	if err != nil {
		return err
	}
	token, tokenerr := GetToken(c)

	if token.Valid && tokenerr == nil{
		return nil
	}else{
		return errors.New("Token error")
	}
}

func FollowService(r *http.Request) (error,string) {

	c, err := r.Cookie("token")

	if err != nil {
		return err, ""
	}

	token, tokenerr := GetToken(c)
	if !token.Valid && tokenerr != nil{
		return errors.New("Token error"), ""
	}

	r.ParseForm()
	userPresent, _ := authRepository.ReturnUser(r.Form["username"][0])
	claims, _ := token.Claims.(jwt.MapClaims)
	followUser, _ := authRepository.ReturnUser(claims["username"].(string))

	if userPresent == followUser{
		return nil, "Cant follow yourself"
	}

	for e := followUser.Followers.Front() ; e != nil ; e = e.Next(){
		k := e.Value.(authmodel.User)
		if userPresent == k{
			return nil, "User already followed"
		}
	}

	if userPresent.Username != "" {
		followUser.Followers.PushBack(userPresent)
		authRepository.SaveUser(followUser)
		return nil, ""
	} else {
		return nil, "Username doesnt exist"
	}
}

func TweetService(r *http.Request) (error,string) {

	c, err := r.Cookie("token")

	if err != nil {
		return err, ""
	}

	token, tokenerr := GetToken(c)
	if !token.Valid && tokenerr != nil{
		return errors.New("Token error"), ""
	}

	r.ParseForm()
	tweetContent := r.Form["tweet"][0]

	if tweetContent != "" {
		claims, _ := token.Claims.(jwt.MapClaims)
		tweetUser := claims["username"].(string)
		profileRepository.SaveTweet(tweetUser,tweetContent)
		return nil, ""
	} else {
		return nil, "Enter tweet content"
	}
}

func FeedService(r *http.Request)(error,string,string){

	c, err := r.Cookie("token")

	if err != nil {
		return err, "",""
	}

	token, tokenerr := GetToken(c)
	if !token.Valid && tokenerr != nil{
		return errors.New("Token error"), "",""
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	feedUser, _ := authRepository.ReturnUser(claims["username"].(string))
	feed := ""

	for e:= feedUser.Followers.Front(); e != nil; e = e.Next(){
		followUser := e.Value.(authmodel.User)
		followUsername := followUser.Username
		feed = feed + GetTopFiveTweets(profileRepository.GetTweetList(followUsername),followUsername)
	}

	if feed != "" {
		return nil, "",feed
	} else {
		return nil, "No feed",""
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
