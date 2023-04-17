package main

import (
	"todo-server/initializers"
	"todo-server/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	// r.Use(cors.Default())
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000/todos"}
	// config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	// config.AllowHeaders = []string{"Content-Type"}

	// r.Use(cors.New(config))
	r.POST("/todos", models.PostTodo)
	r.GET("/todos", models.GetTodo)
	r.POST("/todos/delete", models.DelTodo)
	// r := gin.Default()
	r.Use(CORSMiddleware())
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//Access-Control-Allow-Origin: http://www.example.com
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
