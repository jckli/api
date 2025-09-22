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

func GetMalRedisTokens(redis rueidis.Client) (*MalTokens, error) {
	ctx := context.Background()

	refreshToken, err := redis.Do(ctx, redis.B().Get().Key("mal_refresh_token").Build()).ToString()
	if err != nil {
		if rueidis.IsRedisNil(err) {
			refreshToken = ""
		} else {
			return nil, err
		}
	}

	accessToken, err := redis.Do(ctx, redis.B().Get().Key("mal_access_token").Build()).ToString()
	if err != nil {
		if rueidis.IsRedisNil(err) {
			accessToken = ""
		} else {
			return nil, err
		}
	}

	return &MalTokens{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}

func SetMalRedisTokens(redis rueidis.Client, tokens *MalTokens) error {
	ctx := context.Background()

	cmds := make(rueidis.Commands, 0, 2)
	cmds = append(cmds, redis.B().Set().Key("mal_access_token").Value(tokens.AccessToken).ExSeconds(3540).Build())
	cmds = append(cmds, redis.B().Set().Key("mal_refresh_token").Value(tokens.RefreshToken).Build())

	for _, resp := range redis.DoMulti(ctx, cmds...) {
		if err := resp.Error(); err != nil {
			return err
		}
	}

	return nil
}
