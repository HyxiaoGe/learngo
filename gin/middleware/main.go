package main

import "github.com/gin-gonic/gin"

func main() {

	// 新建一个没有任何默认中间件的路由
	router := gin.New()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置成 release
	router.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic，如果有 panic 的话，会写入 500
	router.Use(gin.Recovery())

	// 你甚至为每个路由添加任意数量的中间件
	//router.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// 认证路由组
	// authorized := router.Group("/", AuthRequired())
	// 和使用以下两行代码的效果完全一样
	//authorized := router.Group("/")
	// 路由组中间件
	// AuthRequired() 中间件
	//authorized.Use(AuthRequired())
	//{
	//	authorized.POST("/login", loginEndpoint)
	//	authorized.POST("/submit", submitEndpoint)
	//	authorized.POST("/read", readEndpoint)
	//
	//	// 嵌套路由组
	//	testing := authorized.Group("testing")
	//	testing.GET("/analytics", analyticsEndpoint)
	//}

	router.Run()
}
