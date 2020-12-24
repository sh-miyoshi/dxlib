package main

import (
	"github.com/sh-miyoshi/dxlib"
)

func main() {
	dxlib.Init("DxLib.dll")

	dxlib.ChangeWindowMode(1)
	dxlib.SetGraphMode(640, 480, 16)
	dxlib.SetOutApplicationLogValidFlag(1)

	dxlib.DxLib_Init()
	dxlib.SetDrawScreen(0xfffffffe)

	for dxlib.ScreenFlip() == 0 && dxlib.ProcessMessage() == 0 && dxlib.ClearDrawScreen() == 0 {
		dxlib.DrawString(10, 10, "Hello, world", dxlib.GetColor(255, 255, 255))
	}

	dxlib.DxLib_End()
}
