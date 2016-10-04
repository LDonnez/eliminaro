package main

import (
	"flag"
	"fmt"
	"os"
)

var dir string
var logDir string

func init() {
	flag.StringVar(&dir, "dir", ".", "dir")
	flag.StringVar(&logDir, "log-dir", "log", "dir")
	flag.Parse()
	fmt.Println("logfiles from " + dir + " will be removed.")
}

func main() {
	logFiles := &LogFiles{Options{ProjectDir: dir, LogDir: logDir}}
	gitBranches := &GitBranches{Options{ProjectDir: dir, LogDir: logDir}}
	Clean(logFiles)
	Clean(gitBranches)
	os.Exit(1)
}
