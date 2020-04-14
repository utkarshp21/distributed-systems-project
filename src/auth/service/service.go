package main

import (
	authmodel "auth/model"
	repository "auth/repository"
	"container/list"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	profileRepository "profile/repository"
	"auth/authpb"
	"context"
	"log"
	"net"
	"google.golang.org/grpc"

)

type server struct {
}

func (*server) Login(ctx context.Context, request *authpb.LoginRequest) (*authpb.LoginResponse, error) {

	user, usernameExists, ctxErr := repository.ReturnUser(request.Username,ctx)
	bkpUser := user

	if ctxErr != nil{
		response := &authpb.LoginResponse{Message:"Request timeout. Try again", Tokenstring: ""}
		return response, nil
	}

	if usernameExists {
		passowrdErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

		if passowrdErr != nil {
			response := &authpb.LoginResponse{Message:"Invalid Password", Tokenstring: ""}
			return response, nil
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username":  user.Username,
			"firstname": user.FirstName,
			"lastname":  user.LastName,
		})

		tokenString, loginErr := token.SignedString([]byte("secret"))

		if loginErr != nil {
			response := &authpb.LoginResponse{Message:"Error while generating token,Try again", Tokenstring: ""}
			return response, nil
		}

		user.Token = tokenString

		ctxErr2 := repository.SaveUser(user,ctx,bkpUser)

		if ctxErr2 != nil{
			response := &authpb.LoginResponse{Message:"Request timeout. Try again", Tokenstring: ""}
			return response, nil
		}

		response := &authpb.LoginResponse{Message:"", Tokenstring: tokenString}
		return response, nil

	}

	response := &authpb.LoginResponse{Message:"Invalid Username", Tokenstring: ""}
	return response, nil

}

func (*server) Register(ctx context.Context, request *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {

	_, usernameExists, ctxErr := repository.ReturnUser(request.Username,ctx)

	if ctxErr != nil{
		response := &authpb.RegisterResponse{Message:"Request timeout. Try again"}
		return response, nil
	}

	if usernameExists {
		response := &authpb.RegisterResponse{Message: "User already exists"}
		return response, nil
	}

	registerFromInput := authmodel.User{
		Username:  request.Username,
		Password:  request.Password,
		FirstName: request.Firstname,
		LastName:  request.Lastname,
		Followers: list.New(),
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(registerFromInput.Password), 5)
	if err != nil {
		response := &authpb.RegisterResponse{Message: "Error While Hashing Password, Try Again"}
		return response, nil
	}
	registerFromInput.Password = string(hash)

	ctxErr2 := repository.SaveUserRegister(registerFromInput,ctx)

	if ctxErr2 != nil{
		response := &authpb.RegisterResponse{Message:"Request timeout. Try again"}
		return response, nil
	}

	ctxErr3 := profileRepository.InitialiseTweets(registerFromInput,ctx)

	if ctxErr3 != nil{
		response := &authpb.RegisterResponse{Message:"Request timeout. Try again"}
		return response, nil
	}

	response := &authpb.RegisterResponse{Message: "",}
	return response, nil
}

func (*server) Logout(ctx context.Context, request *authpb.LogoutRequest) (*authpb.LogoutResponse, error) {
	token, tokenerr := jwt.Parse(request.Tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if !token.Valid || tokenerr != nil {
		response := &authpb.LogoutResponse{Message: "Please login to continue"}
		return response, nil
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	signoutUserName := claims["username"].(string)
	signoutUser, _, ctxErr := repository.ReturnUser(signoutUserName,ctx)
	bkpUser := signoutUser

	if ctxErr != nil{
		response := &authpb.LogoutResponse{Message: "Request timeout. Try again"}
		return response, nil
	}

	if signoutUser.Username != "" {
		signoutUser.Token = ""
		ctxErr2 := repository.SaveUser(signoutUser,ctx,bkpUser)
		if ctxErr2 != nil{
			response := &authpb.LogoutResponse{Message: "Request timeout. Try again"}
			return response, nil
		}
		response := &authpb.LogoutResponse{Message: ""}
		return response, nil
	} else {
		response := &authpb.LogoutResponse{Message: "Please login to continue"}
		return response, nil
	}
}

func (*server) FollowService(ctx context.Context, request *authpb.ProfileRequest) (*authpb.ProfileResponse, error) {

	userPresent, _, ctxErr1 := repository.ReturnUser(request.GetReqparm1(),ctx)

	if ctxErr1 != nil{
		response := &authpb.ProfileResponse{Resparm1: "Request timeout. Try again"}
		return response, nil
	}

	followUser, _, ctxErr2 := repository.ReturnUser(request.GetReqparm2(),ctx)
	bkpUser := followUser

	if ctxErr2 != nil{
		response := &authpb.ProfileResponse{Resparm1: "Request timeout. Try again"}
		return response, nil
	}

	if userPresent == followUser{
		response := &authpb.ProfileResponse{Resparm1: "Cant follow yourself"}
		return response, nil
	}

	for e := followUser.Followers.Front() ; e != nil ; e = e.Next(){
		k := e.Value.(authmodel.User)
		if userPresent == k{
			response := &authpb.ProfileResponse{Resparm1: "User already followed"}
			return response, nil
		}
	}

	if userPresent.Username != "" {
		followUser.Followers.PushBack(userPresent)
		ctxErr3 := repository.SaveUser(followUser,ctx,bkpUser)
		if ctxErr3 != nil{
			response := &authpb.ProfileResponse{Resparm1: "Request timeout. Try again"}
			return response, nil
		}
		response := &authpb.ProfileResponse{Resparm1: ""}
		return response, nil
	} else {
		response := &authpb.ProfileResponse{Resparm1: "Username doesnt exist"}
		return response, nil
	}

}

func (*server) UnfollowService(ctx context.Context, request *authpb.ProfileRequest) (*authpb.ProfileResponse, error) {

	userPresent, _ , ctxErr1 := repository.ReturnUser(request.GetReqparm1(),ctx)

	if ctxErr1 != nil{
		response := &authpb.ProfileResponse{Resparm1: "Request timeout. Try again"}
		return response, nil
	}

	unfollowUser, _, ctxErr2 := repository.ReturnUser(request.GetReqparm2(),ctx)
	bkpUser := unfollowUser

	if ctxErr2 != nil{
		response := &authpb.ProfileResponse{Resparm1: "Request timeout. Try again"}
		return response, nil
	}

	if userPresent == unfollowUser{
		response := &authpb.ProfileResponse{Resparm1: "Cant unfollow yourself"}
		return response, nil
	}

	if userPresent.Username == "" {
		response := &authpb.ProfileResponse{Resparm1: "Username doesnt exist"}
		return response, nil	}

	for e := unfollowUser.Followers.Front() ; e != nil ; e = e.Next(){
		k := e.Value.(authmodel.User)
		if userPresent == k{
			unfollowUser.Followers.Remove(e)
			ctxErr3 := repository.SaveUser(unfollowUser,ctx,bkpUser)
			if ctxErr3 != nil{
				response := &authpb.ProfileResponse{Resparm1: "Request timeout. Try again"}
				return response, nil
			}
			response := &authpb.ProfileResponse{Resparm1: ""}
			return response, nil
		}
	}

	response := &authpb.ProfileResponse{Resparm1: "Follow user first"}
	return response, nil
}

func (*server) TweetService(ctx context.Context, request *authpb.ProfileRequest) (*authpb.ProfileResponse, error) {

	tweetContent := request.GetReqparm1()
	tweetUser := request.GetReqparm2()

	if tweetContent != "" {
		ctxErr := profileRepository.SaveTweet(tweetUser,tweetContent,ctx)
		if ctxErr != nil{
			response := &authpb.ProfileResponse{Resparm1: "Request timeout. Try again"}
			return response, nil
		}
		response := &authpb.ProfileResponse{Resparm1: ""}
		return response, nil
	} else {
		response := &authpb.ProfileResponse{Resparm1: "Enter tweet content"}
		return response, nil
	}
}

func (*server) FeedService(ctx context.Context, request *authpb.FeedRequest) (*authpb.FeedResponse, error) {

	feedUser, _, ctxErr1 := repository.ReturnUser(request.GetReqparm1(),ctx)

	if ctxErr1 != nil{
		response := &authpb.FeedResponse{Resparm1: "Request timeout. Try again",Resparm2: ""}
		return response, nil
	}

	feed := ""

	for e:= feedUser.Followers.Front(); e != nil; e = e.Next(){
		followUser := e.Value.(authmodel.User)
		followUsername := followUser.Username
		tweetList, ctxErr2 := profileRepository.GetTweetList(followUsername,ctx)
		if ctxErr2 != nil{
			response := &authpb.FeedResponse{Resparm1: "Request timeout. Try again",Resparm2: ""}
			return response, nil
		}
		feed = feed + GetTopFiveTweets(tweetList,followUsername)
	}

	if feed != "" {
		response := &authpb.FeedResponse{Resparm1: "",Resparm2: feed}
		return response, nil
	} else {
		response := &authpb.FeedResponse{Resparm1: "No feed",Resparm2: ""}
		return response, nil
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

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	authpb.RegisterRegisterServiceServer(s, &server{})
	authpb.RegisterLoginServiceServer(s, &server{})
	authpb.RegisterLogoutServiceServer(s, &server{})
	authpb.RegisterFollowServiceServer(s, &server{})
	authpb.RegisterUnfollowServiceServer(s, &server{})
	authpb.RegisterTweetServiceServer(s, &server{})
	authpb.RegisterFeedServiceServer(s, &server{})

	s.Serve(lis)
}
