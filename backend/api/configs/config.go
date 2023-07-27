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

	DATABASE_USERNAME string
	DATABASE_PASSWORD string
	DATABASE_ADDRESS  string
	DATABASE_PORTS    int
	DATABASE_TABLE    string
	DATABASE_TIMEOUT  time.Duration

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

	DATABASE_USERNAME = os.Getenv("DATABASE_USERNAME")
	DATABASE_PASSWORD = os.Getenv("DATABASE_PASSWORD")
	DATABASE_ADDRESS = os.Getenv("DATABASE_ADDRESS")
	DATABASE_PORTS = string2int(os.Getenv("DATABASE_PORTS"))
	DATABASE_TABLE = os.Getenv("DATABASE_TABLE")
	DATABASE_TIMEOUT = time.Duration(string2int(os.Getenv("DATABASE_TIMEOUT"))) * time.Second

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
