package defs

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

type Handler interface {
	startGRPCServer(reuseServer bool) (ServerInfo, context.CancelFunc)
}

type ServerInfo struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

func NilServerInfo() ServerInfo {
	return ServerInfo{nil, nil}
}

func NewServerInfo(grpcServer *grpc.Server, listener net.Listener) ServerInfo {
	return ServerInfo{
		grpcServer: grpcServer,
		listener:   listener,
	}
}

func (s *ServerInfo) GetProtocol() string {
	return s.listener.Addr().Network()
}

func (s *ServerInfo) IsProtocolUDS() bool {
	return s.GetProtocol() == "unix"
}

func (s *ServerInfo) IsProtocolTCP() bool {
	return s.GetProtocol() == "tcp"
}

func (s *ServerInfo) GetPort() int {
	if s.listener == nil {
		return 0
	}
	if tcpAddr, ok := s.listener.Addr().(*net.TCPAddr); ok {
		return tcpAddr.Port
	}
	return 0
}
func (s *ServerInfo) GetAddress() string {
	if s.listener == nil {
		return ""
	}
	suffix := ""
	if s.IsProtocolUDS() {
		suffix = "unix://"
	}
	return suffix + s.listener.Addr().String()
}
func (s *ServerInfo) GetGRPCServer() *grpc.Server {
	if s.grpcServer == nil {
		return nil
	}
	return s.grpcServer
}
func (s *ServerInfo) GetListener() net.Listener {
	if s.listener == nil {
		return nil
	}
	return s.listener
}
