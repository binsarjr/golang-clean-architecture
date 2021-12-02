package database

import (
	"context"

	"giapps/newapp/exception"

	log "gopkg.in/inconshreveable/log15.v2"

	"github.com/jackc/pgx/v4/log/log15adapter"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres struct {
	config   *pgxpool.Config
	Uri      string
	MaxConns int32
	MinConns int32
}

func (pgsl *Postgres) ParseConfig(dsn string) {
	poolConfig, err := pgxpool.ParseConfig(dsn)
	exception.PanicIfNeeded(err)
	pgsl.config = poolConfig
}

func (pgsl *Postgres) SetMaxConns(max int32) {
	pgsl.config.MaxConns = max
}

func (pgsl *Postgres) SetMinConns(min int32) {
	pgsl.config.MinConns = min
}

func (pgsl *Postgres) SetLogger(logger *log15adapter.Logger) {
	pgsl.config.ConnConfig.Logger = logger
}

func (pgsl *Postgres) Connect() *pgxpool.Pool {
	db, err := pgxpool.ConnectConfig(context.Background(), pgsl.config)
	exception.PanicIfNeeded(err)
	return db
}

func NewPostgres(postgres Postgres) *pgxpool.Pool {
	postgres.ParseConfig(postgres.Uri)
	postgres.SetLogger(log15adapter.NewLogger(log.New("module", "pgx")))

	if postgres.MaxConns != 0 {
		postgres.SetMaxConns(postgres.MaxConns)
	}
	if postgres.MinConns != 0 {
		postgres.SetMinConns(postgres.MinConns)
	}

	db, err := pgxpool.ConnectConfig(context.Background(), postgres.config)
	exception.PanicIfNeeded(err)

	return db
}
