package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/openai/openai-go/v3"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if os.Getenv("OPENAI_API_KEY") == "" {
		fmt.Fprintln(os.Stderr, "error: OPENAI_API_KEY is not set")
		os.Exit(1)
	}

	model := os.Getenv("OPENAI_MODEL")
	if model == "" {
		model = openai.ChatModelGPT5_4Mini
	}

	agent := &openAIAgent{client: openai.NewClient(), model: model}

	if err := runLoop(ctx, os.Stdin, os.Stdout, agent); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
