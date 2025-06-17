package capitalize

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	pb "github.com/prabhakk-mw/learngo/mw/services/capitalize/pb"
	"google.golang.org/grpc"
)

type capServer struct {
	pb.CapServiceServer
}

func capitalizeCore(payload string) string {
	return strings.ToUpper(payload)
}

func (s *capServer) Capitalize(ctx context.Context, req *pb.CapRequest) (res *pb.CapResponse, err error) {
	capitalizedPayload := capitalizeCore(req.GetPayload())
	res = &pb.CapResponse{Payload: capitalizedPayload}
	log.Printf("Received request: %s, responding with: %s", req.GetPayload(), capitalizedPayload)
	return res, nil
}

func StartCapService(ctx context.Context, ready chan<- int, errChan chan<- error) {

	lis, err := net.Listen("tcp", ":0") // Use the next available port
	if err != nil {
		errChan <- fmt.Errorf("capservice failed to listen: %v", err)
		return
	}

	port := lis.Addr().(*net.TCPAddr).Port

	grpcServer := grpc.NewServer()
	pb.RegisterCapServiceServer(grpcServer, &capServer{})
	ready <- port // Signal that the server is ready to accept connections

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("CapService failed to serve on port %d: %v", port, err)
		errChan <- fmt.Errorf("capservice failed to serve: %v", err)
	}
}

type ServerInfo struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

func StartCapitalizeService(ctx context.Context, serverInfo chan<- ServerInfo, errChan chan<- error) {

	lis, err := net.Listen("tcp", ":0") // Use the next available port
	if err != nil {
		errChan <- fmt.Errorf("capservice failed to listen: %v", err)
		return
	}

	port := lis.Addr().(*net.TCPAddr).Port

	grpcServer := grpc.NewServer()
	pb.RegisterCapServiceServer(grpcServer, &capServer{})

	serverInfo <- ServerInfo{grpcServer, lis}

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("CapService failed to serve on port %d: %v", port, err)
		errChan <- fmt.Errorf("capservice failed to serve: %v", err)
	}
}
