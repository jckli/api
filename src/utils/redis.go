package utils

import (
	"github.com/rueian/rueidis"
	"os"
)

func InitRedis() rueidis.Client {
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{os.Getenv("REDIS_URL")},
		DisableCache: true,
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	if err != nil {
		panic(err)
	}
	return client
}