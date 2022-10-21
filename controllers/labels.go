package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	mysql_username string `envconfig:"MYSQL_USERNAME"`
	mysql_password string `envconfig:"MYSQL_PASSWORD"`
	mysql_host     string `envconfig:"MYSQL_HOST"`
	mysql_port     string `envconfig:"MYSQL_PORT"`
	mysql_database string `envconfig:"MYSQL_DATABASE"`
}

func getLabels(c *gin.Context) {
	var s Configuration
	err := envconfig.Process("conf", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	// "root:example@tcp(mysql:3306)/gogin"
	db, err := sql.Open("mysql", s.mysql_username+":"+s.mysql_password+"@tcp("+s.mysql_host+":"+s.mysql_port+")/"+s.mysql_database)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	c.String(http.StatusOK, version)
}
