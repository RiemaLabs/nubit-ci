package codeformats

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/RiemaLabs/nubit-ci/internal/checkers"
)

type Checker struct{}

func (*Checker) Name() string   { return "go-format" }
func (*Checker) Install() error { return nil }

const formatter = "golang.org/x/tools/cmd/goimports@latest"

func (*Checker) Check() error {
	raw, err := checkers.GoRun(formatter, "-l", ".").CombinedOutput()
	if err != nil {
		fmt.Println(string(raw))
		return err
	}
	if output := string(bytes.TrimSpace(raw)); output != "" {
		var files []string
		for _, s := range strings.Split(output, "\n") {
			line := strings.TrimSpace(s)
			if strings.HasPrefix(line, "go:") {
				continue
			}
			files = append(files, line)
		}
		if len(files) > 0 {
			fmt.Printf("⚠️ %d file(s) not formatted:\n\n", len(files))
			for _, f := range files {
				fmt.Println("  *", f)
			}
			fmt.Println()
			return errors.New("code not formatted")
		}
	}
	return nil
}

func (*Checker) Fix() error {
	return checkers.Tee(checkers.GoRun(formatter, "-w", ".")).Run()
}
