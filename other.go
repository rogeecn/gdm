package gdm

import "github.com/rogeecn/gdm/utils"

func (com *DmSoft) ActiveInputMethod(hwnd int, inputMethod string) bool {
	ret, _ := com.dm.CallMethod("ActiveInputMethod", hwnd, inputMethod)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) CheckInputMethod(hwnd int, inputMethod string) bool {
	ret, _ := com.dm.CallMethod("CheckInputMethod", hwnd, inputMethod)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) FindInputMethod(inputMethod string) bool {
	ret, _ := com.dm.CallMethod("FindInputMethod", inputMethod)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) EnterCri() bool {
	ret, _ := com.dm.CallMethod("EnterCri")
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) InitCri() bool {
	ret, _ := com.dm.CallMethod("InitCri")
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) LeaveCri() bool {
	ret, _ := com.dm.CallMethod("LeaveCri")
	return utils.IsOK(ret.Val)
}
func (com *DmSoft) ReleaseRef() bool {
	ret, _ := com.dm.CallMethod("ReleaseRef")
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetExitThread() bool {
	ret, _ := com.dm.CallMethod("SetExitThread")
	return utils.IsOK(ret.Val)
}
