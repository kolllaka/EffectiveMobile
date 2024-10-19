package api

import (
	"context"

	"github.com/kolllaka/EffectiveMobile/internal/model"
	"github.com/kolllaka/EffectiveMobile/pkg/logging"
)

type methods struct {
	logger logging.Logger
	conf   *model.Config
}

type Methods interface {
	GetSongInfo(ctx context.Context, songName model.AddSong) (model.Song, error)
}

func New(logger logging.Logger, conf *model.Config) Methods {
	return &methods{
		logger: logger,
		conf:   conf,
	}
}
