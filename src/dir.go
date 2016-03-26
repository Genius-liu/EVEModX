// Directory functions
// Copyright 2016 masahoshiro
package evemodx

import (
	"path/filepath"
	"log"
	"strings"
	"io/ioutil"
	"os"
)

// GetCurrentDirectory() returns a string of current directory.
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
	//return dir
}

// GetMods() returns a slice of all Dirs(Mods) in ./mods .
func GetMods() []string {
	
	modReaderDir, _ := ioutil.ReadDir("./mods/")
	var mods []string
	for _, fileInfo := range modReaderDir {
		if fileInfo.IsDir() {
			mods = append(mods, fileInfo.Name())
		}
	}
	return mods
}
