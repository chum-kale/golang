package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Custom claims structure
type CustomClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Secret key (should be securely stored)
var secretKey = []byte("your-256-bit-secret")

// Separate function to verify the signing method
func verifySigningMethod(token *jwt.Token) (interface{}, error) {
	// Check if the signing method is what we expect
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	// Return the secret key for verification
	return secretKey, nil
}

// Function to create a JWT token
func createToken(userID int, username string) (string, error) {
	// Create claims
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// Set token expiration
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			// Set issued at time
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// Optional issuer
			Issuer: "MyApp",
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate the signed token string
	return token.SignedString(secretKey)
}

// Function to validate and parse token
func validateToken(tokenString string) (*CustomClaims, error) {
	// Parse the token using the separate verification function
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, verifySigningMethod)
	if err != nil {
		return nil, err
	}

	// Extract and type assert claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Verify user credentials
	if !validateCredentials(username, password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT
	token, err := createToken(user.ID, user.Username)
	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	// Send token to client
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// Extract token from Authorization header
	tokenString := extractTokenFromHeader(r)

	// Validate token
	claims, err := validateToken(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Access user information from claims
	userID := claims.UserID
	// Proceed with protected operation
}

func main() {
	// Simulate user authentication
	userID := 123
	username := "johndoe"

	// Create a token
	tokenString, err := createToken(userID, username)
	if err != nil {
		fmt.Println("Error creating token:", err)
		return
	}

	fmt.Println("Generated Token:", tokenString)

	// Validate the token
	claims, err := validateToken(tokenString)
	if err != nil {
		fmt.Println("Token validation failed:", err)
		return
	}

	// Access token claims
	fmt.Printf("User ID: %d\n", claims.UserID)
	fmt.Printf("Username: %s\n", claims.Username)
	fmt.Printf("Issued At: %v\n", claims.IssuedAt)
	fmt.Printf("Expires At: %v\n", claims.ExpiresAt)
}
