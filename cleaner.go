package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type Options struct {
	ProjectDir string
	LogDir     string
}

type Cleaner interface {
	clean() error
}

func Clean(c Cleaner) {
	err := c.clean()
	if err != nil {
		fmt.Println(err)
	}
}

type LogFiles struct {
	Options
}

func (lf *LogFiles) clean() error {
	logFiles, err := lf.find()
	if err != nil {
		return err
	}
	if len(logFiles) == 0 {
		return errors.New("No log files found.")
	}
	messages := lf.remove(logFiles)
	for _, message := range messages {
		fmt.Println(message)
	}
	return nil
}

func (lf *LogFiles) find() ([]os.FileInfo, error) {
	logDir := dir + "/" + logDir
	logFiles, err := ioutil.ReadDir(logDir)
	return logFiles, err
}

func (lf *LogFiles) remove(logFiles []os.FileInfo) []string {
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

type GitBranches struct {
	Options
}

func (gb *GitBranches) clean() error {
	messages, err := gb.remove()
	if err != nil {
		return err
	}
	for _, message := range messages {
		fmt.Println(message)
	}
	return nil
}

func (gb *GitBranches) remove() ([]string, error) {
	messages := []string{}
	mergedBranches := exec.Command("git", "branch", "--merged")
	getMergedBranches := exec.Command("grep", "-vE", "master|develop")
	removeMergedBranches := exec.Command("xargs", "-n 1", "git", "branch", "-d")
	out, err := pipeCommands(mergedBranches, getMergedBranches, removeMergedBranches)
	if err != nil {
		return nil, errors.New("Error removing branches. Are you in the default branch? Do you still have unstaged changes?")
	}
	if len(out) != 0 {
		messages = append(messages, string(out))
		messages = append(messages, "Succesfully cleaned up branches.")
		return messages, nil
	}
	messages = append(messages, "No branches had to be cleaned up!")
	return messages, nil
}

func pipeCommands(commands ...*exec.Cmd) ([]byte, error) {
	for i, command := range commands[:len(commands)-1] {
		out, err := command.StdoutPipe()
		if err != nil {
			return nil, err
		}
		command.Start()
		commands[i+1].Stdin = out
	}
	final, err := commands[len(commands)-1].Output()
	if err != nil {
		return nil, err
	}
	return final, nil
}
