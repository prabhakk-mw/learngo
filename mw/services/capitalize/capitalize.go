package capitalize

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/prabhakk-mw/learngo/mw/common/defs"
	pb "github.com/prabhakk-mw/learngo/mw/services/capitalize/pb"
	"google.golang.org/grpc"
)

type capServer struct {
	pb.CapServiceServer
}

func capitalizeCore(payload string) string {
	return strings.ToUpper(payload)
}

func (s *capServer) Capitalize(_ context.Context, req *pb.CapRequest) (res *pb.CapResponse, err error) {
	capitalizedPayload := capitalizeCore(req.GetPayload())
	res = &pb.CapResponse{Payload: capitalizedPayload}
	log.Printf("Received request: %s, responding with: %s", req.GetPayload(), capitalizedPayload)
	return res, nil
}

func StartCapitalizeService(ctx context.Context, serverInfo chan<- defs.ServerInfo, errChan chan<- error) {
	defer log.Println("Shutting down CaptializeService")

	lis, err := net.Listen("tcp", ":0") // Use the next available port
	if err != nil {
		errChan <- fmt.Errorf("capservice failed to listen: %v", err)
		return
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCapServiceServer(grpcServer, &capServer{})

	newServerInfo := defs.NewServerInfo(grpcServer, lis)
	serverInfo <- newServerInfo

	port := newServerInfo.GetPort()

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("CapService failed to serve on port %d: %v", port, err)
			errChan <- fmt.Errorf("capservice failed to serve: %v", err)
		}
	}()

	// Wait for a shutdown signal
	<-ctx.Done()
	log.Printf("Context Shutdown received, stopping GRPCServer:%s", newServerInfo.GetAddress())
	grpcServer.GracefulStop()
}
