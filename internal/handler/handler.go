package handler

import (
	"github.com/kolllaka/EffectiveMobile/internal/model"
	"github.com/kolllaka/EffectiveMobile/internal/service"
	"github.com/kolllaka/EffectiveMobile/pkg/logging"

	"github.com/gin-gonic/gin"
	_ "github.com/kolllaka/EffectiveMobile/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type server struct {
	logger logging.Logger
	conf   *model.Config

	service service.Service
}

func New(logger logging.Logger, conf *model.Config,
	service service.Service) *server {
	return &server{
		logger:  logger,
		conf:    conf,
		service: service,
	}
}

func (s *server) Start() *gin.Engine {
	if s.conf.IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/")
	{
		api.GET("/songs", s.getSongsHandler)
		api.POST("/songs", s.createSongHandler)
		api.PUT("/songs/:id", s.changeSongHandler)
		api.DELETE("/songs/:id", s.deleteSongHandler)
		api.GET("/songs/:id", s.getTextOfSongHandler)
	}

	return router
}
