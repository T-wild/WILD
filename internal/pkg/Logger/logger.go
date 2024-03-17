package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

var (
	ZapLogger *zap.Logger
	ZapSugar  *zap.SugaredLogger

	LogrusLogger *logrus.Logger
)

func InitLogger() error {
	// LogrusLogger = InitLogrusLogger()

	err := InitZapLogger()
	if err != nil {
		fmt.Println("Init zaplogger faild %v\n", err)
		return err
	}
	// ZapSugar = ZapLogger.Sugar()
	return nil
}
