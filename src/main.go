package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world", os.Getenv("ENV"))
}
