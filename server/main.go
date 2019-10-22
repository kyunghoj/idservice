package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/kyunghoj/idservice/idservice"
	"google.golang.org/grpc"
)

type idServiceServer struct {
	pb.UnimplementedIdServiceServer
}

func (s *idServiceServer) CreateNewID(ctx context.Context, req *pb.IdRequest) (*pb.IdResponse, error) {
	fmt.Println("[CreateNewId] " + req.Name)
	return &pb.IdResponse{RetCode: 0, Id: 17109498, ErrorMsg: ""}, nil
}

func (s *idServiceServer) GetID(ctx context.Context, req *pb.IdRequest) (*pb.IdResponse, error) {
	fmt.Println("[GetId] " + req.Name)
	return &pb.IdResponse{RetCode: 0, Id: 17109498, ErrorMsg: ""}, nil
}

func (s *idServiceServer) DeleteID(ctx context.Context, req *pb.IdRequest) (*pb.IdResponse, error) {
	fmt.Println("[DeleteId] " + req.Name)
	return &pb.IdResponse{RetCode: 0, Id: 17109498, ErrorMsg: ""}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":30010")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterIdServiceServer(s, &idServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
