package snowops

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

func ListDatabases(db *sql.DB) {
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		log.Fatalf("Failed to list databases: %v", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatalf("Failed to retrieve column metadata: %v", err)
	}

	fmt.Println("Available Databases:")

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		var dbName, owner string
		for i, col := range columns {
			switch strings.ToLower(col) {
			case "name":
				dbName = string(values[i])
			case "owner":
				owner = string(values[i])
			}
		}

		if dbName != "" {
			fmt.Printf("- %s (Owner: %s)\n", dbName, owner)
		} else {
			fmt.Println("- <unknown database>")
		}
	}
}

func ListSchemas(db *sql.DB, database string) {
	query := fmt.Sprintf("SHOW SCHEMAS IN DATABASE %s", database)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Failed to list schemas: %v", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatalf("Failed to retrieve column metadata: %v", err)
	}

	fmt.Printf("Schemas in %s:\n", database)

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		var schemaName string
		for i, col := range columns {
			if strings.EqualFold(col, "name") {
				schemaName = string(values[i])
				break
			}
		}

		if schemaName != "" {
			fmt.Printf("- %s\n", schemaName)
		} else {
			fmt.Println("- <unknown schema name>")
		}
	}
}

func ListTables(db *sql.DB, database string, schema string) {
	query := fmt.Sprintf("ch %s.%s", database, schema)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Failed to list tables: %v", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatalf("Failed to retrieve column metadata: %v", err)
	}

	fmt.Printf("Tables in schema %s.%s:\n", database, schema)

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		var tableName string
		for i, col := range columns {
			if strings.EqualFold(col, "name") {
				tableName = string(values[i])
			}
		}

		if tableName != "" {
			fmt.Printf("- %s\n", tableName)
		} else {
			fmt.Println("- <unknown table name>")
		}
	}
}

func GetUsers(db *sql.DB) {
	query := fmt.Sprintf("SHOW USERS")
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Failed to get user details")
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatalf("Failed to retrieve column metadata: %v", err)
	}

	fmt.Printf("Snowflake Users:")

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		var userName, loginName, email string

		for i, col := range columns {
			switch strings.ToLower(col) {
			case "name":
				userName = string(values[i])
			case "login_name":
				loginName = string(values[i])
			case "email":
				email = string(values[i])
			}
		}

		if userName != "" {
			fmt.Printf("- %s (Login: %s, Email: %s)\n", userName, loginName, email)
		} else {
			fmt.Println("- <unknown user>")
		}
	}
}

// The query here returns the last 100 queries performed on a
// particular database and schema
// can be edited further for specific user/time/amount related queries.
func ListQueryHistory(db *sql.DB, database string, filename string) {
	// Set the database and schema (INFORMATION_SCHEMA)
	_, err := db.Exec(fmt.Sprintf("USE DATABASE %s", database))
	if err != nil {
		log.Fatalf("Failed to use database %s: %v", database, err)
	}

	_, err = db.Exec("USE SCHEMA INFORMATION_SCHEMA")
	if err != nil {
		log.Fatalf("Failed to use schema INFORMATION_SCHEMA: %v", err)
	}

	// Open the existing file in append mode
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", filename, err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Run QUERY_HISTORY function
	query := `SELECT * FROM TABLE(INFORMATION_SCHEMA.QUERY_HISTORY()) ORDER BY start_time DESC`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Failed to retrieve query history: %v", err)
	}
	defer rows.Close()

	// Get column metadata
	columns, err := rows.Columns()
	if err != nil {
		log.Fatalf("Failed to retrieve column metadata: %v", err)
	}

	fmt.Printf("Query history for database %s:\n", database)

	// Prepare scan targets
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Iterate and write queries
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		var queryText, startTime string
		for i, col := range columns {
			switch {
			case strings.EqualFold(col, "query_text"):
				queryText = string(values[i])
			case strings.EqualFold(col, "start_time"):
				startTime = string(values[i])
			}
		}

		if queryText != "" {
			formatted := fmt.Sprintf("[%s] %s\n", startTime, queryText)
			fmt.Print("🕒 ", formatted)

			_, err := writer.WriteString(formatted)
			if err != nil {
				log.Printf("Failed to write to file: %v", err)
			}
		}
	}
}

// Function to create an empty text file
func CreateTextFile() {
	filename := "query_list.txt"

	// Create the file (empty)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Empty file created successfully:", filename)
}

func GetTableDetails(db *sql.DB, database, schema, table string) {
	// Get column information
	query := fmt.Sprintf(`
        SELECT 
            COLUMN_NAME, 
            DATA_TYPE, 
            CHARACTER_MAXIMUM_LENGTH, 
            IS_NULLABLE, 
            COLUMN_DEFAULT,
            COMMENT
        FROM 
            %s.INFORMATION_SCHEMA.COLUMNS 
        WHERE 
            TABLE_SCHEMA = '%s' 
            AND TABLE_NAME = '%s'
        ORDER BY 
            ORDINAL_POSITION
    `, database, schema, table)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Failed to get table details: %v", err)
	}
	defer rows.Close()

	fmt.Printf("Columns in %s.%s.%s:\n", database, schema, table)
	for rows.Next() {
		var colName, dataType, isNullable, colDefault, comment sql.NullString
		var charMaxLen sql.NullInt64

		err := rows.Scan(&colName, &dataType, &charMaxLen, &isNullable, &colDefault, &comment)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		fmt.Printf("- %s (%s", colName.String, dataType.String)
		if charMaxLen.Valid {
			fmt.Printf("(%d)", charMaxLen.Int64)
		}

		nullable := "NOT NULL"
		if isNullable.String == "YES" {
			nullable = "NULL"
		}

		fmt.Printf(", %s", nullable)

		if colDefault.Valid {
			fmt.Printf(", DEFAULT %s", colDefault.String)
		}

		fmt.Println(")")
	}

	// Get table statistics
	statsQuery := fmt.Sprintf(`
        SELECT 
            TABLE_NAME,
            ROW_COUNT,
            BYTES
        FROM 
            %s.INFORMATION_SCHEMA.TABLES
        WHERE 
            TABLE_SCHEMA = '%s' 
            AND TABLE_NAME = '%s'
    `, database, schema, table)

	var tableName sql.NullString
	var rowCount, bytes sql.NullInt64

	err = db.QueryRow(statsQuery).Scan(&tableName, &rowCount, &bytes)
	if err != nil {
		log.Printf("Failed to get table statistics: %v", err)
	} else {
		fmt.Printf("\nStatistics for %s:\n", tableName.String)
		fmt.Printf("- Row count: %d\n", rowCount.Int64)
		fmt.Printf("- Size: %d bytes (%.2f MB)\n", bytes.Int64, float64(bytes.Int64)/(1024*1024))
	}
}
