package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"

    pb "github.com/klayhu/studyGolang/studyGolang/gRPC/example/proto/search"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
    return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

func (s *SearchService) mustEmbedUnimplementedSearchServiceServer() {}

const PORT = "9001"

func main() {
    server := grpc.NewServer()
    pb.RegisterSearchServiceServer(server, &SearchService{})

    lis, err := net.Listen("tcp", ":"+PORT)
    if err != nil {
        log.Fatalf("net.Listen err: %v", err)
    }

    server.Serve(lis)
}