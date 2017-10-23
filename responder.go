package main

import (
	"fmt"
	"os"
)

func main() {
	name, _ := os.Hostname()

	fmt.Println("Hostname:", name)
}

