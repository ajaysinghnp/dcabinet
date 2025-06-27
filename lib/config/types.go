package config

import "time"

type HTTPServer struct {
	Host    string        `yaml:"host" json:"host" env-default:"localhost" env-description:"Server Host" env-required:"true"`
	Port    string        `yaml:"port" json:"port" env-default:"8000" env-description:"Server Port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" json:"timeout" env-default:"5" env-description:"Request Timeout in seconds" env-required:"true"`
}

type Config struct {
	Env        string     `yaml:"env" json:"env" env-default:"production" env-description:"Environment (development, production, etc.)" env-required:"true"`
	Database   string     `yaml:"database" json:"database" env-default:".storage/database.db" env-description:"Database URL (sqlite3, mysql, etc.)" env-required:"true"`
	HTTPServer HTTPServer `yaml:"http_server" json:"http_server" env-description:"HTTP Server Configuration" env-required:"true"`
}
