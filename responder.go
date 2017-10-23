package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hostname:", Respond())
}


func Respond() (string) {
	name, _ := os.Hostname()
	return name
}

