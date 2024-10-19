package service

import (
	"github.com/kolllaka/EffectiveMobile/internal/api"
	"github.com/kolllaka/EffectiveMobile/internal/model"
	db "github.com/kolllaka/EffectiveMobile/internal/repository/postgresql"
	"github.com/kolllaka/EffectiveMobile/pkg/logging"
)

type Service interface {
	MusicService
}

type service struct {
	logger logging.Logger
	conf   *model.Config

	api   api.Methods
	store db.Store
}

func New(logger logging.Logger, conf *model.Config, api api.Methods, store db.Store) Service {
	return &service{
		logger: logger,
		conf:   conf,

		api:   api,
		store: store,
	}
}
