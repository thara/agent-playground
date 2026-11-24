package main

import "context"

type Agent interface {
	HandlePrompt(ctx context.Context, prompt string) error
}

type nopAgent struct{}

func (a nopAgent) HandlePrompt(ctx context.Context, prompt string) error { return nil }
