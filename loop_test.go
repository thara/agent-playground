package main

import (
	"context"
	"io"
	"slices"
	"strings"
	"testing"
	"time"
)

type recordingAgent struct {
	prompts []string
}

func (a *recordingAgent) HandlePrompt(ctx context.Context, prompt string) error {
	a.prompts = append(a.prompts, prompt)
	return nil
}

func TestRunLoop_PassesNonEmptyPromptsInOrderUntilEOF(t *testing.T) {
	agent := &recordingAgent{}
	in := strings.NewReader("a\n\n   \n b \n")

	err := runLoop(t.Context(), in, io.Discard, agent)

	if err != nil {
		t.Fatalf("runLoop returned error: %v", err)
	}
	want := []string{"a", "b"}
	if !slices.Equal(agent.prompts, want) {
		t.Errorf("prompts = %q, want %q", agent.prompts, want)
	}
}

func TestRunLoop_ReturnsNilWhenCanceledWhileWaitingForInput(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	pr, pw := io.Pipe()
	t.Cleanup(func() { pw.Close() })

	done := make(chan error, 1)
	go func() {
		done <- runLoop(ctx, pr, io.Discard, &recordingAgent{})
	}()

	cancel()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("runLoop returned error: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("runLoop did not return after cancel")
	}
}
