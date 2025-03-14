package servergrpc

import (
	"bank-service/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type AppGRPC struct {
	gRPCServer *grpc.Server
	port       int
}

func RegisterBankServer(s *grpc.Server, bankServer *BankServer) {
	proto.RegisterBankServiceServer(s, bankServer)
}

func NewAppGRPC(bankServer *BankServer, port int) *AppGRPC {
	server := grpc.NewServer()
	RegisterBankServer(server, bankServer)
	return &AppGRPC{
		gRPCServer: server,
		port:       port,
	}
}

func (app *AppGRPC) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", app.port))
	if err != nil {
		return err
	}
	log.Printf("Server listening at %v", lis.Addr())
	err = app.gRPCServer.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

func (app *AppGRPC) Stop() {
	app.gRPCServer.GracefulStop()
}
