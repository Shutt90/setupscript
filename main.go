package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	getHelixConfig()
}

func getHelixConfig() string {
	c := http.Client{Timeout: time.Duration(30) * time.Second}
	res, err := c.Get("https://raw.githubusercontent.com/Shutt90/archdotfiles/main/helix/config.toml")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}
