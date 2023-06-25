package dao

import (
	"context"
	"zero-start/mall/user/internal/model"
)

// 这里主要是实现

type userDao struct{}

func NewUserDao() *userDao {
	return &userDao{}
}

func (u *userDao) SaveUser(ctx context.Context, user *model.User) error {
	return nil
}
