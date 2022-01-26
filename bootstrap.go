//go:build windows
// +build windows

package gdm

import (
	"github.com/pkg/errors"
	"syscall"
	"unsafe"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

var (
	dmReg32         *syscall.LazyDLL
	procSetDllPathA *syscall.LazyProc
	procSetDllPathW *syscall.LazyProc
)

// DmSoft ...
type DmSoft struct {
	dm       *ole.IDispatch
	IUnknown *ole.IUnknown
}

// New return *DmSoft.DmSoft
func New(dllPath string) (dm *DmSoft, err error) {
	dmReg32 = syscall.NewLazyDLL(dllPath + `\DmReg.dll`)

	procSetDllPathA = dmReg32.NewProc("SetDllPathA")
	procSetDllPathW = dmReg32.NewProc("SetDllPathW")

	if !SetDllPathW(dllPath+`\dm.dll`, 1) {
		return nil, errors.New("load dm.dll failed")
	}

	var com DmSoft
	// 创建对象
	ole.CoInitialize(0)
	com.IUnknown, err = oleutil.CreateObject("dm.dmsoft")
	if err != nil {
		return nil, err
	}
	// 查询接口
	com.dm, err = com.IUnknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, err
	}
	return &com, nil
}

// Release 释放
func (com *DmSoft) Release() {
	com.IUnknown.Release()
	com.dm.Release()
	ole.CoUninitialize()
}

// SetDllPathA Ascii
func SetDllPathA(path string, mode int) bool {
	var _p0 *uint16
	_p0, _ = syscall.UTF16PtrFromString(path)
	ret, _, _ := syscall.Syscall(procSetDllPathA.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)
	return ret != 0
}

// SetDllPathW Unicode
func SetDllPathW(path string, mode int) bool {
	var _p0 *uint16
	_p0, _ = syscall.UTF16PtrFromString(path)
	ret, _, _ := syscall.Syscall(procSetDllPathW.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)
	return ret != 0
}
