package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.connect_db()

	router.POST("/posts", controlllers.create_post)
	router.GET("/posts", controllers.find_posts)
	router.GET("/posts/:id", controllers.find_post)
	router.PATCH("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	router.Run("localhost:8080")
}
