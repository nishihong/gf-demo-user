package controller

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/v1"

	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
)

var User = cUser{}

type cUser struct{}

func (c *cUser) SignUp(ctx context.Context, req *v1.UserSignUpReq) (res *v1.UserSignUpRes, err error) {
	err = service.User().Create(ctx, model.UserCreateInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	return
}
func (c *cUser) SignIn(ctx context.Context, req *v1.UserSignInReq) (res *v1.UserSignInRes, err error) {
	err = service.User().SignIn(ctx, model.UserSignInInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	return
}

func (c *cUser) IsSignedIn(ctx context.Context, req *v1.UserIsSignedInReq) (res *v1.UserIsSignedInRes, err error) {
	res = &v1.UserIsSignedInRes{
		OK: service.User().IsSignedIn(ctx),
	}
	return
}

func (c *cUser) SignOut(ctx context.Context, req *v1.UserSignOutReq) (res *v1.UserSignOutRes, err error) {
	err = service.User().SignOut(ctx)
	return
}

func (c *cUser) CheckPassport(ctx context.Context, req *v1.UserCheckPassportReq) (res *v1.UserCheckPassportRes, err error) {
	available, err := service.User().IsPassportAvailable(ctx, req.Passport)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Passport "%s" is already token by others`, req.Passport)
	}
	return
}

func (c *cUser) CheckNickName(ctx context.Context, req *v1.UserCheckNickNameReq) (res *v1.UserCheckNickNameRes, err error) {
	available, err := service.User().IsNicknameAvailable(ctx, req.Nickname)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Nickname "%s" is already token by others`, req.Nickname)
	}
	return
}

func (c *cUser) Profile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
	res = &v1.UserProfileRes{
		User: service.User().GetProfile(ctx),
	}
	return
}

//func (c *cUser) UserList(ctx context.Context, req *v1.UserListReq) (res []*entity.User, err error) {
//	res = service.User().GetList(ctx)
//	return
//}

func (c *cUser) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	res = &v1.UserListRes{
		List: service.User().GetList(ctx),
	}
	return
}

func (c *cUser) UserDelete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserDeleteRes, err error) {
	tes, err := service.User().UserDelete(ctx, req.Passport)
	if err != nil {
		return nil, err
	}
	//???有问题
	if !tes {
		return nil, gerror.Newf(`Passport "%s" is no exist`, req.Passport)
	}
	return
}

func (c *cUser) UserUpdate(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error) {
	tes, err := service.User().UserUpdate(ctx, model.UserUpdateInput{
		Id:       req.Id,
		Passport: req.Passport,
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	//???有问题
	if !tes {
		return nil, gerror.Newf(`Id "%s" is no exist`, req.Passport)
	}
	return
}
