package command

import (
	"bytes"
	"log"
	"os/exec"
)

type Commander interface {
	Execute(name string, args ...string) (string, error)
}

func NewCommander() Commander {
	return &commander{}
}

type commander struct{}

func (c *commander) Execute(name string, args ...string) (message string, err error) {
	cmd := exec.Command(name, args...)
	stdOut := GetStdOut(cmd)
	if err := cmd.Run(); err != nil {
		log.Println(err.Error(), stdOut.String())
	}
	return stdOut.String(), err
}

func GetStdOut(cmd *exec.Cmd) *bytes.Buffer {
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	return &out
}
