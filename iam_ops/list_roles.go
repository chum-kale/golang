package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

type RoleWrapper struct {
	IamClient *iam.IAM
}

func (wrapper *RoleWrapper) ListRoles(maxRoles int) ([]*iam.Role, error) {
	params := &iam.ListRolesInput{
		MaxItems: aws.Int64(int64(maxRoles)),
	}

	result, err := wrapper.IamClient.ListRoles(params)
	if err != nil {
		return nil, err
	}

	return result.Roles, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./program_name <maxRoles>")
		return
	}

	// Parse the maximum number of roles to list from command line arguments
	maxRoles, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error parsing maxRoles:", err)
		return
	}

	// Create an AWS session with shared configuration
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create an IAM client
	iamClient := iam.New(sess)

	// Create a RoleWrapper object
	roleWrapper := &RoleWrapper{IamClient: iamClient}

	// List the roles
	roles, err := roleWrapper.ListRoles(maxRoles)
	if err != nil {
		fmt.Println("Error listing roles:", err)
		return
	}

	fmt.Println("Roles:")
	for _, role := range roles {
		fmt.Println(*role.RoleName)
	}
}
