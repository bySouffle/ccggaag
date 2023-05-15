package client

import (
	"context"
	"fmt"
	pb "github.com/bysouffle/ccggaag/proto"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"log"
)

const PORT = "9001"

func req() {
	conn, err := grpc.Dial("127.0.0.1:"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	jsonstr := fmt.Sprintf("{\"%s\":\"%s\"}", "request", "grpc")
	pbdata := &pb.SearchRequest{}
	err = jsonpb.UnmarshalString(jsonstr, pbdata)
	if err != nil {
		log.Printf("json UnmarshalString err")
		return
	}

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), pbdata)
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	marshaler := jsonpb.Marshaler{}
	response, err := marshaler.MarshalToString(resp)
	if err != nil {
		log.Printf("json trans err")
		return
	}
	log.Printf("resp : %s", response)
}
