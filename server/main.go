package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strconv"
	"strings"

	pb "github.com/kyunghoj/idservice/idservice"
	"google.golang.org/grpc"
)

type idServiceServer struct {
	pb.UnimplementedIdServiceServer
}

/*
func (s *idServiceServer) CreateNewID(ctx context.Context, req *pb.IdRequest) (*pb.IdResponse, error) {
	fmt.Println("[CreateNewId] " + req.Name)
	return &pb.IdResponse{RetCode: 0, Id: 17109498, ErrorMsg: ""}, nil
}
*/

func (s *idServiceServer) GetUID(ctx context.Context, req *pb.IdRequest) (*pb.IdResponse, error) {
	fmt.Println("[GetUID] " + req.Query)
	// id -u req.Query
	cmd := exec.Command("id", "-u", req.Query)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return &pb.IdResponse{RetCode: -1, Id: -1, ErrorMsg: "User not found"}, err
	}
	var uid int32
	i, err := strconv.ParseInt(strings.TrimSuffix(out.String(), "\n"), 10, 32)
	if err != nil {
		log.Fatal(err)
		return &pb.IdResponse{RetCode: -1, Id: -1, ErrorMsg: "User not found"}, err
	}
	uid = int32(i)

	return &pb.IdResponse{RetCode: 0, Id: uid, ErrorMsg: ""}, nil
}

func (s *idServiceServer) GetGID(ctx context.Context, req *pb.IdRequest) (*pb.IdResponse, error) {
	fmt.Println("[GetGID] " + req.Query)
	// getent group req.Query
	cmd := exec.Command("getent", "group", req.Query)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return &pb.IdResponse{RetCode: -1, Id: -1, ErrorMsg: "Group not found"}, err
	}
        // "kyungho.jeon:x:1000:kyungho.jeon"
	var res string = strings.Split(strings.TrimSuffix(out.String(), "\n"), ":")[2]
	var gid int32
	i, err := strconv.ParseInt(res, 10, 32)
	if err != nil {
		log.Fatal(err)
		return &pb.IdResponse{RetCode: -1, Id: -1, ErrorMsg: "Group not found"}, err
	}
	gid = int32(i)

	return &pb.IdResponse{RetCode: 0, Id: gid, ErrorMsg: ""}, nil
}

/*
func (s *idServiceServer) DeleteID(ctx context.Context, req *pb.IdRequest) (*pb.IdResponse, error) {
	fmt.Println("[DeleteId] " + req.Name)
	return &pb.IdResponse{RetCode: 0, Id: 17109498, ErrorMsg: ""}, nil
}
*/

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
