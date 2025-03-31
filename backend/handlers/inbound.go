package handlers

import (
	"net/http"

	"coldchain/backend/models"

	"github.com/gin-gonic/gin"
)

// HandleInbound 处理入库请求
func HandleInbound(c *gin.Context) {
	var inboundData models.InboundRequest

	if err := c.ShouldBindJSON(&inboundData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求数据无效",
			"details": err.Error(),
		})
		return
	}

	// 验证温度范围
	if inboundData.Temperature < -50 || inboundData.Temperature > 50 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "温度超出有效范围",
			"details": "温度必须在 -50°C 到 50°C 之间",
		})
		return
	}

	// TODO: 这里后续会添加数据库操作
	// 现在先返回成功响应
	response := models.InboundResponse{
		Message: "入库成功",
		Data:    inboundData,
	}

	c.JSON(http.StatusOK, response)
}
