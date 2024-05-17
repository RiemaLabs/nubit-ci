package checkers

import (
	"errors"
)

var ErrNoFixes = errors.New("no fixes available")

type Checker interface {
	Name() string
	Install() error
	Check() error
	Fix() error
}
