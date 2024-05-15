package codeformats

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/RiemaLabs/nubit-ci/internal/checkers"
)

type Checker struct{}

func (*Checker) Name() string   { return "code-format" }
func (*Checker) Install() error { return nil }

const formatter = "golang.org/x/tools/cmd/goimports@latest"

func (*Checker) Check() error {
	raw, err := checkers.GoRun(formatter, "-l", ".").CombinedOutput()
	if err != nil {
		fmt.Println(string(raw))
		return err
	}
	if output := string(bytes.TrimSpace(raw)); output != "" {
		files := strings.Split(output, "\n")
		fmt.Printf("⚠️ %d file(s) not formatted:\n\n", len(files))
		for _, f := range files {
			fmt.Println("  *", f)
		}
		fmt.Println()
		return errors.New("code not formatted")
	}
	return nil
}

func (*Checker) Fix() error {
	return checkers.Tee(checkers.GoRun(formatter, "-w", ".")).Run()
}
