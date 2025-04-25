package snowops

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	_ "github.com/snowflakedb/gosnowflake"
)

// OAuth Credentials: Put them in an env file later
const (
	snowflakeTokenURL = "https://<user-id>.snowflakecomputing.com/oauth/token-request"
	clientID          = "clid"   // Replace with your actual Client ID
	clientSecret      = "clisec" // Replace with your actual Client Secret
	refreshToken      = "reftok" // Store securely (e.g., in a config file or environment variable)
)

// Struct to Parse Response
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`    // Expiry time in seconds
	RefreshToken string `json:"refresh_token"` // New refresh token (if rotated)
}

// Function to Refresh Token
func RefreshAccessToken() (string, error) {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", snowflakeTokenURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the response
	var tokenRes TokenResponse
	if err := json.Unmarshal(body, &tokenRes); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Optionally store the new refresh token if Snowflake rotates it
	fmt.Println("New Access Token:", tokenRes.AccessToken)
	fmt.Println("Token Expires In:", tokenRes.ExpiresIn, "seconds")

	return tokenRes.AccessToken, nil
}

// Connect to Snowflake and return DB handle
func ConnectToSnowflake(accessToken string) (*sql.DB, error) {
	account := "userid"
	user := "username"
	database := "dbname"
	schema := "PUBLIC"

	escapedToken := url.QueryEscape(accessToken)

	dsn := fmt.Sprintf("%s@%s/%s/%s?authenticator=OAUTH&token=%s",
		user, account, database, schema, escapedToken)

	db, err := sql.Open("snowflake", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping Snowflake: %w", err)
	}

	fmt.Println("Connected to Snowflake successfully!")
	return db, nil
}

// Function to only create table from csv header
func CreateTablefromCSV(db *sql.DB, csvFilePath, tableName string, uniqueColumns []string) error {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV: %v", err)
	}
	if len(records) < 2 {
		return fmt.Errorf("CSV must have at least a header and one data row")
	}

	headers := records[0]

	// Create table with unique constraint if specified
	var colsDef []string
	for _, col := range headers {
		colsDef = append(colsDef, fmt.Sprintf("\"%s\" VARCHAR(500)", col))
	}

	createStmt := fmt.Sprintf(`CREATE OR REPLACE TABLE %s (%s`, tableName, strings.Join(colsDef, ", "))

	// Add unique constraint if uniqueColumns are provided
	if len(uniqueColumns) > 0 {
		quotedUniqCols := make([]string, len(uniqueColumns))
		for i, col := range uniqueColumns {
			quotedUniqCols[i] = fmt.Sprintf(`"%s"`, col)
		}
		createStmt += fmt.Sprintf(`, UNIQUE(%s)`, strings.Join(quotedUniqCols, ", "))
	}
	createStmt += ")"

	if _, err = db.Exec(createStmt); err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}
	fmt.Println("✅ Table created successfully!")

	return nil
}

// Function to create table from csv headers
// all columns will be varchar
// It also populates the table from the csv file
func createAndPopulateFromCSV(db *sql.DB, csvFilePath, tableName string, uniqueColumns []string) error {
	const batchSize = 100
	const maxWorkers = 10

	file, err := os.Open(csvFilePath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV: %v", err)
	}
	if len(records) < 2 {
		return fmt.Errorf("CSV must have at least a header and one data row")
	}

	headers := records[0]
	dataRows := records[1:]

	// Create table with unique constraint if specified
	var colsDef []string
	for _, col := range headers {
		colsDef = append(colsDef, fmt.Sprintf("\"%s\" VARCHAR(500)", col))
	}

	createStmt := fmt.Sprintf(`CREATE OR REPLACE TABLE %s (%s`, tableName, strings.Join(colsDef, ", "))

	// Add unique constraint if uniqueColumns are provided
	if len(uniqueColumns) > 0 {
		quotedUniqCols := make([]string, len(uniqueColumns))
		for i, col := range uniqueColumns {
			quotedUniqCols[i] = fmt.Sprintf(`"%s"`, col)
		}
		createStmt += fmt.Sprintf(`, UNIQUE(%s)`, strings.Join(quotedUniqCols, ", "))
	}
	createStmt += ")"

	if _, err = db.Exec(createStmt); err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}
	fmt.Println("✅ Table created successfully!")

	// Table creation ends here

	// Quote headers and prepare placeholders
	quotedHeaders := make([]string, len(headers))
	placeholders := make([]string, len(headers))
	for i, h := range headers {
		h = strings.TrimSpace(h)
		h = strings.ReplaceAll(h, `"`, "")
		quotedHeaders[i] = fmt.Sprintf(`"%s"`, h)
		placeholders[i] = "?"
	}

	// Prepare the insert statement - use ON CONFLICT DO NOTHING for duplicate prevention
	var insertSQL string
	if len(uniqueColumns) > 0 {
		insertSQL = fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES (%s) ON CONFLICT DO NOTHING`,
			tableName,
			strings.Join(quotedHeaders, ","),
			strings.Join(placeholders, ","),
		)
	} else {
		insertSQL = fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES (%s)`,
			tableName,
			strings.Join(quotedHeaders, ","),
			strings.Join(placeholders, ","),
		)
	}

	// Split the data into batches
	var batches [][][]string
	for start := 0; start < len(dataRows); start += batchSize {
		end := start + batchSize
		if end > len(dataRows) {
			end = len(dataRows)
		}
		batches = append(batches, dataRows[start:end])
	}

	// Set up worker pool and communication channels
	type batchResult struct {
		batchIndex int
		err        error
	}

	jobs := make(chan int, len(batches))
	results := make(chan batchResult, len(batches))
	var wg sync.WaitGroup

	// Create worker goroutines
	workerCount := min(maxWorkers, len(batches))
	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for batchIndex := range jobs {
				batch := batches[batchIndex]

				tx, err := db.Begin()
				if err != nil {
					results <- batchResult{batchIndex, fmt.Errorf("worker %d: failed to begin transaction: %v", workerID, err)}
					continue
				}

				stmt, err := tx.Prepare(insertSQL)
				if err != nil {
					tx.Rollback()
					results <- batchResult{batchIndex, fmt.Errorf("worker %d: failed to prepare insert: %v", workerID, err)}
					continue
				}

				for i, row := range batch {
					vals := make([]interface{}, len(row))
					for j, v := range row {
						vals[j] = v
					}
					globalRowIndex := batchIndex*batchSize + i + 1
					fmt.Printf("📥 Worker %d inserting row %d\n", workerID, globalRowIndex)

					if _, err := stmt.Exec(vals...); err != nil {
						stmt.Close()
						tx.Rollback()
						results <- batchResult{batchIndex, fmt.Errorf("worker %d: failed to insert row %d: %v", workerID, globalRowIndex, err)}
						break
					}
				}

				stmt.Close()
				if err := tx.Commit(); err != nil {
					results <- batchResult{batchIndex, fmt.Errorf("worker %d: failed to commit batch %d: %v", workerID, batchIndex+1, err)}
					continue
				}

				start := batchIndex*batchSize + 1
				end := min(start+batchSize-1, len(dataRows))
				fmt.Printf("✅ Worker %d committed batch rows %d–%d\n", workerID, start, end)

				results <- batchResult{batchIndex, nil}
			}
		}(w)
	}

	// Send jobs to the workers
	for i := range batches {
		jobs <- i
	}
	close(jobs)

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	errs := []error{}
	for result := range results {
		if result.err != nil {
			errs = append(errs, result.err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors during parallel insertion: %v", errs)
	}

	fmt.Println("🚀 All data inserted successfully in parallel batches!")
	return nil
}

// Helper function min for Go versions before 1.21
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Fetch and display rows from the BIRD_INFO table
func fetchBirdInfo(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM BIRD_INFO LIMIT 10")
	if err != nil {
		return fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	// Get column names
	cols, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("failed to get column names: %w", err)
	}
	fmt.Println(cols)

	// Iterate over rows
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))

		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return fmt.Errorf("failed to scan row: %w", err)
		}

		for _, col := range columns {
			fmt.Printf("%v\t", col)
		}
		fmt.Println()
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error iterating rows: %w", err)
	}
	return nil
}

// Fetch and display rows from the POKEMON_INFO table, pretty-printed
func fetchPokemonInfo(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM POKEMON_INFO LIMIT 10")
	if err != nil {
		return fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	// Get column names
	cols, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("failed to get column names: %w", err)
	}

	// Store all rows in memory to calculate column widths
	var data [][]string
	data = append(data, cols)

	for rows.Next() {
		columnValues := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columnValues {
			columnPointers[i] = &columnValues[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			return fmt.Errorf("failed to scan row: %w", err)
		}

		row := make([]string, len(cols))
		for i, val := range columnValues {
			if val != nil {
				row[i] = fmt.Sprintf("%v", val)
			} else {
				row[i] = "NULL"
			}
		}
		data = append(data, row)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error iterating rows: %w", err)
	}

	// Calculate max width of each column
	colWidths := make([]int, len(cols))
	for _, row := range data {
		for i, col := range row {
			if len(col) > colWidths[i] {
				colWidths[i] = len(col)
			}
		}
	}

	// Print the table
	printRow := func(row []string) {
		for i, col := range row {
			fmt.Printf("%-*s  ", colWidths[i], col)
		}
		fmt.Println()
	}

	// Print header and separator
	printRow(data[0])
	for i := range data[0] {
		fmt.Printf("%s  ", strings.Repeat("-", colWidths[i]))
	}
	fmt.Println()

	// Print data rows
	for _, row := range data[1:] {
		printRow(row)
	}

	return nil
}

// func main() {
// 	// Get fresh access token
// 	accessToken, err := refreshAccessToken()
// 	if err != nil {
// 		log.Fatalf("Failed to refresh access token: %v", err)
// 	}

// 	// Connect to Snowflake
// 	db, err := connectToSnowflake(accessToken)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Fetch and display bird data
// 	if err := fetchBirdInfo(db); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Create and populate pokemon table.
// 	err = createAndPopulateFromCSV(db, "all_pokemon_data.csv", "POKEMON_INFO", []string{"Name"})
// 	if err != nil {
// 		log.Fatalf("Failed to load CSV into Snowflake: %v", err)
// 	}

// 	if err := fetchPokemonInfo(db); err != nil {
// 		log.Fatal(err)
// 	}

// }
