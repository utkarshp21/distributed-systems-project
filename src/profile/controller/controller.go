package controller

import (
	"context"
	"fmt"
	"log"
	jwt "github.com/dgrijalva/jwt-go"
	"html/template"
	"net/http"
	"google.golang.org/grpc"
	"auth/authpb"
)

func redirectToLogin(w http.ResponseWriter){
	t, _ := template.ParseFiles("login.gtpl")
	m := map[string]interface{}{}
	m["Error"] = "Please login to continue!"
	m["Success"] = nil
	log.Println("Please login to continue")
	t.Execute(w, m)
}

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

func FollowHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.Redirect(w, r, "/feed", http.StatusFound)
		return
	}
	m := map[string]interface{}{}
	t, _ := template.ParseFiles("profile.gtpl")

	c, err := r.Cookie("token")

	if err != nil {
		redirectToLogin(w)
		return
	}

	token, tokenerr := GetToken(c)
	if !token.Valid && tokenerr != nil{
		redirectToLogin(w)
		return
	}

	r.ParseForm()
	claims, _ := token.Claims.(jwt.MapClaims)
	userPresentUsername := r.Form["username"][0]
	followuserUsername := claims["username"].(string)

	var opts = grpc.WithInsecure()
	var cc, ccerr = grpc.Dial("localhost:50051", opts)

	if ccerr != nil {
		log.Fatal(ccerr)
	}

	defer cc.Close()

	client := authpb.NewFollowServiceClient(cc)

	request := &authpb.ProfileRequest{Reqparm1 : userPresentUsername, Reqparm2: followuserUsername}
	response, _ := client.FollowService(context.Background(),request)

	if response.GetResparm1() != "" {
		m["Error"] = response.GetResparm1()
		m["Success"] = nil
		log.Println(response.GetResparm1())
		t.Execute(w, m)
		return
	}else {
		m["Error"] = nil
		m["Success"] = "Succesfully followed"
		log.Println("Succesfully followed")
		t.Execute(w, m)
		return
	}
}

func UnfollowHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.Redirect(w, r, "/feed", http.StatusFound)
		return
	}
	m := map[string]interface{}{}
	t, _ := template.ParseFiles("profile.gtpl")

	c, err := r.Cookie("token")

	if err != nil {
		redirectToLogin(w)
		return
	}

	token, tokenerr := GetToken(c)
	if !token.Valid && tokenerr != nil{
		redirectToLogin(w)
		return
	}

	r.ParseForm()
	claims, _ := token.Claims.(jwt.MapClaims)
	userPresentUsername := r.Form["username"][0]
	unfollowuserUsername := claims["username"].(string)

	var opts = grpc.WithInsecure()
	var cc, ccerr = grpc.Dial("localhost:50051", opts)

	if ccerr != nil {
		log.Fatal(ccerr)
	}

	defer cc.Close()

	client := authpb.NewUnfollowServiceClient(cc)

	request := &authpb.ProfileRequest{Reqparm1 : userPresentUsername, Reqparm2: unfollowuserUsername}
	response, _ := client.UnfollowService(context.Background(),request)

	if response.GetResparm1() != "" {
		m["Error"] = response.GetResparm1()
		m["Success"] = nil
		log.Println(response.GetResparm1())
		t.Execute(w, m)
		return
	}else {
		m["Error"] = nil
		m["Success"] = "Succesfully unfollowed"
		log.Println("Succesfully unfollowed")
		t.Execute(w, m)
		return
	}
}

func TweetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/feed", http.StatusFound)
		return
	}

	t, _ := template.ParseFiles("profile.gtpl")
	m := map[string]interface{}{}

	c, err := r.Cookie("token")

	if err != nil {
		redirectToLogin(w)
		return
	}

	token, tokenerr := GetToken(c)
	if !token.Valid && tokenerr != nil{
		redirectToLogin(w)
		return
	}

	r.ParseForm()
	tweetContent := r.Form["tweet"][0]
	claims, _ := token.Claims.(jwt.MapClaims)
	tweetUserUsername := claims["username"].(string)

	var opts = grpc.WithInsecure()
	var cc, ccerr = grpc.Dial("localhost:50051", opts)

	if ccerr != nil {
		log.Fatal(ccerr)
	}

	defer cc.Close()

	client := authpb.NewTweetServiceClient(cc)

	request := &authpb.ProfileRequest{Reqparm1 : tweetContent, Reqparm2: tweetUserUsername}
	response, _ := client.TweetService(context.Background(),request)

	if response.GetResparm1() != "" {
		m["Error"] = response.GetResparm1()
		m["Success"] = nil
		log.Println(response.GetResparm1())
		t.Execute(w, m)
		return
	}else {
		m["Error"] = nil
		m["Success"] = "Succesfully tweeted"
		log.Println("Succesfully tweeted")
		t.Execute(w, m)
		return
	}
}

func FeedHandler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("profile.gtpl")
	m := map[string]interface{}{}

	c, err := r.Cookie("token")

	if err != nil {
		redirectToLogin(w)
		return
	}

	token, tokenerr := GetToken(c)
	if !token.Valid && tokenerr != nil{
		redirectToLogin(w)
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	feedUserUsername := claims["username"].(string)

	var opts = grpc.WithInsecure()
	var cc, ccerr = grpc.Dial("localhost:50051", opts)

	if ccerr != nil {
		log.Fatal(ccerr)
	}

	defer cc.Close()

	client := authpb.NewFeedServiceClient(cc)

	request := &authpb.FeedRequest{Reqparm1 : feedUserUsername}
	response, _ := client.FeedService(context.Background(),request)

	if response.GetResparm1() != "" {
		m["Error"] = response.GetResparm1()
		m["Success"] = nil
		log.Println(response.GetResparm1())
		t.Execute(w, m)
		return
	}else {
		m["Error"] = nil
		m["Success"] = nil
		m["Feed"] = response.GetResparm2()
		log.Println("Feed Succesfull")
		t.Execute(w, m)
		return
	}
}
