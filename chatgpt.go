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

func callChatGptApi(c *gin.Context) {
	var incomingRequest AskRequest

	fmt.Print(incomingRequest.Prompt)

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

	askChatGPT(prompt)
	c.IndentedJSON(http.StatusCreated, "")
}



func askChatGPT(prompt string) {
	openaiCompletionsURL := "https://api.openai.com/v1/completions"

	contextPrompt := fmt.Sprintf("Answer the following prompt with as much uwu as possible and with poop emoji and poop innuendo: %s", prompt)

	values := map[string]string{
		"model": "gpt-3.5-turbo", 
		"prompt": contextPrompt,
		"max_tokens": "1000",
		"temperature": "0.5",
	}
    json_data, err := json.Marshal(values)

    if err != nil {
        log.Fatal(err)
    }

    r, err := http.NewRequest("POST", openaiCompletionsURL, bytes.NewBuffer(json_data))
	if err != nil {
        log.Fatal(err)
    }
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s",os.Getenv("CHATGPT_KEY")))
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
        log.Fatal(err)
    }

	defer resp.Body.Close()

    var res map[string]interface{}

    json.NewDecoder(resp.Body).Decode(&res)

    fmt.Println(res["json"])
}