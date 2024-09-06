// Code generated by hertz generator.

package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	mysql "github.com/huahuoao/hertz_base/biz/dal/mysql/user"
	"github.com/huahuoao/hertz_base/biz/model/app/user"
	"github.com/huahuoao/hertz_base/biz/model/common"
	md5 "github.com/huahuoao/hertz_base/biz/util"
	"gorm.io/gorm"
)

// Register .
// @router /user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserRegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := new(user.UserRegisterResp)
	existUser, err := mysql.GetUserByUsername(req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.String(consts.StatusInternalServerError, "数据库错误: "+err.Error())
		return
	}
	if existUser != nil {
		c.JSON(consts.StatusOK, common.NewResult().Error(301, "用户名已存在"))
		return
	}
	mysql.CreateUser(&common.User{
		UserName: req.Username,
		Password: md5.MD5Hash(req.Password),
	})
	resp.Msg = "注册成功"
	c.JSON(consts.StatusOK, common.NewResult().Success(resp))
}

// ListUsers .
// @router /user/list [GET]
func ListUsers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	users, err := mysql.ListAllUsers()
	if err != nil {
		c.String(consts.StatusInternalServerError, "数据库错误: "+err.Error())
		return
	}
	resp := new(user.UserListResp)
	resp.Users = users
	c.JSON(consts.StatusOK, common.NewResult().Success(resp))
}