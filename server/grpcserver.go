package server

import (
	"context"
	"github.com/ivansukach/profile-service/protocol"
	"github.com/ivansukach/profile-service/repositories"
)

//GRPCServer ...
type Server struct {
	Rps
}

//Add ...
func (s *Server) Create(ctx context.Context, req *protocol.CreateRequest) (*protocol.SuccessResponse, error) {

	return &protocol.SuccessResponse{Success: true}, err
}
