package main

import (
	"coldchain/backend/middleware"
	"coldchain/backend/routes"
)

func main() {
	r := routes.SetupRouter()

	// 添加中间件
	r.Use(middleware.CORSMiddleware())

	// 启动服务器
	r.Run(":8080")
}
