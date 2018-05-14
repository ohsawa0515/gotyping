package typing

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
)

var words = []string{
	"advertisement",
	"lid",
	"southeast",
	"perish",
	"inhabit",
	"extent",
	"room",
	"balance",
	"onto",
	"breeze",
	"protein",
	"genetic",
	"setup",
	"kindergarten",
	"satisfactory",
	"appetite",
	"civilization",
	"bookcase",
	"galaxy",
	"suburban",
	"infectious",
	"jerk",
}

func read(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()

	return ch
}

// Run starts typing game.
func Run(ctx context.Context) {
	qa := NewQA(words)
	ch := read(os.Stdin)
LOOP:
	for {
		question := qa.MakeQuestion()
		fmt.Printf("%s > ", question)

		select {
		case answer, ok := <-ch:
			if !ok {
				break LOOP
			}
			if qa.CheckAnswer(question, answer) {
				fmt.Println("✓ That's right!")
			} else {
				fmt.Println("✗ Oh, bad")
			}
		case <-ctx.Done():
			fmt.Println("Timeout")
			break LOOP
		}
	}

	fmt.Printf("Good: %d\n", qa.Good)
	fmt.Printf("Bad: %d\n", qa.Bad)
}
