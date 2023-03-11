package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goombaio/namegenerator"
)

// Functions related to registering / keeping track of users

type RegisterRequest struct {
	DesiredName string `json:"desired_name"`
}

type RegisterResponse struct {
	Token    string `json:"token"`
	UserName string `json:"username"`
}

type UserInRedis struct {
	UserName string `json:"username"`
}

func RegisterUser(c *gin.Context) {
	fmt.Println("Register user")
	// Generate the user's token
	token := GenerateToken(14)
	name := GenerateUsername()

	// Store the user in Redis
	saveMe, err := json.Marshal(UserInRedis{UserName: name})
	if err != nil {
		fmt.Println(err)
	}
	error := RedisConn.Set(context.Background(), "t-"+token, saveMe, 0).Err()
	if error != nil {
		fmt.Println(err)
	}

	// Send response
	c.IndentedJSON(http.StatusCreated, RegisterResponse{Token: token, UserName: name})

	log.Println("Created a new user with username", name, "and token", token)
}

// GenerateToken generates a random string of length length
func GenerateToken(length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GenerateUsername() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	return nameGenerator.Generate()
}
