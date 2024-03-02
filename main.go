package main

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/jckli/api/src/routes"
	"github.com/jckli/api/src/utils"
	_ "github.com/joho/godotenv/autoload"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting API...")
	redis := utils.InitRedis()
	defer redis.Close()
	fmt.Println("Redis connected")

	r := router.New()
	routes.InitRoutes(r, redis)

	port := os.Getenv("PORT")
	fmt.Println("API on port: " + port)
	log.Fatal(fasthttp.ListenAndServe("0.0.0.0:"+port, r.Handler))
}
