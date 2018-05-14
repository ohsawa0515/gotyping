package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

const TIMEOUT = 5

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

func makeWordTest(n int) string {
	if len(words) < n {
		return words[n]
	}

	return words[0]
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	n := 0

	ch := read(os.Stdin)
LOOP:
	for {
		word := makeWordTest(n)

		select {
		case s, ok := <-ch:
			if !ok {
				break LOOP
			}
			fmt.Println(word, s)
		case <-ctx.Done():
			fmt.Println("Timeout")
			break LOOP
		}
	}
}
