package capservice

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	pb "github.com/prabhakk-mw/learngo/capitalize-spawn-ms/proto"
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

func StartCapServiceOn(port string, errChan chan error) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		errChan <- fmt.Errorf("failed to listen: %v", err)
		return
	}
	grpcServer := grpc.NewServer()
	// Register the service
	pb.RegisterCapServiceServer(grpcServer, &capServer{})
	log.Printf("Cap Microservice Server is listening on port %v", port)
	if err := grpcServer.Serve(lis); err != nil {
		errChan <- fmt.Errorf("failed to serve: %v", err)
	}
	errChan <- nil // Indicate that the server started successfully
}

func StartCapService(ctx context.Context, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	grpcServer := grpc.NewServer()
	// Register the service
	pb.RegisterCapServiceServer(grpcServer, &capServer{})
	log.Printf("Cap Microservice Server is listening on port %v", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	select {}
}
