package capitalize

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
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
	log.Printf("Capitalize[%s]=>[%s]", req.GetPayload(), capitalizedPayload)
	return res, nil
}

func startCapitalizeServiceImpl(ctx context.Context,
	serverInfo chan<- defs.ServerInfo,
	errChan chan<- error,
	protocol string,
	address string) {

	defer log.Println("Shutting down CaptializeService on " + protocol)

	lis, err := net.Listen(protocol, address)
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

func StartCapitalizeService(ctx context.Context, serverInfo chan<- defs.ServerInfo, errChan chan<- error) {
	startCapitalizeServiceImpl(ctx, serverInfo, errChan, "tcp", ":0" /*Use the next available port*/)
}

func StartCapitalizeServiceOnUDS(ctx context.Context, serverInfo chan<- defs.ServerInfo, errChan chan<- error) {
	tempFile, err := os.CreateTemp("", "CapitalizeService-*.sock")
	if err != nil {
		log.Fatal(err)
	}

	sockAddr := tempFile.Name()
	os.Remove(tempFile.Name())

	if _, err := os.Stat(sockAddr); !os.IsNotExist(err) {
		if err := os.RemoveAll(sockAddr); err != nil {
			log.Fatal(err)
		}
	}
	startCapitalizeServiceImpl(ctx, serverInfo, errChan, "unix", sockAddr)
}
