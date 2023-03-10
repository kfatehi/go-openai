package openai

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var req = ChatRequest{
	Model: "gpt-3.5-turbo",
	Messages: []ChatMessage{
		{
			Role:    "user",
			Content: "Hello!",
		},
	},
}

func TestChatCompletionHappyPath(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")

	resp, err := ChatCompletion(apiKey, req)
	if err != nil {
		t.Fatalf("error completing request: %v", err)
	}

	if len(resp.Choices) != 1 {
		t.Fatalf("expected 1 choice, got %d", len(resp.Choices))
	}

	if resp.Choices[0].Content == "" {
		t.Fatalf("expected non-empty text, got empty string")
	}

	log.Println(resp.Choices[0].Content)
}

func TestChatCompletionSadPath(t *testing.T) {
	_, err := ChatCompletion("clearly invalid", req)
	if err == nil {
		t.Fatalf("error: should have received an error due to invalid api key")
	}
}
