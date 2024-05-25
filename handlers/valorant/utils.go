package valorant

import (
	"context"
	"github.com/jckli/valorant.go"
	"github.com/rueian/rueidis"
)

func redisGetAuth(redis rueidis.Client) (*valorant.Auth, error) {
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
	id_token, idTokenErr := redis.Do(redis_ctx, redis.B().Get().Key("valorant_id_token").Build()).ToString()
	if (idTokenErr != nil) || (id_token == "") {
		return nil, idTokenErr
	}
	expires_in, expiresInErr := redis.Do(redis_ctx, redis.B().Get().Key("valorant_expires_in").Build()).ToString()
	if (expiresInErr != nil) || (expires_in == "") {
		return nil, expiresInErr
	}

	auth, err := valorant.CreateAuth(cookies, region, access_token, id_token, expires_in, token, version)
	if err != nil {
		return nil, err
	}
	return auth, nil

}

func redisSetAuth(redis rueidis.Client, a *valorant.Auth) error {
	redis_ctx := context.Background()
	redis.Do(redis_ctx, redis.B().Set().Key("valorant_cookies").Value(a.CookieJar).Build()).Error()
	redis.Do(redis_ctx, redis.B().Set().Key("valorant_region").Value(a.Region).Build()).Error()
	redis.Do(redis_ctx, redis.B().Set().Key("valorant_access_token").Value(a.AccessToken).Build()).Error()
	redis.Do(redis_ctx, redis.B().Set().Key("valorant_token").Value(a.Token).Build()).Error()
	redis.Do(redis_ctx, redis.B().Set().Key("valorant_version").Value(a.Version).Build()).Error()
	redis.Do(redis_ctx, redis.B().Set().Key("valorant_id_token").Value(a.IdToken).Build()).Error()
	redis.Do(redis_ctx, redis.B().Set().Key("valorant_expires_in").Value(a.ExpiresIn).Build()).Error()
	return nil
}
