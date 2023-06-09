package core

import "gorm.io/gorm/logger"

func NewServerConfig(mode string, addr string, port int) *ServerConfig {
	return &ServerConfig{
		Mode: mode,
		Addr: addr,
		Port: port,
	}
}
func LoadServerConfig(file string) *ServerConfig {
	v := NewViperHelper(file)
	mode := v.GetString("server.mode")
	addr := v.GetString("server.addr")
	port := v.GetInt("server.port")
	return NewServerConfig(mode, addr, port)
}

func SaveServerConfig(file string, sc *ServerConfig) {
	v := NewViperHelper(file)
	v.Set("server", sc)
	v.WriteConfig()
}

//=============

type ServerConfig struct {
	Mode string
	Addr string
	Port int
}

//=================================

func SaveRedisConfig(file string, rc *RedisConfig) {
	v := NewViperHelper(file)
	v.Set("redis", rc)
	v.WriteConfig()
}

func LoadRedisConfig(file string) *RedisConfig {
	v := NewViperHelper(file)
	addr := v.GetString("redis.addr")
	port := v.GetInt("redis.port")
	username := v.GetString("redis.username")
	password := v.GetString("redis.password")
	db := v.GetInt("redis.db")
	return NewRedisConfig(addr, port, username, password, db)
}

func NewRedisConfig(addr string, port int, username string, password string, db int) *RedisConfig {
	return &RedisConfig{
		Addr:     addr,
		Port:     port,
		Username: username,
		Password: password,
		Db:       db,
	}
}

type RedisConfig struct {
	Addr     string
	Port     int
	Username string
	Password string
	Db       int
}

func LoadSqliteConfig(file string) *SqliteConfig {
	v := NewViperHelper(file)
	// level := v.GetString("sqlite.level")
	databaseFile := v.GetString("sqlite.file")
	return &SqliteConfig{
		LogLevel: logger.Info,
		File:     databaseFile,
	}
}
func LoadMySqlConfig(file string) *MySqlConfig {
	v := NewViperHelper(file)
	username := v.GetString("mysql.username")
	password := v.GetString("mysql.password")
	host := v.GetString("mysql.host")
	port := v.GetInt("mysql.port")
	database := v.GetString("mysql.database")
	return &MySqlConfig{
		LogLevel: logger.Info,
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}
}

type SqliteConfig struct {
	LogLevel logger.LogLevel
	File     string
}
type MySqlConfig struct {
	LogLevel logger.LogLevel
	Username string
	Password string
	Host     string
	Port     int
	Database string
	args     map[string]string
}
