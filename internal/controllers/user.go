package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"wild/internal/models"
	"wild/internal/pkg/utils"
	"wild/internal/repository/mysql"
	"wild/internal/services"
)

// 参数校验

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	err := c.ShouldBindJSON(&p)

	// 手动对请求参数进行详细的业务规则校验
	if err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Sign up with invalid param", zap.Error(err))
		// 判断 err 是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"status": "fail",
				"msg":    err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "fail",
				"msg":    utils.RemoveTopStruct(errs.Translate(utils.Trans)), // 翻译错误
			})
			return
		}
		return
	}
	zap.L().Info("SignUpHandler")
	// 2. 业务逻辑处理
	err = services.SignUp(p)

	// 3. 返回响应
	if err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	} else {
		ResponseSuccess(c, nil)
	}

}

func LoginHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamLogin)
	err := c.ShouldBindJSON(&p)

	// 手动对请求参数进行详细的业务规则校验
	if err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Sign up with invalid param", zap.Error(err))
		// 判断 err 是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(utils.Trans))) // 翻译错误
		return
	}
	// 2. 业务逻辑处理
	user, token, err := services.Login(p)

	// 3. 返回响应
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, gin.H{
		"user_name": user.Username, // id值大雨1<<53-1 int64类型的最大值是1<<63-1
		"user_id":   fmt.Sprintf("%d", user.Userid),
		"token":     token,
	})
	return

}

func GetUsersHandler(c *gin.Context) {
	users, err := services.GetUsers()
	// 3. 返回响应
	if err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, users)
	return
}
