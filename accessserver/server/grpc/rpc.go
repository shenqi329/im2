package grpc

import (
	"golang.org/x/net/context"
	//"google.golang.org/grpc"
	"im/protocol/client"
)

type Rpc struct{}

func (r *Rpc) Rpc(ctx context.Context, request *client.RpcRequest) (*client.RpcResponse, error) {

	response := &client.RpcResponse{
		Rid:    request.Rid,
		ConnId: request.ConnId,
	}

	return response, nil
}
