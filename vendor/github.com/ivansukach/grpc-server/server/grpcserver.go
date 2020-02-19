package server

import (
	"context"
	"github.com/ivansukach/grpc-server/protocol"
)

//GRPCServer ...
type GRPCServer struct{}

//Add ...
func (s *GRPCServer) GiveResponse(ctx context.Context, req *protocol.GRRequest) (*protocol.GRResponse, error) {

	return &protocol.GRResponse{Res: req.GetReq() + " Pong"}, nil
}
