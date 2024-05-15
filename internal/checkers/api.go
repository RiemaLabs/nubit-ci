package checkers

import (
	"os"
	"os/exec"
)

type Checker interface {
	Name() string
	Install() error
	Check() error
	Fix() error
}

// Tee streams the OS default standard output and error while running the command.
func Tee(c *exec.Cmd) *exec.Cmd {
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c
}

// GoRun runs a Go package with a URL and arguments.
func GoRun(gitURL string, args ...string) *exec.Cmd {
	return exec.Command("go", append([]string{"run", gitURL}, args...)...)
}
