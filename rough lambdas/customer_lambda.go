package customer

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const module = "customerdbmodule"

type CustomerDB struct {
	dbname string  // Name of Customer DB
	db     *sql.DB // PostGres DB object
}

// Structure for DirEntry DB config provided in the startup yaml file.
// These work as the startup parameters to the DBs.
type DbConfig struct {
	PostGresServer string `yaml:"postgresserver"` // Host running DB server
	PostGresPort   int    `yaml:"postgresport"`   // DB Server port number

	User     string `yaml:"user"`     // Username for Database
	Password string `yaml:"password"` // Password for Database

	// Names of DBs from yaml file
	CustomerDBName       string `yaml:"customerdbname"`       // Customer DB Name
	CustomerDBServerPort string `yaml:"customerdbserverport"` // Customer DB HTTP server port

	DataSourceDBName string `yaml:"datasourcedbname"` // Data Source Objects DB Name - FS, NFS, VMDK, Oracle, etc.
	ServicesDBName   string `yaml:"servicesdbname"`   // Service Provider DB Name - AWS, GCP, Azure, etc.
	ActionDBName     string `yaml:"actiondbname"`     // Action DB for actions that can be taken by customer

	PaymentsDBName     string `yaml:"paymentsdbname"`     // Payments DB Name
	ExchangeRateDBName string `yaml:"exchangeratedbname"` // Exchange rate DB Name
}

var customerdb *CustomerDB

func GetEnvFile() (config *infra.Dbconfig, err error) {
	bucket := "oceano-env-files"
	item := "env"

	//new aws session
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	//download to buffer
	//buffer stores content and deletes it during garbage collection
	downloader := s3manager.NewDownloader(sess)

	//create buffer
	buf := aws.NewWriteAtBuffer([]byte{})

	_, err := downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(item),
	})
	if err != nil {
		return nil, err
	}

	// Parse the environment file from the buffer
	envMap, err := godotenv.Parse(bytes.NewReader(buf.Bytes()))
	if err != nil {
		return nil, err
	}

	// Set environment variables
	for key, value := range envMap {
		os.Setenv(key, value)
	}

	config = &infra.DbConfig{}

	//get the env variables (same as openclosedbs)
	config.PostGresServer = os.Getenv("PG_SERVER")
	port_string := os.Getenv("PG_PORT")
	config.PostGresPort, err = strconv.Atoi(port_string)
	if err != nil {
		str := fmt.Sprintf("Incorrect Postgres Port parameter: %s\n", port_string)
		infra.Log(startclosemodule, infra.ERROR, str)
		return &infra.DbConfig{}, errors.New(str)
	}
	config.User = os.Getenv("PG_USER")
	config.Password = os.Getenv("PG_PASSWORD")
	config.CustomerDBServerPort = os.Getenv("CUSTDB_SERVERPORT")
	config.CustomerDBName = os.Getenv("CUST_DBNAME")
	config.ServicesDBName = os.Getenv("SERVICES_DB")
	config.ActionDBName = os.Getenv("ACTION_DB")
	config.DataSourceDBName = os.Getenv("DATA_SOURCEDB")
	config.ExchangeRateDBName = os.Getenv("EXCHANGE_RATEDB")
	config.PaymentsDBName = os.Getenv("PAYMENTS_DB")

	// Structures in config file
	if infra.ExtraDebug() {
		fmt.Printf("Config parameters: %+v\n", config)
	}
	if config.PostGresServer == "" || config.PostGresPort == 0 || config.User == "" ||
		config.CustomerDBServerPort == "" || config.CustomerDBName == "" ||
		config.ServicesDBName == "" || config.ActionDBName == "" ||
		config.DataSourceDBName == "" || config.ExchangeRateDBName == "" ||
		config.PaymentsDBName == "" {
		// Hide the password
		config.User = "xxxxxxxx"
		config.Password = "xxxxxxxx"

		str := fmt.Sprintf("Incorrect Customer DB Config parameters: %+v\n", config)
		infra.Log(startclosemodule, infra.ERROR, str)
		return &infra.DbConfig{}, errors.New(str)
	}
	return config, err
}
