package api

import (
	"context"
	"interview-authserver/config"
	"interview-authserver/internal/api/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	auth.UnimplementedAuthServiceServer
	env *config.EnvConfig
}

func New(env *config.EnvConfig) *server {
	return &server{env: env}
}

func (s *server) Authenticate(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	token, authErr := auth.Authenticate(ctx, req.Username, s.env)
	if authErr != nil {
		return nil, status.Error(codes.Unauthenticated, authErr.Error())
	}

	return &auth.AuthResponse{Token: token}, nil
}
