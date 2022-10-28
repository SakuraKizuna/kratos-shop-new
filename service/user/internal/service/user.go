package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "user/api/user/v1"
	"user/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

// CreateUser create a user
func (u *UserService) CreateUser(ctx context.Context, req *v1.CreateUserInfo) (*v1.UserInfoResponse, error) {
	user, err := u.uc.Create(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
		NickName: req.NickName,
	})
	if err != nil {
		return nil, err
	}

	userInfoRsp := v1.UserInfoResponse{
		Id:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
		Birthday: user.Birthday,
	}

	return &userInfoRsp, nil
}

func (u *UserService) GetUserInfo(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	fmt.Println(req.Id)
	if req.Id == 0 {
		return nil, errors.New("bind err")
	}
	user, err := u.uc.GetUserInfo(ctx, &biz.GetUserReqInfo{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	userInfoRsp := v1.GetUserResponse{
		Id:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
		Birthday: user.Birthday,
	}
	return &userInfoRsp, nil
}

func (u *UserService) GetUserInfoRPC(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	user, err := u.uc.GetUserInfo(ctx, &biz.GetUserReqInfo{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	userInfoRsp := v1.GetUserResponse{
		Id:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
		Birthday: user.Birthday,
	}
	return &userInfoRsp, nil
}
