package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"

	. "github.com/dickeyxxx/hk/util"
)

func main() {
	hkPath := filepath.Join(homeDir(), ".hk", "hk")
	if runtime.GOOS == "windows" {
		hkPath = hkPath + ".exe"
	}
	exists, err := FileExists(hkPath)
	Must(err)
	if !exists {
		downloadHk(hkPath)
	}
	Must(runHk(hkPath, os.Args[1:]))
}

func runHk(hkPath string, args []string) error {
	fmt.Println("Running hk...")
	cmd := exec.Command(hkPath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func hkURL() string {
	return "https://s3.amazonaws.com/dickeyxxx_dev/releases/hk_" + runtime.GOOS + "_" + runtime.GOARCH + ".gz"
}

func homeDir() string {
	user, err := user.Current()
	Must(err)
	return user.HomeDir
}

func downloadHk(hkPath string) {
	fmt.Println("Downloading hk...")
	fmt.Println(filepath.Dir(hkPath))
	Must(os.MkdirAll(filepath.Dir(hkPath), 0777))
	out, err := os.Create(hkPath)
	Must(err)
	defer out.Close()
	if runtime.GOOS != "windows" {
		Must(out.Chmod(0777))
	}
	resp, err := http.Get(hkURL())
	Must(err)
	defer resp.Body.Close()
	uncompressed, err := gzip.NewReader(resp.Body)
	Must(err)
	_, err = io.Copy(out, uncompressed)
	Must(err)
	fmt.Println("Downloaded hk...")
}
