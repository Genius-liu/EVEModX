// Misc
// Copyright 2016 masahoshiro
package evemodx

import (
	"fmt"
	"log"
	"os"
)

// Logger specifies a logger.
var (
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
)

// Separator.
func PrintSprt() {
	logger.Println(fmt.Sprintf("[SPRT] ---------------------------------------------"))
}