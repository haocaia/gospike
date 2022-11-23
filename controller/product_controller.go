package controller

import (
	"github.com/gin-gonic/gin"
	"gospike/datamodels"
	"gospike/repository"
	"net/http"
	"strconv"
)

var productRepository repository.IProduct

func init() {
	productRepository = new(repository.ProductImpl)
}

// ProductInsertHandler 返回插入记录的id
func ProductInsertHandler(c *gin.Context) {
	var product datamodels.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lastInsertId, err := productRepository.Insert(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"result": lastInsertId,
	})
}

// ProductSelectAllHandler result返回查询product数组, len返回数组长度
func ProductSelectAllHandler(c *gin.Context) {
	var productList []*datamodels.Product
	productList, err := productRepository.SelectAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"result": productList,
		"len":    len(productList),
	})
}

// ProductSelectByIdHandler result返回对应id的Product
func ProductSelectByIdHandler(c *gin.Context) {
	var productId int
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	var product *datamodels.Product
	product, err = productRepository.SelectByKey(int64(productId))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"result": product,
	})
}

// ProductUpdateHandler result返回update成功(true)或者失败(false)
func ProductUpdateHandler(c *gin.Context) {
	var product datamodels.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := productRepository.Update(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"result": true,
	})
}

// ProductDeleteHandler result返回删除是否成功
func ProductDeleteHandler(c *gin.Context) {
	var productID int
	productID, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	success := productRepository.Delete(int64(productID))
	c.JSON(http.StatusAccepted, gin.H{
		"result": success,
	})
}
