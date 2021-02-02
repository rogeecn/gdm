package gdm

import (
	"github.com/rogeecn/gdm/utils"
)

func (com *DmSoft) SendShortCut(vKey []int) {
	for _, vkey := range vKey {
		com.KeyDown(vkey)
	}

	for i := len(vKey) - 1; i >= 0; i-- {
		com.KeyUp(vKey[i])
	}
}

func (com *DmSoft) SetKeypadDelay(types string, delay int) bool {
	ret, _ := com.dm.CallMethod("SetKeypadDelay", types, delay)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) GetKeyState(vkCode int) bool {
	ret, _ := com.dm.CallMethod("GetKeyState", vkCode)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) KeyDown(vkCode int) bool {
	ret, _ := com.dm.CallMethod("KeyDown", vkCode)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) KeyPress(vkCode int) bool {
	ret, _ := com.dm.CallMethod("KeyPress", vkCode)
	return utils.IsOK(ret.Val)
}

// KeyPressStr指定的字符串序列，依次按顺序按下其中的字符
func (com *DmSoft) KeyPressStr(keyStr string, delay int) bool {
	ret, _ := com.dm.CallMethod("KeyPressStr", keyStr, delay)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) KeyUp(vkCode int) bool {
	ret, _ := com.dm.CallMethod("KeyUp", vkCode)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) WaitKey(vkCode int, timeOut int) bool {
	ret, _ := com.dm.CallMethod("SetSimMode", vkCode, timeOut)
	return utils.IsOK(ret.Val)
}
