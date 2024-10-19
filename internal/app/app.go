package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kolllaka/EffectiveMobile/internal/api"
	"github.com/kolllaka/EffectiveMobile/internal/config"
	"github.com/kolllaka/EffectiveMobile/internal/handler"
	repo "github.com/kolllaka/EffectiveMobile/internal/repository/postgresql"
	"github.com/kolllaka/EffectiveMobile/internal/service"
	db "github.com/kolllaka/EffectiveMobile/pkg/db/postgresql"
	"github.com/kolllaka/EffectiveMobile/pkg/logging"

	_ "github.com/lib/pq"
)

func Start() {
	logger := logging.GetLogger()

	conf := config.GetConfig(logger)

	logger.SetLevel(conf.LogLevel)

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		conf.Postgres.User, conf.Postgres.Password, conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.DBName, conf.Postgres.SSlmode)

	logger.Debugln(dsn)

	pool, err := db.NewClient(context.TODO(), 5, dsn)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.MigrationsUp(pool, conf.Postgres.MigrationFolder); err != nil {
		logger.Fatalf("Failed to migrate database: %v", err)
	}

	appStore := repo.New(logger, conf, pool)
	appAPI := api.New(logger, conf)
	appService := service.New(logger, conf, appAPI, appStore)
	appHandler := handler.New(logger, conf, appService)
	router := appHandler.Start()

	// go func() {
	// 	if err := router.Run(fmt.Sprintf("%s:%s", conf.HTTP.Host, conf.HTTP.Port)); err != nil {
	// 		logger.Fatalf("error occured while running http server: %s", err.Error())
	// 	}
	// }()
	go http.ListenAndServe(fmt.Sprintf(":%s", conf.HTTP.Port), router)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	<-sc
	logger.Infoln("Stopping...")
	pool.Close()
}
