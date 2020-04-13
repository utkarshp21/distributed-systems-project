package main

import (
	authmodel "auth/model"
	repository "auth/repository"
	"container/list"
	"errors"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"auth/authpb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SignoutService(w http.ResponseWriter, r *http.Request) error {

	c, err := r.Cookie("token")
	if err != nil {
		return err
	}

	tokenString := c.Value
	token, tokenerr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if !token.Valid || tokenerr != nil {
		return tokenerr
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	signoutUserName := claims["username"].(string)
	signoutUser, _ := repository.ReturnUser(signoutUserName)

	if signoutUser.Username != "" {
		signoutUser.Token = ""
		repository.SetCurrentUser(signoutUser.Username, signoutUser)
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   "",
			Expires: time.Unix(0, 0),
		})
		return nil
	} else {
		return errors.New("User unavailable")
	}
}

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

	s.Serve(lis)
}
