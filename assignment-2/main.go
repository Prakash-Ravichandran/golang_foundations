package main

import (
	"fmt"
	"io"
	"os"
)

type myWriter struct{}

func(myWriter) Write(b []byte)(n int, err error) {
	fmt.Println(string(b))
	return len(b), nil
}


func main() {
    filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	mw := myWriter{}

	io.Copy(mw, file)
}