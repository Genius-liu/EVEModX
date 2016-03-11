// EVEModX Main
// Works under Microsoft Windows platform
// Still under HEAVY developing
//
// WARNING: THIS IS A HACK OF EVE ONLINE.
// ANY USE OF THIS CODE MAY BE PUNISHED BY
// THE OFFICIAL. USE AT YOUR OWN RISK.

package main

import (
	"fmt"
	"syscall"
	"unsafe"
	//"strconv"
	"bytes"
	"os"
	"os/exec"
	"log"
	"runtime"
	"path/filepath"
	"strings"
	"io/ioutil"

)

const (
	VERSION = "0.0.1a"
)


var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

type ulong int32

type ulong_ptr uintptr

type PROCESSENTRY32 struct {
	dwSize ulong
	cntUsage ulong
	th32ProcessID ulong
	th32DefaultHeapID ulong_ptr
	th32ModuleID ulong
	cntThreads ulong
	th32ParentProcessID ulong
	pcPriClassBase ulong
	dwFlags ulong
	szExeFile [260]byte
}

func getGamePids() []int {

	kernel32 := syscall.NewLazyDLL("kernel32.dll");
	CreateToolhelp32Snapshot := kernel32.NewProc("CreateToolhelp32Snapshot");
	pHandle ,_ ,_ := CreateToolhelp32Snapshot.Call(uintptr(0x2),uintptr(0x0));
	var results []int
	if int(pHandle) == -1 {
		fmt.Println("Unable to create pHandle")
		os.Exit(1)
	}
	Process32Next := kernel32.NewProc("Process32Next");
	
	for {
		var proc PROCESSENTRY32;
		proc.dwSize = ulong(unsafe.Sizeof(proc));
		if rt, _, _ := Process32Next.Call(uintptr(pHandle), uintptr(unsafe.Pointer(&proc))); int(rt)==1 {
			if string(bytes.TrimRight(proc.szExeFile[0:], "\x00")) == "exefile.exe" {
				//results[0] = int( proc.th32ProcessID)
				results = append(results, int(proc.th32ProcessID))
				//fmt.Println("ProcessName : "+ string(bytes.TrimRight(proc.szExeFile[0:], "\x00")));
				//fmt.Println("ProcessID : "+ strconv.Itoa(int(proc.th32ProcessID)));

			}
		} else {
			break;
		}
	}
	CloseHandle := kernel32.NewProc("CloseHandle");
	_, _, _ = CloseHandle.Call(pHandle);
	return results
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
	//return dir
}

func getMods() []string {

	
	modReaderDir, _ := ioutil.ReadDir("./mods/")
	var mods []string
	for _, fileInfo := range modReaderDir {
		if fileInfo.IsDir() {
			mods = append(mods, fileInfo.Name())
		}
	}
	return mods
}

func callExecutable(pid string, code string) {
	fmt.Println(pid)
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


func main() {

	// MAXPROCS
	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)

	// Get current game pids
	exeFilePid := getGamePids()

	pid := fmt.Sprintf("%d",exeFilePid[0])


	// Get mod directory
	currentModDirectory := getCurrentDirectory() + "/mods/"

	// Get mod list
	mods := getMods()
	


	logger.Println(fmt.Sprintf("[INFO] EVEModX %s start", VERSION))
	logger.Println(fmt.Sprintf("[INFO] CPU number: %d", cpuNum))
	logger.Println(fmt.Sprintf("[INFO] Current mod directory: [%s]", currentModDirectory))
	logger.Println(fmt.Sprintf("[INFO] Existing mods: %s", mods))
	
	// THIS IS FUCKED -> pid := string(exeFilePid[0])
	logger.Println(fmt.Sprintf("[INFO] Using pid %d", exeFilePid[0]))
	

	code := `import sys;sys.path.append('` + currentModDirectory + `');import ` + mods[0] + ``

	logger.Println(fmt.Sprintf("[INFO] Using payload [%s]", code))
	
	callExecutable(pid, code)
	
}

