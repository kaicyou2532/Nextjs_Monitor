package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func isRunning(url string) bool {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	return resp.StatusCode < 500
}

func isProcessRunning(pattern string) bool {
	cmd := exec.Command("pgrep", "-f", pattern)
	return cmd.Run() == nil
}

func startProcess(dir string) error {
	cmd := exec.Command("npm", "run", "dev")
	cmd.Dir = dir
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()
	if err := cmd.Start(); err != nil {
		return err
	}
	go cmd.Wait()
	return nil
}

func main() {
	dir := flag.String("dir", ".", "directory of the web app")
	url := flag.String("url", "http://localhost:3000", "URL to check")
	interval := flag.Duration("interval", time.Minute, "check interval")
	pattern := flag.String("pattern", "npm.*run.*dev", "pgrep pattern for the dev process")
	flag.Parse()

	check := func() {
		if !isRunning(*url) {
			if !isProcessRunning(*pattern) {
				log.Printf("process not running, starting in %s", *dir)
				if err := startProcess(*dir); err != nil {
					log.Printf("failed to start process: %v", err)
				}
			} else {
				log.Printf("process running but health check failed")
			}
		}
	}

	check()
	ticker := time.NewTicker(*interval)
	for range ticker.C {
		check()
	}
}
