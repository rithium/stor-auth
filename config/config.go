package config

import (
	"strconv"
	"os"
)

type HttpServerConfig struct {
	Uri		string
	Port		int
}

type MySQLConfig struct {
	Url		string
	Port		string
	User		string
	Pass		string
	Database	string
}

var (
	HttpServer	HttpServerConfig
	MySQL 		MySQLConfig
)

const ENV_HTTP_URL = "S_URL"
const ENV_HTTP_PORT = "S_PORT"
const ENV_MYSQL_PORT = "S_MYSQL_PASS"
const DEFAULT_HTTP_URL = "0.0.0.0"
const DEFAULT_HTTP_PORT = "80"

func LoadConfig() {
	HttpServer.Uri = getEnv(ENV_HTTP_URL, DEFAULT_HTTP_URL)
	HttpServer.Port, _ = strconv.Atoi(getEnv(ENV_HTTP_PORT, DEFAULT_HTTP_PORT))

	MySQL.Url = "localhost"
	MySQL.Port = "3306"
	MySQL.User = "root"
	MySQL.Pass = getEnv(ENV_MYSQL_PORT, "")
	MySQL.Database = "stor"
}

func getEnv(name string, defaultValue string) (string) {
	if value := os.Getenv(name); value != "" {
		return value
	}

	return defaultValue
}