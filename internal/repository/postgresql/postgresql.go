package db

import (
	"context"

	"github.com/kolllaka/EffectiveMobile/internal/model"
	db "github.com/kolllaka/EffectiveMobile/pkg/db/postgresql"
	"github.com/kolllaka/EffectiveMobile/pkg/logging"
)

type Store interface {
	GetSongs(ctx context.Context, param model.QueryParam) ([]model.Song, error)
	CreateSong(ctx context.Context, newSong model.Song) (string, error)
	ChangeSong(ctx context.Context, song model.Song) error
	DeleteSong(ctx context.Context, id string) error
	GetTextOfSong(ctx context.Context, id string) (string, error)
}
type store struct {
	logger logging.Logger
	conf   *model.Config

	db db.Client
}

func New(logger logging.Logger, conf *model.Config, db db.Client) Store {
	return &store{
		logger: logger,
		conf:   conf,
		db:     db,
	}
}
