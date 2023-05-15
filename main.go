package main

import (
	"context"
	pb "github.com/bysouffle/ccggaag/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type grpcapi interface {
	Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error)
}

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	log.Println(r.GetRequest())
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = server.Serve(lis)
	if err != nil {
		return
	}
}
