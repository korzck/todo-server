package main

import (
	"todo-server/initializers"
	"todo-server/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:1420"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000/todos"}
	// config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	// config.AllowHeaders = []string{"Content-Type"}

	// r.Use(cors.New(config))
	r.POST("/todos", models.PostTodo)
	r.GET("/todos", models.GetTodo)
	r.POST("/todos/delete", models.DelTodo)
	// r := gin.Default()
	r.Run()
}
