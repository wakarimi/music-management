package track

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"music-files/internal/handlers/types"
	"music-files/internal/models"
	"net/http"
)

type readAllResponseItem struct {
	TrackId   int    `json:"trackId"`
	CoverId   *int   `json:"coverId,omitempty"`
	Extension string `json:"extension"`
	Size      int64  `json:"size"`
	Hash      string `json:"hash"`
}

type readAllResponse struct {
	Tracks []readAllResponseItem `json:"tracks"`
}

func (h *Handler) ReadAll(c *gin.Context) {
	log.Debug().Msg("Fetching all tracks")

	var tracks []models.Track

	err := h.TransactionManager.WithTransaction(func(tx *sqlx.Tx) (err error) {
		tracks, err = h.TrackService.ReadAll(tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch all tracks")
		c.JSON(http.StatusInternalServerError, types.Error{
			Error: "Failed to fetch all tracks",
		})
		return
	}

	tracksResponse := make([]readAllResponseItem, 0)
	for _, track := range tracks {
		trackResponse := readAllResponseItem{
			TrackId:   track.TrackId,
			CoverId:   track.CoverId,
			Extension: track.Extension,
			Size:      track.Size,
			Hash:      track.Hash,
		}
		tracksResponse = append(tracksResponse, trackResponse)
	}

	log.Info().Int("tracksCount", len(tracks)).Msg("Track fetched successfully")
	c.JSON(http.StatusOK, readAllResponse{
		Tracks: tracksResponse,
	})
}