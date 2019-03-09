package config

import (
	"fmt"
	"os"
	"strconv"
)

const DefaultNoOfGoroutines = 5
const DefaultNoOfCrawlsPerLink = 10
const DefaultDbUser = "guest"
const DefaultDbPwd = ""
const DefaultDbUrl = "127.0.0.1:3306"
const DefaultDbName = ""
const DefaultFilePath = "defaultFile.dat"

//main config struct
type Config struct {
	CrawlerConfig CrawlerConfig
	StoreConfig   StoreConfig
}

//crawler config struct
type CrawlerConfig struct {
	NumOfGoroutines int
	NumOfCrawls     int
}

//configuration for data storing, db connection settings, file path etc
type StoreConfig struct {
	DbUser   string
	DbPwd    string
	DbUrl    string
	DbName   string
	FilePath string
}

//returns pointer to new config struct
func New() *Config {
	return &Config{
		CrawlerConfig: CrawlerConfig{
			NumOfGoroutines: getEnvAsInt("GOROUTINES", DefaultNoOfGoroutines),
			NumOfCrawls:     getEnvAsInt("NUMOFCRAWLS", DefaultNoOfCrawlsPerLink),
		},
		StoreConfig: StoreConfig{
			DbUser:   getEnv("DBUSER", DefaultDbUser),
			DbPwd:    getEnv("DBPWD", DefaultDbPwd),
			DbUrl:    getEnv("DBURL", DefaultDbUrl),
			DbName:   getEnv("DBNAME", DefaultDbName),
			FilePath: getEnv("FILESTORE", DefaultFilePath),
		},
	}
}

// looks up environment by name, returns default value if not found
func getEnv(envName string, defaultValue string) string {
	value, exists := os.LookupEnv(envName)
	if !exists {
		fmt.Printf("Didn't find env '%s'. Setting default value '%v'\n", envName, defaultValue)
		return defaultValue
	}
	return value
}

// looks up environment by name and converts it to string, if not found returns default value
func getEnvAsInt(envName string, defaultValue int) int {
	value := getEnv(envName, "")
	if value == "" {
		fmt.Printf("Env \"%s\" not found. Setting default value '%v'", envName, defaultValue)
		return defaultValue
	}

	n, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("Failed to convert env \"%s\" value '%v' to int. Setting default value '%v'", envName, value, defaultValue)
		return defaultValue
	}

	return n
}
