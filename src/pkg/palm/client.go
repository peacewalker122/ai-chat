package palm

import (
	"context"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func NewClient(ctx context.Context, api_key string) *genai.GenerativeModel {
	client, err := genai.NewClient(ctx, option.WithAPIKey(api_key))
	if err != nil {
		panic(err)
	}

	// The Gemini 1.5 models are versatile and work with most use cases
	model := client.GenerativeModel("gemini-1.5-flash")

	return model
}
