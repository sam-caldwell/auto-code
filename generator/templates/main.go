// main.go
package main

import (
	"configuration" // update with the actual path
	"fmt"
	"os"
)

func main() {
	config, err := configuration.NewConfiguration()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		os.Exit(1)
	}

	// Program logic...
	fmt.Println("Configuration loaded:", config)
}
