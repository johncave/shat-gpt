package main

import (
	// openapi "github.com/johncave/shat-gpt/chatgpt-go"
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type AskRequest struct {
	Prompt string `json:"prompt"`
}

type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTOptions struct {
	Messages []Message `json:"messages"`
	Model string `json:"model"`
	MaxTokens int `json:"max_tokens"`
	Temperature float32 `json:"temperature"`
}

type ChatGPTUsage struct {
	Prompt int `json:"prompt_tokens"`
	Completion int `json:"completion_tokens"`
	Total int `json:"total_tokens"`
}

type ChatGPTChoices struct {
	Message Message `json:"message"`
	FinishReason string `json:"finish_reason"`
	Index int `json:"index"`
}

type ChatGPTOutput struct {
	ID string `json:"id"`
	Object string `json:"object"`
	Created int `json:"created"`
	Model string `json:"model"`
	Choices []ChatGPTChoices `json:"choices"`
	Usage ChatGPTUsage `json:"usage"`
}

func callChatGptApi(c *gin.Context) {
	var incomingRequest AskRequest

	if err := c.BindJSON(&incomingRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "Invalid chatgpt request"})
		return
	}

	var prompt string
	if incomingRequest.Prompt == "" {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "Prompt body cannot be empty"})
	} else {
		prompt = incomingRequest.Prompt
	}

	response := askChatGPT(prompt)
	fmt.Println(response)

	json_data, err := json.Marshal(response)

    if err != nil {
        log.Fatal(err)
    }
	
	m := message{json_data, "shatgpt"}
	h.broadcast <- m

	c.IndentedJSON(http.StatusCreated, response)
}


func askChatGPT(prompt string) string {
	openaiCompletionsURL := "https://api.openai.com/v1/chat/completions"

	systemContext := "Answer the following prompt with as much uwu as possible and with poop emoji and poop innuendo"

	options := ChatGPTOptions{
		Model: "gpt-3.5-turbo",
		Temperature: 0.5,
		MaxTokens: 1000,
		Messages: []Message{
			Message{
				Role: "system",
				Content: systemContext,
			},
			Message{
				Role: "user",
				Content: prompt,
			},
		},
	}

    json_data, err := json.Marshal(options)

    if err != nil {
        log.Fatal(err)
    }

    r, err := http.NewRequest("POST", openaiCompletionsURL, bytes.NewBuffer(json_data))
	if err != nil {
        log.Fatal(err)
    }

	key := os.Getenv("CHATGPT_KEY")

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s",key))
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
        log.Fatal(err)
    }

	defer resp.Body.Close()

    var res ChatGPTOutput

    json.NewDecoder(resp.Body).Decode(&res)

	messageResponse := res.Choices[0].Message.Content

	return messageResponse
}