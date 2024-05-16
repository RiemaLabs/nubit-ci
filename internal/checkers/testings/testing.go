package testings

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/RiemaLabs/nubit-ci/internal/checkers"
)

type Checker struct{}

func (*Checker) Name() string   { return "go-test" }
func (*Checker) Install() error { return nil }

func (*Checker) Check() error {
	raw, err := exec.Command("go", "test", ".").CombinedOutput()
	output := string(raw)
	if err != nil && !strings.Contains(output, "no Go files in ") {
		fmt.Println(output)
		return err
	}
	return checkers.Tee(exec.Command("go", "test", "./...")).Run()
}

func (*Checker) Fix() error { return checkers.ErrNoFixes }
