package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"

	"wild/configs"
)

func InitLogrusLogger() *logrus.Logger {
	logger := logrus.New()

	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logCfg := configs.Conf.LogConfig
	logFile := &lumberjack.Logger{
		Filename: logCfg.Filename,
		MaxSize:  logCfg.MaxSize,
		Compress: true,
	}

	// multi writer, both file and stdout
	writers := []io.Writer{logFile, os.Stdout}
	logger.SetOutput(io.MultiWriter(writers...))

	return logger
}
