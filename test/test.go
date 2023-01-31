package test

import (
	"fmt"
	"os"
)

func ReadFileTest() {
	path := "test/res.txt"

	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	data := make([]byte, 1024)
	n, err := file.Read(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	if n != 0 {
		fmt.Println(data)
	}
}
