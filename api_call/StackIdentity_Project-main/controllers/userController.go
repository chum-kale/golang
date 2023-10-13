// package controllers

// import (
// 	"fmt"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/awserr"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/iam"
// 	"github.com/gofiber/fiber/v2"
// )

// type UserController struct {
// }

// func NewUserController() *UserController {
// 	return &UserController{}
// }

// func (c *UserController) GetUser(ctx *fiber.Ctx) error {
// 	username := ctx.Params("username")

// 	// Initialize a session in us-west-2 that the SDK will use to load
// 	// credentials from the shared credentials file ~/.aws/credentials.
// 	sess, err := session.NewSession(&aws.Config{
// 		Region: aws.String("us-west-2"),
// 	})

// 	if err != nil {
// 		fmt.Println("Error creating AWS session:", err)
// 		return err
// 	}

// 	// Create an IAM service client.
// 	svc := iam.New(sess)

// 	// Get information about the user.
// 	userInfo, err := svc.GetUser(&iam.GetUserInput{
// 		UserName: &username,
// 	})

// 	if err != nil {
// 		awsErr, isAWSErr := err.(awserr.Error)
// 		if isAWSErr && awsErr.Code() == iam.ErrCodeNoSuchEntityException {
// 			_, err := svc.CreateUser(&iam.CreateUserInput{
// 				UserName: &username,
// 			})

// 			if err != nil {
// 				fmt.Println("CreateUser Error:", err)
// 				return err
// 			}

// 			return ctx.JSON(fiber.Map{"message": "User created:" + username})
// 		} else {
// 			fmt.Println("GetUser Error:", err)
// 			return err
// 		}
// 	} else {
// 		// Build the user information response
// 		// userInfoResponse := fiber.Map{
// 		// 	"Username":         *userInfo.User.UserName,
// 		// 	"UserID":           *userInfo.User.UserId,
// 		// 	"ARN":              *userInfo.User.Arn,
// 		// 	"Path":             *userInfo.User.Path,
// 		// 	"PasswordLastUsed": userInfo.User.PasswordLastUsed,
// 		// 	"Tags":             make([]fiber.Map, len(userInfo.User.Tags)),
// 		// 	// Add more user attributes here as needed
// 		// }
// 		userInfoResponse := fiber.Map{
// 			"Username":         *userInfo.User.UserName,
// 			"UserID":           *userInfo.User.UserId,
// 			"ARN":              *userInfo.User.Arn,
// 			"Path":             *userInfo.User.Path,
// 			"PasswordLastUsed": userInfo.User.PasswordLastUsed,
// 		}

// 		// Check if Tags exist, and add them to the response if present
// 		if userInfo.User.Tags != nil {
// 			userInfoResponse["Tags"] = make([]fiber.Map, len(userInfo.User.Tags))

// 			// Add tags to the "Tags" array in userInfoResponse
// 			for i, tag := range userInfo.User.Tags {
// 				userInfoResponse["Tags"].([]fiber.Map)[i] = fiber.Map{
// 					"Key":   *tag.Key,
// 					"Value": *tag.Value,
// 				}
// 			}
// 		}

// 		// Fetch and print policy names associated with the user.
// 		// fmt.Println("Policy Names:")
// 		// Add code to get and append policy names to userInfoResponse

// 		return ctx.JSON(userInfoResponse)
// 	}
// }

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

func NewUserController() *UserController {
	return &UserController{}
}

type UserResponse struct {
	Username         string `json:"username"`
	UserID           string `json:"userID"`
	ARN              string `json:"arn"`
	Path             string `json:"path"`
	PasswordLastUsed string `json:"passwordLastUsed"`
	Tags             []Tag  `json:"tags,omitempty"`
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// GetUser gets user information by username.
// @Summary Get user information by username
// @Description Retrieves user details by username.
// @Tags Users
// @Accept json
// @Produce json
// @Param username path string true "Username of the user"
// @Success 200 {object} UserResponse "User information"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/user/{username} [get]
func (c *UserController) GetUser(ctx *fiber.Ctx) error {

	username := ctx.Params("username")
	fmt.Println("GetUser method called with username:", username)

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

type User struct {
	Username   string `json:"username"`
	CreateDate string `json:"createDate"`
}

type ListUsersResponse struct {
	Users []User `json:"users"`
}

// ListUsers lists users.
// @Summary List all users
// @Description Retrieves a list of all users.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array}  "List of users"
// @Failure 500 {object}  "Internal server error"
// @Router /users [get]

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

// func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
// 	// Parse user details from request body
// 	var createUserRequest struct {
// 		Username  string `json:"username"`
// 		GroupName string `json:"group"`
// 	}

// 	if err := ctx.BodyParser(&createUserRequest); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Invalid request body",
// 		})
// 	}

// 	username := createUserRequest.Username
// 	groupName := createUserRequest.GroupName

// 	if username == "" || groupName == "" {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Username and group name cannot be empty",
// 		})
// 	}

// 	sess, err := session.NewSession(&aws.Config{
// 		Region: aws.String("us-west-2"),
// 	})

// 	if err != nil {
// 		fmt.Println("Error creating AWS session:", err)
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Error creating AWS session",
// 		})
// 	}

// 	svc := iam.New(sess)

// 	if userExists(svc, username) {
// 		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
// 			"message": "IAM user with username '" + username + "' already exists",
// 		})
// 	}

// 	if !groupExists(svc, groupName) {
// 		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 			"message": "IAM group '" + groupName + "' does not exist",
// 		})
// 	}

// 	createUser(svc, username)
// 	addUserToGroup(svc, username, groupName)

// 	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"message": "IAM user '" + username + "' created and added to group '" + groupName + "' successfully",
// 	})
// }

// ListUsers lists users.
// @Summary List all users
// @Description Retrieves a list of all users.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array}  "List of users"
// @Failure 500 {object}  "Internal server error"
// @Router /users [get]

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	// Parse user details from the request body
	var createUserRequest struct {
		Username  string `json:"username"`
		GroupName string `json:"group"`
	}

	// Read the request body
	body := ctx.Body()

	// Print or log the received request body
	fmt.Println("Received Request Body:", string(body))

	// if err := ctx.BodyParser(&createUserRequest); err != nil {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "Invalid request body",
	// 	})
	// }

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

func userExists(svc *iam.IAM, username string) bool {
	listUsersInput := &iam.ListUsersInput{}

	err := svc.ListUsersPages(listUsersInput, func(output *iam.ListUsersOutput, lastPage bool) bool {
		for _, user := range output.Users {
			if *user.UserName == username {
				return true // User already exists
			}
		}
		return !lastPage
	})

	if err != nil {
		fmt.Println("Error checking if user exists:", err)
		return false
	}

	return false // User does not exist
}

func groupExists(svc *iam.IAM, groupName string) bool {
	listGroupsInput := &iam.ListGroupsInput{}

	err := svc.ListGroupsPages(listGroupsInput, func(output *iam.ListGroupsOutput, lastPage bool) bool {
		for _, group := range output.Groups {
			if *group.GroupName == groupName {
				return true // Group already exists
			}
		}
		return !lastPage
	})

	if err != nil {
		fmt.Println("Error checking if group exists:", err)
		return false
	}

	return true // Group exists
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

// func (c *UserController) CheckUserExists(ctx *fiber.Ctx) error {
// 	username := ctx.Params("username")
// 	sess, err := session.NewSession(&aws.Config{
// 		Region: aws.String("us-west-2"),
// 	})

// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Error creating AWS session",
// 		})
// 	}

// 	svc := iam.New(sess)
// 	if userExists(svc, username) {
// 		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
// 			"message": "IAM user with username '" + username + "' exists",
// 		})
// 	}
// 	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 		"message": "IAM user with username '" + username + "' does not exist",
// 	})
// }

type CheckUserExistsRequest struct {
	Username string `json:"username"`
}

type CheckUserExistsResponse struct {
	Message string `json:"message"`
}

// CheckUserExists checks if a user exists by username.
// @Summary Check if a user exists
// @Description Checks if a user exists by username.
// @Tags Users
// @Accept json
// @Produce json
// @Param username path string true "Username of the user"
// @Success 200 {object}  "User exists"
// @Failure 404 {object}  "User not found"
// @Failure 500 {object}  "Internal server error"
// @Router /users/check/{username} [get]

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
	_, err = svc.GetUser(&iam.GetUserInput{
		UserName: &username,
	})

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

/*
// CheckUserExists checks if a user exists by username.
// @Summary Check if a user exists
// @Description Checks if a user exists by username.
// @Tags Users
// @Accept json
// @Produce json
// @Param username path string true "Username of the user"
// @Success 200 {object}  "User exists"
// @Failure 404 {object}  "User not found"
// @Failure 500 {object}  "Internal server error"
// @Router /users/check/{username} [get]
*/
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
