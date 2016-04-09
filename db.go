package mscore

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/go-sql-driver/mysql"
	"github.com/newsquid/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"os"
)

// The open database connection
var DB *gorm.DB

/*
SSL configuration parameters
*/
type SSLConfig struct {
	ClientCertFile string
	ClientKeyFile  string
	ServerCaFile   string
}

/*
Validates the SSL configuration parameters
*/
func (config *SSLConfig) Complete() bool {
	return config.ClientCertFile != "" &&
		config.ClientKeyFile != "" &&
		config.ServerCaFile != ""
}

/*
Get the open database connection
*/
func GetDB() *gorm.DB {
	return DB
}

/*
Initialize the database object from the given variables
*/
func InitDBFromVariables(database string, conn string, debug bool) *gorm.DB {
	dbConnection, err := gorm.Open(database, conn)
	DB = &dbConnection

	if err != nil {
		log.Panicf("Could not init database: %s", err.Error())
	}

	dbConnection.LogMode(debug)

	return DB
}

/*
Initialize the Gorm Database Object from env.
It uses the environment variables DB and DBCONN for configuration
*/
func InitDB() *gorm.DB {
	database := os.Getenv("DB")
	conn := os.Getenv("DBCONN")
	debugenv := os.Getenv("DBDEBUG")
	debug := false

	// Optional ssl configuration
	tls := os.Getenv("TLS") == "true"

	if database == "" {
		database = "sqlite3"
	}

	if conn == "" {
		conn = "database.db"
	}

	if debugenv == "true" {
		debug = true
	}

	if database == "mysql" && tls {
		config := SSLConfig{
			os.Getenv("CLIENT_CERT_FILE"),
			os.Getenv("CLIENT_KEY_FILE"),
			os.Getenv("SERVER_CA_FILE"),
		}
		if !config.Complete() {
			panic(`Invalid configuration! TLS requested,
				but one or more certificate files are missing`)
		}
		InitCustomTls(config)
	}
	return InitDBFromVariables(database, conn, debug)
}

/*
Initializes the test database
*/
func InitTestDB() *gorm.DB {
	return InitDBFromVariables("sqlite3", "test.db", false)
}

/*
Removes / deletes the test database
*/
func RemoveTestDB() {
	err := DB.Close()
	if err != nil {
		panic(err)
	}
	os.Remove("test.db")
}

/*
Load the certificate files from the given SSL configuration parameters
*/
func InitCustomTls(config SSLConfig) {
	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(config.ServerCaFile)
	if err != nil {
		log.Fatal(err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}
	clientCert := make([]tls.Certificate, 0, 1)
	certs, err := tls.LoadX509KeyPair(config.ClientCertFile, config.ClientKeyFile)
	if err != nil {
		log.Fatal(err)
	}
	clientCert = append(clientCert, certs)
	mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs:      rootCertPool,
		Certificates: clientCert,
		// Google does not provide IP sans in their certificates
		InsecureSkipVerify: true,
	})
}
