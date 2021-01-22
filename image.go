// 图色

package gdm

import (
	"fmt"
	ole "github.com/go-ole/go-ole"
	"github.com/rogeecn/gdm/utils"
	"strings"
)

func (com *DmSoft) AppendPicAddr(picInfo string, addr, size int) string {
	ret, _ := com.dm.CallMethod("AppendPicAddr", picInfo, addr, size)
	return ret.ToString()
}

func (com *DmSoft) BGR2RGB(color string) string {
	ret, _ := com.dm.CallMethod("AppendPicAddr", color)
	return ret.ToString()
}

func (com *DmSoft) Capture(r utils.Rect, file string) int {
	ret, _ := com.dm.CallMethod("Capture", r.Left, r.Top, r.Right, r.Bottom, file)
	return int(ret.Val)
}

func (com *DmSoft) CaptureGif(r utils.Rect, file string, delay, time int) int {
	ret, _ := com.dm.CallMethod("CaptureGif", r.Left, r.Top, r.Right, r.Bottom, file, delay, time)
	return int(ret.Val)
}

func (com *DmSoft) CaptureJpg(r utils.Rect, file string, quality int) int {
	ret, _ := com.dm.CallMethod("CaptureJpg", r.Left, r.Top, r.Right, r.Bottom, file, quality)
	return int(ret.Val)
}

func (com *DmSoft) CapturePng(r utils.Rect, file string) int {
	ret, _ := com.dm.CallMethod("CapturePng", r.Left, r.Top, r.Right, r.Bottom, file)
	return int(ret.Val)
}

func (com *DmSoft) CapturePre(file string) int {
	ret, _ := com.dm.CallMethod("CapturePre", file)
	return int(ret.Val)
}

func (com *DmSoft) CmpColor(pt utils.Point, colors *utils.Colors, sim float32) bool {
	ret, _ := com.dm.CallMethod("CmpColor", pt.X, pt.Y, colors.String(), sim)
	return utils.IsColorOK(ret.Val)
}

func (com *DmSoft) EnableDisplayDebug(enableDebug int) bool {
	ret, _ := com.dm.CallMethod("EnableDisplayDebug", enableDebug)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) EnableFindPicMultithread(enable int) bool {
	ret, _ := com.dm.CallMethod("EnableFindPicMultithread", enable)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) EnableGetColorByCapture(enable int) bool {
	ret, _ := com.dm.CallMethod("EnableGetColorByCapture", enable)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) FindColor(r utils.Rect, colors *utils.Colors, sim float32, dir int) (utils.Point, bool) {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindColor", r.Left, r.Top, r.Right, r.Bottom, colors.String(), sim, dir, &x, &y)

	ptX := int(x.Val)
	ptY := int(y.Val)

	x.Clear()
	y.Clear()

	return utils.Point{ptX, ptY}, utils.IsOK(ret.Val)
}

func (com *DmSoft) FindColorBlock(r utils.Rect, colors *utils.Colors, sim float32, count, width, height int, intX, intY *int) (utils.Point, bool) {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindColorBlock", r.Left, r.Top, r.Right, r.Bottom, colors.String(), sim, count, width, height, &x, &y)

	ptX := int(x.Val)
	ptY := int(y.Val)

	x.Clear()
	y.Clear()

	return utils.Point{ptX, ptY}, utils.IsOK(ret.Val)
}

func (com *DmSoft) FindColorBlockEx(r utils.Rect, colors *utils.Colors, sim float32, count int, s utils.Size) string {
	ret, _ := com.dm.CallMethod("FindColorBlockEx", r.Left, r.Top, r.Right, r.Bottom, colors.String(), sim, count, s.Width, s.Height)
	return ret.ToString()
}

func (com *DmSoft) FindColorE(r utils.Rect, colors *utils.Colors, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindColorE", r.Left, r.Top, r.Right, r.Bottom, colors.String(), sim, dir)
	return ret.ToString()
}

func (com *DmSoft) FindColorEx(r utils.Rect, colors *utils.Colors, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindColorEx", r.Left, r.Top, r.Right, r.Bottom, colors.String(), sim, dir)
	return ret.ToString()
}

func (com *DmSoft) FindMulColor(r utils.Rect, colors *utils.Colors, sim float32) bool {
	ret, _ := com.dm.CallMethod("FindMulColor", r.Left, r.Top, r.Right, r.Bottom, colors.String(), sim)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) FindMultiColor(r utils.Rect, firstColor string, offsetColor string, sim float32, dir int) (utils.Point, bool) {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindMultiColor", r.Left, r.Top, r.Right, r.Bottom, firstColor, offsetColor, sim, dir, &x, &y)

	ptX := int(x.Val)
	ptY := int(y.Val)

	x.Clear()
	y.Clear()

	return utils.Point{ptX, ptY}, utils.IsOK(ret.Val)
}

func (com *DmSoft) FindMultiColorEx(r utils.Rect, firstColor string, offsetColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindMultiColorEx", r.Left, r.Top, r.Right, r.Bottom, firstColor, offsetColor, sim, dir)
	return ret.ToString()
}

func (com *DmSoft) FindPic(r utils.Rect, pics []string, deltaColor string, sim float32, dir int) *utils.FindWordResult {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindPic", r.Left, r.Top, r.Right, r.Bottom, strings.Join(pics, "|"), deltaColor, sim, dir, &x, &y)

	ptX := int(x.Val)
	ptY := int(y.Val)

	x.Clear()
	y.Clear()

	return utils.NewFindWordResult(pics, fmt.Sprintf("%d|%d|%d", ret.Val, ptX, ptY))
}

func (com *DmSoft) FindPicE(r utils.Rect, pics []string, deltaColor string, sim float32, dir int) *utils.FindWordResult {
	ret, _ := com.dm.CallMethod("FindPicE", r.Left, r.Top, r.Right, r.Bottom, strings.Join(pics, "|"), deltaColor, sim, dir)
	return utils.NewFindWordResult(pics, ret.ToString())
}

func (com *DmSoft) FindPicEx(r utils.Rect, pics []string, deltaColor string, sim float32, dir int) *utils.FindWordsResult {
	ret, _ := com.dm.CallMethod("FindPicEx", r.Left, r.Top, r.Right, r.Bottom, strings.Join(pics, "|"), deltaColor, sim, dir)
	return utils.NewFindWordsResult(pics, ret.ToString())
}

// func (com *DmSoft)FindPicMem(r.Left, r.Top, r.Right, r.Bottom, pic_info, delta_color,sim, dir,intX, intY)  int{}
// func (com *DmSoft)FindPicMemE(r.Left, r.Top, r.Right, r.Bottom, pic_info, delta_color,sim, dir,intX, intY)  string{}
// func (com *DmSoft)FindPicMemEx(r.Left, r.Top, r.Right, r.Bottom, pic_info, delta_color,sim, dir,intX, intY)  string{}

// func (com *DmSoft)FindShape(r.Left, r.Top, r.Right, r.Bottom, offset_color,sim, dir,intX,intY) int{}
// func (com *DmSoft)FindShapeE(r.Left, r.Top, r.Right, r.Bottom, offset_color,sim, dir,intX,intY) string{}
// func (com *DmSoft)FindShapeEx(r.Left, r.Top, r.Right, r.Bottom, offset_color,sim, dir,intX,intY) string{}
// func (com *DmSoft)FreePic(pic_name) int{}

//
func (com *DmSoft) GetAveHSV(r utils.Rect) string {
	ret, _ := com.dm.CallMethod("GetAveHSV", r.Left, r.Top, r.Right, r.Bottom)
	return ret.ToString()
}

func (com *DmSoft) GetAveRGB(r utils.Rect) string {
	ret, _ := com.dm.CallMethod("GetAveRGB", r.Left, r.Top, r.Right, r.Bottom)
	return ret.ToString()
}

func (com *DmSoft) GetColor(x, y int) string {
	ret, _ := com.dm.CallMethod("GetColor", x, y)
	return ret.ToString()
}

// func (com *DmSoft)GetColorBGR(x,y)string{}

func (com *DmSoft) GetColorHSV(x, y int) string {
	ret, _ := com.dm.CallMethod("GetColorHSV", x, y)
	return ret.ToString()
}

// func (com *DmSoft)GetColorNum(r.Left, r.Top, r.Right, r.Bottom, color, sim) int{}

func (com *DmSoft) GetPicSize(picName string) string {
	ret, _ := com.dm.CallMethod("GetPicSize", picName)
	return ret.ToString()
}

// func (com *DmSoft)GetScreenData(x1, y1, x2, y2) int{}
// func (com *DmSoft)GetScreenDataBmp(x1,y1,x2,y2,data,size) int{}
// func (com *DmSoft)ImageToBmp(pic_name,bmp_name) int{}

func (com *DmSoft) IsDisplayDead(r utils.Rect, time int) bool {
	ret, _ := com.dm.CallMethod("IsDisplayDead", r.Left, r.Top, r.Right, r.Bottom, time)
	return utils.IsOK(ret.Val)
}

// func (com *DmSoft)LoadPic(pic_name) int{}
// func (com *DmSoft)LoadPicByte(addr,size,pic_name) int{}
// func (com *DmSoft)LoadPic(pic_name) int{}

func (com *DmSoft) MatchPicName(picName string) string {
	ret, _ := com.dm.CallMethod("MatchPicName", picName)
	return ret.ToString()
}

func (com *DmSoft) SetExcludeRegion(mode int, info string) int {
	ret, _ := com.dm.CallMethod("SetExcludeRegion", mode, info)
	return int(ret.Val)
}

func (com *DmSoft) SetPicPwd(pwd string) int {
	ret, _ := com.dm.CallMethod("SetExcludeRegion", pwd)
	return int(ret.Val)
}
