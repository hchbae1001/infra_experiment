package main

import (
	"auth-pod/db"
	"fmt"
)

func main() {
	db.Init()
	fmt.Println("Database initialized and ready to use!")
}
