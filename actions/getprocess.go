package actions

import (
	"os"
	"syscall"
	"unsafe"
	"github.com/TheTitanrain/w32"
)

const MAX_PATH = 260

//https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/ns-tlhelp32-processentry32
type PROCESSENTRY32 struct {
	Size uint32
	CntUsage uint32
	Th32ProcessID uint32
	Th32DefaultHeapID uintptr
	Th32ModuleID uint32
	CntThreads uint32
	Th32ParentProcessID uint32
	PcPriClassBase int32
	DwFlags uint32
	SzExeFile [MAX_PATH]uint16
}

func GetAllProcesses() []uint32{
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	proc32First := kernel32.MustFindProc("Process32FirstW")
	proc32Next := kernel32.MustFindProc("Process32NextW")
	isX86 := kernel32.MustFindProc("IsWow64Process")
	openProc := kernel32.MustFindProc("OpenProcess")

	var x86u bool
	var x86PidArray []uint32
	currentPid := os.Getpid()

	snapshot := w32.CreateToolhelp32Snapshot(w32.TH32CS_SNAPPROCESS, 0) //This windows API gets snapshot of all processes in user space.
	var processEntry PROCESSENTRY32
	processEntry.Size = uint32(unsafe.Sizeof(processEntry))

	/*
	In this block, 32 bit all processes IDs is added to array by using Process32First, Process32Next, IsWow64Process and OpenProcess API calls.
	*/
	_, _, err := proc32First.Call(uintptr(snapshot), uintptr(unsafe.Pointer(&processEntry)))
	if err != nil {
		for {
			_, _, err := proc32Next.Call(uintptr(snapshot), uintptr(unsafe.Pointer(&processEntry)))
			if err == syscall.Errno(0) {
				processHandle, _, _ := openProc.Call(uintptr(0x1F0FFF), 0, uintptr(processEntry.Th32ProcessID))
				_, _, err := isX86.Call(uintptr(processHandle), uintptr(unsafe.Pointer(&x86u)))
				if err == syscall.Errno(0) && x86u == true {
					if processEntry.Th32ProcessID != uint32(currentPid) {
						x86PidArray = append(x86PidArray, processEntry.Th32ProcessID)
					}
				}
			} else {
				break
			}
		}
	}
	return x86PidArray
}
