package main

import (
	authmodel "auth/model"
	repository "auth/repository"
	"container/list"
	"errors"
	"fmt"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	
	"auth/authpb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"log"
	"net"
)


func RegisterService (r *http.Request)(string){

	r.ParseForm()

	_, usernameExists :=  repository.ReturnUser(r.Form["username"][0])

	if usernameExists {
		return "Email already in use!"
	}

	registerFromInput := authmodel.User{
		Username: r.Form["username"][0],
		Password: r.Form["password"][0],
		FirstName: r.Form["firstname"][0],
		LastName: r.Form["lastname"][0],
		Followers: list.New(),
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(registerFromInput.Password), 5)
	if err != nil {
		return "Error While Hashing Password, Try Again"
	}
	registerFromInput.Password = string(hash)

	repository.SaveUser(registerFromInput)
	repository.InitialiseTweets(registerFromInput)
	return ""
}

func LoginService(w http.ResponseWriter ,r *http.Request)(string)  {


	r.ParseForm()

	user, usernameExists :=  repository.ReturnUser(r.Form["username"][0])

	if usernameExists {
		passowrdErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Form["password"][0]))

		if  passowrdErr != nil {
			return "Invalid password"
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username":  user.Username,
			"firstname": user.FirstName,
			"lastname":  user.LastName,
		})

		tokenString, loginErr := token.SignedString([]byte("secret"))

		if loginErr != nil {
			return "Error while generating token,Try again"
		}

		user.Token = tokenString
		repository.SetCurrentUser(user.Username, user)

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
		})

		return ""

	}
	return "Invalid username"
}

func SignoutService(w http.ResponseWriter, r *http.Request) error{

	c, err := r.Cookie("token")
	if err != nil{
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

func (*server) Register(ctx context.Context, request *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	
	_, usernameExists :=  repository.ReturnUser(request.Username)

	if usernameExists {
		// response := &authpb.RegisterResponse{
		// 	Message: "User already exists",
		// }
		st := status.New(codes.InvalidArgument, "User already exists")
		return nil, st.Err()
	}

	registerFromInput := authmodel.User{
		Username: request.Username,
		Password: request.Password,
		FirstName: request.Firstname,
		LastName: request.Lastname,
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


func (*server) Hello(ctx context.Context, request *authpb.HelloRequest) (*authpb.HelloResponse, error) {
	
	log.Printf("Step2!!!")

	firstname := request.Firstname
	lastname := request.Lastname
	response := &authpb.HelloResponse{
		Greeting: "Hello " + firstname + lastname,
	}
	return response, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	authpb.RegisterHelloServiceServer(s, &server{})
	authpb.RegisterRegisterServiceServer(s, &server{})

	s.Serve(lis)
}