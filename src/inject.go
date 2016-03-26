// Injection part of EVEModX
// Copyright 2016 masahoshiro
package evemodx

import (
	"os/exec"
	"fmt"
)

func CallExecutable(pid string, code string) {
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