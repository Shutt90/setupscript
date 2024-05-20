package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"
)

var stuffToInstall = []string{
	"helix",
	"postman",
	"go",
	"vue-language-server",
	"typescript-language-server",
	"vscode-css-language-server",
	"intelephense",
	"vscode-json-language-server",
	"vscode-html-language-server",
	"gopls",
	"bash-language-server",
}

func main() {
	results := []string{}
	wg := sync.WaitGroup{}
	for _, installing := range stuffToInstall {
		go func() {
			wg.Add(1)
			stdin, stdout := bytes.Buffer{}, bytes.Buffer{}
			cmd := exec.Command("brew", "install", installing)
			cmd.Stdout = &stdout
			cmd.Stdin = &stdin
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
			results = append(results, fmt.Sprintf("%s %s", installing, cmd.Stdout))
			wg.Done()
		}()
		wg.Wait()
	}

	os.WriteFile("~/.config/helix/config.toml", getHelixConfig(), 0644)
	os.WriteFile("~/.config/helix/themes/custom_theme.toml", getHelixTheme(), 0644)
}

func getHelixConfig() []byte {
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

	return body
}

func getHelixTheme() []byte {
	c := http.Client{Timeout: time.Duration(30) * time.Second}
	res, err := c.Get("https://raw.githubusercontent.com/Shutt90/archdotfiles/main/helix/themes/custom_theme.toml")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}
