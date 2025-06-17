package defs

import (
	"net"

	"google.golang.org/grpc"
)

type ServerInfo struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

func NewServerInfo(grpcServer *grpc.Server, listener net.Listener) ServerInfo {
	return ServerInfo{
		grpcServer: grpcServer,
		listener:   listener,
	}
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
	return s.listener.Addr().String()
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
