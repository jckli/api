package valorant

import (
	"github.com/jckli/valorant.go"
	"os"
)

func getAuth() (*valorant.Auth, error) {
	auth, err := valorant.New(os.Getenv("VALORANT_USERNAME"), os.Getenv("VALORANT_PASSWORD"))
	if err != nil {
		return nil, err
	}

	return auth, nil
}
