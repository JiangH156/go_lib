package service

import (
	"Go_lib/common"
	"Go_lib/model"
	"Go_lib/repository"
	"Go_lib/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type UserService struct {
	DB *gorm.DB
}

// Register
// @Description 用户注册
// @Author John 2023-05-05 14:49:12
// @Param reader
// @Return lErr
func (u *UserService) Register(reader model.Reader) (lErr *common.LError) {
	//数据验证
	// 用户名
	length := len([]rune(reader.ReaderName))
	if length == 0 || length > 10 {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "数据验证错误",
			Err:      errors.New("请输入正确的用户名"),
		}
	}
	// 密码
	length = len([]rune(reader.Password))
	if length < 4 || length > 20 {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "数据验证错误",
			Err:      errors.New("请输入正确的密码"),
		}
	}
	// 邮箱
	if err := utils.EmailRegexp(reader.Email); err != nil {
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "数据验证错误",
			Err:      err,
		}
	}
	// 手机号
	if err := utils.PhoneRegexp(reader.Phone); err != nil {
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "数据验证错误",
			Err:      err,
		}
	}
	// 判断该手机号已使用
	userRepository := repository.NewUserRepository()
	readerByPhone, err := userRepository.GetReaderByPhone(reader.Phone)

	// 手机号已使用
	if readerByPhone.ReaderId != "" {
		return &common.LError{
			HttpCode: http.StatusConflict,
			Msg:      "手机号已使用",
			Err:      errors.New("手机号已使用"),
		}
	}
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(reader.Password), bcrypt.DefaultCost)
	if err != nil {
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "注册失败",
			Err:      err,
		}
	}
	// 密码加密，数据库添加数据
	reader.Password = string(encryptPassword)
	tx := u.DB.Begin()
	err = userRepository.CreateReader(tx, reader)
	if err != nil {
		tx.Rollback() // 如果操作出错，则回滚事务
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "注册失败",
			Err:      err,
		}
	}
	tx.Commit()
	return nil
}

// LoginAsAdmin
// @Description 管理员登录
// @Author John 2023-05-05 15:17:06
// @Param admin
// @Return lErr
func (u *UserService) LoginAsAdmin(admin model.Admin) (lErr *common.LError) {
	// 数据验证
	if len([]rune(admin.Phone)) < 1 || len([]rune(admin.Phone)) > 20 {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "请输入正确的用户名",
			Err:      errors.New("请输入正确的用户名"),
		}
	}

	userRepository := repository.NewUserRepository()
	loginAdmin, exist := userRepository.GetAdminByPhone(admin.Phone)

	// 验证密码
	if !exist {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "请输入正确的管理员账号",
			Err:      errors.New("请输入正确的管理员账号"),
		}
	}
	// 密码校验失败
	if loginAdmin.Password != admin.Password {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "请输入正确的密码",
			Err:      errors.New("请输入正确的密码"),
		}
	}
	return nil
}
func NewUserService() UserService {
	return UserService{
		DB: common.GetDB(),
	}
}
