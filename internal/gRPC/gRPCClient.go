package servergrpc

import (
	"bank-service/internal/services"
	"bank-service/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func RegisterBankServer(s *grpc.Server, bankService services.BankService) {
	proto.RegisterBankServiceServer(s, &BankServer{bankService: bankService})
}

func NewApp(bankService services.BankService, port int) *App {
	server := grpc.NewServer()
	RegisterBankServer(server, bankService)
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
	err = app.gRPCServer.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

func (app *App) Stop() {
	app.gRPCServer.GracefulStop()
}
