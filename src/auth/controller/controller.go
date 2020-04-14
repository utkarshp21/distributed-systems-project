package controller

import (
	"html/template"
	"net/http"
	"time"
	"auth/authpb"
	"context"
	"log"
	"google.golang.org/grpc"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	_, err := r.Cookie("token")

	if err == nil {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	m := map[string]interface{}{}
	t, _ := template.ParseFiles("register.gtpl")

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

		client := authpb.NewRegisterServiceClient(cc)

		r.ParseForm()
		request := &authpb.RegisterRequest{Firstname: r.Form["firstname"][0], Lastname:r.Form["lastname"][0], Username:r.Form["username"][0], Password:r.Form["password"][0]}
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

	t, _ := template.ParseFiles("login.gtpl")

	_, err := r.Cookie("token")

	if err == nil {
		http.Redirect(w, r, "/profile", http.StatusFound)
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

		client := authpb.NewLoginServiceClient(cc)

		r.ParseForm()
		request := &authpb.LoginRequest{Username: r.Form["username"][0], Password: r.Form["password"][0]}
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
			http.Redirect(w, r, "/profile", http.StatusFound)
			return
		}
	}
}

func SignoutHandler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("login.gtpl")
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

	client := authpb.NewLogoutServiceClient(cc)

	request := &authpb.LogoutRequest{Tokenstring: tokenString}

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
