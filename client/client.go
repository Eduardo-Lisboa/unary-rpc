package client

import (
	"context"
	"fmt"
	"unary-rpc/pb"

	"google.golang.org/grpc"
)

func Run() {
	dial, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer dial.Close()

	userClient := pb.NewUserClient(dial)
	user, err := userClient.AddUser(context.Background(), &pb.AddUserRequest{
		Id:   "1",
		Name: "John Doe",
		Age:  30,
	})
	if err != nil {
		panic(err)

	}

	fmt.Printf("User added: %v\n", user)

	getUserResponse, err := userClient.GetUser(context.Background(), &pb.GetUserRequest{Id: "1"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("User fetched: %v\n", getUserResponse)

}
