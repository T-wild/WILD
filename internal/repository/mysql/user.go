package mysql

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
	"wild/internal/models"
	"wild/internal/pkg/mysql"
)

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
)

// CheckUserExist 检查用户是否存在
func CheckUserExist(username string) (err error) {
	var count int64
	result := mysql.DB.Table("users").Where("username = ?", username).Count(&count)
	// 检查查询结果和错误
	if result.Error != nil {
		zap.L().Error("CheckUserExist error", zap.Error(err))
		return result.Error
	} else {
		if count > 0 {
			return ErrorUserExist
		}
	}
	return
}

// InsertUser 向数据库插入一条用户数据
func InsertUser(user *models.User) (err error) {
	// 执行SQL语句入库
	result := mysql.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}

// Login 登陆
func Login(user *models.User) (err error) {
	oPassword := user.Password // 用户登陆的密码
	result := mysql.DB.Where("username = ?", user.Username).First(&user)

	// 检查查询结果和错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ErrorUserNotExist
	} else if result.Error != nil {
		fmt.Println("Error occurred while fetching user:", result.Error)
		return result.Error
	}
	// 判断密码是否正确
	if oPassword != user.Password {
		return ErrorInvalidPassword
	}

	result = mysql.DB.Model(&user).Update("LastLoginAt", time.Now())
	// 检查查询结果和错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		zap.L().Error("update user last_login_at failed", zap.Error(result.Error))
		return ErrorUserNotExist
	} else if result.Error != nil {
		zap.L().Error("update user last_login_at failed", zap.Error(result.Error))
		fmt.Println("Error occurred while fetching user:", result.Error)
		return result.Error
	}
	return
}

// GetUserById 根据用户id 获取用户名称
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	result := mysql.DB.Where("userid = ?", user.Userid).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrorUserNotExist
	} else if result.Error != nil {
		return nil, result.Error
	}
	return
}

func GetUsers() (users []*models.User, err error) {
	result := mysql.DB.Find(&users)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrorUserNotExist
	} else if result.Error != nil {
		return nil, result.Error
	}
	return
}
