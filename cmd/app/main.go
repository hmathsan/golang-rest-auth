package main

import (
	s "example/rest/test/internal/app"
	db "example/rest/test/internal/app/database"
)

func init() {
	db.InitDB()
}

func main() {
	defer db.Db.Close()
	s.Server()
}
