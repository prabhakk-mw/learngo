package capservice

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	pb "github.com/prabhakk-mw/learngo/capitalize-ms-in-another-module/capservice/pb"
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

func StartCapService(ctx context.Context, port string, ready chan<- struct{}, errChan chan<- error) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		errChan <- fmt.Errorf("capservice failed to listen: %v", err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCapServiceServer(grpcServer, &capServer{})
	close(ready) // Signal that the server is ready to accept connections
	if err := grpcServer.Serve(lis); err != nil {
		errChan <- fmt.Errorf("capservice failed to serve: %v", err)
	}
}
