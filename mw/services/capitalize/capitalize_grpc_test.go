package capitalize

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	pb "github.com/prabhakk-mw/learngo/mw/services/capitalize/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func setupTestServer(ctx context.Context) (*grpc.Server, net.Listener) {

	serverInfoChan := make(chan ServerInfo)
	errChan := make(chan error, 1)
	go StartCapitalizeService(ctx, serverInfoChan, errChan)

	// This line will block until the service responds with the port number.
	// It also marks that the service is ready to accept requests.
	select {
	case serverInfo := <-serverInfoChan:
		return serverInfo.grpcServer, serverInfo.listener

	case err := <-errChan:
		log.Println("Server failed to start:", err)
		return nil, nil

	case <-ctx.Done():
		err := ctx.Err()
		log.Println("Context cancelled while starting the capitalization service:", err)
		return nil, nil
	}
}

func TestCapitalizeService(test *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	grpcServer, lis := setupTestServer(ctx)
	defer grpcServer.Stop()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		test.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewCapServiceClient(conn)

	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "HELLO WORLD"},
		{"go programming", "GO PROGRAMMING"},
		{"capitalize this", "CAPITALIZE THIS"},
		{"", ""},
	}

	req := &pb.CapRequest{
		Payload: "test payload",
	}

	for _, testpoint := range tests {
		req = &pb.CapRequest{
			Payload: testpoint.input,
		}

		resp, err := client.Capitalize(ctx, req)
		if err != nil {
			test.Fatalf("Failed to call Capitalize: %v", err)
		}
		result := resp.GetPayload()
		if result != testpoint.expected {
			test.Errorf("Capitalize(%q) = %q; want %q", testpoint.input, result, testpoint.expected)
		}
		log.Printf("Test passed: received capitalized payload: %s", resp.Payload)
	}
}
