package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"
)

type Question struct {
	Text    string
	Options []string
	Answer  string
}

func runQuiz(ctx context.Context, questions []Question) int {
	score := 0
	reader := bufio.NewReader(os.Stdin)

	for _, question := range questions {
		// check if context is cancelled before each question
		select {
		case <-ctx.Done():
			fmt.Println("\nTime's up!")
			return score
		default:
			// context not cancelled, continue
		}
		fmt.Println("\nQuestion: ", question.Text)
		fmt.Println("Options ")
		for _, option := range question.Options {
			fmt.Println(option)
		}
		fmt.Print("Your answer: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == question.Answer {
			fmt.Println("Correct!")
			score += 1
		} else {
			fmt.Println("Wrong! Correct answer was: " + question.Answer)
		}
	}
	return score
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	questions := []Question{
		{
			Text:    "What is used to declare a variable in Go?",
			Options: []string{"a.var", "b.let", "c.integer", "d.declare"},
			Answer:  "a",
		},
		{
			Text:    "What is used to define a function in Go?",
			Options: []string{"a.var", "b.func", "c.integer", "d.declare"},
			Answer:  "b",
		},
	}

	score := 0
	score = runQuiz(ctx, questions)
	fmt.Printf("\nScore: %d/%d\n", score, len(questions))
}
