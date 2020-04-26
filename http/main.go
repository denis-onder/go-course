package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (l logWriter) Write(buf []byte) (int, error) {
	bw, err := fmt.Println(string(buf))
	if err != nil {
		return 0, err
	}
	return bw, nil
}

func main() {
	lw := logWriter{}
	res, err := http.Get("https://duckduckgo.com/")
	if err != nil {
		fmt.Printf("An error has occured! %v\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	// buf := make([]byte, 32*1024)
	// res.Body.Read(buf)
	// fmt.Println(string(buf))
	io.Copy(lw, res.Body)
}
