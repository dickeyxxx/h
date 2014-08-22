package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func runRubyCli(args ...string) (int, error) {
	downloadRubyCli()
	args = append([]string{rubyBin()}, args...)
	cmd := exec.Command("ruby", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus(), nil
			}
		}
		return -1, err
	}
	return 0, nil
}

func downloadRubyCli() {
	tmp := filepath.Join(homeDir(), "tmp")
	err := os.RemoveAll(tmp)
	must(err)
	err = os.MkdirAll(tmp, 0777)
	must(err)
	resp, err := http.Get("http://assets.heroku.com.s3.amazonaws.com/heroku-client/heroku-client.tgz")
	must(err)
	defer resp.Body.Close()
	uncompressed, err := gzip.NewReader(resp.Body)
	must(err)
	archive := tar.NewReader(uncompressed)
	for {
		hdr, err := archive.Next()
		if err == io.EOF {
			break
		}
		must(err)
		path := filepath.Join(tmp, hdr.Name)
		if hdr.FileInfo().IsDir() {
			err = os.Mkdir(path, 0777)
			must(err)
		} else {
			file, err := os.Create(path)
			_, err = io.Copy(file, archive)
			must(err)
		}
	}
	from := filepath.Join(tmp, "heroku-client")
	err = os.RemoveAll(rubyHome())
	must(err)
	err = os.Rename(from, rubyHome())
	must(err)
	err = os.RemoveAll(tmp)
	must(err)
}

func rubyHome() string {
	return filepath.Join(homeDir(), "ruby")
}

func rubyBin() string {
	return filepath.Join(rubyHome(), "bin", "heroku")
}
