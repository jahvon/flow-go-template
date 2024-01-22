package main

import (
	"fmt"
	"os"
)

func main() {
	ws := os.Getenv("FLOW_WORKSPACE_PATH")

	fmt.Printf("Hello from %s\n", ws)
}
