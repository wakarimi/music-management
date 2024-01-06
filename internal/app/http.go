package app

import (
	"fmt"
	"music-files/internal/config"
)

func (a *App) RegisterRoutes() {
	api := a.router.Group("/api")
	{
		roots := api.Group("/roots")
		{
			roots.POST("", a.handler.AddRoot)
			roots.GET("", a.handler.GetRoots)
			roots.DELETE("/:dirId", a.handler.DeleteRoot)
		}

		dirs := api.Group("/dirs")
		{
			dirs.GET("/:dirId", a.handler.GetDir)
			dirs.GET("/:dirId/content", a.handler.GetDirContent)
			dirs.POST("/:dirId/scan", a.handler.ScanDir)
			dirs.POST("/scan", a.handler.ScanDirs)
		}

		audios := api.Group("/audios")
		{
			audios.GET("", a.handler.GetAudios)
			audios.GET("/all", a.handler.GetAllAudios)
			audios.GET("/:audioId/download", a.handler.DownloadAudio)
			audios.GET("/best-covers", a.handler.GetBestCovers)
			audios.GET("/sha256/:sha256", a.handler.SearchBySha256)
		}

		covers := api.Group("/covers")
		{
			covers.GET("", a.handler.GetCovers)
			covers.GET("/:coverId/image", a.handler.GetStaticImage)
		}
	}
}

func (a *App) StartHTTPServer(cfg config.HTTPConfig) error {
	err := a.router.Run(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return err
	}
	return nil
}
