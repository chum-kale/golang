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

type Policy struct {
	Name           string `json:"name"`
	Arn            string `json:"arn"`
	Description    string `json:"description"`
	Path           string `json:"path"`
	DefaultVersion string `json:"default_version"`
}

// GetPolicies retrieves a list of IAM policies.
// @Summary Get IAM policies
// @Description Retrieves a list of IAM policies.
// @Tags Policies
// @Accept json
// @Produce json
// @Success 200 {array} PolicyInfo "List of IAM policies"
// @Failure 500 {object} ErrorResponse "Internal server error"
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
