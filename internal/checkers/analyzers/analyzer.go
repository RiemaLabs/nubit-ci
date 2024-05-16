package analyzers

import "github.com/RiemaLabs/nubit-ci/internal/checkers"

const analyzer = "github.com/golangci/golangci-lint/cmd/golangci-lint@latest"

type Checker struct{}

func (*Checker) Name() string   { return "go-analyze" }
func (*Checker) Install() error { return nil }

func (*Checker) Check() error {
	return checkers.Tee(checkers.GoRun(analyzer, "run")).Run()
}

func (*Checker) Fix() error {
	return checkers.Tee(checkers.GoRun(analyzer, "run", "--fix")).Run()
}
