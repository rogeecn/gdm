package gdm

import "github.com/rogeecn/gdm/utils"

func (com *DmSoft) EnablePicCache(enable int) bool {
	ret, _ := com.dm.CallMethod("EnablePicCache", enable)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) GetBasePath() string {
	ret, _ := com.dm.CallMethod("GetBasePath")
	return ret.ToString()
}

func (com *DmSoft) GetDmCount() int {
	ret, _ := com.dm.CallMethod("GetDmCount")
	return int(ret.Val)
}

func (com *DmSoft) GetID() int {
	ret, _ := com.dm.CallMethod("GetID")
	return int(ret.Val)
}

func (com *DmSoft) GetLastError() error {
	ret, _ := com.dm.CallMethod("GetLastError")
	return utils.ErrorMap[int(ret.Val)]
}

func (com *DmSoft) GetPath() string {
	ret, _ := com.dm.CallMethod("GetPath")
	return ret.ToString()
}

func (com *DmSoft) Reg(regCode string, verInfo string) error {
	ret, _ := com.dm.CallMethod("Reg", regCode, verInfo)
	return utils.RegErrorMap[int(ret.Val)]
}

func (com *DmSoft) RegEx(regCode, verInfo, ip string) error {
	ret, _ := com.dm.CallMethod("RegEx", regCode, verInfo, ip)
	return utils.RegErrorMap[int(ret.Val)]
}

func (com *DmSoft) RegExNoMac(regCode, verInfo, ip string) error {
	ret, _ := com.dm.CallMethod("RegExNoMac", regCode, verInfo, ip)
	return utils.RegErrorMap[int(ret.Val)]
}

func (com *DmSoft) RegNoMac(regCode, verInfo string) error {
	ret, _ := com.dm.CallMethod("RegNoMac", regCode, verInfo)
	return utils.RegErrorMap[int(ret.Val)]
}

func (com *DmSoft) SetDisplayInput(mode string) bool {
	ret, _ := com.dm.CallMethod("SetDisplayInput", mode)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetEnumWindowDelay(delay int) bool {
	ret, _ := com.dm.CallMethod("SetEnumWindowDelay", delay)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetPath(path string) bool {
	ret, _ := com.dm.CallMethod("SetPath", path)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetShowErrorMsg(show int) bool {
	ret, _ := com.dm.CallMethod("SetShowErrorMsg", show)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SpeedNormalGraphic(enable int) bool {
	ret, _ := com.dm.CallMethod("SpeedNormalGraphic", enable)
	return utils.IsOK(ret.Val)
}

// Ver get version
func (com *DmSoft) Ver() string {
	ver, _ := com.dm.CallMethod("Ver")
	return ver.ToString()
}
