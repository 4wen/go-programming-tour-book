package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("配置文件初始化失败: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("日志组件初始化失败: %v", err)
	}
}

// @title 博客
// @version 1.0
// @description 学习go语言
func main() {
	// gin运行模式
	gin.SetMode(global.ServerSetting.RunMode)

	// 初始化路由对象
	router := routers.NewRouter()
	fmt.Println(global.ServerSetting.ReadTimeout)

	// 自定义http.Server MaxHeaderBytes-->请求头最大字节数
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	global.Logger.Infof("%s: go-programming-tour-book/%s", "4wen", "blog-service")

	// 开启监听
	_ = s.ListenAndServe()
}

// 配置文件初始化
func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

// 数据库连接初始化
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

// 初始化日志
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
