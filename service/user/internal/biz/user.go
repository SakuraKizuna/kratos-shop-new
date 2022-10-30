package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// 定义返回数据结构体
type User struct {
	ID       int64
	Mobile   string
	Password string
	NickName string
	Birthday int64
	Gender   string
	Role     int
}

// 定义返回数据结构体
type UserInfo struct {
	ID       int64
	Mobile   string
	Password string
	NickName string
	Birthday int64
	Gender   string
	Role     int
}

type GetUserReqInfo struct {
	Id int64
}

type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	GetUserInfo(context.Context, *GetUserReqInfo) (*UserInfo, error)
	ListUser(ctx context.Context, pageNum, pageSize int) ([]*User, int, error)
	//UserByMobile(ctx context.Context, mobile string) (*User, error)
	//GetUserById(ctx context.Context, id int64) (*User, error)
	//UpdateUser(context.Context, *User) (bool, error)
	CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}

func (uc *UserUsecase) GetUserInfo(ctx context.Context, ui *GetUserReqInfo) (*UserInfo, error) {
	return uc.repo.GetUserInfo(ctx, ui)
}

func (uc *UserUsecase) List(ctx context.Context, pageNum, pageSize int) ([]*User, int, error) {
	return uc.repo.ListUser(ctx, pageNum, pageSize)
}

//func (uc *UserUsecase) UserByMobile(ctx context.Context, mobile string) (*User, error) {
//	return uc.repo.UserByMobile(ctx, mobile)
//}

//func (uc *UserUsecase) UpdateUser(ctx context.Context, user *User) (bool, error) {
//	return uc.repo.UpdateUser(ctx, user)
//}

func (uc *UserUsecase) CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error) {
	return uc.repo.CheckPassword(ctx, password, encryptedPassword)
}

//func (uc *UserUsecase) UserById(ctx context.Context, id int64) (*User, error) {
//	return uc.repo.GetUserById(ctx, id)
//}
