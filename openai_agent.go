package main

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/responses"
)

type openAIAgent struct {
	client openai.Client
	model  string
}

func (a *openAIAgent) HandlePrompt(ctx context.Context, prompt string) (string, error) {
	resp, err := a.client.Responses.New(ctx, responses.ResponseNewParams{
		Model: a.model,
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(prompt)},
	})
	if err != nil {
		return "", fmt.Errorf("create response: %w", err)
	}
	return resp.OutputText(), nil
}
