package main

import (
	authmodel "auth/model"
	repository "auth/repository"
	"container/list"
	"fmt"
	//authStorage "auth/storage"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	profileRepository "profile/repository"

	"auth/authpb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

)

type server struct {
}

func (*server) Login(ctx context.Context, request *authpb.LoginRequest) (*authpb.LoginResponse, error) {

	user, usernameExists := repository.ReturnUser(request.Username)

	if usernameExists {
		passowrdErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

		if passowrdErr != nil {
			st := status.New(codes.InvalidArgument, "Invalid Password!")
			return nil, st.Err()
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username":  user.Username,
			"firstname": user.FirstName,
			"lastname":  user.LastName,
		})

		tokenString, loginErr := token.SignedString([]byte("secret"))

		if loginErr != nil {
			st := status.New(codes.Unknown, "Error while generating token,Try again")
			return nil, st.Err()
		}

		user.Token = tokenString
		repository.SetCurrentUser(user.Username, user)

		response := &authpb.LoginResponse{
			Message:     "Successfully Logged In!!",
			Tokenstring: tokenString,
		}

		return response, nil

	}

	st := status.New(codes.InvalidArgument, "Invalid username")

	return nil, st.Err()

}

func (*server) Register(ctx context.Context, request *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {

	_, usernameExists := repository.ReturnUser(request.Username)

	if usernameExists {
		st := status.New(codes.InvalidArgument, "User already exists")
		return nil, st.Err()
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
		st := status.New(codes.Unknown, "Error While Hashing Password, Try Again")
		return nil, st.Err()
	}
	registerFromInput.Password = string(hash)

	repository.SaveUser(registerFromInput)
	repository.InitialiseTweets(registerFromInput)

	response := &authpb.RegisterResponse{
		Message: "Successfully registered",
	}

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
		st := status.New(codes.Unknown, "Not Logged In")
		return nil, st.Err()
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	signoutUserName := claims["username"].(string)
	signoutUser, _ := repository.ReturnUser(signoutUserName)

	if signoutUser.Username != "" {
		signoutUser.Token = ""
		repository.SetCurrentUser(signoutUser.Username, signoutUser)
		response := &authpb.LogoutResponse{
			Message: "Successfully LoggedOut",
		}
		return response, nil
	} else {
		st := status.New(codes.Unknown, "User unavailable")
		return nil, st.Err()
	}
}

func (*server) FollowService(ctx context.Context, request *authpb.ProfileRequest) (*authpb.ProfileResponse, error) {

	userPresent, _ := repository.ReturnUser(request.GetReqparm1())
	followUser, _ := repository.ReturnUser(request.GetReqparm2())

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
		repository.SaveUser(followUser)
		response := &authpb.ProfileResponse{Resparm1: ""}
		return response, nil
	} else {
		response := &authpb.ProfileResponse{Resparm1: "Username doesnt exist"}
		return response, nil
	}

}

func (*server) UnfollowService(ctx context.Context, request *authpb.ProfileRequest) (*authpb.ProfileResponse, error) {

	userPresent, _ := repository.ReturnUser(request.GetReqparm1())
	unfollowUser, _ := repository.ReturnUser(request.GetReqparm2())

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
			repository.SaveUser(unfollowUser)
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
		profileRepository.SaveTweet(tweetUser,tweetContent)
		response := &authpb.ProfileResponse{Resparm1: ""}
		return response, nil
	} else {
		response := &authpb.ProfileResponse{Resparm1: "Enter tweet content"}
		return response, nil
	}
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

	s.Serve(lis)
}
