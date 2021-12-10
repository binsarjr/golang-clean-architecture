package database

import (
	"context"

	"giapps/servisin/infrastructure/exception"

	log "gopkg.in/inconshreveable/log15.v2"

	"github.com/jackc/pgx/v4/log/log15adapter"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres struct {
	config     *pgxpool.Config
	Uri        string
	MaxConns   int32
	MinConns   int32
	Connection PostgresConnection
}

type PostgresConnection struct {
	User     string
	Password string
	Host     string
	Port     uint16
	Database string
}

func (pgsl *Postgres) ParseConfig(conn PostgresConnection) {
	poolConfig, err := pgxpool.ParseConfig("")
	exception.PanicIfNeeded(err)
	poolConfig.ConnConfig.Host = conn.Host
	poolConfig.ConnConfig.User = conn.User
	poolConfig.ConnConfig.Password = conn.Password
	poolConfig.ConnConfig.Port = conn.Port
	poolConfig.ConnConfig.Database = conn.Database
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
	postgres.ParseConfig(postgres.Connection)
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
