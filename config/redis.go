package config

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var Rdb *RedisDb

type RedisDb struct {
	rdb *redis.Client
	ctx context.Context
	mux sync.RWMutex
}

func ConnectRDB() error {
	Rdb = new(RedisDb)
	Rdb.ctx = context.Background()

	viper.SetConfigName("redis")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", viper.Get("redis.hostname"), viper.Get("redis.port")),
	})

	Rdb.rdb = rdb
	_, err = Rdb.rdb.Ping().Result()
	if err != nil {
		return fmt.Errorf("连接redis失败:%v", err)
	}
	return nil
}

func SetToken(username, token string) error {
	err := Rdb.rdb.Set(fmt.Sprintf("token_%s", username), token, 10*time.Minute).Err()
	return err
}

func GetToken(username string) (string, error) {
	token, err := Rdb.rdb.Get(fmt.Sprintf("token_%s", username)).Result()
	if err != nil {
		return token, err
	}
	return token, nil
}
