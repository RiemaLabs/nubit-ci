package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/RiemaLabs/nubit-ci/internal/checkers"
	"github.com/RiemaLabs/nubit-ci/internal/checkers/codeformats"
	"github.com/RiemaLabs/nubit-ci/internal/checkers/commits"
	"github.com/RiemaLabs/nubit-ci/internal/logs"
)

var C = []checkers.Checker{
	new(commits.Checker),
	new(codeformats.Checker),
}

func main() {
	var (
		isListCheckers bool
		skips          string
		isFix          bool
	)
	flag.BoolVar(&isListCheckers, "l", false, "list available checkers")
	flag.StringVar(&skips, "skips", "", "comma separated list of skipped rules")
	flag.BoolVar(&isFix, "w", false, "apply fixes (if available(")
	flag.Parse()

	if isListCheckers {
		for _, c := range C {
			fmt.Println(c.Name())
		}
		return
	}

	skipSet := make(map[string]struct{})
	if skips != "" {
		for _, s := range strings.Split(skips, ",") {
			skipSet[strings.TrimSpace(s)] = struct{}{}
		}
	}

	for _, c := range C {
		if _, ok := skipSet[c.Name()]; ok {
			logs.L.Println("Skipping:", c.Name())
			continue
		}

		logs.L.Println("Checking:", c.Name())
		if err := c.Install(); err != nil {
			logs.L.Fatalf("Checker %q install failed: %v", c.Name(), err)
		}
		if err := c.Check(); err != nil {
			if !isFix {
				logs.L.Fatalf("Checker %q run failed: %v", c.Name(), err)
			}
			logs.L.Printf("Checker %q run failed and start fixing: %v", c.Name(), err)
			if err := c.Fix(); err != nil {
				logs.L.Fatalf("Checker %q fix failed: %v", c.Name(), err)
			}
		}
	}
}
