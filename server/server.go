package server

import (
	"context"
	"net"
	"sync"
	"unary-rpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type User struct {
	Id   string
	Name string
	Age  int32
}

func Run() {
	lsiten, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, NewService())
	reflection.Register(s)

	err = s.Serve(lsiten)
	if err != nil {
		panic(err)
	}
}

type UserService struct {
	pb.UnimplementedUserServer

	users map[string]*User
	mu    sync.Mutex // guards
}

func NewService() *UserService {
	return &UserService{
		users: make(map[string]*User),
	}
}

func (us *UserService) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	us.mu.Lock()
	defer us.mu.Unlock()

	user := &User{
		Id:   req.Id,
		Name: req.Name,
		Age:  req.Age,
	}

	us.users[user.Id] = user

	return &pb.AddUserResponse{
		Id:   user.Id,
		Age:  user.Age,
		Name: user.Name,
	}, nil

}

func (us *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	us.mu.Lock()
	defer us.mu.Unlock()

	user, ok := us.users[req.Id]
	if !ok {
		return nil, nil
	}

	return &pb.GetUserResponse{
		Id:   user.Id,
		Age:  user.Age,
		Name: user.Name,
	}, nil

}
