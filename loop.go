package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"
)

func runLoop(ctx context.Context, in io.Reader, out io.Writer, agent Agent) error {
	lines := make(chan string)
	scanErr := make(chan error, 1)

	go func() {
		scanner := bufio.NewScanner(in)
		defer func() {
			scanErr <- scanner.Err()
			close(lines)
		}()
		for scanner.Scan() {
			select {
			case lines <- scanner.Text():
			case <-ctx.Done():
				return
			}
		}
	}()

	for {
		fmt.Fprint(out, ">> ")

		var line string
		select {
		case <-ctx.Done():
			return nil
		case l, ok := <-lines:
			if !ok {
				if err := <-scanErr; err != nil {
					return fmt.Errorf("read input: %w", err)
				}
				return nil
			}
			line = l
		}

		prompt := strings.TrimSpace(line)
		if prompt == "" {
			continue
		}

		resp, err := agent.HandlePrompt(ctx, prompt)
		if err != nil {
			return fmt.Errorf("handle prompt: %w", err)
		}
		fmt.Fprintln(out, resp)
	}
}
