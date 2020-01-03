package load

import (
	"fmt"

	"github.com/YAWAL/ETLUNL/config"
	. "github.com/YAWAL/ETLUNL/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Postgres(conf config.Postgres) (db *gorm.DB, err error) {
	db, err = gorm.Open(conf.Dialect, fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s", conf.User,
		conf.DataBaseName, conf.SSLMode, conf.Password))
	if err != nil {
		Log.Errorf("error during connection to Postgres: %s", err.Error())
		return nil, err
	} else {
		Log.Info("Connection to Postgres has been established")
	}
	return db, err
}


