package gdm

import "strings"

const (
	DX_GRAPHIC_2D          = "dx.graphic.2d"
	DX_GRAPHIC_2D_2        = "dx.graphic.2d.2"
	DX_GRAPHIC_3D          = "dx.graphic.3d"
	DX_GRAPHIC_3D_8        = "dx.graphic.3d.8"
	DX_GRAPHIC_OPENGL      = "dx.graphic.opengl"
	DX_GRAPHIC_OPENGL_ESV2 = "dx.graphic.opengl.esv2"
	DX_GRAPHIC_3D_10PLUS   = "dx.graphic.3d.10plus"
	// mouse
	DX_MOUSE_POSITION_LOCK_API     = "dx.mouse.position.lock.api"
	DX_MOUSE_POSITION_LOCK_MESSAGE = "dx.mouse.position.lock.message"
	DX_MOUSE_FOCUS_INPUT_API       = "dx.mouse.focus.input.api"
	DX_MOUSE_FOCUS_INPUT_MESSAGE   = "dx.mouse.focus.input.message"
	DX_MOUSE_CLIP_LOCK_API         = "dx.mouse.clip.lock.api"
	DX_MOUSE_INPUT_LOCK_API        = "dx.mouse.input.lock.api"
	DX_MOUSE_STATE_API             = "dx.mouse.state.api"
	DX_MOUSE_STATE_MESSAGE         = "dx.mouse.state.message"
	DX_MOUSE_API                   = "dx.mouse.api"
	DX_MOUSE_CURSOR                = "dx.mouse.cursor"
	DX_MOUSE_RAW_INPUT             = "dx.mouse.raw.input"
	DX_MOUSE_INPUT_LOCK_API2       = "dx.mouse.input.lock.api2"
	DX_MOUSE_INPUT_LOCK_API3       = "dx.mouse.input.lock.api3"
	// keypad
	DX_KEYPAD_INPUT_LOCK_API = "dx.keypad.input.lock.api"
	DX_KEYPAD_STATE_API      = "dx.keypad.state.api"
	DX_KEYPAD_API            = "dx.keypad.api"
	DX_KEYPAD_RAW_INPUT      = "dx.keypad.raw.input"
	// public
	DX_PUBLIC_HIDE_DLL            = "dx.public.hide.dll"
	DX_PUBLIC_ACTIVE_API2         = "dx.public.active.api2"
	DX_PUBLIC_INPUT_IME           = "dx.public.input.ime"
	DX_PUBLIC_GRAPHIC_PROTECT     = "dx.public.graphic.protect"
	DX_PUBLIC_DISABLE_WINDOW_SHOW = "dx.public.disable.window.show"
	DX_PUBLIC_ANTI_API            = "dx.public.anti.api"
	DX_PUBLIC_KM_PROTECT          = "dx.public.km.protect"
	DX_PUBLIC_PREVENT_BLOCK       = "dx.public.prevent.block"
	DX_PUBLIC_ORI_PROC            = "dx.public.ori.proc"
	DX_PUBLIC_DOWN_CPU            = "dx.public.down.cpu"
	DX_PUBLIC_FOCUS_MESSAGE       = "dx.public.focus.message"
	DX_PUBLIC_GRAPHIC_SPEED       = "dx.public.graphic.speed"
	DX_PUBLIC_MEMORY              = "dx.public.memory"
	DX_PUBLIC_INJECT_SUPER        = "dx.public.inject.super"
	DX_PUBLIC_HACK_SPEED          = "dx.public.hack.speed"
)

type DXMode struct {
	items []string
}

func NewDXMode() *DXMode {
	return &DXMode{}
}

func (d *DXMode) With(item string) *DXMode {
	d.items = append(d.items, item)
	return d
}

func (d *DXMode) String() string {
	return strings.Join(d.items, "|")
}
