package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		db_host := os.Getenv("MYSQL_HOST")
		db_port := os.Getenv("MYSQL_PORT")
		db_username := os.Getenv("MYSQL_USER")
		db_password := os.Getenv("MYSQL_PASSWORD")
		db_name := os.Getenv("MYSQL_DATABASE")
		db_status := "Connected to database"

		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			db_username, db_password, db_host, db_port, db_name)
		_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			db_status = err.Error()
		}

		if os.Getenv("db_host") == "" {
			db_status = "Database invalid configuration"
		}

		c.JSON(http.StatusOK, gin.H{
			"version":     "v1.0",
			"description": "sample golang application",
			"database": gin.H{
				"info": gin.H{
					"host":     db_host,
					"port":     db_port,
					"username": db_username,
					"password": db_password,
					"db_name":  db_name,
				},
				"status": db_status,
			},
		})
	})
	r.Run("0.0.0.0:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
