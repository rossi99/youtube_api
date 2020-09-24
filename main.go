package main

import (
	"fmt"
	"youtube_api/router"
)

func main() {
	fmt.Println("Welcome to the Server!")

	e := router.New()

	e.Start(":8000")
}
