package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		return
	}

	str := string(data)
	fmt.Println(str)
}
