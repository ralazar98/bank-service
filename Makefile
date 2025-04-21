gen:
	mockgen -source=D:\GO\bank-service\proto\bank-service_grpc.pb.go -destination=mocks/mock_bank.go -package=mocks

mockgen -source=d:\\go\bank-service\internal\rabbit\rabbit.go -destination=mocks/rabbit_mock.go -package=mocks
