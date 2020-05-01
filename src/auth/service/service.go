package main

import (
	authmodel "auth/model"
	repository "auth/repository"
	"container/list"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	profileRepository "profile/repository"
	"auth/proto"
	"context"
	"log"
	"net"
	"google.golang.org/grpc"

)

type server struct {
}

func (*server) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {

	user, usernameExists, ctxErr := repository.ReturnUser(request.Username,ctx)
	bkpUser := user

	if ctxErr != nil{
		response := &proto.LoginResponse{Message:"Request timeout. Try again", Tokenstring: ""}
		return response, nil
	}

	if usernameExists {
		passowrdErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

		if passowrdErr != nil {
			response := &proto.LoginResponse{Message:"Invalid Password", Tokenstring: ""}
			return response, nil
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username":  user.Username,
			"firstname": user.FirstName,
			"lastname":  user.LastName,
		})

		tokenString, loginErr := token.SignedString([]byte("secret"))

		if loginErr != nil {
			response := &proto.LoginResponse{Message:"Error while generating token,Try again", Tokenstring: ""}
			return response, nil
		}

		user.Token = tokenString

		ctxErr2 := repository.SaveUser(user,ctx,bkpUser)

		if ctxErr2 != nil{
			response := &proto.LoginResponse{Message:"Request timeout. Try again", Tokenstring: ""}
			return response, nil
		}

		response := &proto.LoginResponse{Message:"", Tokenstring: tokenString}
		return response, nil

	}

	response := &proto.LoginResponse{Message:"Invalid Username", Tokenstring: ""}
	return response, nil

}

func (*server) Register(ctx context.Context, request *proto.RegisterRequest) (*proto.RegisterResponse, error) {

	_, usernameExists, ctxErr := repository.ReturnUser(request.Username,ctx)

	if ctxErr != nil{
		response := &proto.RegisterResponse{Message:"Request timeout. Try again"}
		return response, nil
	}

	if usernameExists {
		response := &proto.RegisterResponse{Message: "User already exists"}
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
		response := &proto.RegisterResponse{Message: "Error While Hashing Password, Try Again"}
		return response, nil
	}
	registerFromInput.Password = string(hash)

	ctxErr2 := repository.SaveUserRegister(registerFromInput,ctx)

	if ctxErr2 != nil{
		response := &proto.RegisterResponse{Message:"Request timeout. Try again"}
		return response, nil
	}

	ctxErr3 := profileRepository.InitialiseTweets(registerFromInput,ctx)

	if ctxErr3 != nil{
		response := &proto.RegisterResponse{Message:"Request timeout. Try again"}
		return response, nil
	}

	response := &proto.RegisterResponse{Message: "",}
	return response, nil
}

func (*server) Logout(ctx context.Context, request *proto.LogoutRequest) (*proto.LogoutResponse, error) {
	token, tokenerr := jwt.Parse(request.Tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if !token.Valid || tokenerr != nil {
		response := &proto.LogoutResponse{Message: "Please login to continue"}
		return response, nil
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	signoutUserName := claims["username"].(string)
	signoutUser, _, ctxErr := repository.ReturnUser(signoutUserName,ctx)
	bkpUser := signoutUser

	if ctxErr != nil{
		response := &proto.LogoutResponse{Message: "Request timeout. Try again"}
		return response, nil
	}

	if signoutUser.Username != "" {
		signoutUser.Token = ""
		ctxErr2 := repository.SaveUser(signoutUser,ctx,bkpUser)
		if ctxErr2 != nil{
			response := &proto.LogoutResponse{Message: "Request timeout. Try again"}
			return response, nil
		}
		response := &proto.LogoutResponse{Message: ""}
		return response, nil
	} else {
		response := &proto.LogoutResponse{Message: "Please login to continue"}
		return response, nil
	}
}

func (*server) FollowService(ctx context.Context, request *proto.ProfileRequest) (*proto.ProfileResponse, error) {

	userPresent, _, ctxErr1 := repository.ReturnUser(request.GetReqparm1(),ctx)

	if ctxErr1 != nil{
		response := &proto.ProfileResponse{Resparm1: "Request timeout. Try again"}
		return response, nil
	}

	followUser, _, ctxErr2 := repository.ReturnUser(request.GetReqparm2(),ctx)
	bkpUser := followUser

	if ctxErr2 != nil{
		response := &proto.ProfileResponse{Resparm1: "Request timeout. Try again"}
		return response, nil
	}

	if userPresent == followUser{
		response := &proto.ProfileResponse{Resparm1: "Cant follow yourself"}
		return response, nil
	}

	for e := followUser.Followers.Front() ; e != nil ; e = e.Next(){
		k := e.Value.(authmodel.User)
		if userPresent == k{
			response := &proto.ProfileResponse{Resparm1: "User already followed"}
			return response, nil
		}
	}

	if userPresent.Username != "" {
		followUser.Followers.PushBack(userPresent)
		ctxErr3 := repository.SaveUser(followUser,ctx,bkpUser)
		if ctxErr3 != nil{
			response := &proto.ProfileResponse{Resparm1: "Request timeout. Try again"}
			return response, nil
		}
		response := &proto.ProfileResponse{Resparm1: ""}
		return response, nil
	} else {
		response := &proto.ProfileResponse{Resparm1: "Username doesnt exist"}
		return response, nil
	}

}

func (*server) UnfollowService(ctx context.Context, request *proto.ProfileRequest) (*proto.ProfileResponse, error) {

	userPresent, _ , ctxErr1 := repository.ReturnUser(request.GetReqparm1(),ctx)

	if ctxErr1 != nil{
		response := &proto.ProfileResponse{Resparm1: "Request timeout. Try again"}
		return response, nil
	}

	unfollowUser, _, ctxErr2 := repository.ReturnUser(request.GetReqparm2(),ctx)
	bkpUser := unfollowUser

	if ctxErr2 != nil{
		response := &proto.ProfileResponse{Resparm1: "Request timeout. Try again"}
		return response, nil
	}

	if userPresent == unfollowUser{
		response := &proto.ProfileResponse{Resparm1: "Cant unfollow yourself"}
		return response, nil
	}

	if userPresent.Username == "" {
		response := &proto.ProfileResponse{Resparm1: "Username doesnt exist"}
		return response, nil	}

	for e := unfollowUser.Followers.Front() ; e != nil ; e = e.Next(){
		k := e.Value.(authmodel.User)
		if userPresent == k{
			unfollowUser.Followers.Remove(e)
			ctxErr3 := repository.SaveUser(unfollowUser,ctx,bkpUser)
			if ctxErr3 != nil{
				response := &proto.ProfileResponse{Resparm1: "Request timeout. Try again"}
				return response, nil
			}
			response := &proto.ProfileResponse{Resparm1: ""}
			return response, nil
		}
	}

	response := &proto.ProfileResponse{Resparm1: "Follow user first"}
	return response, nil
}

func (*server) TweetService(ctx context.Context, request *proto.ProfileRequest) (*proto.ProfileResponse, error) {

	tweetContent := request.GetReqparm1()
	tweetUser := request.GetReqparm2()

	if tweetContent != "" {
		ctxErr := profileRepository.SaveTweet(tweetUser,tweetContent,ctx)
		if ctxErr != nil{
			response := &proto.ProfileResponse{Resparm1: "Request timeout. Try again"}
			return response, nil
		}
		response := &proto.ProfileResponse{Resparm1: ""}
		return response, nil
	} else {
		response := &proto.ProfileResponse{Resparm1: "Enter tweet content"}
		return response, nil
	}
}

func (*server) FeedService(ctx context.Context, request *proto.FeedRequest) (*proto.FeedResponse, error) {

	feedUser, _, ctxErr1 := repository.ReturnUser(request.GetReqparm1(),ctx)

	if ctxErr1 != nil{
		response := &proto.FeedResponse{Resparm1: "Request timeout. Try again",Resparm2: ""}
		return response, nil
	}

	feed := ""

	for e:= feedUser.Followers.Front(); e != nil; e = e.Next(){
		followUser := e.Value.(authmodel.User)
		followUsername := followUser.Username
		tweetList, ctxErr2 := profileRepository.GetTweetList(followUsername,ctx)
		if ctxErr2 != nil{
			response := &proto.FeedResponse{Resparm1: "Request timeout. Try again",Resparm2: ""}
			return response, nil
		}
		feed = feed + GetTopFiveTweets(tweetList,followUsername)
	}

	if feed != "" {
		feed = feed[:len(feed)-1]
		response := &proto.FeedResponse{Resparm1: "",Resparm2: feed}
		return response, nil
	} else {
		response := &proto.FeedResponse{Resparm1: "No feed",Resparm2: ""}
		return response, nil
	}
}

func GetTopFiveTweets(tweetList *list.List,followUsername string)(string){
	numOfTweets := 5
	feed := ""
	for k := tweetList.Back(); k != nil && numOfTweets > 0; k = k.Prev() {
		numOfTweets = numOfTweets - 1
		feed = feed + k.Value.(string) + ","
	}

	if feed != ""{
		feed = feed[:len(feed)-1]
		feed = followUsername + "^" + feed + "$"
	}
	return feed

}

func (*server) UserListService(ctx context.Context, request *proto.FeedRequest) (*proto.FeedResponse, error) {

	userNameList, ctxErr1 := profileRepository.GetUsers(ctx)

	if ctxErr1 != nil{
		response := &proto.FeedResponse{Resparm1: "Request timeout. Try again",Resparm2: ""}
		return response, nil
	}
	userNameList += "$"
	presentUser, _, ctxErr2 := repository.ReturnUser(request.GetReqparm1(),ctx)

	if ctxErr2 != nil{
		response := &proto.FeedResponse{Resparm1: "Request timeout. Try again",Resparm2: ""}
		return response, nil
	}

	for e:= presentUser.Followers.Front(); e != nil; e = e.Next(){
		followUser := e.Value.(authmodel.User)
		userNameList += followUser.Username + ","
	}
	if userNameList[len(userNameList)-1] == byte(','){
		userNameList = userNameList[:len(userNameList)-1]
	}

	response := &proto.FeedResponse{Resparm1: "",Resparm2: userNameList}
	return response, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	proto.RegisterTwitterServer(s, &server{})

	s.Serve(lis)
}
