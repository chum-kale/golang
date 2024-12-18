package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"

	as "github.com/aerospike/aerospike-client-go"
)

var aerospikeClient *as.Client

var host = "172.17.0.2"
var port = 3000
var namespace = "dataoceanodedupdb"
var set = "random"

type AerospikeRecord struct {
	Key  string
	Bins map[string]interface{}
}

// Store the record key globally so it can be reused
var globalRecordKey *as.Key

func InitAerospikeClient() (*as.Client, error) {
	client, err := as.NewClient(host, port)
	if err != nil {
		log.Fatalf("Failed to connect to Aerospike: %v", err)
		return nil, err
	}

	aerospikeClient = client

	return client, nil
}

func CloseAerospikeClient() {
	if aerospikeClient != nil {
		aerospikeClient.Close() // No need to assign to err since Close() doesn't return a value
		log.Println("Aerospike client closed successfully.")
	} else {
		log.Println("Aerospike client was not initialized or already closed.")
	}
}

func GetRecordKey() (*as.Key, error) {
	if aerospikeClient == nil {
		log.Fatalf("Aerospike not connected")
	}

	// Generate and store the key once globally
	if globalRecordKey == nil {
		recordKey, err := as.NewKey(namespace, set, generateKey())
		if err != nil {
			log.Fatalf("Failed to create recordkey: %v", err)
			return nil, err
		}
		globalRecordKey = recordKey
	}

	return globalRecordKey, nil
}

// we will be using uuid as primary key for our storageids
func generateKey() string {
	aerospikeKey := uuid.New().String()
	return aerospikeKey
}

func CreateRecord(name, city string, age int) error {
	recordKey, err := GetRecordKey()
	if err != nil {
		return err
	}

	// Define a write policy with SendKey enabled
	// This tells aerospike to store userkey as well,
	// Normally, it aint do that
	writePolicy := as.NewWritePolicy(0, 0) // Example write policy
	writePolicy.SendKey = true             // Ensure user-defined key is stored

	// Define bins (columns) for the record
	bins := as.BinMap{
		"name": name,
		"age":  age,
		"city": city,
	}

	// Write the record into Aerospike
	err = aerospikeClient.Put(writePolicy, recordKey, bins)
	if err != nil {
		return fmt.Errorf("failed to create record: %v", err)
	}
	fmt.Println("Record created successfully!")
	return nil
}

// Read a record from Aerospike
func ReadRecord() error {
	recordKey, err := GetRecordKey()
	if err != nil {
		return err
	}

	readPolicy := as.NewPolicy()
	record, err := aerospikeClient.Get(readPolicy, recordKey)
	if err != nil {
		return fmt.Errorf("failed to read record: %v", err)
	}

	fmt.Printf("Record fetched: %v\n", record.Bins)
	return nil
}

// // Update a record in Aerospike
// func UpdateRecord() error {
// 	recordKey, err := GetRecordKey()
// 	if err != nil {
// 		return err
// 	}

// 	// First, get the current record
// 	record, err := aerospikeClient.Get(nil, recordKey)
// 	if err != nil {
// 		return fmt.Errorf("failed to retrieve record: %v", err)
// 	}

// 	// Increment the age
// 	if age, ok := record.Bins["age"].(int); ok {
// 		// Create a BinMap with the new age
// 		bins := as.BinMap{
// 			"age": age + increment, // Add the increment to the current age
// 	}

// 	// // Modify bins (columns) for the update
// 	// bins := as.BinMap{
// 	// 	"age": 31, // Update the age
// 	// 	//"city": "San Francisco", // Update the city
// 	// }

// 	// Write the updated record into Aerospike
// 	err = aerospikeClient.Put(nil, recordKey, bins)
// 	if err != nil {
// 		return fmt.Errorf("failed to update record: %v", err)
// 	}

// 	fmt.Println("Record updated successfully!")
// 	return nil
// 	}
// }

func UpdateRecord(increment int) error {
	recordKey, err := GetRecordKey()
	if err != nil {
		return err
	}

	// First, get the current record
	record, err := aerospikeClient.Get(nil, recordKey)
	if err != nil {
		return fmt.Errorf("failed to retrieve record: %v", err)
	}

	// Increment the age
	if currentAge, exists := record.Bins["age"]; exists {
		newAge := currentAge.(int) + increment // Directly increment the current age

		// Create a BinMap with the new age
		bins := as.BinMap{
			"age": newAge, // Updated age
		}

		// Write the updated record into Aerospike
		err = aerospikeClient.Put(nil, recordKey, bins)
		if err != nil {
			return fmt.Errorf("failed to update record: %v", err)
		}

		fmt.Println("Record updated successfully!")
		return nil
	}

	return fmt.Errorf("age bin does not exist")
}

// Delete a record from Aerospike
func DeleteRecord() error {
	recordKey, err := GetRecordKey()
	if err != nil {
		return err
	}

	existed, err := aerospikeClient.Delete(nil, recordKey)
	if err != nil {
		return fmt.Errorf("failed to delete record: %v", err)
	}

	if existed {
		fmt.Println("Record deleted successfully!")
	} else {
		fmt.Println("Record did not exist.")
	}
	return nil
}

// Check if a particular record exists in Aerospike
func RecordExists() (bool, error) {
	recordKey, err := GetRecordKey()
	if err != nil {
		return false, err
	}

	// Use Exists() to check if the record exists
	exists, err := aerospikeClient.Exists(nil, recordKey)
	if err != nil {
		return false, fmt.Errorf("error checking record existence: %v", err)
	}

	if exists {
		fmt.Println("Record exists.")
	} else {
		fmt.Println("Record does not exist.")
	}

	return exists, nil
}

// CollectAllEntries retrieves all entries from the specified namespace and set
func CollectAllEntries() ([]*as.Record, error) {
	if aerospikeClient == nil {
		log.Fatalf("Aerospike not connected")
	}

	// Define a Scan Policy to avoid unsupported features
	scanPolicy := as.NewScanPolicy()
	scanPolicy.SendKey = false // Avoid sending user keys, as your setup may not support it
	scanPolicy.ConcurrentNodes = false

	// Scan the set in the namespace
	recordset, err := aerospikeClient.ScanAll(scanPolicy, namespace, set)
	if err != nil {
		return nil, fmt.Errorf("failed to scan the set: %v", err)
	}
	defer recordset.Close()

	// Collect all records
	var records []*as.Record
	for rec := range recordset.Results() {
		if rec.Err != nil {
			return nil, fmt.Errorf("error while scanning records: %v", rec.Err)
		}
		records = append(records, rec.Record)
	}

	fmt.Printf("Collected %d records from Aerospike\n", len(records))
	return records, nil
}

func GetBinContent() (map[string]interface{}, error) {
	// Check if the record exists first
	exists, err := RecordExists()
	if err != nil {
		return nil, fmt.Errorf("error checking record existence: %v", err)
	}

	if !exists {
		return nil, fmt.Errorf("record does not exist")
	}

	// Fetch the record's bin content
	recordKey, err := GetRecordKey()
	if err != nil {
		return nil, fmt.Errorf("error getting record key: %v", err)
	}

	// Define the policy for the Get operation
	readPolicy := as.NewPolicy()
	record, err := aerospikeClient.Get(readPolicy, recordKey)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch record: %v", err)
	}

	// Return only the bin values (the values inside the bins)
	return record.Bins, nil
}

// func ViewAllEntries() error {
// 	// Ensure the Aerospike client is initialized
// 	if aerospikeClient == nil {
// 		return fmt.Errorf("Aerospike client not initialized")
// 	}

// 	// Create a Statement for querying the namespace and set
// 	statement := as.NewStatement(namespace, set)

// 	// Execute the query to get all records in the specified namespace and set
// 	recordset, err := aerospikeClient.Query(nil, statement)
// 	if err != nil {
// 		return fmt.Errorf("failed to query all records: %v", err)
// 	}
// 	defer recordset.Close()

// 	fmt.Println("Listing all entries in the database:")

// 	// Loop through all records retrieved by the query
// 	for rec := range recordset.Results() {
// 		if rec.Err != nil {
// 			return fmt.Errorf("error while querying records: %v", rec.Err)
// 		}

// 		// Print each record's key and bins
// 		fmt.Printf("Record Key: %v\n", rec.Record.Key.Value())
// 		fmt.Printf("Bins: %v\n", rec.Record.Bins)
// 	}

// 	fmt.Println("All entries displayed.")
// 	return nil
// }

func ViewAllEntries() error {
	// Ensure the Aerospike client is initialized
	if aerospikeClient == nil {
		return fmt.Errorf("Aerospike client not initialized")
	}

	// Define a Scan Policy
	scanPolicy := as.NewScanPolicy()
	scanPolicy.ConcurrentNodes = true // Enables concurrent scanning across nodes for better performance

	// Execute the scan to get all records in the specified namespace and set
	recordset, err := aerospikeClient.ScanAll(scanPolicy, namespace, set)
	if err != nil {
		return fmt.Errorf("failed to scan all records: %v", err)
	}
	defer recordset.Close()

	//fmt.Println("Listing all entries in the database:")
	fmt.Printf("recordset: %v\n", recordset)

	// Loop through all records retrieved by the scan

	for rec := range recordset.Results() {
		if rec == nil {
			return fmt.Errorf("received nil record")
		}
		if rec.Err != nil {
			return fmt.Errorf("error while reading record: %v", rec.Err)
		}

		fmt.Printf("Record Key: %v\n", rec.Record.Key.Value())
		fmt.Printf("Bins: %v\n", rec.Record.Bins)
		fmt.Println("All entries displayed.")
	}

	//fmt.Println("All entries displayed.")
	return nil
}

func ViewRecordsWithPrefix(prefix string) error {
	// Ensure Aerospike client is initialized
	if aerospikeClient == nil {
		return fmt.Errorf("Aerospike client not initialized")
	}

	// Define a Scan Policy
	scanPolicy := as.NewScanPolicy()
	scanPolicy.ConcurrentNodes = false // Set concurrency

	// Initialize the statement for querying
	statement := as.NewStatement(namespace, set)
	recordset, err := aerospikeClient.Query(nil, statement)
	if err != nil {
		return fmt.Errorf("failed to scan all records: %v", err)
	}
	defer recordset.Close()

	// Track if any records with the prefix are found
	found := false

	// Loop over all scanned records
	for rec := range recordset.Results() {
		if rec.Err != nil {
			return fmt.Errorf("error while scanning records: %v", rec.Err)
		}

		// Check if the primary key is not nil and contains the prefix
		if rec.Record.Key != nil && rec.Record.Key.Value() != nil {
			userKey := rec.Record.Key.Value().String()
			if strings.HasPrefix(userKey, prefix) {
				// If a matching record is found, print it
				fmt.Printf("Found record with prefix '%s': %v\n", prefix, rec.Record.Bins)
				found = true
			}
		} else {
			fmt.Println("Record Key is nil; skipping this record.")
		}
	}

	// If no records were found with the given prefix, return an error
	if !found {
		return fmt.Errorf("no records found with prefix '%s'", prefix)
	}

	return nil
}

func storeAll(keys [][]byte, attrs [][]byte) error {
	if aerospikeClient == nil {
		return fmt.Errorf("Aerospike client not initialized")
	}

	// Ensure keys and attrs have the same length
	if len(keys) != len(attrs) {
		return fmt.Errorf("keys and attributes arrays must be the same length")
	}

	for i, keyBytes := range keys {
		// Generate Aerospike key
		key, err := as.NewKey(namespace, set, keyBytes)
		if err != nil {
			return fmt.Errorf("failed to create key: %v", err)
		}

		// Parse the JSON-encoded bins from attrs[i]
		var binsData map[string]interface{}
		if err := json.Unmarshal(attrs[i], &binsData); err != nil {
			return fmt.Errorf("failed to decode bins for key %v: %v", keyBytes, err)
		}

		// Write the record with the bins into Aerospike
		err = aerospikeClient.Put(nil, key, binsData)
		if err != nil {
			return fmt.Errorf("failed to store record for key %v: %v", keyBytes, err)
		}
		fmt.Printf("Stored record with key: %v and bins: %v\n", keyBytes, binsData)
	}

	return nil
}

// Func that returns all records in a struct format
func ViewAllEntriesRet() ([]AerospikeRecord, error) {
	if aerospikeClient == nil {
		return nil, fmt.Errorf("Aerospike client not initialized")
	}

	scanPolicy := as.NewScanPolicy()
	scanPolicy.ConcurrentNodes = true

	recordset, err := aerospikeClient.ScanAll(scanPolicy, namespace, set)
	if err != nil {
		return nil, fmt.Errorf("failed to scan all records: %v", err)
	}
	defer recordset.Close()

	var records []AerospikeRecord

	for rec := range recordset.Results() {
		if rec == nil {
			return nil, fmt.Errorf("received nil record")
		}
		if rec.Err != nil {
			return nil, fmt.Errorf("error while reading record: %v", rec.Err)
		}

		recordKey := rec.Record.Key.Value().String()
		records = append(records, AerospikeRecord{
			Key:  recordKey,
			Bins: rec.Record.Bins,
		})
	}

	return records, nil
}

func TruncateSet(namespace, set string) error {
	// Ensure the Aerospike client is initialized
	if aerospikeClient == nil {
		return fmt.Errorf("Aerospike client not initialized")
	}

	// Use the Truncate function to delete all records in the namespace and set
	err := aerospikeClient.Truncate(nil, namespace, set, nil)
	if err != nil {
		return fmt.Errorf("failed to truncate set: %v", err)
	}

	fmt.Printf("All entries in set '%s' of namespace '%s' have been deleted.\n", set, namespace)
	return nil
}

func main() {
	// Initialize the Aerospike client
	_, err := InitAerospikeClient()
	if err != nil {
		log.Fatalf("Could not initialize Aerospike client: %v", err)
	}
	defer CloseAerospikeClient()

	// Create a new record
	err = CreateRecord("part", "sinhagad", 22)
	if err != nil {
		log.Printf("Error creating record: %v", err)
	} else {
		fmt.Println("Record created successfully!")
	}

	err = CreateRecord("james", "mumbai", 67)
	if err != nil {
		log.Printf("Error creating record: %v", err)
	} else {
		fmt.Println("Record created successfully!")
	}

	err = CreateRecord("sunil", "bihar", 89)
	if err != nil {
		log.Printf("Error creating record: %v", err)
	} else {
		fmt.Println("Record created successfully!")
	}

	err = CreateRecord("dhoni", "chennai", 45)
	if err != nil {
		log.Printf("Error creating record: %v", err)
	} else {
		fmt.Println("Record created successfully!")
	}

	// Check if the record exists
	exists, err := RecordExists()
	if err != nil {
		log.Printf("Error checking record existence: %v", err)
	} else if exists {
		fmt.Println("Record exists in the database.")
	}

	// Read the record
	err = ReadRecord()
	if err != nil {
		log.Printf("Error reading record: %v", err)
	} else {
		fmt.Println("Record read successfully!")
	}

	// Update the record
	err = UpdateRecord(-30)
	if err != nil {
		log.Printf("Error updating record: %v", err)
	} else {
		fmt.Println("Record updated successfully!")
	}

	// Read the record again after updating
	err = ReadRecord()
	if err != nil {
		log.Printf("Error reading updated record: %v", err)
	} else {
		fmt.Println("Updated record read successfully!")
	}

	// View all entries (this will print all entries from the Aerospike database)
	err = ViewAllEntries()
	if err != nil {
		log.Printf("Error viewing all entries: %v", err)
	} else {
		fmt.Println("All entries viewed successfully!")
	}

	// View records with a specific prefix (empty string will return all records)
	err = ViewRecordsWithPrefix("")
	if err != nil {
		log.Printf("Error viewing records with prefix: %v", err)
	} else {
		fmt.Println("Records with prefix viewed successfully!")
	}

	// // Finally, delete the record
	// err = DeleteRecord()
	// if err != nil {
	// 	log.Printf("Error deleting record: %v", err)
	// } else {
	// 	fmt.Println("Record deleted successfully!")
	// }

	// Check if the record exists after deletion
	exists, err = RecordExists()
	if err != nil {
		log.Printf("Error checking record existence: %v", err)
	} else if !exists {
		fmt.Println("Record does not exist in the database anymore.")
	}
}
