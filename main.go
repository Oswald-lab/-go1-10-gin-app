package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type echoRequest struct {
	Message string `json:"message" binding:"required"`
}

func newRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from Gin"})
	})
	r.GET("/items/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"id": c.Param("id")})
	})
	r.POST("/echo", func(c *gin.Context) {
		var req echoRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"echo": req.Message})
	})
	return r
}

func main() {
	if err := newRouter().Run(":8080"); err != nil {
		panic(err)
	}
}
