package logs

import (
	"log"
	"os"
)

var L = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
