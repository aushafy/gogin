package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type Label struct {
	ID     int64
	Book   string
	Labels string
}

func GetLabels(c *gin.Context) {
	cfg := mysql.Config{
		User:                 os.Getenv("MYSQL_USERNAME"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("MYSQL_HOST"),
		DBName:               os.Getenv("MYSQL_DATABASE"),
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var query string

	err2 := db.QueryRow("SELECT book FROM labels LIMIT 1").Scan(&query)
	if err2 != nil {
		log.Fatal(err2)
	}

	c.JSON(http.StatusOK, gin.H{"book_name": query})
}
