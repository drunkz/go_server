package system

import (
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	Kernel32Dll  = "kernel32.dll"
	User32Dll    = "user32.dll"
	SC_CLOSE     = 0xF060
	SC_MAXIMIZE  = 0xF030
	MF_BYCOMMAND = 0
	FALSE        = 0
)

type Module struct {
	dllHandle syscall.Handle
	procMap   map[string]uintptr
}

var modulesMap map[string]*Module
var windowHandle uintptr
var stdInputHandle windows.Handle

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
	windowHandle = r1

	stdHandle, e := windows.GetStdHandle(windows.STD_INPUT_HANDLE)
	if e != nil {
		return err
	}
	stdInputHandle = stdHandle
	return nil
}

func DisableQuickEdit() error {
	var mode uint32
	if err := windows.GetConsoleMode(stdInputHandle, &mode); err != nil {
		return err
	}
	mode &^= windows.ENABLE_QUICK_EDIT_MODE
	if err := windows.SetConsoleMode(stdInputHandle, mode); err != nil {
		return err
	}
	return nil
}

func RemoveMenu() error {
	r1, _, err := GetModule(User32Dll).Call("GetSystemMenu", windowHandle, FALSE)
	if err != 0 {
		return err
	}
	_, _, err = GetModule(User32Dll).Call("RemoveMenu", r1, SC_CLOSE, MF_BYCOMMAND)
	if err != 0 {
		return err
	}
	_, _, err = GetModule(User32Dll).Call("RemoveMenu", r1, SC_MAXIMIZE, MF_BYCOMMAND)
	if err != 0 {
		return err
	}
	return nil
}
