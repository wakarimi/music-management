package api

import (
	"github.com/gin-gonic/gin"
	"music-files/internal/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/music-files-service")
	{

		dirs := api.Group("/dirs")
		{
			dirs.GET("/", handlers.DirGetAll)
			dirs.POST("/", handlers.DirAdd)
			dirs.DELETE("/{dir_id}")
			dirs.POST("/{dir_id}/scan")
			dirs.POST("/scan-all")
		}

		tracks := api.Group("/tracks")
		{
			tracks.GET("/{track_id}")
			tracks.GET("/")
			tracks.GET("/{track_id}/download")
		}

		covers := api.Group("/covers")
		{
			covers.GET("/{cover_id}")
			covers.GET("/{cover_id}/download")
		}
	}

	return r
}
