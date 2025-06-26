package comms

import (
	"context"
	"log"
	"time"

	"github.com/prabhakk-mw/learngo/mw/common/defs"
	"github.com/prabhakk-mw/learngo/mw/gateway/internal/utils"
	pb "github.com/prabhakk-mw/learngo/mw/services/capitalize/pb"
)

func callCapGrpcService(client pb.CapServiceClient, payload string) (capitalizedPayload *pb.CapResponse, err error) {

	// Create a context that expects the grpc server to respond within 1 second.
	grpcCallCtx, grpcCallCancel := context.WithTimeout(context.Background(), time.Second)
	defer grpcCallCancel()

	res, err := client.Capitalize(grpcCallCtx, &pb.CapRequest{Payload: payload})
	if err != nil {
		log.Printf("Could not call Capitalize: %v\n", err)
		return nil, err
	}

	return res, nil
}

func (handlers *defs.Handlers) startGRPCServer(reuseServer bool) (serverInfo defs.ServerInfo, cancel context.CancelFunc) {

	if reuseServer {
		if handlers.grpcServerInfo == nil {
			// The lifetime of the reusable server is tied to the root context.
			serverInfo = utils.StartGRPCServer(handlers.RootCtx, "capitalize")
			handlers.grpcServerInfo = &serverInfo
		} else {
			serverInfo = *handlers.grpcServerInfo
		}
		cancel = nil

	} else {
		// The context is per request now, this ensures that the server is shutdown when the request is complete.
		ctx, cancelFcn := context.WithTimeout(handlers.RootCtx, 3*time.Second)
		serverInfo = utils.StartGRPCServer(ctx, "capitalize")
		cancel = cancelFcn
	}
	return serverInfo, cancel

}
