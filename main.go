package main

import (
	"fmt"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error retrieving home directory:", err)
		return
	}
	fmt.Println("Home Directory:", pwd)
}
