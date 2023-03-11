package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

var RedisConn *redis.Client

func main() {
	go h.run()

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	log.Println(path)

	redisAddress, present := os.LookupEnv("REDIS_URL")
	if !present {
		redisAddress = "localhost:6379"
	}

	RedisConn = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ping, err := RedisConn.Ping(context.TODO()).Result()
	if err != nil {
		log.Println("error pinging redis", err, ping)
	}

	router := gin.New()
	router.LoadHTMLFiles("./shatgpt-frontend/dist")
	router.StaticFile("/favicon.ico", "favicon.ico")
	router.Use(static.Serve("/", static.LocalFile("./shatgpt-frontend/dist", true)))

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "./shatgpt-frontend/dist", nil)
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, roomId)
	})

	// API Routes
	router.POST("/api/register", RegisterUser)

	listenAddress, present := os.LookupEnv("PORT")
	if !present {
		listenAddress = "0.0.0.0:8080"
	} else {
		listenAddress = "0.0.0.0:" + os.Getenv("PORT")
	}

	router.Run(listenAddress)
}
