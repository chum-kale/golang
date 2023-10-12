package controllers

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/gofiber/fiber/v2"
)

type PolicyController struct {
}

func NewPolicyController() *PolicyController {
	return &PolicyController{}
}

// @Summary Get a list of policies
// @ID get-policies
// @Produce json
// @Success 200 {array} PolicyInfo
// @Router /api/policies [get]
func (c *PolicyController) GetPolicies(ctx *fiber.Ctx) error {
	policies, err := getPolicies()
	if err != nil {
		fmt.Println("Error listing policies:", err)
		return err
	}

	return ctx.JSON(policies)
}

func getPolicies() ([]PolicyInfo, error) {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})

	if err != nil {
		return nil, err
	}

	// Create an IAM service client.
	svc := iam.New(sess)

	// List IAM policies.
	input := &iam.ListPoliciesInput{}
	result, err := svc.ListPolicies(input)

	if err != nil {
		return nil, err
	}

	policyInfoList := make([]PolicyInfo, 0)
	for _, policy := range result.Policies {
		policyInfo := PolicyInfo{
			Name: *policy.PolicyName,
			Arn:  *policy.Arn,
		}
		policyInfoList = append(policyInfoList, policyInfo)
	}

	return policyInfoList, nil
}

type PolicyInfo struct {
	Name string `json:"name"`
	Arn  string `json:"arn"`
}
