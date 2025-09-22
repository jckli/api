package main

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/jckli/api/routes"
	"github.com/jckli/api/utils"
	_ "github.com/joho/godotenv/autoload"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting API...")
	if os.Getenv("REDIS_URL") == "" {
		panic("REDIS_URL not set")
	}
	if os.Getenv("SPOTIFY_CLIENT_ID") == "" {
		panic("SPOTIFY_CLIENT_ID not set")
	}
	if os.Getenv("SPOTIFY_CLIENT_SECRET") == "" {
		panic("SPOTIFY_CLIENT_SECRET not set")
	}
	if os.Getenv("SPOTIFY_REFRESH_TOKEN") == "" {
		panic("SPOTIFY_REFRESH_TOKEN not set")
	}
	if os.Getenv("VALORANT_USERNAME") == "" {
		panic("VALORANT_USERNAME not set")
	}
	if os.Getenv("VALORANT_PASSWORD") == "" {
		panic("VALORANT_PASSWORD not set")
	}
	if os.Getenv("ONEDRIVE_CLIENT_ID") == "" {
		panic("ONEDRIVE_CLIENT_ID not set")
	}
	if os.Getenv("ONEDRIVE_CLIENT_SECRET") == "" {
		panic("ONEDRIVE_CLIENT_SECRET not set")
	}
	if os.Getenv("ONEDRIVE_REFRESH_TOKEN") == "" {
		panic("ONEDRIVE_REFRESH_TOKEN not set")
	}
	if os.Getenv("MAL_REFRESH_TOKEN") == "" {
		panic("MAL_REFRESH_TOKEN not set")
	}
	if os.Getenv("MAL_CLIENT_ID") == "" {
		panic("MAL_CLIENT_ID not set")
	}
	if os.Getenv("MAL_CLIENT_SECRET") == "" {
		panic("MAL_CLIENT_SECRET not set")
	}

	redis := utils.InitRedis()
	defer redis.Close()
	fmt.Println("Redis connected")

	client := &fasthttp.Client{}

	r := router.New()
	routes.InitRoutes(r, redis, client)

	port := os.Getenv("PORT")
	fmt.Println("API on port: " + port)
	log.Fatal(fasthttp.ListenAndServe("0.0.0.0:"+port, r.Handler))
}
