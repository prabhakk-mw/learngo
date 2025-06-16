package capservice

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

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

func TestService(ctx context.Context, ready chan<- struct{}, done chan<- struct{}) {
	log.Println("Mark TestService as ready.")
	close(ready)
	count := 0
	// ticker := make(chan struct{})
	for {
		log.Println("running")
		count++
		if count > 5 {
			log.Println("TestService has run 5 times, exiting...")
			close(done)
			return
		}
		// Sleep for 3 seconds or until context is done
		select {
		case <-ctx.Done():
			return
		case <-time.After(1 * time.Second):
		}
	}
}
