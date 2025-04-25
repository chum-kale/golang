package main

import (
	"log"

	"github.com/chum-kale/golang/snowflake/snowops"
)

func main() {
	// Step 1: Refresh token to get access token
	accessToken, err := snowops.RefreshAccessToken()
	if err != nil {
		log.Fatalf("Failed to refresh access token: %v", err)
	}

	// Step 2: Connect to Snowflake with the access token
	db, err := snowops.ConnectToSnowflake(accessToken)
	if err != nil {
		log.Fatalf("Failed to connect to Snowflake: %v", err)
	}
	defer db.Close()

	// Step 3: Call your list functions
	snowops.CreateTextFile()
	println()

	snowops.GetUsers(db)
	println()

	snowops.ListQueryHistory(db, "BIRD_DATA", "query_list.txt")
	println()

	snowops.ListDatabases(db)
	println()

	snowops.ListSchemas(db, "BIRD_DATA")
	println()

	snowops.ListTables(db, "BIRD_DATA", "PUBLIC")
}
