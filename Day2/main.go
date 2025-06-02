package main

import (
	"E_commerce_System/config"
	"E_commerce_System/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库连接
	config.ConnectDB()

	r := gin.Default()
	err := r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.0/24"})
	if err != nil {
		panic(err)
	}

	r.GET("/", func(c *gin.Context) {
		clientIP := c.ClientIP()
		c.JSON(200, gin.H{
			"message":   "Welcome!",
			"client_ip": clientIP,
		})
	})

	// 路由
	r.GET("/products", handlers.GetProducts)
	r.POST("/products", handlers.AddProduct)
	r.DELETE("/products", handlers.DelProduct)
	r.PUT("/products", handlers.UpdateProduct)

	// 启动服务
	_ = r.Run(":8090")
}
