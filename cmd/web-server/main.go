package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"

	pb "simple-go-application/internal/grpc" // import the generated code
)

const targetEnvKey = "TARGET"
const listenPort = "8080"

func main() {
	fmt.Printf("Start on port %v and the target is %v", listenPort, os.Getenv(targetEnvKey))
	http.HandleFunc("/", getGreeting)

	err := http.ListenAndServe(":"+listenPort, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func getGreeting(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")

	conn, err := grpc.Dial(os.Getenv(targetEnvKey), grpc.WithInsecure())
	if err != nil {
		http.Error(w, fmt.Sprintf("did not connect: %v", err), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	name := "Alice"
	req := &pb.GreetRequest{Name: name}
	res, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Fprintf(w, res.GetMessage())
}
