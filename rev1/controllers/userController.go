package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

type UserResponse struct {
	// User's username
	Username string `json:"username"`

	// User's ID
	UserID string `json:"userId"`

	// User's ARN
	ARN string `json:"arn"`

	// User's path
	Path string `json:"path"`

	// Date when the user's password was last used
	PasswordLastUsed string `json:"passwordLastUsed"`

	// Tags associated with the user (if applicable)
	Tags []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"tags,omitempty"`
}

// Define the UserListResponse type
type UserListResponse struct {
	Users []struct {
		Username   string `json:"username"`
		CreateDate string `json:"createDate"`
	} `json:"users"`
}

// Define the CreateUserRequest type
type CreateUserRequest struct {
	Username  string `json:"username"`
	GroupName string `json:"group"`
}

// Define the SuccessResponse type
type SuccessResponse struct {
	Message string `json:"message"`
}

// Define the ErrorResponse type
type ErrorResponse struct {
	Message string `json:"message"`
}

func NewUserController() *UserController {
	return &UserController{}
}

// @Summary Get a user by username
// @ID get-user-by-username
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} UserResponse
// @Router /api/user/{username} [get]
func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	username := ctx.Params("username")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})

	if err != nil {
		fmt.Println("Error creating AWS session:", err)
		return err
	}

	svc := iam.New(sess)

	userInfo, err := svc.GetUser(&iam.GetUserInput{
		UserName: &username,
	})

	if err != nil {
		awsErr, isAWSErr := err.(awserr.Error)
		if isAWSErr && awsErr.Code() == iam.ErrCodeNoSuchEntityException {
			_, err := svc.CreateUser(&iam.CreateUserInput{
				UserName: &username,
			})

			if err != nil {
				fmt.Println("CreateUser Error:", err)
				return err
			}

			return ctx.JSON(fiber.Map{"message": "User created:" + username})
		} else {
			fmt.Println("GetUser Error:", err)
			return err
		}
	} else {
		userInfoResponse := fiber.Map{
			"Username":         *userInfo.User.UserName,
			"UserID":           *userInfo.User.UserId,
			"ARN":              *userInfo.User.Arn,
			"Path":             *userInfo.User.Path,
			"PasswordLastUsed": userInfo.User.PasswordLastUsed,
		}

		// Check if Tags exist, and add them to the response if present
		if userInfo.User.Tags != nil {
			userInfoResponse["Tags"] = make([]fiber.Map, len(userInfo.User.Tags))

			// Add tags to the "Tags" array in userInfoResponse
			for i, tag := range userInfo.User.Tags {
				userInfoResponse["Tags"].([]fiber.Map)[i] = fiber.Map{
					"Key":   *tag.Key,
					"Value": *tag.Value,
				}
			}
		}

		return ctx.JSON(userInfoResponse)
	}
}

// @Summary List users
// @ID list-users
// @Produce json
// @Success 200 {array} UserListResponse
// @Router /api/users [get]
func (c *UserController) ListUsers(ctx *fiber.Ctx) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})

	if err != nil {
		fmt.Println("Error creating AWS session:", err)
		return err
	}

	svc := iam.New(sess)

	result, err := svc.ListUsers(&iam.ListUsersInput{
		MaxItems: aws.Int64(10),
	})

	if err != nil {
		fmt.Println("Error listing users:", err)
		return err
	}

	users := make([]fiber.Map, len(result.Users))
	for i, user := range result.Users {
		if user != nil {
			users[i] = fiber.Map{
				"Username":   *user.UserName,
				"CreateDate": user.CreateDate.String(),
			}
		}
	}

	return ctx.JSON(users)
}

// @Summary Create a user
// @ID create-user
// @Accept json
// @Param request body CreateUserRequest true "Create User Request"
// @Success 201 {object} SuccessResponse
// @Router /api/user [post]
func (c *UserController) CreateUser(ctx *fiber.Ctx) error {

	var createUserRequest struct {
		Username  string `json:"username"`
		GroupName string `json:"group"`
	}

	body := ctx.Body()

	if err := json.Unmarshal(body, &createUserRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	username := createUserRequest.Username
	groupName := createUserRequest.GroupName

	if username == "" || groupName == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username and group name cannot be empty",
		})
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})

	if err != nil {
		fmt.Println("Error creating AWS session:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating AWS session",
		})
	}

	svc := iam.New(sess)

	_, err = svc.GetUser(&iam.GetUserInput{
		UserName: &username,
	})

	if err == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "IAM user with username '" + username + "' already exists",
		})
	} else if awsErr, isAWSErr := err.(awserr.Error); isAWSErr && awsErr.Code() != iam.ErrCodeNoSuchEntityException {
		fmt.Println("Error checking if user exists:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error checking IAM user existence",
		})
	}

	if !groupExists(svc, groupName) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "IAM group '" + groupName + "' does not exist",
		})
	}

	createUser(svc, username)
	addUserToGroup(svc, username, groupName)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "IAM user '" + username + "' created and added to group '" + groupName + "' successfully",
	})
}

// @Summary Check if a user exists
// @ID check-user-exists
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} SuccessResponse
// @Success 404 {object} ErrorResponse
// @Router /api/user/{username}/exists [get]
func (c *UserController) CheckUserExists(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating AWS session",
		})
	}

	svc := iam.New(sess)
	input := &iam.GetUserInput{
		UserName: aws.String(username),
	}
	_, err = svc.GetUser(input)

	if err != nil {
		awsErr, isAWSErr := err.(awserr.Error)
		if isAWSErr && awsErr.Code() == iam.ErrCodeNoSuchEntityException {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "IAM user with username '" + username + "' does not exist",
			})
		} else {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error checking IAM user existence",
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "IAM user with username '" + username + "' exists",
	})
}

// @Summary Check if a group exists
// @ID check-group-exists
// @Produce json
// @Param group path string true "Group Name"
// @Success 200 {object} SuccessResponse
// @Success 404 {object} ErrorResponse
// @Router /api/group/{group}/exists [get]
func (c *UserController) CheckGroupExists(ctx *fiber.Ctx) error {
	groupName := ctx.Params("group")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating AWS session",
		})
	}

	svc := iam.New(sess)
	if groupExists(svc, groupName) {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "IAM group '" + groupName + "' exists",
		})
	}
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "IAM group '" + groupName + "' does not exist",
	})
}

func createUser(svc *iam.IAM, username string) {
	createUserInput := &iam.CreateUserInput{
		UserName: aws.String(username),
	}

	_, err := svc.CreateUser(createUserInput)
	if err != nil {
		fmt.Println("Error creating IAM user:", err)
		return
	}

	fmt.Printf("IAM user '%s' created successfully.\n", username)
}

func addUserToGroup(svc *iam.IAM, username, groupName string) {
	addUserToGroupInput := &iam.AddUserToGroupInput{
		GroupName: aws.String(groupName),
		UserName:  aws.String(username),
	}

	_, err := svc.AddUserToGroup(addUserToGroupInput)
	if err != nil {
		fmt.Println("Error adding user to group:", err)
		return
	}

	fmt.Printf("IAM user '%s' added to group '%s' successfully.\n", username, groupName)
}
