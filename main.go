package main

import (
	"fmt"
	"log"
	"os"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	_ "github.com/joho/godotenv/autoload"
	"github.com/jckli/api/src/routes"
	"github.com/jckli/api/src/utils"
)

func main() {
	redis := utils.InitRedis()
	defer redis.Close()

	r := router.New()
	routes.InitRoutes(r, redis)

	port := os.Getenv("PORT")
	fmt.Println("API on port: " + port)
	log.Fatal(fasthttp.ListenAndServe(":" + port, r.Handler))
}

