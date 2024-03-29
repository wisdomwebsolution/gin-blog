package main

import (
	"gin-blog/controllers/auth"
	"gin-blog/controllers/post"
	"gin-blog/controllers/tag"
	"gin-blog/controllers/user"
	"gin-blog/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	docs "gin-blog/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": "This is only for test"})
}

// @title		Gin-Blog API
// @version 	1.0
// @description	This is a backend for simple blog site built with Gin-Gonic
// @BasePath 	/api

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	apiRouter := router.Group("/api")

	authorized := apiRouter.Group("/")

	// apiRouter.GET("/test", Test)
	apiRouter.GET("/users", user.Index)
	apiRouter.POST("/register", auth.Register)
	apiRouter.POST("/login", auth.Login)
	apiRouter.GET("/tags", tag.Index)

	authorized.Use(middlewares.AuthenticationMiddleware())
	authorized.GET("/test", Test)
	authorized.GET("/me", auth.Currentuser)
	authorized.POST("/posts", post.Create)
	authorized.PUT("/posts/:id", post.Update)
	authorized.DELETE("/posts/:id", post.Delete)

	posts := apiRouter.Group("/posts")
	posts.GET("/", post.Index)
	posts.GET("/:id", post.Show)

	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run("localhost:8080")
}
