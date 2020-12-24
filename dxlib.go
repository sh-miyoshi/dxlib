package dxlib

//go:generate go run cmd/generate.go -i dxlib.go -o dxlib_gen.go
//go:generate gofmt -w dxlib_gen.go

// DXライブラリリファレンス
// https://dxlib.xsrv.jp/dxfunc.html

// 使用必須関数
//dxlib int DxLib_Init()
//dxlib int DxLib_End()
//dxlib int ProcessMessage()

// 図形描画関数
//dxlib int DrawLine(int x1, int y1, int x2, int y2, unsigned int color)
//dxlib int DrawLineAA(float x1, float y1, float x2, float y2, unsigned int color)
//dxlib int DrawBox(int x1, int y1, int x2, int y2, unsigned int color, int fillFlag)
//dxlib int DrawBoxAA(float x1, float y1, float x2, float y2, unsigned int color, int fillFlag)
//dxlib int DrawCircle(int x, int y, int r, unsigned int color, int fillFlag)
//dxlib int DrawCircleAA(float x, float y, float r, int posnum, unsigned int color, int fillFlag)
//dxlib int DrawOval(int x, int y, int rx, int ry, unsigned int color, int fillFlag)
//dxlib int DrawOvalAA(float x, float y, float rx, float ry, int posnum, unsigned int color, int fillFlag)
//dxlib int DrawTriangle(int x1, int y1, int x2, int y2, int x3, int y3, unsigned int color, int fillFlag)
//dxlib int DrawTriangleAA(float x1, float y1, float x2, float y2, float x3, float y3, unsigned int color, int fillFlag)
//dxlib int DrawPixel(int x, int y, unsigned int color)
//dxlib unsigned int GetPixel(int x, int y)

// グラフィックデータ制御関数
//dxlib int LoadGraphScreen(int x, int y, char *graphName, int transFlag)
//dxlib int LoadGraph(char *fileName)
//dxlib int LoadDivGraph(char *fileName, int allnum, int xnum, int ynum, int xsize, int ysize, int *handleBuf)
//dxlib int MakeGraph(int sizeX, int sizeY)
//dxlib int MakeScreen(int sizeX, int sizeY, int useAlphaChannel)
//dxlib int SetCreateDrawValidGraphMultiSample(int samples, int quality)
//dxlib int SetCreateGraphColorBitDepth(int bitDepth)
//dxlib int SetDrawValidFloatTypeGraphCreateFlag(int flag)
//dxlib int SetCreateDrawValidGraphChannelNum(int channelNum)
//dxlib int SetUsePremulAlphaConvertLoad(int useFlag)
//dxlib int DrawGraph(int x, int y, int grHandle, int transFlag)
//dxlib int DrawTurnGraph(int x, int y, int grHandle, int transFlag)
//dxlib int DrawExtendGraph(int x1, int y1, int x2, int y2, int grHandle, int transFlag)
//dxlib int DrawRotaGraph(int x, int y, double extRate, double angle, int grHandle, int transFlag, int turnFlag)
//dxlib int DrawRotaGraph2(int x, int y, int cx, int cy, double extRate, double angle, int grHandle, int transFlag, int turnFlag)
//dxlib int DrawRotaGraph3(int x, int y, int cx, int cy, double extRateX, double extRateY, double angle, int grHandle, int transFlag, int turnFlag)
//dxlib int DrawModiGraph(int x1, int y1, int x2, int y2, int x3, int y3, int x4, int y4, int grHandle, int transFlag)
//dxlib int DrawRectGraph(int destX, int destY, int srcX, int srcY, int width, int height, int graphHandle, int transFlag, int turnFlag)
//dxlib int DerivationGraph(int srcX, int srcY, int width, int height, int srcGraphHandle)
//dxlib int GetDrawScreenGraph(int x1, int y1, int x2, int y2, int grHandle)
//dxlib int GetGraphiteSize(int grHandle, int *sizeXBuf, int *sizeYBuf)
//dxlib int InitGraph()
//dxlib int DeleteGraph(int grHandle)
//dxlib int SetDrawMode(int drawMode)
//dxlib int SetDrawBlendMode(int blendMode, int pal)
//dxlib int SetDrawBright(int redBright, int greenBright, int blueBright)
//dxlib int SetTransColor(int red, int green, int blue)
//dxlib int LoadBlendGraph(char *fileName)
//dxlib int DrawBlendGraph(int x, int y, int grHandle, int transFlag, int blendGraph, int borderParam, int borderRange)
//TODO int GraphFilter(int grHandle, int filterType, ...)
//TODO int GraphFilterBlt(int srcGrHandle, int destGrHandle, int filterType, ...)
//TODO int GraphFilterRectBlt(int srcGrHandle, int destGrHandle, int srcX1, int srcY1, int srcX2, int srcY2, int destX, int destY, int filterType, ...)
//TODO int GraphBlend(int grHandle, int blendGrHandle, int blendRatio, int blendType, ...)
//TODO int GraphBlendBlt(int srcGrHandle, int blendGrHandle, int destGrHandle, int blendRatio, int blendType, ...)
//TODO int GraphBlendRectBlt( int srcGrHandle, int blendGrHandle, int destGrHandle, int srcX1, int srcY1, int srcX2, int srcY2, int blendX, int blendY, int destX, int destY, int blendRatio, int blendType, ...)

// 文字描画関係関数
//dxlib int DrawString(int x, int y, char *str, unsigned int color)
//TODO int DrawFormatString(int x, int y, unsigned int color, char *format, ...)
//dxlib int GetDrawStringWidth(char *str, int strLen)
//TODO int GetDrawFormatStringWidth(char *format, ...)
//dxlib int SetFontSize(int fontSize)
//dxlib int SetFontThickness(int tinckPal)
//dxlib int ChangeFont(char *fontName)
//dxlib int ChangeFontType(int fontType)
//dxlib int CreateFontToHandle(char *fontName, int size, int thick, int fontType)
//dxlib int LoadFontDataToHandle(char *fileName, int edgeSize)
//dxlib int DeleteFontToHandle(int fontHandle)
//dxlib int SetFontCacheUsePremulAlphaFlag(int flag)
//dxlib int DrawStringToHandle(int x, int y, char *str, unsigned int color, int fontHandle)
//TODO int DrawFormatStringToHandle(int x, int y, unsigned int color, int fonthandle, char *format, ...)
//dxlib int GetDrawStringWidthToHandle(char *str, int strLen, int fontHandle)
//TODO int GetDrawFormatStringWidthToHandle(int fontHandle, char *format, ...)
//dxlib int GetFontStateToHandle(char *fontName, int *size, int *thick, int fontHandle)
//dxlib int InitFontToHandle()

// 簡易画面出力関数
//TODO int printfDx(char *format, ...)
//dxlib int clsDx()

// その他画面操作系関数
//dxlib int SetGraphMode(int sizeX, int sizeY, int colorBitNum)
//dxlib int SetFullScreenResolutionMode(int resolutionMode)
//dxlib int SetFullScreenScalingMode(int scalingMode)
//dxlib int GetScreenState(int *sizeX, int *sizeY, int *colorBitDepth)
//dxlib int SetDrawArea(int x1, int y1, int x2, int y2)
//dxlib int ClearDrawScreen()
//dxlib int SetBackgroundColor(int red, int green, int blue)
//dxlib unsigned int GetColor(int red, int green, int blue)
//dxlib int SetDrawScreen(int drawScreen)
//dxlib int ScreenFlip()
//dxlib int SetFullSceneAntiAliasingMode(int samples, int quality)

// 動画関係関数
//dxlib int PlayMovie(char *fileName, int exRate, int playType)
//dxlib int PlayMovieToGraph(int graphHandle)
//dxlib int PauseMovieToGraph(int graphHandle)
//dxlib int SeekMovieToGraph(int graphHandle, int time)
//dxlib int TellMovieToGraph(int graphHandle)
//dxlib int GetMovieStateToGraph(int graphHandle)

// マスク関係関数
//TODO

// 入力関係の関数
//TODO

// 音利用関数
//TODO

// 音楽再生関数
//TODO

// ウエイト関係の関数
//TODO

// 時間関係の関数
//TODO

// 乱数取得関数
//TODO

// ウインドウモード関係
//dxlib int ChangeWindowMode(int flag)
//TODO

// 通信関係
//TODO

// ファイル読み込み関係
//TODO

// ドット単位で画像にアクセスしたい関係
//TODO

// 非同期読み込み関係
//TODO

// 文字関係関数
//TODO

// マイナー関数
//dxlib int SetOutApplicationLogValidFlag(int flag)
//TODO
