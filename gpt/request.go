package gpt

import (
	"GoGPT/errorHandler"
	"GoGPT/logger"
	"GoGPT/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

const URL = "https://api.openai.com/v1/chat/completions"

var mu sync.Mutex
var conversationHistory = make(map[string][]map[string]interface{})

func GetResponse(config models.Bot, input string) string {
	mu.Lock()
	defer mu.Unlock()

	if len(conversationHistory[config.ID]) == 0 || (len(conversationHistory[config.ID]) > 0 && conversationHistory[config.ID][0]["content"] != config.Role) {
		systemMessage := map[string]interface{}{
			"role":    "system",
			"content": config.Role,
		}
		if len(conversationHistory[config.ID]) > 0 {
			conversationHistory[config.ID][0] = systemMessage
		} else {
			conversationHistory[config.ID] = append([]map[string]interface{}{systemMessage}, conversationHistory[config.ID]...)
		}
	}

	conversationHistory[config.ID] = append(conversationHistory[config.ID], map[string]interface{}{
		"role":    "user",
		"content": input,
	})

	payload := formatPayload(config.ID)
	response := sendRequest(config, payload)

	conversationHistory[config.ID] = append(conversationHistory[config.ID], map[string]interface{}{
		"role":    "assistant",
		"content": response,
	})

	return response
}

func formatPayload(chatID string) []byte {
	messages := conversationHistory[chatID]
	payload := map[string]interface{}{
		"model":       "gpt-4-turbo-preview",
		"messages":    messages,
		"temperature": 0.7,
	}
	payloadBytes, err := json.Marshal(payload)
	errorHandler.Handle(err)
	return payloadBytes
}

func sendRequest(config models.Bot, payload []byte) string {
	req, err := http.NewRequest("POST", URL, bytes.NewReader(payload))
	errorHandler.Handle(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+API_KEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	errorHandler.Handle(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		errorHandler.Handle(err)
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	errorHandler.Handle(err)

	var response models.ChatCompletionResponse
	err = json.Unmarshal(body, &response)
	errorHandler.Handle(err)

	logger.Log(config.ShortName, "System", "Tokens: "+strconv.Itoa(response.Usage.TotalTokens))
	fmt.Println("Tokens used: " + strconv.Itoa(response.Usage.TotalTokens))

	return response.Choices[0].Message.Content
}

func ClearHistory(chatID string) {
	mu.Lock()
	defer mu.Unlock()
	delete(conversationHistory, chatID)
}
