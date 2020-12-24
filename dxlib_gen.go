// Code generated by 'go generate'; DO NOT EDIT.

package dxlib

import (
	"syscall"
	"unsafe"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var (
	dx_DxLib_Init                           *syscall.LazyProc
	dx_DxLib_End                            *syscall.LazyProc
	dx_ProcessMessage                       *syscall.LazyProc
	dx_DrawLine                             *syscall.LazyProc
	dx_DrawLineAA                           *syscall.LazyProc
	dx_DrawBox                              *syscall.LazyProc
	dx_DrawBoxAA                            *syscall.LazyProc
	dx_DrawCircle                           *syscall.LazyProc
	dx_DrawCircleAA                         *syscall.LazyProc
	dx_DrawOval                             *syscall.LazyProc
	dx_DrawOvalAA                           *syscall.LazyProc
	dx_DrawTriangle                         *syscall.LazyProc
	dx_DrawTriangleAA                       *syscall.LazyProc
	dx_DrawPixel                            *syscall.LazyProc
	dx_GetPixel                             *syscall.LazyProc
	dx_LoadGraphScreen                      *syscall.LazyProc
	dx_LoadGraph                            *syscall.LazyProc
	dx_LoadDivGraph                         *syscall.LazyProc
	dx_MakeGraph                            *syscall.LazyProc
	dx_MakeScreen                           *syscall.LazyProc
	dx_SetCreateDrawValidGraphMultiSample   *syscall.LazyProc
	dx_SetCreateGraphColorBitDepth          *syscall.LazyProc
	dx_SetDrawValidFloatTypeGraphCreateFlag *syscall.LazyProc
	dx_SetCreateDrawValidGraphChannelNum    *syscall.LazyProc
	dx_SetUsePremulAlphaConvertLoad         *syscall.LazyProc
	dx_DrawGraph                            *syscall.LazyProc
	dx_DrawTurnGraph                        *syscall.LazyProc
	dx_DrawExtendGraph                      *syscall.LazyProc
	dx_DrawRotaGraph                        *syscall.LazyProc
	dx_DrawRotaGraph2                       *syscall.LazyProc
	dx_DrawRotaGraph3                       *syscall.LazyProc
	dx_DrawModiGraph                        *syscall.LazyProc
	dx_DrawRectGraph                        *syscall.LazyProc
	dx_DerivationGraph                      *syscall.LazyProc
	dx_GetDrawScreenGraph                   *syscall.LazyProc
	dx_GetGraphiteSize                      *syscall.LazyProc
	dx_InitGraph                            *syscall.LazyProc
	dx_DeleteGraph                          *syscall.LazyProc
	dx_SetDrawMode                          *syscall.LazyProc
	dx_SetDrawBlendMode                     *syscall.LazyProc
	dx_SetDrawBright                        *syscall.LazyProc
	dx_SetTransColor                        *syscall.LazyProc
	dx_LoadBlendGraph                       *syscall.LazyProc
	dx_DrawBlendGraph                       *syscall.LazyProc
	dx_DrawString                           *syscall.LazyProc
	dx_GetDrawStringWidth                   *syscall.LazyProc
	dx_SetFontSize                          *syscall.LazyProc
	dx_SetFontThickness                     *syscall.LazyProc
	dx_ChangeFont                           *syscall.LazyProc
	dx_ChangeFontType                       *syscall.LazyProc
	dx_CreateFontToHandle                   *syscall.LazyProc
	dx_LoadFontDataToHandle                 *syscall.LazyProc
	dx_DeleteFontToHandle                   *syscall.LazyProc
	dx_SetFontCacheUsePremulAlphaFlag       *syscall.LazyProc
	dx_DrawStringToHandle                   *syscall.LazyProc
	dx_GetDrawStringWidthToHandle           *syscall.LazyProc
	dx_GetFontStateToHandle                 *syscall.LazyProc
	dx_InitFontToHandle                     *syscall.LazyProc
	dx_clsDx                                *syscall.LazyProc
	dx_SetGraphMode                         *syscall.LazyProc
	dx_SetFullScreenResolutionMode          *syscall.LazyProc
	dx_SetFullScreenScalingMode             *syscall.LazyProc
	dx_GetScreenState                       *syscall.LazyProc
	dx_SetDrawArea                          *syscall.LazyProc
	dx_ClearDrawScreen                      *syscall.LazyProc
	dx_SetBackgroundColor                   *syscall.LazyProc
	dx_GetColor                             *syscall.LazyProc
	dx_SetDrawScreen                        *syscall.LazyProc
	dx_ScreenFlip                           *syscall.LazyProc
	dx_SetFullSceneAntiAliasingMode         *syscall.LazyProc
	dx_PlayMovie                            *syscall.LazyProc
	dx_PlayMovieToGraph                     *syscall.LazyProc
	dx_PauseMovieToGraph                    *syscall.LazyProc
	dx_SeekMovieToGraph                     *syscall.LazyProc
	dx_TellMovieToGraph                     *syscall.LazyProc
	dx_GetMovieStateToGraph                 *syscall.LazyProc
	dx_ChangeWindowMode                     *syscall.LazyProc
	dx_SetOutApplicationLogValidFlag        *syscall.LazyProc
)

func Init(dllFile string) {
	mod := syscall.NewLazyDLL(dllFile)

	dx_DxLib_Init = mod.NewProc("dx_DxLib_Init")
	dx_DxLib_End = mod.NewProc("dx_DxLib_End")
	dx_ProcessMessage = mod.NewProc("dx_ProcessMessage")
	dx_DrawLine = mod.NewProc("dx_DrawLine")
	dx_DrawLineAA = mod.NewProc("dx_DrawLineAA")
	dx_DrawBox = mod.NewProc("dx_DrawBox")
	dx_DrawBoxAA = mod.NewProc("dx_DrawBoxAA")
	dx_DrawCircle = mod.NewProc("dx_DrawCircle")
	dx_DrawCircleAA = mod.NewProc("dx_DrawCircleAA")
	dx_DrawOval = mod.NewProc("dx_DrawOval")
	dx_DrawOvalAA = mod.NewProc("dx_DrawOvalAA")
	dx_DrawTriangle = mod.NewProc("dx_DrawTriangle")
	dx_DrawTriangleAA = mod.NewProc("dx_DrawTriangleAA")
	dx_DrawPixel = mod.NewProc("dx_DrawPixel")
	dx_GetPixel = mod.NewProc("dx_GetPixel")
	dx_LoadGraphScreen = mod.NewProc("dx_LoadGraphScreen")
	dx_LoadGraph = mod.NewProc("dx_LoadGraph")
	dx_LoadDivGraph = mod.NewProc("dx_LoadDivGraph")
	dx_MakeGraph = mod.NewProc("dx_MakeGraph")
	dx_MakeScreen = mod.NewProc("dx_MakeScreen")
	dx_SetCreateDrawValidGraphMultiSample = mod.NewProc("dx_SetCreateDrawValidGraphMultiSample")
	dx_SetCreateGraphColorBitDepth = mod.NewProc("dx_SetCreateGraphColorBitDepth")
	dx_SetDrawValidFloatTypeGraphCreateFlag = mod.NewProc("dx_SetDrawValidFloatTypeGraphCreateFlag")
	dx_SetCreateDrawValidGraphChannelNum = mod.NewProc("dx_SetCreateDrawValidGraphChannelNum")
	dx_SetUsePremulAlphaConvertLoad = mod.NewProc("dx_SetUsePremulAlphaConvertLoad")
	dx_DrawGraph = mod.NewProc("dx_DrawGraph")
	dx_DrawTurnGraph = mod.NewProc("dx_DrawTurnGraph")
	dx_DrawExtendGraph = mod.NewProc("dx_DrawExtendGraph")
	dx_DrawRotaGraph = mod.NewProc("dx_DrawRotaGraph")
	dx_DrawRotaGraph2 = mod.NewProc("dx_DrawRotaGraph2")
	dx_DrawRotaGraph3 = mod.NewProc("dx_DrawRotaGraph3")
	dx_DrawModiGraph = mod.NewProc("dx_DrawModiGraph")
	dx_DrawRectGraph = mod.NewProc("dx_DrawRectGraph")
	dx_DerivationGraph = mod.NewProc("dx_DerivationGraph")
	dx_GetDrawScreenGraph = mod.NewProc("dx_GetDrawScreenGraph")
	dx_GetGraphiteSize = mod.NewProc("dx_GetGraphiteSize")
	dx_InitGraph = mod.NewProc("dx_InitGraph")
	dx_DeleteGraph = mod.NewProc("dx_DeleteGraph")
	dx_SetDrawMode = mod.NewProc("dx_SetDrawMode")
	dx_SetDrawBlendMode = mod.NewProc("dx_SetDrawBlendMode")
	dx_SetDrawBright = mod.NewProc("dx_SetDrawBright")
	dx_SetTransColor = mod.NewProc("dx_SetTransColor")
	dx_LoadBlendGraph = mod.NewProc("dx_LoadBlendGraph")
	dx_DrawBlendGraph = mod.NewProc("dx_DrawBlendGraph")
	dx_DrawString = mod.NewProc("dx_DrawString")
	dx_GetDrawStringWidth = mod.NewProc("dx_GetDrawStringWidth")
	dx_SetFontSize = mod.NewProc("dx_SetFontSize")
	dx_SetFontThickness = mod.NewProc("dx_SetFontThickness")
	dx_ChangeFont = mod.NewProc("dx_ChangeFont")
	dx_ChangeFontType = mod.NewProc("dx_ChangeFontType")
	dx_CreateFontToHandle = mod.NewProc("dx_CreateFontToHandle")
	dx_LoadFontDataToHandle = mod.NewProc("dx_LoadFontDataToHandle")
	dx_DeleteFontToHandle = mod.NewProc("dx_DeleteFontToHandle")
	dx_SetFontCacheUsePremulAlphaFlag = mod.NewProc("dx_SetFontCacheUsePremulAlphaFlag")
	dx_DrawStringToHandle = mod.NewProc("dx_DrawStringToHandle")
	dx_GetDrawStringWidthToHandle = mod.NewProc("dx_GetDrawStringWidthToHandle")
	dx_GetFontStateToHandle = mod.NewProc("dx_GetFontStateToHandle")
	dx_InitFontToHandle = mod.NewProc("dx_InitFontToHandle")
	dx_clsDx = mod.NewProc("dx_clsDx")
	dx_SetGraphMode = mod.NewProc("dx_SetGraphMode")
	dx_SetFullScreenResolutionMode = mod.NewProc("dx_SetFullScreenResolutionMode")
	dx_SetFullScreenScalingMode = mod.NewProc("dx_SetFullScreenScalingMode")
	dx_GetScreenState = mod.NewProc("dx_GetScreenState")
	dx_SetDrawArea = mod.NewProc("dx_SetDrawArea")
	dx_ClearDrawScreen = mod.NewProc("dx_ClearDrawScreen")
	dx_SetBackgroundColor = mod.NewProc("dx_SetBackgroundColor")
	dx_GetColor = mod.NewProc("dx_GetColor")
	dx_SetDrawScreen = mod.NewProc("dx_SetDrawScreen")
	dx_ScreenFlip = mod.NewProc("dx_ScreenFlip")
	dx_SetFullSceneAntiAliasingMode = mod.NewProc("dx_SetFullSceneAntiAliasingMode")
	dx_PlayMovie = mod.NewProc("dx_PlayMovie")
	dx_PlayMovieToGraph = mod.NewProc("dx_PlayMovieToGraph")
	dx_PauseMovieToGraph = mod.NewProc("dx_PauseMovieToGraph")
	dx_SeekMovieToGraph = mod.NewProc("dx_SeekMovieToGraph")
	dx_TellMovieToGraph = mod.NewProc("dx_TellMovieToGraph")
	dx_GetMovieStateToGraph = mod.NewProc("dx_GetMovieStateToGraph")
	dx_ChangeWindowMode = mod.NewProc("dx_ChangeWindowMode")
	dx_SetOutApplicationLogValidFlag = mod.NewProc("dx_SetOutApplicationLogValidFlag")

}

func DxLib_Init() int {
	if dx_DxLib_Init == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DxLib_Init.Call()
	return int(res)
}

func DxLib_End() int {
	if dx_DxLib_End == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DxLib_End.Call()
	return int(res)
}

func ProcessMessage() int {
	if dx_ProcessMessage == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_ProcessMessage.Call()
	return int(res)
}

func DrawLine(x1 int, y1 int, x2 int, y2 int, color uint) int {
	if dx_DrawLine == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawLine.Call(pint(x1), pint(y1), pint(x2), pint(y2), puint(color))
	return int(res)
}

func DrawLineAA(x1 float32, y1 float32, x2 float32, y2 float32, color uint) int {
	if dx_DrawLineAA == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawLineAA.Call(pfloat32(x1), pfloat32(y1), pfloat32(x2), pfloat32(y2), puint(color))
	return int(res)
}

func DrawBox(x1 int, y1 int, x2 int, y2 int, color uint, fillFlag int) int {
	if dx_DrawBox == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawBox.Call(pint(x1), pint(y1), pint(x2), pint(y2), puint(color), pint(fillFlag))
	return int(res)
}

func DrawBoxAA(x1 float32, y1 float32, x2 float32, y2 float32, color uint, fillFlag int) int {
	if dx_DrawBoxAA == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawBoxAA.Call(pfloat32(x1), pfloat32(y1), pfloat32(x2), pfloat32(y2), puint(color), pint(fillFlag))
	return int(res)
}

func DrawCircle(x int, y int, r int, color uint, fillFlag int) int {
	if dx_DrawCircle == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawCircle.Call(pint(x), pint(y), pint(r), puint(color), pint(fillFlag))
	return int(res)
}

func DrawCircleAA(x float32, y float32, r float32, posnum int, color uint, fillFlag int) int {
	if dx_DrawCircleAA == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawCircleAA.Call(pfloat32(x), pfloat32(y), pfloat32(r), pint(posnum), puint(color), pint(fillFlag))
	return int(res)
}

func DrawOval(x int, y int, rx int, ry int, color uint, fillFlag int) int {
	if dx_DrawOval == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawOval.Call(pint(x), pint(y), pint(rx), pint(ry), puint(color), pint(fillFlag))
	return int(res)
}

func DrawOvalAA(x float32, y float32, rx float32, ry float32, posnum int, color uint, fillFlag int) int {
	if dx_DrawOvalAA == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawOvalAA.Call(pfloat32(x), pfloat32(y), pfloat32(rx), pfloat32(ry), pint(posnum), puint(color), pint(fillFlag))
	return int(res)
}

func DrawTriangle(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int, color uint, fillFlag int) int {
	if dx_DrawTriangle == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawTriangle.Call(pint(x1), pint(y1), pint(x2), pint(y2), pint(x3), pint(y3), puint(color), pint(fillFlag))
	return int(res)
}

func DrawTriangleAA(x1 float32, y1 float32, x2 float32, y2 float32, x3 float32, y3 float32, color uint, fillFlag int) int {
	if dx_DrawTriangleAA == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawTriangleAA.Call(pfloat32(x1), pfloat32(y1), pfloat32(x2), pfloat32(y2), pfloat32(x3), pfloat32(y3), puint(color), pint(fillFlag))
	return int(res)
}

func DrawPixel(x int, y int, color uint) int {
	if dx_DrawPixel == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawPixel.Call(pint(x), pint(y), puint(color))
	return int(res)
}

func GetPixel(x int, y int) uint {
	if dx_GetPixel == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_GetPixel.Call(pint(x), pint(y))
	return uint(res)
}

func LoadGraphScreen(x int, y int, graphName string, transFlag int) int {
	if dx_LoadGraphScreen == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_LoadGraphScreen.Call(pint(x), pint(y), pstring(graphName), pint(transFlag))
	return int(res)
}

func LoadGraph(fileName string) int {
	if dx_LoadGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_LoadGraph.Call(pstring(fileName))
	return int(res)
}

func LoadDivGraph(fileName string, allnum int, xnum int, ynum int, xsize int, ysize int, handleBuf *int) int {
	if dx_LoadDivGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_LoadDivGraph.Call(pstring(fileName), pint(allnum), pint(xnum), pint(ynum), pint(xsize), pint(ysize), ppint(handleBuf))
	return int(res)
}

func MakeGraph(sizeX int, sizeY int) int {
	if dx_MakeGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_MakeGraph.Call(pint(sizeX), pint(sizeY))
	return int(res)
}

func MakeScreen(sizeX int, sizeY int, useAlphaChannel int) int {
	if dx_MakeScreen == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_MakeScreen.Call(pint(sizeX), pint(sizeY), pint(useAlphaChannel))
	return int(res)
}

func SetCreateDrawValidGraphMultiSample(samples int, quality int) int {
	if dx_SetCreateDrawValidGraphMultiSample == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetCreateDrawValidGraphMultiSample.Call(pint(samples), pint(quality))
	return int(res)
}

func SetCreateGraphColorBitDepth(bitDepth int) int {
	if dx_SetCreateGraphColorBitDepth == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetCreateGraphColorBitDepth.Call(pint(bitDepth))
	return int(res)
}

func SetDrawValidFloatTypeGraphCreateFlag(flag int) int {
	if dx_SetDrawValidFloatTypeGraphCreateFlag == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetDrawValidFloatTypeGraphCreateFlag.Call(pint(flag))
	return int(res)
}

func SetCreateDrawValidGraphChannelNum(channelNum int) int {
	if dx_SetCreateDrawValidGraphChannelNum == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetCreateDrawValidGraphChannelNum.Call(pint(channelNum))
	return int(res)
}

func SetUsePremulAlphaConvertLoad(useFlag int) int {
	if dx_SetUsePremulAlphaConvertLoad == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetUsePremulAlphaConvertLoad.Call(pint(useFlag))
	return int(res)
}

func DrawGraph(x int, y int, grHandle int, transFlag int) int {
	if dx_DrawGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawGraph.Call(pint(x), pint(y), pint(grHandle), pint(transFlag))
	return int(res)
}

func DrawTurnGraph(x int, y int, grHandle int, transFlag int) int {
	if dx_DrawTurnGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawTurnGraph.Call(pint(x), pint(y), pint(grHandle), pint(transFlag))
	return int(res)
}

func DrawExtendGraph(x1 int, y1 int, x2 int, y2 int, grHandle int, transFlag int) int {
	if dx_DrawExtendGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawExtendGraph.Call(pint(x1), pint(y1), pint(x2), pint(y2), pint(grHandle), pint(transFlag))
	return int(res)
}

func DrawRotaGraph(x int, y int, extRate float64, angle float64, grHandle int, transFlag int, turnFlag int) int {
	if dx_DrawRotaGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawRotaGraph.Call(pint(x), pint(y), pfloat64(extRate), pfloat64(angle), pint(grHandle), pint(transFlag), pint(turnFlag))
	return int(res)
}

func DrawRotaGraph2(x int, y int, cx int, cy int, extRate float64, angle float64, grHandle int, transFlag int, turnFlag int) int {
	if dx_DrawRotaGraph2 == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawRotaGraph2.Call(pint(x), pint(y), pint(cx), pint(cy), pfloat64(extRate), pfloat64(angle), pint(grHandle), pint(transFlag), pint(turnFlag))
	return int(res)
}

func DrawRotaGraph3(x int, y int, cx int, cy int, extRateX float64, extRateY float64, angle float64, grHandle int, transFlag int, turnFlag int) int {
	if dx_DrawRotaGraph3 == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawRotaGraph3.Call(pint(x), pint(y), pint(cx), pint(cy), pfloat64(extRateX), pfloat64(extRateY), pfloat64(angle), pint(grHandle), pint(transFlag), pint(turnFlag))
	return int(res)
}

func DrawModiGraph(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int, x4 int, y4 int, grHandle int, transFlag int) int {
	if dx_DrawModiGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawModiGraph.Call(pint(x1), pint(y1), pint(x2), pint(y2), pint(x3), pint(y3), pint(x4), pint(y4), pint(grHandle), pint(transFlag))
	return int(res)
}

func DrawRectGraph(destX int, destY int, srcX int, srcY int, width int, height int, graphHandle int, transFlag int, turnFlag int) int {
	if dx_DrawRectGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawRectGraph.Call(pint(destX), pint(destY), pint(srcX), pint(srcY), pint(width), pint(height), pint(graphHandle), pint(transFlag), pint(turnFlag))
	return int(res)
}

func DerivationGraph(srcX int, srcY int, width int, height int, srcGraphHandle int) int {
	if dx_DerivationGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DerivationGraph.Call(pint(srcX), pint(srcY), pint(width), pint(height), pint(srcGraphHandle))
	return int(res)
}

func GetDrawScreenGraph(x1 int, y1 int, x2 int, y2 int, grHandle int) int {
	if dx_GetDrawScreenGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_GetDrawScreenGraph.Call(pint(x1), pint(y1), pint(x2), pint(y2), pint(grHandle))
	return int(res)
}

func GetGraphiteSize(grHandle int, sizeXBuf *int, sizeYBuf *int) int {
	if dx_GetGraphiteSize == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_GetGraphiteSize.Call(pint(grHandle), ppint(sizeXBuf), ppint(sizeYBuf))
	return int(res)
}

func InitGraph() int {
	if dx_InitGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_InitGraph.Call()
	return int(res)
}

func DeleteGraph(grHandle int) int {
	if dx_DeleteGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DeleteGraph.Call(pint(grHandle))
	return int(res)
}

func SetDrawMode(drawMode int) int {
	if dx_SetDrawMode == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetDrawMode.Call(pint(drawMode))
	return int(res)
}

func SetDrawBlendMode(blendMode int, pal int) int {
	if dx_SetDrawBlendMode == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetDrawBlendMode.Call(pint(blendMode), pint(pal))
	return int(res)
}

func SetDrawBright(redBright int, greenBright int, blueBright int) int {
	if dx_SetDrawBright == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetDrawBright.Call(pint(redBright), pint(greenBright), pint(blueBright))
	return int(res)
}

func SetTransColor(red int, green int, blue int) int {
	if dx_SetTransColor == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetTransColor.Call(pint(red), pint(green), pint(blue))
	return int(res)
}

func LoadBlendGraph(fileName string) int {
	if dx_LoadBlendGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_LoadBlendGraph.Call(pstring(fileName))
	return int(res)
}

func DrawBlendGraph(x int, y int, grHandle int, transFlag int, blendGraph int, borderParam int, borderRange int) int {
	if dx_DrawBlendGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawBlendGraph.Call(pint(x), pint(y), pint(grHandle), pint(transFlag), pint(blendGraph), pint(borderParam), pint(borderRange))
	return int(res)
}

func DrawString(x int, y int, str string, color uint) int {
	if dx_DrawString == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawString.Call(pint(x), pint(y), pstring(str), puint(color))
	return int(res)
}

func GetDrawStringWidth(str string, strLen int) int {
	if dx_GetDrawStringWidth == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_GetDrawStringWidth.Call(pstring(str), pint(strLen))
	return int(res)
}

func SetFontSize(fontSize int) int {
	if dx_SetFontSize == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetFontSize.Call(pint(fontSize))
	return int(res)
}

func SetFontThickness(tinckPal int) int {
	if dx_SetFontThickness == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetFontThickness.Call(pint(tinckPal))
	return int(res)
}

func ChangeFont(fontName string) int {
	if dx_ChangeFont == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_ChangeFont.Call(pstring(fontName))
	return int(res)
}

func ChangeFontType(fontType int) int {
	if dx_ChangeFontType == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_ChangeFontType.Call(pint(fontType))
	return int(res)
}

func CreateFontToHandle(fontName string, size int, thick int, fontType int) int {
	if dx_CreateFontToHandle == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_CreateFontToHandle.Call(pstring(fontName), pint(size), pint(thick), pint(fontType))
	return int(res)
}

func LoadFontDataToHandle(fileName string, edgeSize int) int {
	if dx_LoadFontDataToHandle == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_LoadFontDataToHandle.Call(pstring(fileName), pint(edgeSize))
	return int(res)
}

func DeleteFontToHandle(fontHandle int) int {
	if dx_DeleteFontToHandle == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DeleteFontToHandle.Call(pint(fontHandle))
	return int(res)
}

func SetFontCacheUsePremulAlphaFlag(flag int) int {
	if dx_SetFontCacheUsePremulAlphaFlag == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetFontCacheUsePremulAlphaFlag.Call(pint(flag))
	return int(res)
}

func DrawStringToHandle(x int, y int, str string, color uint, fontHandle int) int {
	if dx_DrawStringToHandle == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_DrawStringToHandle.Call(pint(x), pint(y), pstring(str), puint(color), pint(fontHandle))
	return int(res)
}

func GetDrawStringWidthToHandle(str string, strLen int, fontHandle int) int {
	if dx_GetDrawStringWidthToHandle == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_GetDrawStringWidthToHandle.Call(pstring(str), pint(strLen), pint(fontHandle))
	return int(res)
}

func GetFontStateToHandle(fontName string, size *int, thick *int, fontHandle int) int {
	if dx_GetFontStateToHandle == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_GetFontStateToHandle.Call(pstring(fontName), ppint(size), ppint(thick), pint(fontHandle))
	return int(res)
}

func InitFontToHandle() int {
	if dx_InitFontToHandle == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_InitFontToHandle.Call()
	return int(res)
}

func clsDx() int {
	if dx_clsDx == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_clsDx.Call()
	return int(res)
}

func SetGraphMode(sizeX int, sizeY int, colorBitNum int) int {
	if dx_SetGraphMode == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetGraphMode.Call(pint(sizeX), pint(sizeY), pint(colorBitNum))
	return int(res)
}

func SetFullScreenResolutionMode(resolutionMode int) int {
	if dx_SetFullScreenResolutionMode == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetFullScreenResolutionMode.Call(pint(resolutionMode))
	return int(res)
}

func SetFullScreenScalingMode(scalingMode int) int {
	if dx_SetFullScreenScalingMode == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetFullScreenScalingMode.Call(pint(scalingMode))
	return int(res)
}

func GetScreenState(sizeX *int, sizeY *int, colorBitDepth *int) int {
	if dx_GetScreenState == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_GetScreenState.Call(ppint(sizeX), ppint(sizeY), ppint(colorBitDepth))
	return int(res)
}

func SetDrawArea(x1 int, y1 int, x2 int, y2 int) int {
	if dx_SetDrawArea == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetDrawArea.Call(pint(x1), pint(y1), pint(x2), pint(y2))
	return int(res)
}

func ClearDrawScreen() int {
	if dx_ClearDrawScreen == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_ClearDrawScreen.Call()
	return int(res)
}

func SetBackgroundColor(red int, green int, blue int) int {
	if dx_SetBackgroundColor == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetBackgroundColor.Call(pint(red), pint(green), pint(blue))
	return int(res)
}

func GetColor(red int, green int, blue int) uint {
	if dx_GetColor == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_GetColor.Call(pint(red), pint(green), pint(blue))
	return uint(res)
}

func SetDrawScreen(drawScreen int) int {
	if dx_SetDrawScreen == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetDrawScreen.Call(pint(drawScreen))
	return int(res)
}

func ScreenFlip() int {
	if dx_ScreenFlip == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_ScreenFlip.Call()
	return int(res)
}

func SetFullSceneAntiAliasingMode(samples int, quality int) int {
	if dx_SetFullSceneAntiAliasingMode == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetFullSceneAntiAliasingMode.Call(pint(samples), pint(quality))
	return int(res)
}

func PlayMovie(fileName string, exRate int, playType int) int {
	if dx_PlayMovie == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_PlayMovie.Call(pstring(fileName), pint(exRate), pint(playType))
	return int(res)
}

func PlayMovieToGraph(graphHandle int) int {
	if dx_PlayMovieToGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_PlayMovieToGraph.Call(pint(graphHandle))
	return int(res)
}

func PauseMovieToGraph(graphHandle int) int {
	if dx_PauseMovieToGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_PauseMovieToGraph.Call(pint(graphHandle))
	return int(res)
}

func SeekMovieToGraph(graphHandle int, time int) int {
	if dx_SeekMovieToGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SeekMovieToGraph.Call(pint(graphHandle), pint(time))
	return int(res)
}

func TellMovieToGraph(graphHandle int) int {
	if dx_TellMovieToGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_TellMovieToGraph.Call(pint(graphHandle))
	return int(res)
}

func GetMovieStateToGraph(graphHandle int) int {
	if dx_GetMovieStateToGraph == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_GetMovieStateToGraph.Call(pint(graphHandle))
	return int(res)
}

func ChangeWindowMode(flag int) int {
	if dx_ChangeWindowMode == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_ChangeWindowMode.Call(pint(flag))
	return int(res)
}

func SetOutApplicationLogValidFlag(flag int) int {
	if dx_SetOutApplicationLogValidFlag == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_SetOutApplicationLogValidFlag.Call(pint(flag))
	return int(res)
}

func ppint(i *int) uintptr {
	return uintptr(unsafe.Pointer(i))
}

func pint(i int) uintptr {
	return uintptr(i)
}

func puint(ui uint) uintptr {
	return uintptr(ui)
}

func pstring(str string) uintptr {
	sjisStr, _, err := transform.String(japanese.ShiftJIS.NewEncoder(), str)
	if err != nil {
		panic(err)
	}
	pbyte, err := syscall.BytePtrFromString(sjisStr)
	if err != nil {
		panic(err)
	}
	return uintptr(unsafe.Pointer(pbyte))
}

func pfloat32(f float32) uintptr {
	return uintptr(f)
}

func pfloat64(f float64) uintptr {
	return uintptr(f)
}
