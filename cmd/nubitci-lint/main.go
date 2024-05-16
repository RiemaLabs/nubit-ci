package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/RiemaLabs/nubit-ci/internal/checkers"
	"github.com/RiemaLabs/nubit-ci/internal/checkers/analyzers"
	"github.com/RiemaLabs/nubit-ci/internal/checkers/codeformats"
	"github.com/RiemaLabs/nubit-ci/internal/checkers/commits"
	"github.com/RiemaLabs/nubit-ci/internal/logs"
)

var C = []checkers.Checker{
	new(commits.Checker),
	new(codeformats.Checker),
	new(analyzers.Checker),
}

func main() {
	var (
		isListCheckers, isFix bool
		skips, only           string
	)
	flag.BoolVar(&isListCheckers, "l", false, "list available checkers")
	flag.BoolVar(&isFix, "w", false, "apply fixes (if available(")
	flag.StringVar(&skips, "skips", "", "comma separated list of skipped rules")
	flag.StringVar(&only, "only", "", "run this rule only")
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

		if only != "" {
			if only != c.Name() {
				continue
			} else {
				logs.L.Println("Run only:", c.Name())
			}
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
