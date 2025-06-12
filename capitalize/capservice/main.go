package main

import (
	"context"
	"log"
	"net"
	"strings"

	pb "github.com/prabhakk-mw/learngo/capitalize/proto"
	"google.golang.org/grpc"
)

type capServer struct {
	pb.CapServiceServer
}

func (s *capServer) Capitalize(ctx context.Context, req *pb.CapRequest) (res *pb.CapResponse, err error) {
	capitalizedPayload := strings.ToUpper(req.GetPayload())
	res = &pb.CapResponse{Payload: capitalizedPayload}
	log.Printf("Received request: %s, responding with: %s", req.GetPayload(), capitalizedPayload)
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	// Register the service
	pb.RegisterCapServiceServer(grpcServer, &capServer{})
	log.Printf("Server is listening on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
