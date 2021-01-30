package gdm

import (
	"github.com/rogeecn/gdm/utils"
)

type VKEY int

const (
	VK_LBUTTON             VKEY = 0x01
	VK_RBUTTON             VKEY = 0x02
	VK_CANCEL              VKEY = 0x03
	VK_MBUTTON             VKEY = 0x04
	VK_XBUTTON1            VKEY = 0x05
	VK_XBUTTON2            VKEY = 0x06
	VK_BACK                VKEY = 0x08
	VK_TAB                 VKEY = 0x09
	VK_CLEAR               VKEY = 0x0C
	VK_RETURN              VKEY = 0x0D
	VK_SHIFT               VKEY = 0x10
	VK_CONTROL             VKEY = 0x11
	VK_MENU                VKEY = 0x12
	VK_PAUSE               VKEY = 0x13
	VK_CAPITAL             VKEY = 0x14
	VK_KANA                VKEY = 0x15
	VK_HANGEUL             VKEY = 0x15
	VK_HANGUL              VKEY = 0x15
	VK_JUNJA               VKEY = 0x17
	VK_FINAL               VKEY = 0x18
	VK_HANJA               VKEY = 0x19
	VK_KANJI               VKEY = 0x19
	VK_ESCAPE              VKEY = 0x1B
	VK_CONVERT             VKEY = 0x1C
	VK_NONCONVERT          VKEY = 0x1D
	VK_ACCEPT              VKEY = 0x1E
	VK_MODECHANGE          VKEY = 0x1F
	VK_SPACE               VKEY = 0x20
	VK_PRIOR               VKEY = 0x21
	VK_NEXT                VKEY = 0x22
	VK_END                 VKEY = 0x23
	VK_HOME                VKEY = 0x24
	VK_LEFT                VKEY = 0x25
	VK_UP                  VKEY = 0x26
	VK_RIGHT               VKEY = 0x27
	VK_DOWN                VKEY = 0x28
	VK_SELECT              VKEY = 0x29
	VK_PRINT               VKEY = 0x2A
	VK_EXECUTE             VKEY = 0x2B
	VK_SNAPSHOT            VKEY = 0x2C
	VK_INSERT              VKEY = 0x2D
	VK_DELETE              VKEY = 0x2E
	VK_HELP                VKEY = 0x2F
	VK_LWIN                VKEY = 0x5B
	VK_RWIN                VKEY = 0x5C
	VK_APPS                VKEY = 0x5D
	VK_SLEEP               VKEY = 0x5F
	VK_NUMPAD0             VKEY = 0x60
	VK_NUMPAD1             VKEY = 0x61
	VK_NUMPAD2             VKEY = 0x62
	VK_NUMPAD3             VKEY = 0x63
	VK_NUMPAD4             VKEY = 0x64
	VK_NUMPAD5             VKEY = 0x65
	VK_NUMPAD6             VKEY = 0x66
	VK_NUMPAD7             VKEY = 0x67
	VK_NUMPAD8             VKEY = 0x68
	VK_NUMPAD9             VKEY = 0x69
	VK_MULTIPLY            VKEY = 0x6A
	VK_ADD                 VKEY = 0x6B
	VK_SEPARATOR           VKEY = 0x6C
	VK_SUBTRACT            VKEY = 0x6D
	VK_DECIMAL             VKEY = 0x6E
	VK_DIVIDE              VKEY = 0x6F
	VK_F1                  VKEY = 0x70
	VK_F2                  VKEY = 0x71
	VK_F3                  VKEY = 0x72
	VK_F4                  VKEY = 0x73
	VK_F5                  VKEY = 0x74
	VK_F6                  VKEY = 0x75
	VK_F7                  VKEY = 0x76
	VK_F8                  VKEY = 0x77
	VK_F9                  VKEY = 0x78
	VK_F10                 VKEY = 0x79
	VK_F11                 VKEY = 0x7A
	VK_F12                 VKEY = 0x7B
	VK_F13                 VKEY = 0x7C
	VK_F14                 VKEY = 0x7D
	VK_F15                 VKEY = 0x7E
	VK_F16                 VKEY = 0x7F
	VK_F17                 VKEY = 0x80
	VK_F18                 VKEY = 0x81
	VK_F19                 VKEY = 0x82
	VK_F20                 VKEY = 0x83
	VK_F21                 VKEY = 0x84
	VK_F22                 VKEY = 0x85
	VK_F23                 VKEY = 0x86
	VK_F24                 VKEY = 0x87
	VK_NUMLOCK             VKEY = 0x90
	VK_SCROLL              VKEY = 0x91
	VK_OEM_NEC_EQUAL       VKEY = 0x92
	VK_OEM_FJ_JISHO        VKEY = 0x92
	VK_OEM_FJ_MASSHOU      VKEY = 0x93
	VK_OEM_FJ_TOUROKU      VKEY = 0x94
	VK_OEM_FJ_LOYA         VKEY = 0x95
	VK_OEM_FJ_ROYA         VKEY = 0x96
	VK_LSHIFT              VKEY = 0xA0
	VK_RSHIFT              VKEY = 0xA1
	VK_LCONTROL            VKEY = 0xA2
	VK_RCONTROL            VKEY = 0xA3
	VK_LMENU               VKEY = 0xA4
	VK_RMENU               VKEY = 0xA5
	VK_BROWSER_BACK        VKEY = 0xA6
	VK_BROWSER_FORWARD     VKEY = 0xA7
	VK_BROWSER_REFRESH     VKEY = 0xA8
	VK_BROWSER_STOP        VKEY = 0xA9
	VK_BROWSER_SEARCH      VKEY = 0xAA
	VK_BROWSER_FAVORITES   VKEY = 0xAB
	VK_BROWSER_HOME        VKEY = 0xAC
	VK_VOLUME_MUTE         VKEY = 0xAD
	VK_VOLUME_DOWN         VKEY = 0xAE
	VK_VOLUME_UP           VKEY = 0xAF
	VK_MEDIA_NEXT_TRACK    VKEY = 0xB0
	VK_MEDIA_PREV_TRACK    VKEY = 0xB1
	VK_MEDIA_STOP          VKEY = 0xB2
	VK_MEDIA_PLAY_PAUSE    VKEY = 0xB3
	VK_LAUNCH_MAIL         VKEY = 0xB4
	VK_LAUNCH_MEDIA_SELECT VKEY = 0xB5
	VK_LAUNCH_APP1         VKEY = 0xB6
	VK_LAUNCH_APP2         VKEY = 0xB7
	VK_OEM_1               VKEY = 0xBA
	VK_OEM_PLUS            VKEY = 0xBB
	VK_OEM_COMMA           VKEY = 0xBC
	VK_OEM_MINUS           VKEY = 0xBD
	VK_OEM_PERIOD          VKEY = 0xBE
	VK_OEM_2               VKEY = 0xBF
	VK_OEM_3               VKEY = 0xC0
	VK_OEM_4               VKEY = 0xDB
	VK_OEM_5               VKEY = 0xDC
	VK_OEM_6               VKEY = 0xDD
	VK_OEM_7               VKEY = 0xDE
	VK_OEM_8               VKEY = 0xDF
	VK_OEM_AX              VKEY = 0xE1
	VK_OEM_102             VKEY = 0xE2
	VK_ICO_HELP            VKEY = 0xE3
	VK_ICO_00              VKEY = 0xE4
	VK_PROCESSKEY          VKEY = 0xE5
	VK_ICO_CLEAR           VKEY = 0xE6
	VK_OEM_RESET           VKEY = 0xE9
	VK_OEM_JUMP            VKEY = 0xEA
	VK_OEM_PA1             VKEY = 0xEB
	VK_OEM_PA2             VKEY = 0xEC
	VK_OEM_PA3             VKEY = 0xED
	VK_OEM_WSCTRL          VKEY = 0xEE
	VK_OEM_CUSEL           VKEY = 0xEF
	VK_OEM_ATTN            VKEY = 0xF0
	VK_OEM_FINISH          VKEY = 0xF1
	VK_OEM_COPY            VKEY = 0xF2
	VK_OEM_AUTO            VKEY = 0xF3
	VK_OEM_ENLW            VKEY = 0xF4
	VK_OEM_BACKTAB         VKEY = 0xF5
	VK_ATTN                VKEY = 0xF6
	VK_CRSEL               VKEY = 0xF7
	VK_EXSEL               VKEY = 0xF8
	VK_EREOF               VKEY = 0xF9
	VK_PLAY                VKEY = 0xFA
	VK_ZOOM                VKEY = 0xFB
	VK_NONAME              VKEY = 0xFC
	VK_PA1                 VKEY = 0xFD
	VK_OEM_CLEAR           VKEY = 0xFE
)

func (com *DmSoft) SendShortCut(vKey []VKEY) {
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

func (com *DmSoft) GetKeyState(vkCode VKEY) bool {
	ret, _ := com.dm.CallMethod("GetKeyState", vkCode)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) KeyDown(vkCode VKEY) bool {
	ret, _ := com.dm.CallMethod("KeyDown", vkCode)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) KeyPress(vkCode VKEY) bool {
	ret, _ := com.dm.CallMethod("KeyPress", vkCode)
	return utils.IsOK(ret.Val)
}

// KeyPressStr指定的字符串序列，依次按顺序按下其中的字符
func (com *DmSoft) KeyPressStr(keyStr string, delay int) bool {
	ret, _ := com.dm.CallMethod("KeyPressStr", keyStr, delay)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) KeyUp(vkCode VKEY) bool {
	ret, _ := com.dm.CallMethod("KeyUp", vkCode)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) WaitKey(vkCode VKEY, timeOut int) bool {
	ret, _ := com.dm.CallMethod("SetSimMode", vkCode, timeOut)
	return utils.IsOK(ret.Val)
}
