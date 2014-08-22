package main

import (
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	hkPath := filepath.Join(HomeDir(), ".hk", "hk")
	if runtime.GOOS == "windows" {
		hkPath = hkPath + ".exe"
	}
	exists, err := FileExists(hkPath)
	Must(err)
	if !exists {
		updateHk(hkPath)
	}
	err = run(hkPath, os.Args)
	Must(err)
}
