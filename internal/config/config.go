package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)


var cfg *Config

func GetConfig() Config {
	if cfg != nil{
		return *cfg
	}

	return Config{}
}

type Grpc struct {
	Port              int    `yaml:"port"`
	MaxConnectionIdle int64  `yaml:"maxConnectionIdle"`
	Timeout           int64  `yaml:"timeout"`
	MaxConnectionAge  int64  `yaml:"maxConnectionAge"`
	Host              string `yaml:"host"`
}

type Gateway struct {
	Port               int      `yaml:"port"`
	Host               string   `yaml:"host"`
	AllowedCORSOrigins []string `yaml:"allowedCorsOrigins"`
}

type Swagger struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Dist	 string `yaml:"dist"`
	Filepath string `yaml:"filepath"`
}

type Config struct {
	Grpc Grpc `yaml:"grpc"`
	Gateway Gateway `yaml:"gateway"`
	Swagger Swagger `yaml:"swagger"`
	ToDoService string `yaml:"to_do_service"`
	DB DB `yaml:"db"`
}


type DB struct {
	DSN        string `yaml:"DSN"`
	DSND		string `yaml:"DSND"`
	MaxOpenConns int `yaml:"maxOpenConns"`
	MaxIdleConns int `yaml:"MaxIdleConns"`
	ConnMaxIdleTime time.Duration `yaml:"connMaxIdleTime"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
}

func (db *DB) GetDSN() string {
	x := os.Getenv("DOCK")
	if x != ""{
		return db.DSND
	}
	return db.DSN
}

func (db *DB) GetMaxOpenConns() int {
	return db.MaxOpenConns
}

func (db *DB) GetMaxIdleConns() int {
	return db.MaxIdleConns
}

func (db *DB) GetConnMaxIdleTime() time.Duration {
	return db.ConnMaxIdleTime
}

func (db *DB) GetConnMaxLifetime() time.Duration {
	return db.ConnMaxLifetime
}

func ReadConfigYaml(configYaml string) error {
	if cfg != nil{
		return nil
	}

	file,err := os.Open(configYaml)
	if err != nil{
		return err
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil{
		return err
	}

	return nil
	

}