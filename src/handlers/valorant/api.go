package valorant

import (
	"os"
	"fmt"
	"encoding/json"
	"github.com/jckli/valorant.go/v2"
)

func getAuth() (*val.AuthBody, error) {
	auth, err := val.Authentication(os.Getenv("VALORANT_USERNAME"), os.Getenv("VALORANT_PASSWORD"))
	if err != nil {
		return nil, err
	}
	return auth, nil

}

func getMmr(auth *val.AuthBody, puuid string) (*MMRFetchPlayerResponse, error) {
	resp, err := val.FetchGet("/mmr/v1/players/" + puuid, "pd", auth)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 400 || resp.StatusCode == 403 {
		return nil, fmt.Errorf("bad_claims")
	}
	defer resp.Body.Close()
	body := new(MMRFetchPlayerResponse)
	json.NewDecoder(resp.Body).Decode(body)

	return body, nil
}