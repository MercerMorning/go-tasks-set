package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type Request struct {
	Payload string
}

type Client interface {
	SendRequest(ctx context.Context, request Request) error
	WithLimiter(ctx context.Context, requests []Request)
}

type client struct {
}

func (c client) SendRequest(ctx context.Context, request Request) error {
	time.Sleep(100 * time.Microsecond)
	fmt.Println("sending request", request.Payload)
	return nil
}

func (c client) WithLimiter(ctx context.Context, requests []Request) {
}

func main() {
	ctx := context.Background()
	c := client{}
	requests := make([]Request, 1000)
	for i := 0; i < 1000; i++ {
		requests[i] = Request{Payload: strconv.Itoa(i)}
	}
	c.WithLimiter(ctx, requests)
}
