package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func Handler(ctx context.Context, req Request) (Response, error) {
	message := fmt.Sprintf("Hello, %s! Welcome to AWS Lambda with Go.", req.Name)
	response := Response{
		Message: message,
	}
	return response, nil
}

func main() {
	lambda.Start(Handler)
}
