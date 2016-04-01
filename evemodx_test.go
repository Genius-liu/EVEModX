//go:generate goversioninfo -icon=icon.ico

// EVEModX Main
// Works under Microsoft Windows platform
// Still under HEAVY developing
//
// WARNING: THIS IS A HACK OF EVE ONLINE.
// ANY USE OF THIS CODE MAY BE PUNISHED BY
// THE OFFICIAL. USE AT YOUR OWN RISK.

// author masahoshiro@github

// THIS CODE IS SHIT. I'M GOING TO REWRITE IT!

package main

import (
	"fmt"
	"runtime"
	"time"
	"os"
	"strings"
	"strconv"
	emx "evemodx"
)

const (
	VERSION = "0.0.2"
)


var (
	pidIndex int
	exeFilePid []int
	modIndex string
	modIndexString []string
	modIndexInt []int
	pid string
	payload string
	code string
) 


func main() {

	// Set MAXPROCS
	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)

	// Get mod directory
	currentModDirectory := emx.GetCurrentDirectory() + "/mods/"

	// Get mod list
	mods := emx.GetMods()
	
	emx.Logger.Println(fmt.Sprintf("[INFO] EVEModX %s starting", VERSION))
	emx.Logger.Println(fmt.Sprintf("[INFO] CPU number: %d", cpuNum))
	emx.Logger.Println(fmt.Sprintf("[INFO] Current mod directory: [%s]", currentModDirectory))
	emx.Logger.Println(fmt.Sprintf("[INFO] Existing mods: %s", mods))
	emx.Logger.Println(fmt.Sprintf("[INFO] Awaiting for game process..."))
	emx.PrintSprt()




	injectAllMod := emx.ReadConf("common.injectallmod")
	injectAllExe := emx.ReadConf("common.injectallexe")

	LABEL1:	
	// Get current game pids
	exeFilePid = emx.GetGamePids()

/*	for i, pid := range exeFilePid {
		emx.Logger.Printf(fmt.Sprintf("[INFO] EXEFILE %d: %d", i + 1 , pid))
	}
	emx.PrintSprt()
*/
	if len(exeFilePid) > 0 {
		var i int
		var modName, modIndexSingle string
		if len(exeFilePid) == 1 {
			emx.Logger.Printf(fmt.Sprintf("[INFO] Using pid: %d as game process", exeFilePid[0]))
			pidIndex = 0

			goto LABEL2
		}

		if injectAllExe == "true" {
			emx.Logger.Printf(fmt.Sprintf("[INFO] Inject all game process" ))
			goto LABEL2
		}

		//emx.Logger.Printf(fmt.Sprintf("[INFO] Listing current game process"))
		for i, pid := range exeFilePid {
			emx.Logger.Printf(fmt.Sprintf("[INFO] EXEFILE %d: %d", i + 1 , pid))
		}
		emx.PrintSprt()
		emx.Logger.Printf(fmt.Sprintf("[PRMT] Please input pid index (0~%d ,0 for all): ", len(exeFilePid) - 1 ))
		emx.Logger.Printf(fmt.Sprintf("[INFO] You should enter it after character choosing"))

		fmt.Scanln(&pidIndex)

		if pidIndex == 0 {
			emx.Logger.Printf(fmt.Sprintf("[INFO] Inject all game process" ))
			injectAllExe = "true"
			goto LABEL2
		}

		pidIndex--

		LABEL2:
		emx.PrintSprt()
		for i, modName = range mods {
			emx.Logger.Printf(fmt.Sprintf("[INFO] MOD %d: %s", i + 1, modName))
		}
		emx.PrintSprt()

		if injectAllMod == "true" {
			modIndex = "0"
			emx.Logger.Printf(fmt.Sprintf("[INFO] Use all mods as config set"))
			goto LABEL4
		}

		emx.Logger.Printf(fmt.Sprintf("[PRMT] Please input mods index (eg: 1,3;0 for all):" ))
		emx.Logger.Printf(fmt.Sprintf("[INFO] You should enter it after character choosing"))

		fmt.Scanln(&modIndex)

		LABEL4:

		if modIndex == "0"{

			var mod string
			for _, mod = range mods {
				payload = payload + "import " + mod + ";"
			}

			goto LABEL3
		}
		modIndexString = strings.Split(modIndex, ",")

		for i, modIndexSingle = range modIndexString {
			b, error := strconv.Atoi(modIndexSingle)
			b = b - 1
			if error != nil{
				emx.Logger.Printf(fmt.Sprintf("[ERRO] Cannot convert input to slice"))
				os.Exit(1)
			}
			modIndexInt = append(modIndexInt, b )
		}

		//emx.Logger.Printf(fmt.Sprintf("%d", modIndexInt ))

	} else {
		time.Sleep(3 * time.Second)
		goto LABEL1
	}

	// Build payload
	
	payload = ""

	for i := 0; i < len(modIndexInt); i++ {
		payload = payload + "import " + mods[modIndexInt[i]] + ";"
	}

	LABEL3:
	emx.PrintSprt()
	code = `import sys;sys.path.append('` + currentModDirectory + `');` + payload + ``
	emx.Logger.Println(fmt.Sprintf("[INFO] Using payload [%s]", code))

	if injectAllExe == "true" {

		emx.Logger.Println(fmt.Sprintf("[INFO] Inject all pids"))

		for i, _ := range exeFilePid {
			pid = fmt.Sprintf("%d",exeFilePid[i])
			emx.Logger.Println(fmt.Sprintf("[INFO] Executing injection for %d", exeFilePid[i]))
			emx.Inject(pid, code)

		}
		os.Exit(1)
	}	

	emx.Logger.Println(fmt.Sprintf("[INFO] Executing injection for %d", exeFilePid[pidIndex]))
	pid = fmt.Sprintf("%d",exeFilePid[pidIndex])
	emx.Inject(pid, code)
}

