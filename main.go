package main

import (
	"github.com/gin-gonic/gin"
	"gospike/controller"
	"gospike/utils"
)

func main() {
	r := gin.New()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())

	// 初始化路由
	init_router(r)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

	utils.Close()
}

func init_router(r *gin.Engine) {
	productGroup := r.Group("/product")
	{ //设置/product后跟的路由
		productGroup.POST("/insert", controller.ProductInsertHandler)
		productGroup.GET("/selectAll", controller.ProductSelectAllHandler)
		productGroup.GET("/selectById/:productId", controller.ProductSelectByIdHandler)
		productGroup.POST("/update", controller.ProductUpdateHandler)
		productGroup.POST("/delete/:productId", controller.ProductDeleteHandler)
	}
}
