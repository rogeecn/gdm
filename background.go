package gdm

import (
	"github.com/rogeecn/gdm/utils"
)

const (
	DISPLAY_NORMAL = "normal"
	DISPLAY_GDI    = "gdi"
	DISPLAY_GDI2   = "gdi2"
	DISPLAY_DX2    = "dx2"
	DISPLAY_DX3    = "dx3"
	DISPLAY_DX     = "dx"
)

const (
	MOUSE_NORMAL   = "normal"
	MOUSE_WINDOWS  = "windows"
	MOUSE_WINDOWS2 = "windows2"
	MOUSE_WINDOWS3 = "windows3"
	MOUSE_DX       = "dx"
	MOUSE_DX2      = "dx2"
)

const (
	KEYPAD_NORMAL  = "normal"
	KEYPAD_WINDOWS = "windows"
	KEYPAD_DX      = "dx"
)

const (
	MODE_0   = 0
	MODE_2   = 2
	MODE_11  = 11
	MODE_13  = 13
	MODE_101 = 101
	MODE_103 = 103
)

func (com *DmSoft) BindWindow(hwnd int, display string, mouse string, keypad string, mode int) bool {
	ret, _ := com.dm.CallMethod("BindWindow", hwnd, display, mouse, keypad, mode)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) BindWindowEx(hwnd int, display string, mouse string, keypad string, public string, mode int) bool {
	ret, _ := com.dm.CallMethod("BindWindowEx", hwnd, display, mouse, keypad, public, mode)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) DownCPU(rate int) bool {
	ret, _ := com.dm.CallMethod("DownCpu", rate)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) EnableBind(enable int) bool {
	ret, _ := com.dm.CallMethod("EnableBind", enable)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) EnableFakeActive(enable int) bool {
	ret, _ := com.dm.CallMethod("EnableFakeActive", enable)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) EnableIme(enable int) bool {
	ret, _ := com.dm.CallMethod("EnableIme", enable)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) EnableKeypadMsg(enable int) bool {
	ret, _ := com.dm.CallMethod("EnableKeypadMsg", enable)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) EnableKeypadPatch(enable int) bool {
	ret, _ := com.dm.CallMethod("EnableKeypadPatch", enable)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) EnableKeypadSync(enable, timeOut int) bool {
	ret, _ := com.dm.CallMethod("EnableKeypadSync", enable, timeOut)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) EnableMouseMsg(enable int) bool {
	ret, _ := com.dm.CallMethod("EnableMouseMsg", enable)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) EnableMouseSync(enable, timeOut int) bool {
	ret, _ := com.dm.CallMethod("EnableMouseSync", enable, timeOut)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) EnableRealKeypad(enable int) bool {
	ret, _ := com.dm.CallMethod("EnableRealKeypad", enable)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) EnableRealMouse(enable, mousedelay, mousestep int) bool {
	ret, _ := com.dm.CallMethod("EnableRealMouse", enable, mousedelay, mousestep)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) EnableSpeedDx(enable int) bool {
	ret, _ := com.dm.CallMethod("EnableSpeedDx", enable)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) ForceUnBindWindow() bool {
	ret, _ := com.dm.CallMethod("ForceUnBindWindow")
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) GetBindWindow() bool {
	ret, _ := com.dm.CallMethod("GetBindWindow")
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) GetFps() int {
	ret, _ := com.dm.CallMethod("GetFps")
	return int(ret.Val)
}

func (com *DmSoft) HackSpeed(rate int) int {
	ret, _ := com.dm.CallMethod("HackSpeed", rate)
	return int(ret.Val)
}

func (com *DmSoft) IsBind(hwnd int) bool {
	ret, _ := com.dm.CallMethod("IsBind", hwnd)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) LockDisplay(lock int) bool {
	ret, _ := com.dm.CallMethod("LockDisplay", lock)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) LockInput(lock int) bool {
	ret, _ := com.dm.CallMethod("LockInput", lock)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) LockMouseRect(r utils.Rect) bool {
	ret, _ := com.dm.CallMethod("LockMouseRect", r.Left, r.Top, r.Right, r.Bottom)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetAero(enable int) bool {
	ret, _ := com.dm.CallMethod("SetAero", enable)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetDisplayDelay(time int) bool {
	ret, _ := com.dm.CallMethod("SetDisplayDelay", time)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetDisplayRefreshDelay(time int) bool {
	ret, _ := com.dm.CallMethod("SetDisplayRefreshDelay", time)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SwitchBindWindow(hwnd int) bool {
	ret, _ := com.dm.CallMethod("SwitchBindWindow", hwnd)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) UnBindWindow() bool {
	ret, _ := com.dm.CallMethod("UnBindWindow")
	return utils.IsOK(ret.Val)
}
