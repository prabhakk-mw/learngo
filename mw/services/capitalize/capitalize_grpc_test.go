package capitalize

import (
	"context"
	"io"
	"log"
	"testing"
	"time"

	"github.com/prabhakk-mw/learngo/mw/common/defs"
	pb "github.com/prabhakk-mw/learngo/mw/services/capitalize/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var testpoints = []struct {
	input    string
	expected string
}{
	{"hello world", "HELLO WORLD"},
	{"go programming", "GO PROGRAMMING"},
	{"capitalize this", "CAPITALIZE THIS"},
	{"", ""},
}

var req = &pb.CapRequest{
	Payload: "test payload",
}

func setupTestServer(ctx context.Context, onUDS bool) defs.ServerInfo {

	serverInfoChan := make(chan defs.ServerInfo)
	errChan := make(chan error, 1)

	if onUDS {
		go StartCapitalizeServiceOnUDS(ctx, serverInfoChan, errChan)
	} else {
		go StartCapitalizeService(ctx, serverInfoChan, errChan)
	}

	// This line will block until the service responds with the port number.
	// It also marks that the service is ready to accept requests.
	select {
	case serverInfo := <-serverInfoChan:
		return serverInfo

	case err := <-errChan:
		log.Println("Server failed to start:", err)
		return defs.NilServerInfo()

	case <-ctx.Done():
		err := ctx.Err()
		log.Println("Context cancelled while starting the capitalization service:", err)
		return defs.NilServerInfo()
	}
}

func TestCapitalizeService(test *testing.T) {
	// redirect log output to discard to avoid cluttering test output
	log.SetOutput(io.Discard)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	servers := []defs.ServerInfo{
		setupTestServer(ctx, false /*onUDS*/),
		setupTestServer(ctx, true /*onUDS*/),
	}
	for _, server := range servers {
		defer server.GetGRPCServer().Stop()
	}

	for _, server := range servers {

		log.Println("Capitalize service started at address:", server.GetAddress())

		conn, err := grpc.NewClient(server.GetAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			test.Fatalf("Failed to connect to server: %v", err)
		}
		defer conn.Close()

		client := pb.NewCapServiceClient(conn)

		for _, testpoint := range testpoints {
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
			protocol := "TCP"
			if server.IsProtocolUDS() {
				protocol = "UDS"
			}
			test.Logf("grpc over %s Capitalize(%q) => %q; expected %q", protocol, testpoint.input, result, testpoint.expected)
		}
	}
}
