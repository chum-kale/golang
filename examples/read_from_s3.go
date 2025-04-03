package main

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"log"
)

// reads config file from s3
func readConfigS3() (config *infra.DbConfig, err error) {
	bucket := "oceano-config"
	item := "albservice.yaml"

	//creating an aws session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	//create service client
	svc := s3.New(sess)

	//gwet file from s3
	result, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(item),
	})
	if err != nil {
		log.Fatalf("Failed to get object: %v", err)
	}
	defer result.Body.Close()

	//reading content
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(result.Body)
	if err != nil {
		log.Fatalf("Failed to read object content: %v", err)
	}

}
