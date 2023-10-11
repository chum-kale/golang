package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	svc := iam.New(sess)

	result, err := svc.ListUsers(&iam.ListUsersInput{
		MaxItems: aws.Int64(10),
	})

	if err != nil {
		fmt.Println("Error listing users:", err)
		return
	}

	for i, user := range result.Users {
		if user == nil {
			continue
		}
		fmt.Printf("%d user %s created %v\n", i, *user.UserName, user.CreateDate)
	}
}
