package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func runRubyCli(args ...string) (int, error) {
	downloadRuby()
	cmd := exec.Command(rubyExe(), args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return 0, cmd.Run()
}

func rubyExe() string {
	return filepath.Join(homeDir(), "heroku.exe")
}

func downloadRuby() {
	out, err := os.Create(rubyExe())
	must(err)
	defer out.Close()
	resp, err := http.Get("https://s3.amazonaws.com/dickeyxxx_dev/heroku.exe")
	must(err)
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	must(err)
}
