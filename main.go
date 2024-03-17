package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wild/configs"
	"wild/internal/pkg/Logger"
	"wild/internal/pkg/mysql"
	"wild/internal/pkg/redis"
	"wild/internal/pkg/snowflake"
	"wild/internal/pkg/utils"
	"wild/internal/routers"
)

func init() {
	// config
	config := "configs/config.yaml"
	err := configs.Init(config)
	if err != nil {
		fmt.Printf("init config failed, err:%v\n", err)
	}

	// log
	err = logger.InitLogger()
	if err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
	}

	// mysql
	err = mysql.InitMysql()
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
	}

	// redis
	err = redis.InitRedis()
	if err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
	}

	// snowflake
	err = snowflake.InitSnowFlake()
	if err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
	}

	// translation
	// 初始化 gin 框架内置的校验器使用的翻译器
	if err = utils.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err: %v \n", err)
	}
}

func main() {

	defer func() {
		e := redis.CloseRedis()
		if e != nil {
			fmt.Printf("close redis failed, err: %v \n", e)
		}
	}()

	r := routers.Setup()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("server.port")),
		Handler: r,
	}

	go func() {

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}

	log.Println("Server exiting")
}
