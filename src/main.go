package main

import (
	"context"
	"handler/proverbs"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	jst *time.Location
	// invokedTime time.Time
)

type Response struct {
	proverbs.Proverb `json:"proverb"`
	Timestamp        time.Time `json:"timestamp"`
}

func handler(ctx context.Context) (Response, error) {
	invokedTime := time.Now().In(jst)
	return Response{
		Proverb:   proverbs.GetRandomProverb(),
		Timestamp: invokedTime,
	}, nil
}

func main() {
	jst = time.FixedZone("Asia/Tokyo", 9*60*60)
	// invokedTime = time.Now().In(jst)

	lambda.Start(handler)
}
