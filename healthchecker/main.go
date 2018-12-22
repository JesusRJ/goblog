package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := flag.String("port", "80", "Port on localhost to check")
	path := flag.String("path", "/health", "Path to check health. Should init with a '/'")
	flag.Parse()

	fmt.Println("http://127.0.0.1:" + *port + *path)

	resp, err := http.Get("http://127.0.0.1:" + *port + *path)

	if err != nil || resp.StatusCode != 200 {
		os.Exit(1)
	}

	os.Exit(0)
}
