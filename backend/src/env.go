package main

import (
	"os"
	"strconv"
	"time"
)

var (
	API_PORT string

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

	REDIS_ADDRESS string
	REDIS_PORT    int

	SESSION_NAME       string
	SESSION_EXPIRATION int
	SESSION_DURATION   time.Duration
)

func init() {
	API_PORT = os.Getenv("API_PORT")

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
	SQL_RETRYMAX = string2int(os.Getenv("SQL_RETRYMAX"))
	SQL_RETRYINTERVAL = time.Duration(string2int(os.Getenv("SQL_RETRYINTERVAL"))) * time.Second

	REDIS_ADDRESS = os.Getenv("REDIS_ADDRESS")
	REDIS_PORT = string2int(os.Getenv("REDIS_PORT"))

	SESSION_NAME = os.Getenv("SESSION_NAME")
	SESSION_EXPIRATION = string2int(os.Getenv("SESSION_EXPIRATION"))
	SESSION_DURATION = time.Duration(SESSION_EXPIRATION) * time.Second
}

func string2int(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		i = 0
	}
	return i
}
