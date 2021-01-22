package gdm

import "github.com/rogeecn/gdm/utils"

func (com *DmSoft) DmGuard(enable int, t string) int {
	ret, _ := com.dm.CallMethod("DmGuard", enable, t)
	return int(ret.Val)
}

func (com *DmSoft) DmGuardParams(cmd, subCmd, param string) string {
	ret, _ := com.dm.CallMethod("DmGuardParams", cmd, subCmd, param)
	return ret.ToString()
}
func (com *DmSoft) UnLoadDriver() bool {
	ret, _ := com.dm.CallMethod("UnLoadDriver")
	return utils.IsOK(ret.Val)
}
