package services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go.uber.org/zap"
	"time"
	"wild/internal/models"
	"wild/internal/pkg/auth"
	"wild/internal/pkg/snowflake"
	"wild/internal/repository/mysql"
)

// 存放业务逻辑的代码
const secret = "wild"

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	h.Sum([]byte(oPassword))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		zap.L().Error("CheckUserExist error", zap.Error(err))
		return err
	}
	// 2. 生成UID
	userID := snowflake.GenID()
	// 构造一个 User 实例
	user := models.User{
		Userid:   userID,
		Username: p.Username,
		// 密码加密🔐
		Password:    encryptPassword(p.Password),
		LastLoginAt: time.Now(),
	}
	fmt.Println(time.Now())
	fmt.Println(user)
	// 3. 保存进数据库
	err = mysql.InsertUser(&user)
	return
}

func Login(p *models.ParamLogin) (user *models.User, token string, err error) {
	user = &models.User{
		Username: p.Username,
		Password: encryptPassword(p.Password),
	}
	// 传递的是指针，就能拿到user.UserID
	if err = mysql.Login(user); err != nil {
		return
	}
	// 生成JWT
	token = auth.GenerateToken(user.Userid, user.Username)

	return
}

func GetUsers() (users []*models.User, err error) {
	users, err = mysql.GetUsers()
	return
}
