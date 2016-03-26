package evemodx

import (
	"syscall"
	"fmt"
	//strconv
	"os"
	"bytes"
	"unsafe"
)

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

// GetGamePids returns a slice of all current exefile.exe's
// PID using WinAPI.
func GetGamePids() []int {

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
