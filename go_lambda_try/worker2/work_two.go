package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Type string `json:"type"`
}

type Response struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, request Request) (Response, error) {
	return Response{Message: "Handled request of type 2"}, nil
}

func main() {
	lambda.Start(handler)
}
