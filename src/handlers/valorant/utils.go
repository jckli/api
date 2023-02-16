package valorant

import (
	"context"
	"github.com/jckli/valorant.go/v2"
	"github.com/rueian/rueidis"
)

func redisGetAuth(redis rueidis.Client) (*val.AuthBody, error) {
	redis_ctx := context.Background()
	cookies, cookiesErr := redis.Do(redis_ctx, redis.B().Get().Key("valorant_cookies").Build()).ToString()
	if (cookiesErr != nil) || (cookies == "") {
		return nil, cookiesErr
	}
	region, regionErr := redis.Do(redis_ctx, redis.B().Get().Key("valorant_region").Build()).ToString()
	if regionErr != nil {
		return nil, regionErr
	}
	access_token, acErr := redis.Do(redis_ctx, redis.B().Get().Key("valorant_access_token").Build()).ToString()
	if (acErr != nil) || (access_token == "") {
		return nil, acErr
	}
	token, tokenErr := redis.Do(redis_ctx, redis.B().Get().Key("valorant_token").Build()).ToString()
	if (tokenErr != nil) || (token == "") {
		return nil, tokenErr
	}
	version, versionErr := redis.Do(redis_ctx, redis.B().Get().Key("valorant_version").Build()).ToString()
	if (versionErr != nil) || (version == "") {
		return nil, versionErr
	}
	puuid, puuidErr := redis.Do(redis_ctx, redis.B().Get().Key("valorant_puuid").Build()).ToString()
	if puuidErr != nil {
		return nil, puuidErr
	}
	return &val.AuthBody{
		Cookies:     cookies,
		Region:      region,
		AccessToken: access_token,
		Token:       token,
		Version:     version,
		Puuid:       puuid,
	}, nil
}

func redisSetAuth(redis rueidis.Client, auth *val.AuthBody) error {
	redis_ctx := context.Background()
	redis.Do(redis_ctx, redis.B().Set().Key("valorant_cookies").Value(auth.Cookies).Build()).Error()
	redis.Do(redis_ctx, redis.B().Set().Key("valorant_region").Value(auth.Region).Build()).Error()
	redis.Do(redis_ctx, redis.B().Set().Key("valorant_access_token").Value(auth.AccessToken).Build()).Error()
	redis.Do(redis_ctx, redis.B().Set().Key("valorant_token").Value(auth.Token).Build()).Error()
	redis.Do(redis_ctx, redis.B().Set().Key("valorant_version").Value(auth.Version).Build()).Error()
	redis.Do(redis_ctx, redis.B().Set().Key("valorant_puuid").Value(auth.Puuid).Build()).Error()
	return nil
}
	