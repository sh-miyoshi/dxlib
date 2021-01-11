package main

import (
	"runtime"

	"github.com/sh-miyoshi/dxlib"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	dxlib.Init("DxLib.dll")

	dxlib.ChangeWindowMode(dxlib.TRUE)
	dxlib.SetGraphMode(640, 480, 16, 60)
	dxlib.SetOutApplicationLogValidFlag(dxlib.TRUE)

	dxlib.DxLib_Init()
	dxlib.SetDrawScreen(dxlib.DX_SCREEN_BACK)

	for dxlib.ScreenFlip() == 0 && dxlib.ProcessMessage() == 0 && dxlib.ClearDrawScreen() == 0 {
		dxlib.DrawString(10, 10, "Hello, world", dxlib.GetColor(255, 255, 255), 0)
	}

	dxlib.DxLib_End()
}
