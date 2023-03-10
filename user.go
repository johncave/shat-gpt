package main

import (
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

type User struct {
	UserName string `json:"username"`
}

func RegisterUser(c *gin.Context) {
	fmt.Println("Register user")
	// Generate the user's token
	token := GenerateToken(14)

	var incomingRequest RegisterRequest
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&incomingRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "Invalid register request"})
		return
	}
	var name string
	if incomingRequest.DesiredName == "" {
		name = GenerateUsername()
	} else {
		name = incomingRequest.DesiredName
	}

	// Store the user in Redis
	saveMe, err := json.Marshal(User{UserName: name})
	if err != nil {
		fmt.Println(err)
	}

	err = redisSet("t-"+token, saveMe)

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

func LookupToken(token string) (User, error) {
	val, err := redisGet("t-" + token)
	if err != nil {
		return User{}, err
	}
	var u User
	json.Unmarshal([]byte(val), &u)
	return u, nil

}
