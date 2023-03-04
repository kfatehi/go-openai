package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type OpenAIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param,omitempty"`
	Code    string `json:"code,omitempty"`
}

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatChoice struct {
	Index        int `json:"index"`
	ChatMessage  `json:"message"`
	FinishReason string `json:"finish_reason"`
}

type ChatResponse struct {
	ID        string       `json:"id"`
	Object    string       `json:"object"`
	Created   int          `json:"created"`
	Choices   []ChatChoice `json:"choices"`
	ChatUsage `json:"usage"`
}

type ChatUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func ChatCompletion(apiKey string, request ChatRequest) (ChatResponse, error) {
	url := "https://api.openai.com/v1/chat/completions"

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return ChatResponse{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequest))
	if err != nil {
		return ChatResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ChatResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var openaiErr OpenAIError
		err := json.NewDecoder(resp.Body).Decode(&openaiErr)
		if err != nil {
			return ChatResponse{}, err
		}
		return ChatResponse{}, fmt.Errorf("%s: %s", resp.Status, openaiErr.Message)
	}

	var response ChatResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return ChatResponse{}, err
	}

	return response, err
}
