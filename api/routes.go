package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"music-files/internal/context"
	"music-files/internal/database/repository/cover_repo"
	"music-files/internal/database/repository/dir_repo"
	"music-files/internal/database/repository/track_repo"
	"music-files/internal/handler/dir_handler"
	"music-files/internal/middleware"
	"music-files/internal/service"
	"music-files/internal/service/cover_service"
	"music-files/internal/service/dir_service"
	"music-files/internal/service/track_service"
)

func SetupRouter(ac *context.AppContext) (r *gin.Engine) {
	log.Debug().Msg("Router setup")
	gin.SetMode(gin.ReleaseMode)

	r = gin.New()
	r.Use(middleware.ZerologMiddleware(log.Logger))

	coverRepo := cover_repo.NewRepository()
	trackRepo := track_repo.NewRepository()
	dirRepo := dir_repo.NewRepository()
	txManager := service.NewTransactionManager(*ac.Db)

	coverService := cover_service.NewService(coverRepo)
	trackService := track_service.NewService(trackRepo)
	dirService := dir_service.NewService(dirRepo, *coverService, *trackService)

	dirHandler := dir_handler.NewHandler(*dirService, txManager)

	api := r.Group("/api")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		dirs := api.Group("dirs")
		{
			dirs.GET("/roots")
			dirs.GET("/:dirId/content", dirHandler.ReadContent)
			dirs.GET("/:dirId/tracks-in-tree")
			dirs.POST("", dirHandler.Create)
			dirs.POST("/:dirId/scan", dirHandler.Scan)
			dirs.POST("/scan-all")
			dirs.DELETE("/:dirId")
		}

		tracks := api.Group("tracks")
		{
			tracks.GET("/:trackId")
			tracks.GET("")
			tracks.GET("/:trackId/download")
		}

		covers := api.Group("covers")
		{
			covers.GET("/:coverId")
			covers.GET("/:coverId/download")
		}
	}

	return r
}
