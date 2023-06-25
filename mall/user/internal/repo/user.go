package repo

import (
	"context"
	"zero-start/mall/user/internal/model"
)

// 这里主要写接口

type UserRepo interface {
	SaveUser(context.Context, *model.User) error
}
