package main

//used to store and exchange info across services and access them in different servicess

import (
	"context"
	"fmt"
)

// Define keys for the values we want to store in the context
type contextKey string

const (
	requestIDKey contextKey = "requestID"
	userIDKey    contextKey = "userID"
)

// Function that fetches values from the context
func processRequest(ctx context.Context) {
	// Retrieve values from the context using the keys
	requestID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		requestID = "unknown"
	}
	userID, ok := ctx.Value(userIDKey).(int)
	if !ok {
		userID = 0
	}

	fmt.Printf("Processing request %s for user %d\n", requestID, userID)
}

func main() {
	// Create a background context
	ctx := context.Background()

	// Attach values to the context using WithValue
	ctx = context.WithValue(ctx, requestIDKey, "abc123")
	ctx = context.WithValue(ctx, userIDKey, 42)

	// Pass the context with values to the function
	processRequest(ctx)
}
