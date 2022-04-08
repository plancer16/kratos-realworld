package service

import (
	"context"
	"errors"

	v1 "realworld_02/api/realworld/v1"
	"realworld_02/internal/biz"
)

// GreeterService is a greeter service.
type RealworldService struct {
	v1.UnimplementedRealworldServer

	uc *biz.UserUsecase
}

func NewRealworldService(uc *biz.UserUsecase) *RealworldService {
	return &RealworldService{uc: uc}
}

func (s *RealworldService) Register(ctx context.Context, in *v1.RegisterRequest) (*v1.UserReply, error) {
	u, err := s.uc.Register(ctx, in.User.Email, in.User.Username, in.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email:    u.Email,
			Token:    u.Token,
			Username: u.Username,
			Bio:      u.Bio,
			Image:    u.Image,
		},
	}, nil
}

func (s *RealworldService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.UserReply, error) {
	if in.User.Email == "" {
		return nil, errors.New("email cannot be empty")
	}
	u, err := s.uc.Login(ctx, in.User.Email, in.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email:    u.Email,
			Token:    u.Token,
			Username: u.Username,
			Bio:      u.Bio,
			Image:    u.Image,
		},
	}, nil
}
