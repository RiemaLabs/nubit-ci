package analyzers

import (
	"github.com/RiemaLabs/nubit-ci/internal/executl"
)

const analyzer = "github.com/golangci/golangci-lint/cmd/golangci-lint@latest"

var args = []string{
	"run",
	// GitHub Actions analyze the modcache which is weird.
	"--exclude-dirs=(^|/)mod($|/)",
	"--exclude-dirs-use-default",
	"--timeout=10m",
	"--tests",
}

type Checker struct{}

func (*Checker) Name() string   { return "go-analyze" }
func (*Checker) Install() error { return nil }

func (*Checker) Check() error {
	return executl.Tee(executl.GoRun(analyzer, args...)).Run()
}

func (*Checker) Fix() error {
	return executl.Tee(executl.GoRun(analyzer, append(args, "--fix")...)).Run()
}
