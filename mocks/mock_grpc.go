package mocks

import (
	proto "bank-service/proto"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockBankServiceClient is a mock of BankServiceClient interface.
type MockBankServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockBankServiceClientMockRecorder
}

// MockBankServiceClientMockRecorder is the mock recorder for MockBankServiceClient.
type MockBankServiceClientMockRecorder struct {
	mock *MockBankServiceClient
}

// NewMockBankServiceClient creates a new mock instance.
func NewMockBankServiceClient(ctrl *gomock.Controller) *MockBankServiceClient {
	mock := &MockBankServiceClient{ctrl: ctrl}
	mock.recorder = &MockBankServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankServiceClient) EXPECT() *MockBankServiceClientMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockBankServiceClient) CreateAccount(ctx context.Context, in *proto.CreateRequest, opts ...grpc.CallOption) (*proto.CreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAccount", varargs...)
	ret0, _ := ret[0].(*proto.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockBankServiceClientMockRecorder) CreateAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockBankServiceClient)(nil).CreateAccount), varargs...)
}

// GetBalance mocks base method.
func (m *MockBankServiceClient) GetBalance(ctx context.Context, in *proto.GetBalanceRequest, opts ...grpc.CallOption) (*proto.GetBalanceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBalance", varargs...)
	ret0, _ := ret[0].(*proto.GetBalanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockBankServiceClientMockRecorder) GetBalance(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockBankServiceClient)(nil).GetBalance), varargs...)
}

// UpdateBalance mocks base method.
func (m *MockBankServiceClient) UpdateBalance(ctx context.Context, in *proto.UpdateBalanceRequest, opts ...grpc.CallOption) (*proto.UpdateBalanceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateBalance", varargs...)
	ret0, _ := ret[0].(*proto.UpdateBalanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBalance indicates an expected call of UpdateBalance.
func (mr *MockBankServiceClientMockRecorder) UpdateBalance(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBalance", reflect.TypeOf((*MockBankServiceClient)(nil).UpdateBalance), varargs...)
}

// MockBankServiceServer is a mock of BankServiceServer interface.
type MockBankServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockBankServiceServerMockRecorder
}

// MockBankServiceServerMockRecorder is the mock recorder for MockBankServiceServer.
type MockBankServiceServerMockRecorder struct {
	mock *MockBankServiceServer
}

// NewMockBankServiceServer creates a new mock instance.
func NewMockBankServiceServer(ctrl *gomock.Controller) *MockBankServiceServer {
	mock := &MockBankServiceServer{ctrl: ctrl}
	mock.recorder = &MockBankServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankServiceServer) EXPECT() *MockBankServiceServerMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockBankServiceServer) CreateAccount(arg0 context.Context, arg1 *proto.CreateRequest) (*proto.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(*proto.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockBankServiceServerMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockBankServiceServer)(nil).CreateAccount), arg0, arg1)
}

// GetBalance mocks base method.
func (m *MockBankServiceServer) GetBalance(arg0 context.Context, arg1 *proto.GetBalanceRequest) (*proto.GetBalanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetBalanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockBankServiceServerMockRecorder) GetBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockBankServiceServer)(nil).GetBalance), arg0, arg1)
}

// UpdateBalance mocks base method.
func (m *MockBankServiceServer) UpdateBalance(arg0 context.Context, arg1 *proto.UpdateBalanceRequest) (*proto.UpdateBalanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBalance", arg0, arg1)
	ret0, _ := ret[0].(*proto.UpdateBalanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBalance indicates an expected call of UpdateBalance.
func (mr *MockBankServiceServerMockRecorder) UpdateBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBalance", reflect.TypeOf((*MockBankServiceServer)(nil).UpdateBalance), arg0, arg1)
}

// mustEmbedUnimplementedBankServiceServer mocks base method.
func (m *MockBankServiceServer) mustEmbedUnimplementedBankServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedBankServiceServer")
}

// mustEmbedUnimplementedBankServiceServer indicates an expected call of mustEmbedUnimplementedBankServiceServer.
func (mr *MockBankServiceServerMockRecorder) mustEmbedUnimplementedBankServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedBankServiceServer", reflect.TypeOf((*MockBankServiceServer)(nil).mustEmbedUnimplementedBankServiceServer))
}

// MockUnsafeBankServiceServer is a mock of UnsafeBankServiceServer interface.
type MockUnsafeBankServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeBankServiceServerMockRecorder
}

// MockUnsafeBankServiceServerMockRecorder is the mock recorder for MockUnsafeBankServiceServer.
type MockUnsafeBankServiceServerMockRecorder struct {
	mock *MockUnsafeBankServiceServer
}

// NewMockUnsafeBankServiceServer creates a new mock instance.
func NewMockUnsafeBankServiceServer(ctrl *gomock.Controller) *MockUnsafeBankServiceServer {
	mock := &MockUnsafeBankServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeBankServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeBankServiceServer) EXPECT() *MockUnsafeBankServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedBankServiceServer mocks base method.
func (m *MockUnsafeBankServiceServer) mustEmbedUnimplementedBankServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedBankServiceServer")
}

// mustEmbedUnimplementedBankServiceServer indicates an expected call of mustEmbedUnimplementedBankServiceServer.
func (mr *MockUnsafeBankServiceServerMockRecorder) mustEmbedUnimplementedBankServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedBankServiceServer", reflect.TypeOf((*MockUnsafeBankServiceServer)(nil).mustEmbedUnimplementedBankServiceServer))
}
