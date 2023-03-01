// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"zero-start/mall/user/internal/logic"
	"zero-start/mall/user/internal/svc"
	"zero-start/mall/user/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *user.IDRequest) (*user.UserResponse, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}