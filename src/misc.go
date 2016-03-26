package evemodx

import (
	"fmt"
	"log"
	"os"
)

// Logger specifies a logger.
var (
	Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
)

// PrintSprt is a separator.
func PrintSprt() {
	Logger.Println(fmt.Sprintf("[SPRT] ---------------------------------------------"))
}