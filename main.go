package main

import (
	"gogin/app"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app.StartApp()
}
