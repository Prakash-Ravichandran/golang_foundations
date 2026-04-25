package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type customWriter struct{}

func (customWriter) Write(p []byte) (n int, err error) {
	fmt.Println("Written using custom writer implementation");
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := customWriter{}

	fmt.Println("Response:", res)
	io.Copy(lw, res.Body)
}
