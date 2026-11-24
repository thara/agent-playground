package main

import "context"

type Agent interface {
	HandlePrompt(ctx context.Context, prompt string) (string, error)
}
