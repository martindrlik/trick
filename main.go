package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func main() {
	jar, err := cookiejar.New(nil)
	check(err)
	client := &http.Client{
		Jar:     jar,
		Timeout: time.Second,
	}
	do(client, "/login")
	do(client, "/create")
	do(client, "/list")
}

func do(client *http.Client, path string) {
	req, err := http.NewRequest(http.MethodPost, "http://:8080"+path, nil)
	check(err)
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()
	fmt.Printf("%8s %s\n", path, res.Status)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
