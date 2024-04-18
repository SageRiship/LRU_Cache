package main

import (
	"apica-backend/internal/router"
	"fmt"
)

func main() {
	e := router.SetupRouter()
	fmt.Println("starting service")

	e.Start(":8080")
}
