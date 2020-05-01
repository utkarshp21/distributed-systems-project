package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"backend/proto"
	"context"
	"log"
	"google.golang.org/grpc"
	jwt "github.com/dgrijalva/jwt-go"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	_, err := r.Cookie("token")

	if err == nil {
		http.Redirect(w, r, "/feed", http.StatusFound)
		return
	}

	m := map[string]interface{}{}
	t, _ := template.ParseFiles("../static/register.gtpl")

	if r.Method == "GET" {
		t.Execute(w, m)
		return
	} else {
		var opts = grpc.WithInsecure()
		var cc, ccerr = grpc.Dial("localhost:50051", opts)

		if ccerr != nil {
			log.Fatal(ccerr)
		}

		defer cc.Close()

		client := proto.NewTwitterClient(cc)

		r.ParseForm()
		request := &proto.RegisterRequest{Firstname: r.Form["firstname"][0], Lastname:r.Form["lastname"][0], Username:r.Form["username"][0], Password:r.Form["password"][0]}
		response, _ := client.Register(context.Background(), request)

		if response.Message != "" {
			m["Error"] = response.Message
			log.Println(response.Message)
			t.Execute(w, m)
			return
		} else {
			log.Println("User Registered succesfully")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("../static/login.gtpl")

	_, err := r.Cookie("token")

	if err == nil {
		http.Redirect(w, r, "/feed", http.StatusFound)
		return
	}

	m := map[string]interface{}{}

	if r.Method == "GET" {
		t.Execute(w, nil)
		return
	} else {

		var opts = grpc.WithInsecure()
		var cc, ccerr = grpc.Dial("localhost:50051", opts)

		if ccerr != nil {
			log.Fatal(ccerr)
		}

		defer cc.Close()

		client := proto.NewTwitterClient(cc)

		r.ParseForm()
		request := &proto.LoginRequest{Username: r.Form["username"][0], Password: r.Form["password"][0]}
		response , _ := client.Login(context.Background(), request)

		if response.Message != "" {
			m["Error"] = response.Message
			log.Println(response.Message)
			t.Execute(w, m)
			return
		} else {
			http.SetCookie(w, &http.Cookie{
				Name:  "token",
				Value: response.Tokenstring,
			})
			log.Println("Login successful")
			http.Redirect(w, r, "/feed", http.StatusFound)
			return
		}
	}
}

func SignoutHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.Redirect(w, r, "/feed", http.StatusFound)
		return
	}

	t, _ := template.ParseFiles("../static/login.gtpl")
	m := map[string]interface{}{}

	c, cerr := r.Cookie("token")
	if cerr != nil {
		m["Error"] = "Please login to continue"
		m["Success"] = nil
		log.Println("Please login to continue")
		t.Execute(w, m)
		return
	}

	tokenString := c.Value

	var opts = grpc.WithInsecure()
	var cc, ccerr = grpc.Dial("localhost:50051", opts)

	if ccerr != nil {
		log.Fatal(ccerr)
	}

	defer cc.Close()

	client := proto.NewTwitterClient(cc)

	request := &proto.LogoutRequest{Tokenstring: tokenString}

	response, _ := client.Logout(context.Background(), request)

	if response.Message != "" {
		m["Error"] = response.Message
		m["Success"] = nil
		log.Println(response.Message)
		t.Execute(w, m)
		return
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   "",
			Expires: time.Unix(0, 0),
		})
		log.Println("Logout successful")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}

func redirectToLogin(w http.ResponseWriter){
	t, _ := template.ParseFiles("../static/login.gtpl")
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

func UserListHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../static/userlist.gtpl")
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

	client := proto.NewTwitterClient(cc)

	request := &proto.FeedRequest{Reqparm1 : feedUserUsername}
	response, _ := client.UserListService(context.Background(),request)

	if response.GetResparm1() != "" {
		m["Error"] = response.GetResparm1()
		m["Success"] = nil
		m["List"] = nil
		log.Println(response.GetResparm1())
		t.Execute(w, m)
		return
	}else {
		m["Error"] = nil
		m["Success"] = nil
		m["List"] = response.GetResparm2()
		log.Println("User List Succesfull")
		t.Execute(w, m)
		return
	}
}

func FollowHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.Redirect(w, r, "/userlist", http.StatusFound)
		return
	}
	m := map[string]interface{}{}
	t, _ := template.ParseFiles("../static/userlist.gtpl")

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

	client := proto.NewTwitterClient(cc)

	request := &proto.ProfileRequest{Reqparm1 : userPresentUsername, Reqparm2: followuserUsername}
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
		http.Redirect(w, r, "/userlist", http.StatusFound)
		return
	}
	m := map[string]interface{}{}
	t, _ := template.ParseFiles("../static/userlist.gtpl")

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

	client := proto.NewTwitterClient(cc)

	request := &proto.ProfileRequest{Reqparm1 : userPresentUsername, Reqparm2: unfollowuserUsername}
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

	t, _ := template.ParseFiles("../static/profile.gtpl")
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

	client := proto.NewTwitterClient(cc)

	request := &proto.ProfileRequest{Reqparm1 : tweetContent, Reqparm2: tweetUserUsername}
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

	t, _ := template.ParseFiles("../static/profile.gtpl")
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

	client := proto.NewTwitterClient(cc)

	request := &proto.FeedRequest{Reqparm1 : feedUserUsername}
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