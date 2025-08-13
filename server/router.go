package main

import (
	"Blog_Backend/server/api"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

const (
	ErrorCodeAuthFailed          = "AUTH_FAILED"
	ErrorCodeInternalServerError = "INTERNAL_SERVER_ERROR"
)

type ErrorResponse struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Authorization 中间件
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 处理中间件内部错误
				HandleError(c, ErrorCodeInternalServerError, "鉴权失败", err)
			}
		}()

		// 模拟鉴权逻辑，这里假设鉴权失败
		if c.Query("token") == "" {
			panic("鉴权失败") // 模拟错误
		}

		c.Next() // 继续执行后续 handler
	}
}

func HandleError(c *gin.Context, code, msg string, err interface{}) {
	// 记录日志，可以根据 err 的类型做不同的处理
	fmt.Println("Error:", err)
	c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func RouterInit(router *gin.Engine) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	router.Use(gin.Recovery())
	router.Use(Authorization())
	//初始化日志
	router = loggerConfig(router)
	//绑定验证器
	//错误处理中间件
	//路由组
	router = routerGroup(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	<-ctx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func routerGroup(router *gin.Engine) *gin.Engine {
	router.GET("/test", func(c *gin.Context) {})
	//路由组
	//V1/article Get Delete Post Update
	api.InitArticlesRouter(router)
	api.InitUsersRouter(router)
	return router
}

func loggerConfig(router *gin.Engine) *gin.Engine {
	// 默认 gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	return router
}
