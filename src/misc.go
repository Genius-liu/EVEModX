// Misc
// Copyright 2016 masahoshiro
package evemodx

import (
	"fmt"
	"log"
	"os"
)

var (
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
)

func PrintSprt() {
	logger.Println(fmt.Sprintf("[SPRT] ---------------------------------------------"))
}