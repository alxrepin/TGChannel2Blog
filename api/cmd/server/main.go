// Package main Posts API
//
//	@title			Posts API
//	@version		1.0
//	@description	API for managing posts
//	@host			localhost:8080
//	@BasePath		/api/v1
package main

import (
	"log"

	_ "app/docs"
	"app/internal/application"
	"app/internal/application/usecase/channel"
	"app/internal/application/usecase/posts"
	"app/internal/infrastructure/postgres"
	v1 "app/internal/presentation/http/api/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config, err := application.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	postRepo := postgres.NewPostRepository(config.DB)
	channelRepo := postgres.NewChannelRepository(config.DB)
	listPostsUseCase := posts.NewListPostsUseCase(postRepo)
	getChannelUseCase := channel.NewGetChannelUseCase(channelRepo)

	postsController := v1.NewPostsController(listPostsUseCase)
	channelController := v1.NewChannelController(getChannelUseCase)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")
	api.GET("/posts", postsController.ListPosts)
	api.GET("/channel", channelController.GetChannel)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	} else {
		log.Println("Server starting on :8080")
	}
}
