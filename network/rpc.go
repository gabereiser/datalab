package network

import (
	"context"

	rpc "github.com/gabereiser/datalab/rpc"
	ath "github.com/gabereiser/datalab/rpc/auth"
	wb "github.com/gabereiser/datalab/rpc/workbook"
)

type AuthRpcService struct {
	rpc.UnimplementedAuthServer
}

type WorkbookRpcService struct {
	rpc.UnimplementedWorkbookServer
}

func (AuthRpcService) Login(context context.Context, in *ath.LoginRequest) (*ath.LoginResponse, error) {
	return nil, nil
}

func (WorkbookRpcService) CreateWorkbook(context context.Context, in *wb.CreateWorkbookRequest) (*rpc.IDResponse, error) {
	return nil, nil
}
