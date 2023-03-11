package system

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

const (
	Kernel32Dll = "kernel32.dll"
	User32Dll   = "user32.dll"
)

type Module struct {
	dllHandle syscall.Handle
	procMap   map[string]uintptr
}

var modulesMap map[string]*Module
var consoleHandle uintptr

func NewModule(dllName string, procNames ...string) (*Module, error) {
	var module *Module = nil
	moduleHandle, err := syscall.LoadLibrary(dllName)
	if err != nil {
		return module, err
	}
	for _, procName := range procNames {
		procAddr, err := syscall.GetProcAddress(moduleHandle, procName)
		if err != nil {
			return module, err
		}
		if module == nil {
			module = new(Module)
			module.procMap = make(map[string]uintptr)
		}
		module.procMap[procName] = procAddr
	}
	module.dllHandle = moduleHandle
	if modulesMap == nil {
		modulesMap = make(map[string]*Module)
	}
	modulesMap[dllName] = module
	return module, err
}

func FreeModule() {
	for _, module := range modulesMap {
		syscall.FreeLibrary(module.dllHandle)
	}
	for k := range modulesMap {
		delete(modulesMap, k)
	}
}

func GetModule(dllName string) *Module {
	for name, module := range modulesMap {
		if strings.EqualFold(name, dllName) {
			return module
		}
	}
	return nil
}

func (m *Module) Call(procName string, args ...uintptr) (r1, r2 uintptr, err syscall.Errno) {
	return syscall.SyscallN(m.procMap[procName], args...)
}

func SetTitle(title string) error {
	lpConsoleTitle, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return err
	}
	GetModule(Kernel32Dll).Call("SetConsoleTitleW", uintptr(unsafe.Pointer(lpConsoleTitle)))
	return nil
}

func InitConsoleHandle() error {
	r1, _, err := GetModule(Kernel32Dll).Call("GetConsoleWindow")
	if err != 0 {
		return err
	}
	consoleHandle = r1
	fmt.Println(consoleHandle)
	return nil
}

func DisableQuickEdit() error {
	var lpMode uintptr = 0
	r1, _, err := GetModule(Kernel32Dll).Call("GetConsoleMode", consoleHandle, lpMode)
	if err != 0 {
		fmt.Println(err)
		return err
	}
	fmt.Println(r1)
	return nil
}
