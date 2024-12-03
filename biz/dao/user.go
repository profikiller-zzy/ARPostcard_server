package dao

import (
	"ARPostcard_server/biz/consts"
	"ARPostcard_server/biz/infra"
	"ARPostcard_server/biz/model"
	"context"
	"errors"
	"github.com/RanFeng/ierror"
	"github.com/RanFeng/ilog"
	"gorm.io/gorm"
)

func QueryUserByUserName(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{}
	err := infra.MysqlDB.WithContext(ctx).Debug().
		Where("user_name = ?", username).
		First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		ilog.EventError(ctx, err, "dao_get_user_by_username_error", "username", username)
		return nil, ierror.NewIError(consts.DBError, err.Error())
	}

	return user, nil
}

// QueryUserByID 通过ID查询用户
func QueryUserByID(ctx context.Context, id int) (*model.User, error) {
	user := &model.User{}
	err := infra.MysqlDB.WithContext(ctx).Debug().
		Where("id = ?", id).
		First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		ilog.EventError(ctx, err, "dao_get_user_by_id_error", "id", id)
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

func PGetUser(ctx context.Context, page, pageSize int) ([]*model.UserReq, int64, error) {
	var users []*model.User
	var total int64
	err := infra.MysqlDB.WithContext(ctx).Debug().
		Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&users).Error
	if err != nil {
		ilog.EventError(ctx, err, "dao_user_list_error")
		return nil, 0, ierror.NewIError(consts.DBError, err.Error())
	}
	err = infra.MysqlDB.WithContext(ctx).Debug().
		Model(&model.User{}).
		Count(&total).Error
	if err != nil {
		ilog.EventError(ctx, err, "dao_user_list_error")
		return nil, 0, ierror.NewIError(consts.DBError, err.Error())
	}
	// 将User转化为UserReq
	var userReqs []*model.UserReq
	for _, user := range users {
		userReqs = append(userReqs, &model.UserReq{
			MODEL: model.MODEL{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
			UserName: user.UserName,
			Role:     user.Role,
		})
	}

	return userReqs, total, nil
}

// UpdateUser 更新用户信息
func UpdateUser(ctx context.Context, user *model.User) error {
	err := infra.MysqlDB.WithContext(ctx).Debug().
		Save(user).Error
	if err != nil {
		ilog.EventError(ctx, err, "dao_update_user_error", "user", user)
		return ierror.NewIError(consts.DBError, err.Error())
	}
	return nil
}

// DeleteUsers 删除用户
func DeleteUsers(ctx context.Context, idList []int) error {
	err := infra.MysqlDB.WithContext(ctx).Debug().
		Where("id in (?)", idList).
		Delete(&model.User{}).Error
	if err != nil {
		ilog.EventError(ctx, err, "dao_delete_user_error", "idList", idList)
		return ierror.NewIError(consts.DBError, err.Error())
	}
	return nil
}
