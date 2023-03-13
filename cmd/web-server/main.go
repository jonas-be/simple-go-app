package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"

	pb "simple-go-application/internal/grpc" // import the generated code
)

func main() {
	http.HandleFunc("/", getGreeting)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func getGreeting(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
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
