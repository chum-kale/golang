package controllers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/gofiber/fiber/v2"
)

type RoleController struct {
}

func NewRoleController() *RoleController {
	return &RoleController{}
}

// @Summary List roles
// @ID list-roles
// @Produce json
// @Param maxRoles path integer true "Maximum number of roles to retrieve"
// @Success 200 {array} string
// @Router /api/roles/{maxRoles} [get]
func (c *RoleController) ListRoles(ctx *fiber.Ctx) error {
	maxRolesStr := ctx.Params("maxRoles")
	maxRoles, err := strconv.Atoi(maxRolesStr)
	if err != nil {
		fmt.Println("Error parsing maxRoles:", err)
		return err
	}

	// Create an AWS session with shared configuration
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create an IAM client
	iamClient := iam.New(sess)

	// List the roles
	roles, err := listRoles(iamClient, maxRoles)
	if err != nil {
		fmt.Println("Error listing roles:", err)
		return err
	}

	return ctx.JSON(roles)
}

func listRoles(iamClient *iam.IAM, maxRoles int) ([]string, error) {
	params := &iam.ListRolesInput{
		MaxItems: aws.Int64(int64(maxRoles)),
	}

	result, err := iamClient.ListRoles(params)
	if err != nil {
		return nil, err
	}

	roleNames := make([]string, 0)
	for _, role := range result.Roles {
		roleNames = append(roleNames, *role.RoleName)
	}

	return roleNames, nil
}
