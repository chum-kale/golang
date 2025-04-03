package main

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gopkg.in/yaml.v2"
)

// Define a random struct to be populated
type DbConfig struct {
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

func getConfigFromS3(region, bucket, item string) (*DbConfig, error) {
	// Create a context
	ctx := context.TODO()

	// Load the AWS configuration
	conf, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	// Create an S3 service client
	svc := s3.NewFromConfig(conf)

	// Get the object from S3
	result, err := svc.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(item),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get object %s from bucket %s: %w", item, bucket, err)
	}
	defer result.Body.Close()

	// Read the content of the object
	var buf bytes.Buffer
	_, err = buf.ReadFrom(result.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read object body: %w", err)
	}
	body := buf.String()

	// Unmarshal the YAML file into the struct
	var cfg DbConfig
	err = yaml.Unmarshal([]byte(body), &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return &cfg, nil
}

func main() {
	// Define the S3 parameters
	region := "region_name"
	bucket := "bucket_name"
	item := "item_name"

	// Get the configuration from S3
	cfg, err := getConfigFromS3(region, bucket, item)
	if err != nil {
		log.Fatalf("Error getting config from S3: %v", err)
	}

	// Print the retrieved configuration
	fmt.Printf("Database: %s\n", cfg.Database)
	fmt.Printf("User: %s\n", cfg.User)
	fmt.Printf("Password: %s\n", cfg.Password)
	fmt.Printf("Host: %s\n", cfg.Host)
	fmt.Printf("Port: %d\n", cfg.Port)
}
