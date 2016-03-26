// Copyright 2016 The EVEModX Authors.
// Use of this code may be punished by the official.
// USE AT YOUR OWN RISK!

// Package evemodx(or emx) implements python injection
// into EVE Online game process (exefile.exe).
// 
// https://github.com/EVEModX/EVEModX
package evemodx

import (
	"os/exec"
	"fmt"
)

// Inject injects python code into process specified
// by PID.
// Haven't test whether it will work when there are
// line breaks in the code.
func Inject(pid string, code string) {
	//fmt.Println(pid)
	args := []string{pid, code}
	cmd := exec.Command("inject_python_32.exe", args...)
	//output, err := cmd.CombinedOutput()
	_, err := cmd.CombinedOutput()
	//fmt.Println(cmd)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//logger.Println("Call exetuable result:\n%v\n\n%v\n\n%v", string(output), cmd.Stdout, cmd.Stderr)
}
