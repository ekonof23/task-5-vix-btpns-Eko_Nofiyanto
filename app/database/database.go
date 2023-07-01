package database

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Inject is a middleware function to inject the database connection into the context
func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

// GetDB returns the database connection from the context
func GetDB(c *gin.Context) *gorm.DB {
	db, _ := c.Get("db")
	return db.(*gorm.DB)
}

// GetDSN returns the Data Source Name for the database connection
func GetDSN() string {
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "3306")
	user := GetEnv("DB_USER", "root")
	password := GetEnv("DB_PASSWORD", "")
	dbName := GetEnv("DB_NAME", "my_database")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbName)
}

// GetEnv returns the value of the specified environment variable or fallback to a default value
func GetEnv(key, fallback string) string {
	value, exists := getEnv(key)
	if exists {
		return value
	}
	return fallback
}

// getEnv returns the value of the specified environment variable
func getEnv(key string) (string, bool) {
	value, exists := goDotEnvVariable(key)
	return value, exists
}

// goDotEnvVariable retrieves the value of the specified environment variable
func goDotEnvVariable(key string) (string, bool) {
	return "", false
}