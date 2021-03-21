package main

import (
	"runtime"

	"github.com/sh-miyoshi/dxlib"
)

// All calls to DX Library must happen from the same thread
// that creates the DX Library object so make sure to add
// this code in your main package:
func init() {
	runtime.LockOSThread()
}

func main() {
	// Initialize dxlib functions with DxLib.dll
	dxlib.Init("DxLib.dll")

	// Run as window mode
	dxlib.ChangeWindowMode(dxlib.TRUE)
	// Set screen size: 640 * 480 and color bit num: 16
	// optional argumentcan set by struct
	dxlib.SetGraphMode(640, 480, dxlib.SetGraphModeOption{ColorBitNum: dxlib.IntPtr(16)})
	// Output DxLib log
	dxlib.SetOutApplicationLogValidFlag(dxlib.TRUE)

	dxlib.DxLib_Init()
	dxlib.SetDrawScreen(dxlib.DX_SCREEN_BACK)

	for dxlib.ScreenFlip() == 0 && dxlib.ProcessMessage() == 0 && dxlib.ClearDrawScreen() == 0 {
		dxlib.DrawString(10, 10, "Hello, world", dxlib.GetColor(255, 255, 255))
	}

	dxlib.DxLib_End()
}
