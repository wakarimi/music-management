package audio_service

import (
	"github.com/rs/zerolog/log"
	"os"
)

func (s Service) IsExistsOnDisk(path string) (bool, error) {
	log.Debug().Str("path", path).Msg("Checking audio on disk")

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Error().Str("path", path).Msg("Audio not found on disk")
		return false, nil
	} else if err != nil {
		log.Error().Str("path", path).Msg("Failed to check audio existence")
		return false, err
	}

	return true, nil
}
