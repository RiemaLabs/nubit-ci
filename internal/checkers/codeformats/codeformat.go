package codeformats

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/RiemaLabs/nubit-ci/internal/checkers"
)

type Checker struct{}

func (*Checker) Name() string   { return "code-format" }
func (*Checker) Install() error { return nil }

const formatter = "github.com/golangci/golangci-lint/cmd/golangci-lint@latest"

func (*Checker) Check() error {
	raw, err := checkers.GoRun(formatter, "run", "--show-stats").CombinedOutput()
	if err != nil {
		fmt.Println(string(raw))
		return err
	}
	fmt.Println()
	if output := string(bytes.TrimSpace(raw)); output != "" {
		lines := strings.Split(output, "\n")
		cnt := -1
		errInfo := ""
		statsInfo := ""
		for i, s := range lines {
			line := strings.TrimSpace(s)
			if strings.Contains(line, " issues") {
				parts := strings.Split(line, " ")
				cnt, _ = strconv.Atoi(parts[0])
				if cnt != 0 {
					errInfo = strings.Join(lines[:i], "\n")
					statsInfo = strings.Join(lines[i:len(lines)-1], "\n")
				} else {
					statsInfo = strings.Join(lines[i:], "\n")
				}
				break
			}
		}
		if cnt == 0 {
			fmt.Println("✅ No issue found! Code passed the lint check!")
		} else if cnt > 0 {
			fmt.Printf("⚠️ %d issue(s) found by lint:\n", cnt)
			fmt.Println("Error info: ")
			fmt.Println(errInfo)
			fmt.Println("Statistic info")
			fmt.Println(statsInfo)
			return errors.New("code lint check failed")
		} else {
			return errors.New("unexpected errors occurred when run golangci-lint")
		}
	}
	return nil
}

func (*Checker) Fix() error { return nil }
