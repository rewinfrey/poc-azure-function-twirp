package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	pb "github.com/github/functions-demo/pkg/proto"
)

func getPort() string {
	port := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
	return port
}

func main() {
	client := pb.NewHelloWorldProtobufClient("http://localhost:7071", &http.Client{})

	resp, err := client.Hello(context.Background(), &pb.HelloReq{Subject: "World"})
	if err == nil {
		fmt.Println(resp.Text) // prints "Hello World"
	}
}
