package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	awsLambda "github.com/aws/aws-sdk-go/service/lambda"
)

type Request struct {
	Type string `json:"type"`
}

type Response struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, request Request) (Response, error) {
	svc := awsLambda.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
	})))

	var functionName string
	switch request.Type {
	case "type1":
		functionName = "trial-worker-one"
	case "type2":
		functionName = "trial-worker-two"
	default:
		return Response{Message: "Invalid request type"}, nil
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return Response{}, err
	}

	result, err := svc.Invoke(&awsLambda.InvokeInput{
		FunctionName: aws.String(functionName),
		Payload:      payload,
	})
	if err != nil {
		return Response{}, err
	}

	var response Response
	err = json.Unmarshal(result.Payload, &response)
	if err != nil {
		return Response{}, err
	}

	return response, nil
}

func main() {
	lambda.Start(handler)
}
