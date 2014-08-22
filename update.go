package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func autoupdatePath() string {
	return filepath.Join(homeDir(), "autoupdate.log")
}

func shouldAutoupdate() bool {
	if f, err := os.Stat(autoupdatePath()); err == nil {
		return f.ModTime().Add(4 * time.Second).Before(time.Now())
	}
	return true
}

func autoupdate() {
	err := os.MkdirAll(filepath.Base(autoupdatePath()), 0777)
	must(err)
	file, err := os.OpenFile(autoupdatePath(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	must(err)
	defer file.Close()
	logger := log.New(file, "", log.LstdFlags)
	logger.Println("checking for update")
}
