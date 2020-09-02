package adder

import (
	"context"

	"github.com/HunterGooD/gRPC_test/pkg/api"
)

// GRPCServer ...
type GRPCServer struct {
}

// Add ...
func (g *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
