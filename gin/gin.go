package gin

import (
	"github.com/gin-gonic/gin"
	"go_code/Study/Collection/gin/middleware"
	"go_code/Study/Collection/gin/routers/hello"
	"go_code/Study/Collection/gin/routers/user"
	"log"
	"net/http"
	"sync"
	"time"
)

// debug端口服务
func debugRouter() http.Handler {
	// 创建一个新的路由引擎
	r := gin.Default()

	// 使用全局中间件
	r.Use(middleware.StatCost())

	// 加载多个模块的路由配置
	var debugOptions optionsHolder
	debugOptions.Include(user.Routers)

	// 初始化路由
	debugOptions.Init(r)

	return r
}

// app端口服务
func appRouter() http.Handler {
	// 创建一个新的路由引擎
	r := gin.Default()

	// 使用全局中间件
	r.Use(middleware.Cors())

	// 加载多个模块的路由配置
	var appOptions optionsHolder
	appOptions.Include(user.Routers, hello.Routers)

	// 初始化路由
	appOptions.Init(r)

	return r
}

// Start 启动服务
func Start() {
	// 借助goroutine启动两个服务
	// 定义一个等待组，用于等待goroutine完成
	var wg sync.WaitGroup

	// app端口服务
	appServer := http.Server{
		Addr:         ":8080",
		Handler:      appRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := appServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("App Server ListenAndServe: %v", err)
		}
	}()

	// debug端口服务
	debugServer := http.Server{
		Addr:         ":8081",
		Handler:      debugRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := debugServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Debug Server ListenAndServe: %v", err)
		}
	}()

	// 等待两个goroutine完成
	wg.Wait()
}
