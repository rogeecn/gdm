package gdm

import (
	"github.com/rogeecn/draw"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	reg := os.Getenv("DM_REG")
	ver := os.Getenv("DM_VER")
	t.Log("REG: ", reg)
	t.Log("REG_Ver: ", ver)

	pwd, _ := os.Getwd()
	dm, err := New(pwd)
	if err != nil {
		t.Fatal(err)
	}
	defer dm.Release()

	t.Log("Ver: ", dm.Ver())
	err = dm.RegNoMac(reg, ver)
	if err != nil {
		t.Fatal("RegNoMac: ", err)
	}
	t.Log("MachineCodeNoMac: ", dm.GetMachineCodeNoMac())
	t.Log("MachineCode: ", dm.GetMachineCode())
	dm.MoveTo(draw.NewPoint(10, 10))
}
