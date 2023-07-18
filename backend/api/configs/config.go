package config

import (
	"os"
	"strconv"
	"time"
)

var (
	API_PORT          string
	SSL_CERTIFICATION string
	SSL_PRIVATE_KEY   string

	JWT_SIGNING_KEY string
	JWT_CLAIMS_KEY  string

	SQL_DRIVER         string
	SQL_USERNAME       string
	SQL_PASSWORD       string
	SQL_ADDRESS        string
	SQL_PORT           int
	SQL_DATABASE       string
	SQL_TIMEOUT        time.Duration
	SQL_MAXLIFETIME    int
	SQL_MAXOPENCONNECT int
	SQL_MAXIDLECONNECT int
	SQL_RETRYMAX       int
	SQL_RETRYINTERVAL  time.Duration

	REDIS_CACHE_ADDRESS    string
	REDIS_CACHE_PORT       int
	REDIS_CACHE_EXPIRATION int
	REDIS_CACHE_DURATION   time.Duration
)

func init() {
	API_PORT = os.Getenv("API_PORT")
	SSL_CERTIFICATION = os.Getenv("SSL_CERTIFICATION")
	SSL_PRIVATE_KEY = os.Getenv("SSL_PRIVATE_KEY")

	JWT_SIGNING_KEY = os.Getenv("JWT_SIGNING_KEY")
	JWT_CLAIMS_KEY = os.Getenv("JWT_CLAIMS_KEY")

	SQL_DRIVER = os.Getenv("SQL_DRIVER")
	SQL_USERNAME = os.Getenv("SQL_USERNAME")
	SQL_PASSWORD = os.Getenv("SQL_PASSWORD")
	SQL_ADDRESS = os.Getenv("SQL_ADDRESS")
	SQL_PORT = string2int(os.Getenv("SQL_PORT"))
	SQL_DATABASE = os.Getenv("SQL_DATABASE")
	SQL_TIMEOUT = time.Duration(string2int(os.Getenv("SQL_TIMEOUT"))) * time.Second
	SQL_MAXLIFETIME = string2int(os.Getenv("SQL_MAXLIFETIME"))
	SQL_MAXOPENCONNECT = string2int(os.Getenv("SQL_MAXOPENCONNECT"))
	SQL_MAXIDLECONNECT = string2int(os.Getenv("SQL_MAXIDLECONNECT"))

	REDIS_CACHE_ADDRESS = os.Getenv("REDIS_CACHE_ADDRESS")
	REDIS_CACHE_PORT = string2int(os.Getenv("REDIS_CACHE_PORT"))
	REDIS_CACHE_EXPIRATION = string2int(os.Getenv("SESSION_EXPIRATION"))
	REDIS_CACHE_DURATION = time.Duration(REDIS_CACHE_EXPIRATION) * time.Second
}

func string2int(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		i = 0
	}
	return i
}
