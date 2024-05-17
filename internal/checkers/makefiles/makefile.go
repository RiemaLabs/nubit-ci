package makefiles

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/RiemaLabs/nubit-ci/internal/checkers"
)

var expectedRules = []string{
	"build",
	"ci-fix",
	"ci",
	"clean",
}

type prefix struct {
	prefix string
	found  bool
}

type Checker struct {
	prefixes []*prefix
}

func (*Checker) Name() string { return "makefile" }

func (c *Checker) Install() error {
	for _, rule := range expectedRules {
		c.prefixes = append(c.prefixes, &prefix{prefix: rule + ":"})
		c.prefixes = append(c.prefixes, &prefix{prefix: ".PHONY: " + rule})
	}
	return nil
}

func (c *Checker) Check() error {
	f, err := os.Open("Makefile")
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		for _, p := range c.prefixes {
			if !strings.HasPrefix(line, p.prefix) {
				continue
			}
			p.found = true
			break
		}
	}

	var missing []string
	for _, p := range c.prefixes {
		if !p.found {
			missing = append(missing, fmt.Sprintf("%q", p.prefix))
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing rules: %v", strings.Join(missing, ", "))
	}
	return nil
}

func (*Checker) Fix() error { return checkers.ErrNoFixes }
