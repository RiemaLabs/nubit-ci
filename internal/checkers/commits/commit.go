package commits

import (
	"os/exec"

	"github.com/RiemaLabs/nubit-ci/internal/checkers"
)

type Checker struct{}

func (*Checker) Name() string { return "commit-message" }

func (*Checker) Install() error {
	return checkers.Tee(exec.Command(
		"npm",
		"i",
		"-g",
		"@commitlint/cli",
		"@commitlint/config-conventional",
	)).Run()
}

func (*Checker) Check() error {
	return checkers.Tee(exec.Command(
		"npx",
		"commitlint",
		"--from=origin/main",
		"--extends=@commitlint/config-conventional",
		"-V",
	)).Run()
}

func (*Checker) Fix() error { return nil }
