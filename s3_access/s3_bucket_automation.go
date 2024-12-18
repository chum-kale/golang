package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func createS3Bucket(bucketName string) error {
	// Load the AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	// Create an S3 client
	client := s3.NewFromConfig(cfg)

	// Create the S3 bucket
	_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: &bucketName,
	})
	if err != nil {
		return fmt.Errorf("failed to create bucket: %w", err)
	}

	fmt.Printf("Bucket %q successfully created.\n", bucketName)
	return nil
}

func main() {
	// Check if a bucket name is provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <bucket-name>")
		os.Exit(1)
	}

	bucketName := os.Args[1]

	// Call the function to create the bucket
	if err := createS3Bucket(bucketName); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
