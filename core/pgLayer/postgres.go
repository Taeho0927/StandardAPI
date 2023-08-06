package pgLayer

import (
	"github.com/Taeho0927/goLogger/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github/Taeho0927/Standard_API/conf"
	"github/Taeho0927/Standard_API/middlewares"
)

type PGORM struct {
	*gorm.DB
}

func PGConn() (*PGORM, error) {
	dsn := conf.SettingPostgresDsn()
	pg, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	logger.Debug(middlewares.SetLogger("PGConn", "PostgreSQL Connection Made"))
	return &PGORM{
		DB: pg,
	}, nil
}
