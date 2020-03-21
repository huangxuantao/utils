package dal_util

import (
	"fmt"
	"github.com/go-redis/redis"
	"utils/encrypt_util"
	"utils/log_util"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type RD struct {
	Client *redis.Client
}

func connRD(cfg *RedisConfig) (*RD, error) {
	RD := new(RD)
	var err error

	password := ""
	if cfg.Password != "" {
		password, err = encrypt_util.DesDecrypt(cfg.Password)
		if err != nil {
			return nil, err
		}
	}

	RD.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: password,
		DB:       cfg.DB,
	})

	return RD, nil
}

func GetRD(cfg *RedisConfig) (*RD, error) {
	RD, err := connRD(cfg)
	if err != nil {
		return nil, err
	}

	_, err = RD.Client.Ping().Result()
	if err != nil {
		return nil, err
	}
	log_util.Logger.Debugf("Redis/Connect: Connected to %s", cfg.Host)
	return RD, nil

}
