package gdm

import "log"

// 内存操作暂时不需要
const (
	INT32 = 0
	INT16 = 1
	INT8  = 2
)

const (
	STR_GBK     = 0
	STR_UNICODE = 1
)

func (com *DmSoft) ReadInt(hwnd int, addr string, intType int) int64 {
	log.Println(hwnd, addr, intType)
	ret, _ := com.dm.CallMethod("ReadInt", hwnd, addr, intType)
	return ret.Val
}

func (com *DmSoft) ReadString(hwnd int, addr string, strType int) string {
	ret, _ := com.dm.CallMethod("ReadString", hwnd, addr, strType)
	return string(ret.Val)
}

func (com *DmSoft) ReadDouble(hwnd int, addr string) float64 {
	ret, _ := com.dm.CallMethod("ReadDouble", hwnd, addr)
	return float64(ret.Val)
}

func (com *DmSoft) ReadFloat(hwnd int, addr string) float32 {
	ret, _ := com.dm.CallMethod("ReadFloat", hwnd, addr)
	return float32(ret.Val)
}

// string DoubleToData(value)
// string FindData(hwnd, addr_range, data)
// string FindDataEx(hwnd, addr_range, data,step,multi_thread,mode)
// string FindDouble(hwnd, addr_range, double_value_min, double_value_max)
// string FindDoubleEx(hwnd, addr_range, double_value_min, double_value_max,step,multi_thread,mode)
// string FindFloat(hwnd, addr_range, float_value_min, float_value_max)
// string FindFloatEx(hwnd, addr_range, float_value_min, float_value_max,step,multi_thread,mode)
// string FindInt(hwnd, addr_range, int_value_min, int_value_max,type)
// string FindIntEx(hwnd, addr_range, int_value_min, int_value_max,type,step,multi_thread,mode)
// string FindString(hwnd, addr_range, string_value,type)
// string FindStringEx(hwnd, addr_range, string_value,type,step,multi_thread,mode)
// string FloatToData(value)
// long FreeProcessMemory(hwnd)
// string GetCommandLine(hwnd)
// LONGLONG GetModuleBaseAddr(hwnd,module)
// long GetModuleSize(hwnd,module)
// LONGLONG GetRemoteApiAddress(hwnd,base_addr,fun_name)
// long Int64ToInt32(value)
// string IntToData(value,type)
// long OpenProcess(pid)
// string ReadData(hwnd,addr,len)
// string ReadDataAddr(hwnd,addr,len)
// long ReadDataAddrToBin(hwnd,addr,len)
// long ReadDataToBin(hwnd,addr,len)
// double ReadDouble(hwnd,addr)
// double ReadDoubleAddr(hwnd,addr)
// float ReadFloat(hwnd,addr)
// float ReadFloatAddr(hwnd,addr)
// LONGLONG ReadInt(hwnd,addr,type)
// LONGLONG ReadIntAddr(hwnd,addr,type)
// string ReadString(hwnd,addr,type,len)
// string ReadStringAddr(hwnd,addr,type,len)
// long SetMemoryFindResultToFile(file)
// long SetMemoryHwndAsProcessId(en)
// string StringToData(value,type)
// long TerminateProcess(pid)
// LONGLONG VirtualAllocEx(hwnd,addr,size,type)
// long VirtualFreeEx(hwnd,addr)
// long VirtualProtectEx(hwnd,addr,size,type,old_protect)
// string VirtualQueryEx(hwnd,addr,pmbi)
// long WriteData(hwnd,addr,data)
// long WriteDataAddr(hwnd,addr,data)
// long WriteDataAddrFromBin(hwnd,addr,data,len)
// long WriteDataFromBin(hwnd,addr,data,len)
// long WriteDouble(hwnd,addr,v)
// long WriteDoubleAddr(hwnd,addr,v)
// long WriteFloat(hwnd,addr,v)
// long WriteFloatAddr(hwnd,addr,v)
// long WriteInt(hwnd,addr,type,v)
// long WriteIntAddr(hwnd,addr,type,v)
// long WriteString(hwnd,addr,type,v)
// long WriteStringAddr(hwnd,addr,type,v)
