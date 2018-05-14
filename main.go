package main

import (
	"context"
	"time"

	"github.com/ohsawa0515/gotyping/typing"
)

const TIMEOUT = 10

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()

	typing.Run(ctx)
}
