package servergrpc

import (
	"bank-service/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func RegisterBankServer(s *grpc.Server, bankServer *BankServer) {
	proto.RegisterBankServiceServer(s, bankServer)
}

func NewApp(bankServer *BankServer, port int) *App {
	server := grpc.NewServer()
	RegisterBankServer(server, bankServer)
	return &App{
		gRPCServer: server,
		port:       port,
	}
}

func (app *App) Run() error {
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

func (app *App) Stop() {
	app.gRPCServer.GracefulStop()
}
