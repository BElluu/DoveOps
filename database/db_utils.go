package database

import (
	"strconv"
	"sync"

	"github.com/belluu/doveops/configuration"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
  once sync.Once
)

func GetDB() *gorm.DB{
  once.Do(func() {
    var err error
    config := configuration.Get()
    conn := "host=" + config.DatabaseConfig.Host + " port=" + strconv.Itoa(config.DatabaseConfig.Port) + " user=" + config.DatabaseConfig.User + " dbname=" + config.DatabaseConfig.Database + " password=" + config.DatabaseConfig.Password + " sslmode=disable TimeZone=" + config.DatabaseConfig.TimeZone
    db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
    if err != nil {
      log.Error().Msg("Error connecting to database.")
      panic(err)
    }
  })
  return db
}
