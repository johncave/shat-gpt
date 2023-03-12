package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/dayvson/go-leaderboard"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

var RedisConn *redis.Client
var GlobalLeaderBoard leaderboard.Leaderboard

func main() {
	go h.run()

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	log.Println(path)

	redisAddress, present := os.LookupEnv("REDIS_URL")
	if !present {
		redisAddress = "redis://localhost:6379"
	}

	redOpts, err := redis.ParseURL(redisAddress)
	if err != nil {
		log.Println("Error parsing Redis URL")
		return
	}
	RedisConn = redis.NewClient(redOpts)
	ping, err := RedisConn.Ping(context.TODO()).Result()
	if err != nil {
		log.Println("error pinging redis", err, ping)
		return
	}

	GlobalLeaderBoard = leaderboard.NewLeaderboard(leaderboard.RedisSettings{
		Host:     redOpts.Addr,
		Password: redOpts.Password,
	}, "highscores", 50)

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

	router.POST("/api/chatgpt", callChatGptApi)

	listenAddress, present := os.LookupEnv("PORT")
	if !present {
		listenAddress = "0.0.0.0:8080"
	} else {
		listenAddress = "0.0.0.0:" + os.Getenv("PORT")
	}

	// Print out the scoreboard periodically
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for range ticker.C {
			lb := GlobalLeaderBoard.GetLeaders(1)
			log.Print()
			var lbUsers []WSLeaderboardUser
			if lb[0].Score != 0 {
				for _, s := range lb {
					if s.Score > 0 {
						lbUsers = append(lbUsers, WSLeaderboardUser{UserName: s.Name, Score: s.Score, Rank: s.Rank})
						fmt.Printf("%02d  %05d %s\n", s.Rank, s.Score, s.Name)
					}

				}
				wsout := WSLeaderboard{EventName: "leaderboard", Timestamp: time.Now(), Leaderboard: lbUsers}
				outjson, _ := json.MarshalIndent(wsout, "", "  ")

				m := message{outjson, "shatgpt"}
				h.broadcast <- m
			} else {
				fmt.Println("No users")
			}

		}
	}()

	router.Run(listenAddress)
}

func ReverseProxy(target string) gin.HandlerFunc {
	url, err := url.Parse(target)
	if err != nil {
		log.Println(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
