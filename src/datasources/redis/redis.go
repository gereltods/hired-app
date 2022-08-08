package redis

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	redisHost = "REDIS_SERVER"
	redisPORT = "REDIS_PORT"
	redisPwd  = "REDIS_PASSWORD"
)

type Params struct {
	Id        string
	Data      string
	AtExpires int64
}

type (
	redisService   struct{}
	redisInterface interface {
		SetRedis(Params) error
		GetRedis(string) (string, error)
		DeleteRedis(string) error
	}
)

var (
	client       *redis.Client
	RedisService redisInterface = &redisService{}
	ctx                         = context.Background()
	rHost                       = os.Getenv(redisHost)
	rPort                       = os.Getenv(redisPORT)
	rPwd                        = os.Getenv(redisPwd)
)

func init() {
	if rHost == "" {
		rHost = "103.50.205.205"
		rPort = "16379"
		rPwd = ""
	}

	client = redis.NewClient(&redis.Options{
		Addr:     rHost + ":" + rPort,
		Password: rPwd,
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	log.Println("Redis successfully configured")
}

// save Redis

func (cache *redisService) SetRedis(p Params) error {

	if p.AtExpires != 0 {
		at := time.Unix(p.AtExpires, 0) // converting Unix to UTC(to Time object)
		now := time.Now()

		atCreated, err := client.Set(ctx, p.Id, p.Data, at.Sub(now)).Result()
		if err != nil {
			return err
		}

		if atCreated == "0" {
			return errors.New("no record inserted")
		}
	} else {
		atCreated, err := client.Set(ctx, p.Id, p.Data, 0).Result()
		if err != nil {
			fmt.Printf("error: %+v", err)
			return err
		}

		if atCreated == "0" {
			return errors.New("no record inserted")
		}
	}
	return nil
}

func (cache *redisService) GetRedis(id string) (string, error) {
	info, err := client.Get(ctx, id).Result()
	if err != nil {
		return "", nil
	}
	return info, nil
}

func (cache *redisService) DeleteRedis(id string) error {
	deleteAt, err := client.Del(ctx, id).Result()
	if err != nil {
		return err
	}
	if deleteAt != 1 {
		return errors.New("something wrong")
	}
	return nil
}
