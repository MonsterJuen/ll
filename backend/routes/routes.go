package routes

import (
	"coldchain/backend/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// API 路由组
	api := r.Group("/api")
	{
		// 入库相关路由
		api.POST("/inbound", handlers.HandleInbound)
	}

	return r
}
