package main

import (
	"github.com/gin-gonic/gin"
	"gospike/datamodels"
	"gospike/repository"
	"gospike/utils"
	"log"
	"net/http"
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
		productGroup.POST("/insert", func(c *gin.Context) {
			var product datamodels.Product
			if err := c.BindJSON(&product); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			impl := &repository.ProductImpl{}
			lastInsertId, err := impl.Insert(&product)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			log.Println(lastInsertId)
		})

		productGroup.GET("/selectAll", func(c *gin.Context) {
			var productList []*datamodels.Product
			impl := &repository.ProductImpl{}
			productList, err := impl.SelectAll()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			log.Println(len(productList))
			c.JSON(http.StatusAccepted, gin.H{"result len": len(productList)})
		})
	}
}
