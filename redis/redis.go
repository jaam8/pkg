package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string `yaml:"host" env:"HOST" validate:"required,ip|hostname"`
	Port     uint16 `yaml:"port" env:"PORT" validate:"required,port"`
	Username string `yaml:"username" env:"USERNAME" validate:"required"`
	Password string `yaml:"password" env:"PASSWORD" validate:"required,min=4,max=128"`

	MaxRetries    int `yaml:"max_retries" env:"MAX_RETRIES" validate:"required,gt=0"`
	PoolSize      int `yaml:"pool_size" env:"POOL_SIZE" validate:"required,gt=0"`
	DialTimeoutS  int `yaml:"dial_timeout_s" env:"DIAL_TIMEOUT_S" validate:"required,gt=0"`
	ReadTimeoutS  int `yaml:"read_timeout_s" env:"READ_TIMEOUT_S" validate:"required,gt=0"`
	WriteTimeoutS int `yaml:"write_timeout_s" env:"WRITE_TIMEOUT_S" validate:"required,gt=0"`
}

// New try to connect to Redis and get the client
func New(ctx context.Context, config Config, redisDB int) (*redis.Client, error) {
	option := &redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		Username:     config.Username,
		Password:     config.Password,
		DB:           redisDB,
		MaxRetries:   config.MaxRetries,
		DialTimeout:  time.Duration(config.DialTimeoutS) * time.Second,
		ReadTimeout:  time.Duration(config.ReadTimeoutS) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeoutS) * time.Second,
		PoolSize:     config.PoolSize,
	}
	client := redis.NewClient(option)
	res := client.Ping(ctx)
	if res.Err() != nil {
		return nil, res.Err()
	}
	
	return client, nil
}
