package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

const TIMEOUT = 10

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

type QA struct {
	Good      int
	Bad       int
	Counter   int
	Questions []string
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

func (qa *QA) makeQuestion() string {
	if len(qa.Questions) < qa.Counter+1 {
		qa.Counter = 0
	} else {
		qa.Counter++
	}

	return qa.Questions[qa.Counter]
}

func (qa *QA) checkAnswer(question, answer string) bool {
	if question == answer {
		qa.Good++
		return true
	} else {
		qa.Bad++
		return false
	}
}

func NewQA(questions []string) *QA {
	return &QA{
		Good:      0,
		Bad:       0,
		Counter:   -1,
		Questions: questions,
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()

	qa := NewQA(words)
	ch := read(os.Stdin)
LOOP:
	for {
		question := qa.makeQuestion()
		fmt.Printf("%s > ", question)

		select {
		case answer, ok := <-ch:
			if !ok {
				break LOOP
			}
			if qa.checkAnswer(question, answer) {
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
