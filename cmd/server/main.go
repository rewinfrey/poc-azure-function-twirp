package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	pb "github.com/github/functions-demo/pkg/proto"
)

type HelloWorldServer struct{}

func (s *HelloWorldServer) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	fmt.Println("inside hello")
	return &pb.HelloResp{Text: "Hello " + req.Subject}, nil
}

func getPort() string {
	port := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
	return port
}

// Run the implementation in a local server
func main() {
	twirpHandler := pb.NewHelloWorldServer(&HelloWorldServer{})
	// You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
	mux := http.NewServeMux()
	// The generated code includes a method, PathPrefix(), which
	// can be used to mount your service on a mux.
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	http.ListenAndServe(getPort(), mux)
}
