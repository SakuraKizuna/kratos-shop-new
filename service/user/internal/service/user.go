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

// GetUserList .
func (u *UserService) GetUserList(ctx context.Context, req *v1.PageInfo) (*v1.UserListResponse, error) {
	list, total, err := u.uc.List(ctx, int(req.Pn), int(req.PSize))
	if err != nil {
		return nil, err
	}
	rsp := &v1.UserListResponse{}
	rsp.Total = int32(total)

	for _, user := range list {
		userInfoRsp := UserResponse(user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}

	return rsp, nil
}

func UserResponse(user *biz.User) v1.UserInfoResponse {
	userInfoRsp := v1.UserInfoResponse{
		Id:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
		Birthday: user.Birthday,
	}
	//if user.Birthday != 0 {
	//	userInfoRsp.Birthday = uint64(user.Birthday.Unix())
	//}
	return userInfoRsp
}

//// GetUserByMobile .
//func (u *UserService) GetUserByMobile(ctx context.Context, req *v1.MobileRequest) (*v1.UserInfoResponse, error) {
//	user, err := u.uc.UserByMobile(ctx, req.Mobile)
//	if err != nil {
//		return nil, err
//	}
//	rsp := UserResponse(user)
//	return &rsp, nil
//}

// CheckPassword .
func (u *UserService) CheckPassword(ctx context.Context, req *v1.PasswordCheckInfo) (*v1.CheckResponse, error) {
	check, err := u.uc.CheckPassword(ctx, req.Password, req.EncryptedPassword)
	if err != nil {
		return nil, err
	}
	return &v1.CheckResponse{Success: check}, nil
}
