package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var dir string

func init() {
	flag.StringVar(&dir, "dir", "log", "dir")
	flag.Parse()
	if dir == "log" {
		dir = "."
	}
	fmt.Println("logfiles from " + dir + " will be removed.")
}

func main() {
	logDir := dir + "/log"
	logFiles, _ := ioutil.ReadDir(logDir)

	if len(logFiles) == 0 {
		fmt.Println("No log files found")
		os.Exit(3)
	}

	messages := removeLogFiles(logFiles, logDir)
	for _, message := range messages {
		fmt.Println(message)
	}
}

func removeLogFiles(logFiles []os.FileInfo, logDir string) []string {
	messages := []string{}
	for _, f := range logFiles {
		err := os.RemoveAll(filepath.Join(logDir, f.Name()))
		if err != nil {
			messages = append(messages, err.Error())
			continue
		}
		messages = append(messages, f.Name()+" is removed.")
	}
	return messages
}
