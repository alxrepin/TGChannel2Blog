package app

import (
	"app/internal/bootstrap/config"
	"app/internal/context/application/usecase/channel"
	"app/internal/context/application/usecase/posts"
	"app/internal/context/infrastructure/postgres"
	v1 "app/internal/context/presentation/http/api/v1"
	"context"
	"errors"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	db     *postgres.Client
	server *http.Server
}

func New(ctx context.Context, cfg *config.Config) *App {
	db := postgres.MustNewClient(ctx, cfg.Database.URL)

	// Init repo
	postRepo := postgres.NewPostRepository(db.Pool)
	channelRepo := postgres.NewChannelRepository(db.Pool)

	// Init usecase
	listPostsUseCase := posts.NewListPostsUseCase(postRepo)
	getChannelUseCase := channel.NewGetChannelUseCase(channelRepo)

	// Init controller
	postsController := v1.NewPostsController(listPostsUseCase)
	channelController := v1.NewChannelController(getChannelUseCase)

	// Init router
	router := gin.Default()
	router.Use(cors.Default())

	// Init api
	api := router.Group("/api/v1")
	api.GET("/posts", postsController.ListPosts)
	api.GET("/channel", channelController.GetChannel)

	// Init docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Init server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return &App{
		db:     db,
		server: server,
	}
}

func (app *App) Run() error {
	if err := app.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (app *App) Stop(ctx context.Context) error {
	if err := app.server.Shutdown(ctx); err != nil {
		return err
	}

	app.db.Close()

	return nil
}
