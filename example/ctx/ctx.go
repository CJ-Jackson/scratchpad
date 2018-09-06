package ctx

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	"github.com/cjtoolkit/ctx"
)

type Config struct {
	DbRsn string `json:"DbRsn"`
}

func GetConfig(context ctx.BackgroundContext) Config {
	type ConfigContext struct{}
	return context.Persist(ConfigContext{}, func() (interface{}, error) {
		return initConfig(), nil
	}).(Config)
}

func initConfig() (config Config) {
	file, err := os.Open("setting.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func GetDatabaseConnection(context ctx.BackgroundContext) *sql.DB {
	type DatabaseContext struct{}
	return context.Persist(DatabaseContext{}, func() (interface{}, error) {
		return initDatabaseConnection(context)
	}).(*sql.DB)
}

func initDatabaseConnection(context ctx.BackgroundContext) (*sql.DB, error) {
	return sql.Open("postgres", GetConfig(context).DbRsn)
}
