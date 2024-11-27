package dao

import (
	"ARPostcard_server/biz/consts"
	"ARPostcard_server/biz/infra"
	"ARPostcard_server/biz/model"
	"context"
	"github.com/RanFeng/ierror"
	"github.com/RanFeng/ilog"
)

func QueryUserByUserName(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{}
	err := infra.MysqlDB.WithContext(ctx).Debug().
		Where("user_name = ? OR email = ? OR tel = ? ", username).
		First(user).Error
	if err != nil {
		ilog.EventError(ctx, err, "dao_get_user_by_username_error", "username", username)
		return nil, ierror.NewIError(consts.DBError, err.Error())
	}

	return user, nil
}

func CreateUser(ctx context.Context, username, password string) error {
	user := &model.User{
		UserName: username,
		Password: password,
		Role:     consts.PermissionAdmin,
	}
	err := infra.MysqlDB.WithContext(ctx).Debug().
		Create(user).Error
	if err != nil {
		ilog.EventError(ctx, err, "dao_create_user_error", "username", username)
		return ierror.NewIError(consts.DBError, err.Error())
	}
	return nil
}
