package core

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

var RunningServer map[int]*GinEngine

func init() {
	if nil == RunningServer {
		RunningServer = make(map[int]*GinEngine)
	}
}

func NewGinEngine(engine *gin.Engine, config *ServerConfig) *GinEngine {
	server := RunningServer[config.Port]
	if nil != server {
		fmt.Printf("Port %d has running server\n", config.Port)
		return server
	}
	gin.SetMode(config.Mode)
	server = &GinEngine{
		engine: engine,
		config: config,
	}
	RunningServer[config.Port] = server
	return server
}

func NewGinDefaultEngine(config *ServerConfig) *GinEngine {
	return NewGinEngine(gin.Default(), config)
}

/////

type GinEngine struct {
	config *ServerConfig
	engine *gin.Engine

	quitSignal chan os.Signal
}

func (engine *GinEngine) AddMiddleware(middlewares ...gin.HandlerFunc) {
	engine.engine.Use(middlewares...)
}

func (engine *GinEngine) AddRouter(call func(engine *gin.Engine)) {
	call(engine.engine)
}

func (engine *GinEngine) Start() {
	addr := fmt.Sprintf("%s:%d", engine.config.Addr, engine.config.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: engine.engine,
	}
	go func() {
		fmt.Printf("Server will be start at : http://%s:%d\n", engine.config.Addr, engine.config.Port)
		err := srv.ListenAndServe()
		if nil != err && err != http.ErrServerClosed {
			fmt.Printf("listen err :%s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	engine.quitSignal = make(chan os.Signal)
	signal.Notify(engine.quitSignal, os.Interrupt)
	<-engine.quitSignal
	fmt.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server Shutdown error : %s\n", err)
	}
	log.Println("Server exiting")

}

func (engine *GinEngine) Stop() {
	quit := engine.quitSignal
	if nil == quit {
		fmt.Printf("Server is not running at http://%s:%d\n", engine.config.Addr, engine.config.Port)
		return
	}
	engine.quitSignal <- os.Interrupt
	signal.Stop(engine.quitSignal)
}
