package utils

import (
	"context"
	"os"

	"github.com/rueian/rueidis"
)

func InitRedis() rueidis.Client {
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{os.Getenv("REDIS_URL")},
	})
	if err != nil {
		panic(err)
	}
	redis_ctx := context.Background()

	_, err = client.Do(redis_ctx, client.B().Get().Key("spotify_access_token").Build()).ToString()
	if err != nil {
		client.Do(redis_ctx, client.B().Set().Key("spotify_access_token").Value("").Build()).Error()
	}

	return client
}

func GetOnedriveRedisTokens(redis rueidis.Client) (*OnedriveTokens, error) {
	redis_ctx := context.Background()

	refreshToken, err := redis.Do(redis_ctx, redis.B().Get().Key("onedrive_refresh_token").Build()).
		ToString()
	if err != nil {
		if err == rueidis.Nil {
			refreshToken = ""
		} else {
			return nil, err
		}
	}

	accessToken, err := redis.Do(redis_ctx, redis.B().Get().Key("onedrive_access_token").Build()).
		ToString()
	if err != nil {
		if err == rueidis.Nil {
			accessToken = ""
		} else {
			return nil, err
		}
	}

	return &OnedriveTokens{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}

func SetOnedriveRedisTokens(redis rueidis.Client, tokens *OnedriveTokens) error {
	redis_ctx := context.Background()
	err := redis.Do(redis_ctx, redis.B().Set().Key("onedrive_refresh_token").Value(tokens.RefreshToken).Build()).
		Error()
	if err != nil {
		return err
	}
	err = redis.Do(redis_ctx, redis.B().Set().Key("onedrive_access_token").Value(tokens.AccessToken).Build()).
		Error()
	if err != nil {
		return err
	}
	err = redis.Do(redis_ctx, redis.B().Expire().Key("onedrive_access_token").Seconds(tokens.AccessTokenExpiry).Build()).
		Error()
	if err != nil {
		return err
	}

	return nil
}
