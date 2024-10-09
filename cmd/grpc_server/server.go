package main

import (
	"context"
	"log"

	pb "github.com/Danya97i/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedChatV1Server
}

func (s *server) Create(_ context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	log.Println("create chat request", req)
	return nil, nil
}

func (s *server) Delete(_ context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	log.Println("create chat request", req)
	return nil, nil
}

func (s *server) SendMesage(_ context.Context, req *pb.SendMessageRequest) (*emptypb.Empty, error) {
	log.Println("send message request", req)
	return nil, nil
}
