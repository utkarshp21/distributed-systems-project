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
		// log.Printf(r.Form["username"][1])
		request := &authpb.RegisterRequest{Firstname: r.Form["firstname"][0], Lastname:r.Form["lastname"][0], Username:r.Form["username"][0], Password:r.Form["password"][0]}

		//request := &authpb.RegisterRequest{Firstname: "Utkarsh", Lastname: "Prakash", Username: "up@gmail.com", Password: "up"}

		_, err := client.Register(context.Background(), request)

		if err != nil {
			log.Printf("Receive Error Regiseter response => [%v]", err)
			m["Error"] = err
			log.Println(err)
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
		// log.Printf("Username", r.Form["username"][0])

		//request := &authpb.LoginRequest{Username: "up@gmail.com", Password: "up"}

		response, errMsg := client.Login(context.Background(), request)

		if errMsg != nil {
			m["Error"] = errMsg
			log.Println(errMsg)
			t.Execute(w, m)
			return
		} else {
			http.SetCookie(w, &http.Cookie{
				Name:  "token",
				Value: response.Tokenstring,
			})
			log.Println(response.Message)
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
		m["Error"] = "Please login to continue!"
		m["Success"] = nil
		log.Println(cerr)
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

	response, err := client.Logout(context.Background(), request)

	if err != nil {
		m["Error"] = "Please login to continue!"
		m["Success"] = nil
		log.Println(err)
		t.Execute(w, m)
		return
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   "",
			Expires: time.Unix(0, 0),
		})
		log.Println(response.Message)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}
