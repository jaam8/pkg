package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Config struct {
	Host     string `yaml:"host" env:"HOST" validate:"required,ip|hostname"`
	Port     uint16 `yaml:"port" env:"PORT" validate:"required,port"`
	User     string `yaml:"user" env:"USER" validate:"required"`
	Password string `yaml:"password" env:"PASSWORD" validate:"required,min=4,max=128"`
	Database string `yaml:"db" env:"DB" validate:"required"`
	SSL      string `yaml:"ssl" env:"SSL" validate:"required,oneof=enable disable"`
	MaxConns int32  `yaml:"max_conns" env:"MAX_CONNS" validate:"required,min=1"`
	MinConns int32  `yaml:"min_conns" env:"MIN_CONNS" validate:"required,min=1"`
}

func New(ctx context.Context, config Config) (*pgxpool.Pool, error) {
	dsn := config.GetDsn()
	dsn += fmt.Sprintf("&pool_max_conns=%d&pool_min_conns=%d",
		config.MaxConns,
		config.MinConns,
	)

	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %v", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %v", err)
	}

	return conn, nil
}

func (c *Config) GetDsn() string {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)
	return dsn
}
