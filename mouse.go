package gdm

import (
	"errors"
	ole "github.com/go-ole/go-ole"
	"github.com/rogeecn/draw"
	"github.com/rogeecn/gdm/utils"
)

const (
	SIMMODE_NORMAL       = 0
	SIMMODE_HARDWARE     = 1
	SIMMODE_HARDWARE_PS2 = 2
	SIMMODE_HARDWARE3    = 3
)

func (com *DmSoft) EnableMouseAccuracy(enable int) int {
	ret, _ := com.dm.CallMethod("EnableMouseAccuracy", enable)
	return int(ret.Val)
}

func (com *DmSoft) GetCursorPos() (*draw.Point, bool) {
	intx := ole.NewVariant(ole.VT_I4, 0)
	inty := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("GetCursorPos", &intx, &inty)
	ptX, ptY := int(intx.Val), int(inty.Val)

	intx.Clear()
	inty.Clear()

	return draw.NewPoint(ptX, ptY), utils.IsOK(ret.Val)
}

func (com *DmSoft) GetCursorShape() string {
	ret, _ := com.dm.CallMethod("GetCursorShape")
	return ret.ToString()
}

func (com *DmSoft) GetCursorShapeEx(types int) string {
	ret, _ := com.dm.CallMethod("GetCursorShapeEx", types)
	return ret.ToString()
}

func (com *DmSoft) GetCursorSpot() (*draw.Point, bool) {
	ret, _ := com.dm.CallMethod("GetCursorSpot")

	if len(ret.ToString()) == 0 {
		return EmptyPoint, false
	}

	pt, err := draw.NewPointFromString(ret.ToString())
	if err != nil {
		return EmptyPoint, false
	}
	return pt, true
}

func (com *DmSoft) GetMouseSpeed() int {
	ret, _ := com.dm.CallMethod("GetMouseSpeed")
	return int(ret.Val)
}

// LeftClick 按下鼠标左键
func (com *DmSoft) LeftClick() bool {
	ret, _ := com.dm.CallMethod("LeftClick")
	return utils.IsOK(ret.Val)
}

// LeftDoubleClick 双击鼠标左键
func (com *DmSoft) LeftDoubleClick() bool {
	ret, _ := com.dm.CallMethod("LeftDoubleClick")
	return utils.IsOK(ret.Val)
}

// LeftDown 按住鼠标左键
func (com *DmSoft) LeftDown() bool {
	ret, _ := com.dm.CallMethod("LeftDown")
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) LeftUp() bool {
	ret, _ := com.dm.CallMethod("LeftUp")
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) MiddleClick() bool {
	ret, _ := com.dm.CallMethod("MiddleClick")
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) MiddleDown() bool {
	ret, _ := com.dm.CallMethod("MiddleDown")
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) MiddleUp() bool {
	ret, _ := com.dm.CallMethod("MiddleUp")
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) MoveR(rx, ry int) bool {
	ret, _ := com.dm.CallMethod("MoveR", rx, ry)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) MoveTo(pt *draw.Point) bool {
	ret, _ := com.dm.CallMethod("MoveTo", pt.X, pt.Y)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) MoveToEx(pt *draw.Point, s *draw.Size) *draw.Point {
	ret, _ := com.dm.CallMethod("MoveToEx", pt.X, pt.Y, s.Width, s.Height)
	pt, err := draw.NewPointFromString(ret.ToString())
	if err != nil {
		return EmptyPoint
	}
	return pt
}

func (com *DmSoft) RightClick() bool {
	ret, _ := com.dm.CallMethod("RightClick")
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) RightDown() bool {
	ret, _ := com.dm.CallMethod("RightDown")
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) RightUp() bool {
	ret, _ := com.dm.CallMethod("RightUp")
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetMouseDelay(types string, delay int) bool {
	ret, _ := com.dm.CallMethod("SetMouseDelay", types, delay)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetMouseSpeed(speed int) int {
	ret, _ := com.dm.CallMethod("SetMouseSpeed", speed)
	return int(ret.Val)
}

func (com *DmSoft) SetSimMode(mode int) error {
	ret, _ := com.dm.CallMethod("SetSimMode", mode)

	simErrors := map[int64]error{
		0:   errors.New("插件没注册"),
		1:   nil,
		-1:  errors.New("32位系统不支持"),
		-2:  errors.New("驱动释放失败."),
		-3:  errors.New("驱动加载失败.可能是权限不够. 参考UAC权限设置. 或者是被安全软件拦截.如果是WIN10 1607之后的系统，出现这个错误，可参考这里"),
		-10: errors.New("设置失败"),
		-7:  errors.New("系统版本不支持. 可以用winver命令查看系统内部版本号. 驱动只支持正式发布的版本，所有预览版本都不支持."),
	}
	return simErrors[ret.Val]
}

func (com *DmSoft) WheelDown() bool {
	ret, _ := com.dm.CallMethod("WheelDown")
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) WheelUp() bool {
	ret, _ := com.dm.CallMethod("WheelUp")
	return utils.IsOK(ret.Val)
}
