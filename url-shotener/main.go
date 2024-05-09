package main

import (
	"example.com/api"
	"example.com/db"
)

func main() {
	db.Run()
	api.Run()

}