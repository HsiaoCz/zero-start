package svc

import (
	"zero-start/mall/user/internal/config"
	"zero-start/mall/user/internal/repo"
)

// ServiceContext logic 依赖的资源池

type ServiceContext struct {
	Config config.Config
	Usre   repo.UserRepo
}

func NewServiceContext(c config.Config, u repo.UserRepo) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Usre:   u,
	}
}
