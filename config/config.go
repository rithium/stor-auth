package config

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

func LoadConfig() {
	HttpServer.Uri = "0.0.0.0"
	HttpServer.Port = 11001

	MySQL.Url = "localhost"
	MySQL.Port = "3306"
	MySQL.User = "root"
	MySQL.Pass = "toor"
	MySQL.Database = "stor"
}