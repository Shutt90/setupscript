package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"time"
)

var stuffToInstall = []string{
	"helix",
	"postman",
	"go",
	"git",
	"helix",
	"nnn",
	"vue-language-server",
	"typescript-language-server",
	"vscode-css-language-server",
	"intelephense",
	"vscode-json-language-server",
	"vscode-html-language-server",
	"gopls",
	"bash-language-server",
}

var cmd *exec.Cmd

func main() {
	cmd = exec.Command("/bin/bash -c \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\"")
	stdin, stdout := bytes.Buffer{}, bytes.Buffer{}
	cmd.Stdout = &stdout
	cmd.Stdin = &stdin
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

	buf := bufio.NewReader(cmd.Stdin)
	fmt.Printf("> ")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sentence))

	results := []string{}
	ch := make(chan struct{}, runtime.NumCPU()-2)
	for _, installing := range stuffToInstall {
		go func(installing string) {
			cmd = exec.Command("brew", "install", installing)
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
			results = append(results, fmt.Sprintf("%s: %s", installing, cmd.Stdout))
			ch <- struct{}{}
		}(installing)
	}

	if !slices.Contains(results, "helix") {
		os.WriteFile("~/.config/helix/config.toml", getHelixConfig(), 0644)
		os.WriteFile("~/.config/helix/themes/custom_theme.toml", getHelixTheme(), 0644)
	}
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
