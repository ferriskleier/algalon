package models

import "encoding/json"

type ChatCompletionResponse struct {
	ID                string       `json:"id"`
	Object            string       `json:"object"`
	Created           int64        `json:"created"`
	Model             string       `json:"model"`
	Choices           []Choice     `json:"choices"`
	Usage             Usage        `json:"usage"`
	SystemFingerprint *interface{} `json:"system_fingerprint"`
}

type Choice struct {
	Index        int              `json:"index"`
	Message      Message          `json:"message"`
	LogProbs     *json.RawMessage `json:"logprobs"`
	FinishReason string           `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
