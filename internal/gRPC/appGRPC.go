package servergrpc

import (
	"bank-service/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type AppGRPC struct {
	GRPCServer *grpc.Server
	port       int
}

func NewAppGRPC(bankServer *BankServer, port int) *AppGRPC {
	server := grpc.NewServer()
	proto.RegisterBankServiceServer(server, bankServer)
	return &AppGRPC{
		GRPCServer: server,
		port:       port,
	}
}

func (app *AppGRPC) Run() error {
	log.Println("Starting server...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", app.port))
	if err != nil {
		return err
	}
	log.Printf("Server listening at %v", lis.Addr())
	err = app.GRPCServer.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

func (app *AppGRPC) Stop() {
	app.GRPCServer.GracefulStop()
}
