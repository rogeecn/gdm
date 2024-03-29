// 文字识别

package gdm

import (
	"fmt"
	ole "github.com/go-ole/go-ole"
	"github.com/rogeecn/draw"
	"github.com/rogeecn/gdm/color"
	"github.com/rogeecn/gdm/utils"
	"strings"
)

// AddDict 给指定的字库中添加一条字库信息
func (com *DmSoft) AddDict(index int, dictInfo string) bool {
	ret, _ := com.dm.CallMethod("AddDict", index, dictInfo)
	return utils.IsOK(ret.Val)
}

// ClearDict 清空指定的字库
func (com *DmSoft) ClearDict(index int) bool {
	ret, _ := com.dm.CallMethod("ClearDict", index)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) EnableShareDict(enable int) bool {
	ret, _ := com.dm.CallMethod("EnableShareDict", enable)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) FetchWord(r *draw.Rect, color *color.Colors, word string) string {
	ret, _ := com.dm.CallMethod("FetchWord", r.Left(), r.Top(), r.Right(), r.Bottom(), color.String(), word)
	return ret.ToString()
}

func (com *DmSoft) FindStr(r *draw.Rect, str []string, colors *color.Colors, sim float32) (*draw.Point, bool) {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindStr", r.Left(), r.Top(), r.Right(), r.Bottom(), strings.Join(str, "|"), colors.String(), sim, &x, &y)

	ptX := int(x.Val)
	ptY := int(y.Val)

	x.Clear()
	y.Clear()

	return draw.NewPoint(ptX, ptY), utils.IsFindStrOK(ret.Val)
}

func (com *DmSoft) FindStrE(r *draw.Rect, str []string, colors *color.Colors, sim float32) *utils.FindItemResult {
	ret, _ := com.dm.CallMethod("FindStrE", r.Left(), r.Top(), r.Right(), r.Bottom(), strings.Join(str, "|"), colors.String(), sim)
	// log.Debugf("FindStrE result: %s", ret.ToString())
	return utils.NewFindItemResult(str, ret.ToString())
}

func (com *DmSoft) FindStrEx(r *draw.Rect, str []string, colors *color.Colors, sim float32) *utils.FindItemsResult {
	ret, _ := com.dm.CallMethod("FindStrEx", r.Left(), r.Top(), r.Right(), r.Bottom(), strings.Join(str, "|"), colors.String(), sim)
	return utils.NewFindItemsResult(str, ret.ToString())
}

func (com *DmSoft) FindStrFast(r *draw.Rect, str []string, colors *color.Colors, sim float32, intX *int, intY *int) *utils.FindItemResult {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindStrFast", r.Left(), r.Top(), r.Right(), r.Bottom(), strings.Join(str, "|"), colors.String(), sim, &x, &y)

	ptX := int(x.Val)
	ptY := int(y.Val)

	x.Clear()
	y.Clear()

	return utils.NewFindItemResult(str, fmt.Sprintf("%s|%d|%d", ret.ToString(), ptX, ptY))
}

func (com *DmSoft) FindStrFastE(r *draw.Rect, str []string, colors *color.Colors, sim float32) *utils.FindItemResult {
	ret, _ := com.dm.CallMethod("FindStrFastE", r.Left(), r.Top(), r.Right(), r.Bottom(), strings.Join(str, "|"), colors.String(), sim)
	return utils.NewFindItemResult(str, ret.ToString())
}

func (com *DmSoft) FindStrFastEx(r *draw.Rect, str []string, colors *color.Colors, sim float32) *utils.FindItemsResult {
	ret, _ := com.dm.CallMethod("FindStrFastEx", r.Left(), r.Top(), r.Right(), r.Bottom(), strings.Join(str, "|"), colors.String(), sim)
	return utils.NewFindItemsResult(str, ret.ToString())
}

func (com *DmSoft) FindStrWithFont(r *draw.Rect, str []string, colors *color.Colors, sim float32, fontName string, fontSize, flag int) *utils.FindItemResult {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindStrWithFont", r.Left(), r.Top(), r.Right(), r.Bottom(), strings.Join(str, "|"), colors.String(), sim, fontName, fontSize, flag, &x, &y)

	ptX := int(x.Val)
	ptY := int(y.Val)

	x.Clear()
	y.Clear()

	return utils.NewFindItemResult(str, fmt.Sprintf("%s|%d|%d", ret.ToString(), ptX, ptY))
}

func (com *DmSoft) FindStrWithFontE(r *draw.Rect, str []string, colors *color.Colors, sim float32, fontName string, fontSize, flag int) *utils.FindItemResult {
	ret, _ := com.dm.CallMethod("FindStrWithFontE", r.Left(), r.Top(), r.Right(), r.Bottom(), strings.Join(str, "|"), colors.String(), sim, fontName, fontSize, flag)
	return utils.NewFindItemResult(str, ret.ToString())
}

func (com *DmSoft) FindStrWithFontEx(r *draw.Rect, str []string, colors *color.Colors, sim float32, fontName string, fontSize, flag int) *utils.FindItemsResult {
	ret, _ := com.dm.CallMethod("FindStrWithFontEx", r.Left(), r.Top(), r.Right(), r.Bottom(), strings.Join(str, "|"), colors.String(), sim, fontName, fontSize, flag)
	return utils.NewFindItemsResult(str, ret.ToString())
}

func (com *DmSoft) GetDict(index, fontIndex int) string {
	ret, _ := com.dm.CallMethod("GetDict", index, fontIndex)
	return ret.ToString()
}

func (com *DmSoft) GetDictCount(index int) int {
	ret, _ := com.dm.CallMethod("GetDictCount", index)
	return int(ret.Val)
}

func (com *DmSoft) GetDictInfo(str, fontName string, fontSize, flag int) string {
	ret, _ := com.dm.CallMethod("GetDictInfo", str, fontName, fontSize, flag)
	return ret.ToString()
}

func (com *DmSoft) GetNowDict() int {
	ret, _ := com.dm.CallMethod("GetNowDict")
	return int(ret.Val)
}

func (com *DmSoft) GetWordResultCount(str string) int {
	ret, _ := com.dm.CallMethod("GetWordResultCount", str)
	return int(ret.Val)
}

func (com *DmSoft) Ocr(r *draw.Rect, colors *color.Colors, sim float32) string {
	ret, _ := com.dm.CallMethod("Ocr", r.Left(), r.Top(), r.Right(), r.Bottom(), colors.String(), sim)
	return ret.ToString()
}

func (com *DmSoft) OcrEx(r *draw.Rect, colors *color.Colors, sim float32) *utils.OcrResuts {
	ret, _ := com.dm.CallMethod("OcrEx", r.Left(), r.Top(), r.Right(), r.Bottom(), colors.String(), sim)
	return utils.NewOcrResuts(ret.ToString())
}

func (com *DmSoft) OcrInFile(r *draw.Rect, picName, colors *color.Colors, sim float32) string {
	ret, _ := com.dm.CallMethod("OcrInFile", r.Left(), r.Top(), r.Right(), r.Bottom(), picName, colors.String(), sim)
	return ret.ToString()
}

func (com *DmSoft) SaveDict(index int, file string) bool {
	ret, _ := com.dm.CallMethod("SaveDict", index, file)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetColGapNoDict(colGap int) bool {
	ret, _ := com.dm.CallMethod("SetColGapNoDict", colGap)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetDefaultDict(dict string) {
	com.SetDict(0, dict)
	com.UseDict(0)
}

func (com *DmSoft) SetDict(index int, file string) bool {
	ret, _ := com.dm.CallMethod("SetDict", index, file)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetDictMem(index, addr, size int) bool {
	ret, _ := com.dm.CallMethod("SetDictMem", index, addr, size)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetDictPwd(pwd string) bool {
	ret, _ := com.dm.CallMethod("SetDictPwd", pwd)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetExactOcr(exactOcr int) bool {
	ret, _ := com.dm.CallMethod("SetExactOcr", exactOcr)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetMinColGap(minColGap int) bool {
	ret, _ := com.dm.CallMethod("SetMinColGap", minColGap)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetMinRowGap(minRowGap int) bool {
	ret, _ := com.dm.CallMethod("SetMinRowGap", minRowGap)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetRowGapNoDict(RowGap int) bool {
	ret, _ := com.dm.CallMethod("SetRowGapNoDict", RowGap)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetWordGap(wordGap int) bool {
	ret, _ := com.dm.CallMethod("SetWordGap", wordGap)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetWordGapNoDict(wordGap int) bool {
	ret, _ := com.dm.CallMethod("SetWordGapNoDict", wordGap)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetWordLineHeight(lineHeight int) bool {
	ret, _ := com.dm.CallMethod("SetWordLineHeight", lineHeight)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) SetWordLineHeightNoDict(lineHeight int) bool {
	ret, _ := com.dm.CallMethod("SetWordLineHeightNoDict", lineHeight)
	return utils.IsOK(ret.Val)
}

func (com *DmSoft) UseDict(index int) bool {
	ret, _ := com.dm.CallMethod("UseDict", index)
	return utils.IsOK(ret.Val)
}
