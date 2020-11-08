package sql

import (
	"log"
	"os"
	"time"

	"github.com/naufalfmm/project-iot/resource/config"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newLog(config *config.EnvConfig) logger.Interface {
	loggerConf := logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Silent,
		Colorful:      config.PostgresLogColorful,
	}

	if config.PostgresLogMode {
		loggerConf.LogLevel = logger.Info
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		loggerConf,
	)

	return newLogger
}

func New(config *config.EnvConfig) (*gorm.DB, error) {
	log := newLog(config)

	postgreOpen := postgres.Open(config.PostgresDbURI)
	gormConf := gorm.Config{
		Logger: log,
	}

	db, err := gorm.Open(postgreOpen, &gormConf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open db")
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "failed to open db")
	}

	sqlDb.SetMaxIdleConns(config.PostgresMaxIdleConnection)
	sqlDb.SetMaxOpenConns(config.PostgresMaxOpenConnection)
	sqlDb.SetConnMaxLifetime(config.PostgresConnMaxLifetime)

	return db, nil
}
