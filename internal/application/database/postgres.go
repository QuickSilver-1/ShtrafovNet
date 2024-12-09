package database

import (
	"auction/internal/application/config"
	"auction/internal/application/logger"
	repo "auction/internal/infrastructure/repository/postgres"
)

// DB глобальная переменная для хранения подключения к базе данных
var (
    DB = repo.NewDb(
        config.AppConfig.PgHost, config.AppConfig.PgName,
        config.AppConfig.PgUser, config.AppConfig.PgPass,
        config.AppConfig.PgPort, logger.Log,
    )
)
