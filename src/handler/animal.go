package handler

import (
	"demogo/src/model"
	"demogo/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	var param model.CreateParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"message":"参数不正确",
		})
		return
	}

	if err := service.DefaultAnimalServer.Create(param); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success":false,
			"message":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":true,
		"message":"创建成功",
	})
}