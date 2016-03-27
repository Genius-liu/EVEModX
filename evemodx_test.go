//go:generate goversioninfo -icon=icon.ico

// EVEModX Main
// Works under Microsoft Windows platform
// Still under HEAVY developing
//
// WARNING: THIS IS A HACK OF EVE ONLINE.
// ANY USE OF THIS CODE MAY BE PUNISHED BY
// THE OFFICIAL. USE AT YOUR OWN RISK.

// author masahoshiro@github

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
	importMods string
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

	injectAll := emx.ReadConf("common.injectall")

	LABEL1:	
	// Get current game pids
	exeFilePid = emx.GetGamePids()

	if len(exeFilePid) > 0 {
		var pid, i int
		var modName, modIndexSingle string
		if len(exeFilePid) == 1 {
			emx.Logger.Printf(fmt.Sprintf("[INFO] Using pid: %d as game process", exeFilePid[0]))
			pidIndex = 0

			goto LABEL2
		}

		//emx.Logger.Printf(fmt.Sprintf("[INFO] Listing current game process"))
		for i, pid = range exeFilePid {
			emx.Logger.Printf(fmt.Sprintf("[INFO] EXEFILE %d: %d", i, pid))
		}
		emx.PrintSprt()
		emx.Logger.Printf(fmt.Sprintf("[PRMT] Please input pid index (0~%d ,default for 0): ", len(exeFilePid) - 1 ))
		emx.Logger.Printf(fmt.Sprintf("[INFO] You should enter it after character choosing"))

		fmt.Scanln(&pidIndex)

		LABEL2:
		emx.PrintSprt()
		for i, modName = range mods {
			emx.Logger.Printf(fmt.Sprintf("[INFO] MOD %d: %s", i + 1, modName))
		}
		emx.PrintSprt()

		if injectAll == "true" {
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
				importMods = importMods + "import " + mod + ";"
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

	// FUCKED -> LABEL2:

	

/*	// [OLD]
	// Build import string
	importMods := ""

	var mod string
	for _, mod = range mods {
		importMods = importMods + "import " + mod + ";"
	}
*/

	// Build import string
	importMods = ""

	for i := 0; i < len(modIndexInt); i++ {
		importMods = importMods + "import " + mods[modIndexInt[i]] + ";"
	}

	LABEL3:

	pid = fmt.Sprintf("%d",exeFilePid[pidIndex])

	emx.PrintSprt()

	// THIS IS FUCKED -> pid := string(exeFilePid[0])
	emx.Logger.Println(fmt.Sprintf("[INFO] Using pid %d", exeFilePid[pidIndex]))

	code = `import sys;sys.path.append('` + currentModDirectory + `');` + importMods + ``
	
	emx.Logger.Println(fmt.Sprintf("[INFO] Using payload [%s]", code))
	emx.Logger.Println(fmt.Sprintf("[INFO] Executing injection"))

	emx.Inject(pid, code)

}

