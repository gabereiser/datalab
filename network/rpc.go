package network

import (
	"context"

	"github.com/gabereiser/datalab/rpc"
)

type AuthRpcService struct {
	rpc.UnimplementedAuthServer
}

type WorkbookRpcService struct {
	rpc.UnimplementedWorkbookServer
}

func (AuthRpcService) Login(context context.Context, in *rpc.LoginRequest) (*rpc.LoginResponse, error) {
	return nil, nil
}

func (WorkbookRpcService) CreateWorkbook(context context.Context, in *rpc.CreateWorkbookRequest) (*rpc.IDResponse, error) {
	return nil, nil
}
