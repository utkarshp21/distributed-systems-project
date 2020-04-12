
package main

import (
	"context"
	"fmt"
	"auth/authpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello client ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	// client := authpb.NewHelloServiceClient(cc)

	// request := &authpb.HelloRequest{Firstname: "Utkarsh", Lastname:"Prakash"}
	// resp, _ := client.Hello(context.Background(), request)
	// fmt.Printf("Receive Hello response => [%v]", resp.Greeting)

	client := authpb.NewRegisterServiceClient(cc)
	
	request := &authpb.RegisterRequest{Firstname: "Utkarsh", Lastname:"Prakash", Username:"up@gmail.com", Password:"up"}
	
	resp, err := client.Register(context.Background(), request)

	if err != nil{
		fmt.Printf("Receive Error Regiseter response => [%v]", err)
	}else{
		fmt.Printf("Receive  Regiseter response => [%v]", resp.Message)
	}
}