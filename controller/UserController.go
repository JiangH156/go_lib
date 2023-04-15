package controller

import (
	"Go_lib/common"
	"Go_lib/model"
	"Go_lib/response"
	"Go_lib/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	DB *gorm.DB
}

// Register
// @Description 用户注册
// @Author John 2023-04-14 15:22:14
// @Param ctx
func (u *UserController) Register(ctx *gin.Context) {
	//数据接收
	userName := ctx.PostForm("userName")
	email := ctx.PostForm("email")
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")
	//数据验证
	// 用户名
	length := len([]rune(userName))
	if length == 0 || length > 10 {
		fmt.Println("用户名输出格式错误，长度为1~10位")
		response.Fail(ctx, nil, "用户名输出格式错误，长度为1~10位")
		return
	}
	// 密码
	length = len([]rune(password))
	if length < 4 || length > 20 {
		fmt.Println("密码输出格式错误，长度为4~20位")
		response.Fail(ctx, nil, "密码输出格式错误，长度为4~20位")
		return
	}
	// 邮箱
	if err := utils.EmailRegexp(email); err != nil {
		fmt.Println(err.Error())
		response.Fail(ctx, nil, err.Error())
		return
	}
	// 手机号
	if err := utils.PhoneRegexp(phone); err != nil {
		fmt.Println(err.Error())
		response.Fail(ctx, nil, err.Error())
		return
	}
	// 判断该手机号是否存在
	var reader model.Reader
	u.DB.Where("phone = ?", phone).First(&reader)
	// 不存在该用户
	if reader.ReaderId != "" {
		fmt.Println("用户已存在")
		response.Response(ctx, http.StatusConflict, 409, nil, "用户已存在")
		return
	}
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("加密失败")
		response.Response(ctx, http.StatusInternalServerError, 500, nil, err.Error())
		return
	}
	reader = model.Reader{
		ReaderName: userName,
		Password:   string(encryptPassword),
		Email:      email,
		Phone:      phone,
	}
	tx := u.DB.Begin()
	if err := tx.Create(&reader).Error; err != nil {
		tx.Rollback() // 如果操作出错，则回滚事务
		response.Response(ctx, http.StatusInternalServerError, 500, nil, err.Error())
		return
	}
	// 如果所有操作都成功，则提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback() // 如果提交出错，则回滚事务
		response.Response(ctx, http.StatusInternalServerError, 500, nil, err.Error())
		return
	}
	response.Success(ctx, nil, "注册成功")
}

// Login
// @Description 用户登录
// @Author John 2023-04-14 15:22:25
// @Param ctx
func (u *UserController) Login(ctx *gin.Context) {
	isAdmin := ctx.DefaultPostForm("isAdmin", "false")

	//判断是否管理员
	if isAdmin == "true" {
		u.loginAsAdmin(ctx)
	} else {
		u.loginAsReader(ctx)
	}

}

// loginAsAdmin
// @Description 管理员登陆
// @Author John 2023-04-15 10:57:33
// @Param ctx
func (u *UserController) loginAsAdmin(ctx *gin.Context) {

}

// loginAsReader
// @Description 读者登录
// @Author John 2023-04-15 10:59:19
// @Param ctx
func (u *UserController) loginAsReader(ctx *gin.Context) {

}

func NewUserController() UserController {
	return UserController{DB: common.GetDB()}
}
