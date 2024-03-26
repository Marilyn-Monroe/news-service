package configuration

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"math/rand"
	"news-service/exception"
	"strconv"
	"time"
)

func NewDatabase(config Config) *reform.DB {
	username := config.Get("POSTGRES_USER")
	password := config.Get("POSTGRES_PASSWORD")
	host := config.Get("POSTGRES_HOST")
	port := config.Get("POSTGRES_PORT")
	dbName := config.Get("POSTGRES_DB")
	maxPoolOpen, err := strconv.Atoi(config.Get("POSTGRES_POOL_MAX_CONN"))
	maxPoolIdle, err := strconv.Atoi(config.Get("POSTGRES_POOL_IDLE_CONN"))
	maxPollLifeTime, err := strconv.Atoi(config.Get("POSTGRES_POOL_LIFE_TIME"))
	exception.PanicLogging(err)

	db, err := sql.Open("postgres", "postgres://"+username+":"+password+"@"+host+":"+port+"/"+dbName+"?sslmode=disable")
	exception.PanicLogging(err)

	db.SetMaxOpenConns(maxPoolOpen)
	db.SetMaxIdleConns(maxPoolIdle)
	db.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)

	sqlDB := reform.NewDB(db, postgresql.Dialect, nil)

	return sqlDB
}
