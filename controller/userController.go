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
		fmt.Println("请输入正确的用户名")
		response.Fail(ctx, gin.H{
			"status": 400,
			"msg":    "请输入正确的用户名",
		})
		return
	}
	// 密码
	length = len([]rune(password))
	if length < 4 || length > 20 {
		fmt.Println("请输入正确的手机号")
		response.Fail(ctx, gin.H{
			"status": 400,
			"msg":    "请输入正确的手机号",
		})
		return
	}
	// 邮箱
	if err := utils.EmailRegexp(email); err != nil {
		fmt.Println(err.Error())
		response.Fail(ctx, gin.H{
			"status": 400,
			"msg":    err.Error(),
		})
		return
	}
	// 手机号
	if err := utils.PhoneRegexp(phone); err != nil {
		fmt.Println(err.Error())
		response.Fail(ctx, gin.H{
			"status": 400,
			"msg":    err.Error(),
		})
		return
	}
	// 判断该手机号是否存在
	var reader model.Reader
	u.DB.Where("phone = ?", phone).First(&reader)
	// 不存在该用户
	if reader.ReaderId != "" {
		fmt.Println("用户已存在")
		response.Response(ctx, http.StatusConflict, gin.H{
			"status": 409,
			"msg":    "用户已存在",
		})
		return
	}
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("注册失败")
		response.Response(ctx, http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
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
		response.Response(ctx, http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
	// 如果所有操作都成功，则提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback() // 如果提交出错，则回滚事务
		response.Response(ctx, http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
	response.Success(ctx, gin.H{
		"msg":    "注册成功",
		"status": 200,
	})
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
	// 数据接收
	// 作为管理时，前端接收的数据phone为名称，不需要手机号验证
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	// 数据验证
	if len([]rune(phone)) < 1 || len([]rune(phone)) > 20 {
		fmt.Println("请输入正确的用户名")
		response.Response(ctx, http.StatusBadRequest, gin.H{
			"status": 400,
			"msg":    "请输入正确的用户名",
		})
		return
	}

	// 验证密码
	loginAdmin, exist := getAdmin(u.DB, phone)
	if !exist {
		fmt.Println("请输入正确的管理员账号")
		response.Response(ctx, http.StatusConflict, gin.H{
			"status": 409,
			"msg":    "请输入正确的管理员账号",
		})
		return
	}
	// 密码校验失败
	if loginAdmin.Password != password {
		fmt.Println("请输入正确的密码")
		response.Response(ctx, http.StatusBadRequest, gin.H{
			"status": 400,
			"msg":    "请输入正确的密码",
		})
		return
	}
	response.Success(ctx, gin.H{
		"msg":      "管理员登录成功",
		"status":   200,
		"userName": loginAdmin.Phone,
		"isAdmin":  true,
	})
}

// loginAsReader
// @Description 读者登录
// @Author John 2023-04-15 10:59:19
// @Param ctx
func (u *UserController) loginAsReader(ctx *gin.Context) {
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	// 手机号匹配
	if err := utils.PhoneRegexp(phone); err != nil {
		fmt.Println("请输入正确的手机号")
		response.Response(ctx, http.StatusBadRequest, gin.H{
			"status": 400,
			"msg":    "请输入正确的手机号",
		})
		return
	}

	loginReader, exist := getReader(u.DB, phone)
	if !exist {
		fmt.Println("账号密码错误或该用户未注册")
		response.Response(ctx, http.StatusBadRequest, gin.H{
			"status": 400,
			"msg":    "账号密码错误或该用户未注册",
		})
		return
	}
	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(loginReader.Password), []byte(password)); err != nil {
		fmt.Println("账号密码错误或该用户未注册")
		response.Response(ctx, http.StatusBadRequest, gin.H{
			"status": 400,
			"msg":    "账号密码错误或该用户未注册",
		})
		return
	}
	response.Success(ctx, gin.H{
		"msg":         "读者登录成功",
		"status":      200,
		"readerId":    loginReader.ReaderId,
		"readerName":  loginReader.ReaderName,
		"readerPhone": loginReader.Phone,
		"borrowTimes": loginReader.BorrowTimes,
		"ovdTimes":    loginReader.OvdTimes,
		"email":       loginReader.Email,
		"isAdmin":     false,
	})

}

// getAdmin
// @Description 通过手机号获取管理员
// @Author John 2023-04-15 11:44:57
// @Param db
// @Param phone 手机号
// @Return model.Admin
// @Return bool 是否存在手机号对应的管理员
func getAdmin(db *gorm.DB, phone string) (model.Admin, bool) {
	var admin = model.Admin{}
	db.Where("phone = ?", phone).First(&admin)
	return admin, admin.Id != 0
}

func getReader(db *gorm.DB, phone string) (model.Reader, bool) {
	var reader = model.Reader{}
	db.Where("phone = ?", phone).First(&reader)
	return reader, reader.ReaderId != ""
}

// NewUserController
// @Description UserController的构造器
// @Author John 2023-04-16 15:22:31
// @Return UserController
func NewUserController() UserController {
	return UserController{DB: common.GetDB()}
}
