package main

import (
	"context"
	"errors"
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
		var apiErr *openai.Error
		if errors.As(err, &apiErr) {
			return "", fmt.Errorf("create response: %s: %s: %w", apiErr.Code, apiErr.Message, err)
		}
		return "", fmt.Errorf("create response: %w", err)
	}
	return resp.OutputText(), nil
}
