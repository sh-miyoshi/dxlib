package dxlib

//go:generate go run cmd/generate.go -i dxlib.go -o dxlib_gen.go
//go:generate gofmt -w dxlib_gen.go

// DXライブラリリファレンス
// https://dxlib.xsrv.jp/dxfunc.html

// 基本は先頭にdxlibをつけてC言語の関数名を書く
// ただし、配列を扱いたい場合はarray<型>(arrayintなど)のように書く
// commentから始まる行はコメント、ext_dxlibから始まる行はgenerator側で関数の作成が必要
// そのほかの行は無視される

// 使用必須関数
//dxlib int DxLib_Init()
//dxlib int DxLib_End()
//dxlib int ProcessMessage()

// 図形描画関数
//comment; DrawLine; 線を描画\n引数\n  x1, y1: 線の起点座標\n  x2, y2: 線の終点座標\n  color: 線の色\n  thickness: 文字の太さ(デフォルト: 1)
//dxlib int DrawLine(int x1, int y1, int x2, int y2, unsigned int color, int thickness)
//comment; DrawLineAA; 線を描画(アンチエイリアス効果付き)\n引数\n  x1, y1: 線の起点座標\n  x2, y2: 線の終点座標\n  color: 線の色\n  thickness: 文字の太さ(デフォルト: 1.0)
//dxlib int DrawLineAA(float x1, float y1, float x2, float y2, unsigned int color, float thickness)
//comment; DrawBox; 四角形を描画\n引数\n  x1, y1: 四角形の左上の頂点座標\n  x2, y2: 四角形の右下＋１の頂点座標\n  color: 四角形の色\n  fillFlag: 四角の中身を塗りつぶすか(TRUEで塗りつぶし)
//dxlib int DrawBox(int x1, int y1, int x2, int y2, unsigned int color, int fillFlag)
//comment; DrawBoxAA; 四角形を描画(アンチエイリアス効果付き)\n引数\n  x1, y1: 四角形の左上の頂点座標\n  x2, y2: 四角形の右下＋１の頂点座標\n  color: 四角形の色\n  fillFlag: 四角の中身を塗りつぶすか(TRUEで塗りつぶし)\n  lineThickness: 文字の太さ(デフォルト: 1.0)
//dxlib int DrawBoxAA(float x1, float y1, float x2, float y2, unsigned int color, int fillFlag, float lineThickness)
//comment; DrawCircle; 円を描画\n引数\n  x, y: 円の中心座標\n  r: 半径\n  color: 円の色\n  fillFlag: 塗りつぶすか(TRUEで塗りつぶし)\n  lineThickness: 文字の太さ(デフォルト: 1)
//dxlib int DrawCircle(int x, int y, int r, unsigned int color, int fillFlag, int lineThickness)
//comment; DrawCircleAA; 円を描画(アンチエイリアス効果付き)\n引数\n  x, y: 円の中心座標\n  r: 半径\n  posnum: 円を形成する頂点の数\n  color: 円の色\n  fillFlag: 塗りつぶすか(TRUEで塗りつぶし)\n  lineThickness: 文字の太さ(デフォルト: 1.0)
//dxlib int DrawCircleAA(float x, float y, float r, int posnum, unsigned int color, int fillFlag, float lineThickness)
//comment; DrawOval; 楕円を描画\n引数\n  x, y: 楕円の中心座標\n  rx, ry: 描く楕円のX軸に対する半径とY軸に対する半径\n  color: 円の色\n  fillFlag: 塗りつぶすか(TRUEで塗りつぶし)\n  lineThickness: 文字の太さ(デフォルト: 1.0)
//dxlib int DrawOval(int x, int y, int rx, int ry, unsigned int color, int fillFlag, int lineThickness)
//comment; DrawOvalAA; 楕円を描画(アンチエイリアス効果付き)\n引数\n  x, y: 楕円の中心座標\n  rx, ry: 描く楕円のX軸に対する半径とY軸に対する半径\n  posnum: 円を形成する頂点の数\n  color: 円の色\n  fillFlag: 塗りつぶすか(TRUEで塗りつぶし)\n  lineThickness: 文字の太さ(デフォルト: 1.0)
//dxlib int DrawOvalAA(float x, float y, float rx, float ry, int posnum, unsigned int color, int fillFlag, float lineThickness)
//comment; DrawTriangle; 三角形の描画\n引数\n  x1, y1, x2, y2, x3, y3: 三角形を描く３つの座標\n  color: 三角形の色\n  fillFlag: 塗りつぶすか(TRUEで塗りつぶし)
//dxlib int DrawTriangle(int x1, int y1, int x2, int y2, int x3, int y3, unsigned int color, int fillFlag)
//comment; DrawTriangle; 三角形の描画(アンチエイリアス効果付き)\n引数\n  x1, y1, x2, y2, x3, y3: 三角形を描く３つの座標\n  color: 三角形の色\n  fillFlag: 塗りつぶすか(TRUEで塗りつぶし)\n  lineThickness: 文字の太さ(デフォルト: 1.0)
//dxlib int DrawTriangleAA(float x1, float y1, float x2, float y2, float x3, float y3, unsigned int color, int fillFlag, float lineThickness)
//comment; DrawPixel; 点を描画\n引数\n  x, y: 座標\n  color: 点の色
//dxlib int DrawPixel(int x, int y, unsigned int color)
//comment; DrawPixel; 指定点の色を取得\n引数\n  x, y: 座標
//dxlib unsigned int GetPixel(int x, int y)

// グラフィックデータ制御関数
//dxlib int LoadGraphScreen(int x, int y, char *graphName, int transFlag)
// TODO notUse3DFlag false
//dxlib int LoadGraph(char *fileName, int notUse3DFlag)
//dxlib int LoadDivGraph(char *fileName, int allnum, int xnum, int ynum, int xsize, int ysize, arrayint handleBuf, int notUse3DFlag)
//dxlib int MakeGraph(int sizeX, int sizeY, int notUse3DFlag)
//dxlib int MakeScreen(int sizeX, int sizeY, int useAlphaChannel)
//dxlib int SetCreateDrawValidGraphMultiSample(int samples, int quality)
//dxlib int SetCreateGraphColorBitDepth(int bitDepth)
//dxlib int SetDrawValidFloatTypeGraphCreateFlag(int flag)
//dxlib int SetCreateDrawValidGraphChannelNum(int channelNum)
//dxlib int SetUsePremulAlphaConvertLoad(int useFlag)
//dxlib int DrawGraph(int x, int y, int grHandle, int transFlag)
//dxlib int DrawTurnGraph(int x, int y, int grHandle, int transFlag)
//dxlib int DrawExtendGraph(int x1, int y1, int x2, int y2, int grHandle, int transFlag)
//dxlib int DrawRotaGraph(int x, int y, double extRate, double angle, int grHandle, int transFlag, int reverseXFlag, int reverseYFlag)
//dxlib int DrawRotaGraph2(int x, int y, int cx, int cy, double extRate, double angle, int grHandle, int transFlag, int reverseXFlag, int reverseYFlag)
//dxlib int DrawRotaGraph3(int x, int y, int cx, int cy, double extRateX, double extRateY, double angle, int grHandle, int transFlag, int reverseXFlag, int reverseYFlag)
//dxlib int DrawModiGraph(int x1, int y1, int x2, int y2, int x3, int y3, int x4, int y4, int grHandle, int transFlag)
//dxlib int DrawRectGraph(int destX, int destY, int srcX, int srcY, int width, int height, int graphHandle, int transFlag, int reverseXFlag, int reverseYFlag)
//dxlib int DerivationGraph(int srcX, int srcY, int width, int height, int srcGraphHandle)
// UseClientFlag = TRUE
//dxlib int GetDrawScreenGraph(int x1, int y1, int x2, int y2, int grHandle, int useClientFlag)
//dxlib int GetGraphiteSize(int grHandle, int *sizeXBuf, int *sizeYBuf)
//dxlib int InitGraph(int logOutFlag)
//dxlib int DeleteGraph(int grHandle)
//dxlib int SetDrawMode(int drawMode)
//dxlib int SetDrawBlendMode(int blendMode, int pal)
//dxlib int SetDrawBright(int redBright, int greenBright, int blueBright)
//dxlib int SetTransColor(int red, int green, int blue)
//dxlib int LoadBlendGraph(char *fileName)
//dxlib int DrawBlendGraph(int x, int y, int grHandle, int transFlag, int blendGraph, int borderParam, int borderRange)

// 文字描画関係関数
// edgeColor = 0
//dxlib int DrawString(int x, int y, char *str, unsigned int color, unsigned int edgeColor)
// vericalFlag = FALSE
//dxlib int GetDrawStringWidth(char *str, int strLen, int vericalFlag)
//dxlib int SetFontSize(int fontSize)
//dxlib int SetFontThickness(int tinckPal)
// charSet = -1
//dxlib int ChangeFont(char *fontName, int charSet)
//dxlib int ChangeFontType(int fontType)
//dxlib int CreateFontToHandle(char *fontName, int size, int thick, int fontType, int charSet, int edgeSize, int italic, int handle)
//dxlib int LoadFontDataToHandle(char *fileName, int edgeSize)
//dxlib int DeleteFontToHandle(int fontHandle)
//dxlib int SetFontCacheUsePremulAlphaFlag(int flag)
//dxlib int DrawStringToHandle(int x, int y, char *str, unsigned int color, int fontHandle, int edgeColor, int verticalFlag)
//dxlib int GetDrawStringWidthToHandle(char *str, int strLen, int fontHandle, int verticalFlag)
//dxlib int GetFontStateToHandle(char *fontName, int *size, int *thick, int fontHandle, int *fontType, int *charSet, int *edgeSize, int *italic)
//dxlib int InitFontToHandle()

// 簡易画面出力関数
//dxlib int clsDx()

// その他画面操作系関数
// refreshRate = 60
//dxlib int SetGraphMode(int sizeX, int sizeY, int colorBitNum, int refreshRate)
//dxlib int SetFullScreenResolutionMode(int resolutionMode)
//dxlib int SetFullScreenScalingMode(int scalingMode, int fitScaling)
//dxlib int GetScreenState(int *sizeX, int *sizeY, int *colorBitDepth)
//dxlib int SetDrawArea(int x1, int y1, int x2, int y2)
//dxlib int ClearDrawScreen()
//dxlib int SetBackgroundColor(int red, int green, int blue, int alpha)
//dxlib unsigned int GetColor(int red, int green, int blue)
//dxlib int SetDrawScreen(int drawScreen)
//dxlib int ScreenFlip()
//dxlib int SetFullSceneAntiAliasingMode(int samples, int quality)

// 動画関係関数
//dxlib int PlayMovie(char *fileName, int exRate, int playType)
//dxlib int PlayMovieToGraph(int graphHandle, int playType, int sysPlay)
//dxlib int PauseMovieToGraph(int graphHandle, int sysPause)
//dxlib int SeekMovieToGraph(int graphHandle, int time)
//dxlib int TellMovieToGraph(int graphHandle)
//dxlib int GetMovieStateToGraph(int graphHandle)

// マスク関係関数
//dxlib int CreateMaskScreen()
//dxlib int DeleteMaskScreen()
//dxlib int LoadMask(char *fileName)
//dxlib int LoadDivMask(char *fileName, int allnum, int xnum, int ynum, int xsize, int ysize, arrayint handleBuf)
//dxlib int DrawMask(int x, int y, int maskHandle, int transMode)
//dxlib int DrawFillMask(int x1, int y1, int x2, int y2, int maskHandle)
//dxlib int DeleteMask(int maskHandle)
//dxlib int InitMask()
//dxlib int FillMaskScreen(int flag)
//dxlib int SetUseMaskScreenFlag(int validFlag)
//dxlib int MakeMask(int width, int height)
//dxlib int GetMaskSize(int *widthBuf, int *heightBuf, int maskHandle)

// 入力関係の関数
//dxlib int GetJoypadNum()
//dxlib int GetJoypadInputState(int inputType)
//dxlib int GetJoypadAnalogInput(int *xbuf, int *ybuf, int inputType)
//TODO int GetJoypadDirectInputState(int inputType, DINPUT_JOYSTATE *dinputState)
//TODO int GetJoypadXInputState(int inputType, XINPUT_STATE *xinputState)
//dxlib int SetJoypadDeadZone(int inputType, double zone)
//dxlib int StartJoypadVibration(int inputType, int power, int time)
//dxlib int StopJoypadVibration(int inputType, int effectIndex)
//dxlib int SetMouseDispFlag(int dispFlag)
//dxlib int GetMousePoint(int *xbuf, int *ybuf)
//dxlib int SetMousePoint(int pointX ,int pointY)
//dxlib int GetMouseInput()
//dxlib int GetMouseInputLog2(int *button, int *clickX, int *clickY, int *logType, int logDelete)
//dxlib int GetMouseWheelRotVol(int counterReset)
//dxlib int GetTouchInputNum()
//dxlib int GetTouchInput(int inputNo, int *positionX, int *positionY, int *id, int *device)
//dxlib int CheckHitKeyAll(int checkType)
//dxlib int CheckHitKey(int keyCode)
//dxlib int GetHitKeyStateAll(arraychar keyStateBuf)
//dxlib char GetInputChar(int deleteFlag)
//dxlib char GetInputCharWait(int deleteFlag)
//dxlib int ClearInputCharBuf()
//dxlib int KeyInputString(int x, int y, int charMaxLength, char *strBuffer, int cancelValidFlag)
//dxlib int KeyInputSingleCharString(int x, int y, int charMaxLength, char *strBuffer, int cancelValidFlag)
//dxlib int KeyInputNumber(int x, int y, int maxNum, int minNum, int cancelValidFlag)
//dxlib int SetKeyInputStringColor(int nmlStr, int nmlCur, int imeStrBack, int imeCur, int imeLine, int imeSelectStr, int imeModeStr, int nmlStrE, int imeSelectStrE, int imeModeStrE, int imeSelectWinE, int imeSelectWinF, int selectStrBackColor, int selectStrColor, int selectStrEdgeColor, int imeStr, int imeStrE)
//dxlib int MakeKeyInput(int maxStrLength, int cancelValidFlag, int singleCharOnlyFlag, int numCharOnlyFlag, int doubleCharOnlyFlag, int enableNewLineFlag)
//dxlib int DeleteKeyInput(int inputHandle)
//dxlib int InitKeyInput()
//dxlib int SetActiveKeyInput(int inputHandle)
//dxlib int CheckKeyInput(int inputHandle)
//dxlib int DrawKeyInputString(int x, int y, int inputHandle, int drawCandidateList)
//dxlib int DrawKeyInputModeString(int x, int y)
//dxlib int SetKeyInputString(char *str, int inputHandle)
//dxlib int SetKeyInputNumber(int number, int inputHandle)
//dxlib int GetKeyInputNumber(int inputHandle)

// 音利用関数
//dxlib int PlaySoundFile(char *fileName, int playType)
//dxlib int CheckSoundFile()
//dxlib int StopSoundFile()
//dxlib int LoadSoundMem(char *fileName, int bufferNum, int unionHandle)
//dxlib int PlaySoundMem(int soundHandle, int playType, int topPositionFlag)
//dxlib int CheckSoundMem(int soundHandle)
//dxlib int StopSoundMem(int soundHandle)
//dxlib int DeleteSoundMem(int soundHandle, int logOutFlag)
//dxlib int InitSoundMem(int logOutFlag)
//dxlib int ChangePanSoundMem(int panPal, int soundHandle)
//dxlib int ChangeVolumeSoundMem(int volumePan, int soundHandle)
//dxlib int ChangeNextPlayPanSoundMem(int panPal, int soundHandle)
//dxlib int ChangeNextPlayVolumeSoundMem(int volumePal, int soundHandle)
//dxlib int SetFrequencySoundMem(int frequencyPal, int soundHandle)
//dxlib int SetLoopPosSoundMem(int loopTime, int soundHandle)
//dxlib int SetLoopSamplePosSoundMem(int loopSamplePosition, int soundHandle)
//dxlib int SetCurrentPositionSoundMem(int samplePosition, int soundHandle)
//dxlib int DuplicateSoundMem(int srcSoundHandle, int bufferNum)
//dxlib int SetCreateSoundPitchRate(float cents)
//dxlib int SetCreateSoundTimeStretchRate(float rate)
//TODO int Set3DPositionSoundMem(VECTOR position, int soundHandle)
//dxlib int Set3DRadiusSoundMem(float radius, int soundHandle)
//TODO int Set3DVelocitySoundMem(VECTOR velocity, int soundHandle)
//TODO int SetNextPlay3DPositionSoundMem(VECCTOR position, int soundHandle)
//dxlib int SetNextPlay3DRadiusSoundMem(float radius, int soundHandle)
//TODO int SetNextPlay3DVelocitySoundMem(VECTOR velocity, int soundHandle)
//TODO int Set3DReverbParamSoundMem(const SOUND3D_REVERB_PARAM *param, int soundHandle)
//dxlib int Set3DPresetReverbParamSoundMem(int presetNo, int soundHandle)
//TODO int Get3DPresetReverbParamSoundMem(SOUND3D_REVERB_PARAM *paramBuffer, int presetNo)
//TODO int Set3DReverbParamSoundMemAll(const SOUND3D_REVERB_PARAM *param, int playSoundOnly)
//dxlib int Set3DPresetReverbParamSoundMemAll(int presetNo, int playSoundOnly)
//dxlib int SetCreate3DSoundFlag(int flag)
//dxlib int SetEnableXAudioFlag(int flag)
//dxlib int Set3DSoundOneMetre(float distance)
//TODO int Set3DSoundListenerPosAndFrontPos_UpVecY(VECTOR position, VECTOR frontPosition)
//TODO int Set3DSoundListenerPosAndFrontPosAndUpVec(VECTOR position, VECTOR frontPosition, VECTOR upVector)
//TODO int Set3DSoundListenerVelocity(VECTOR velocity)

// 音楽再生関数
//dxlib int PlayMusic(char *fileName, int playType)
//dxlib int CheckMusic()
//dxlib int StopMusic()
//dxlib int SetVolumeMusic(int volume)

// ウエイト関係の関数
//dxlib int WaitTimer(int waitTime)
//dxlib int WaitVSync(int syncNum)
//dxlib int WaitKey()

// 時間関係の関数
//dxlib int GetNowCount(int useRDTSCFlag)
//dxlib LONGLONG GetNowHiPerformanceCount(int useRDTSCFlag)
//TODO int GetDateTime(DATEDATA *dataBuf)

// 乱数取得関数
//dxlib int GetRand(int randMax)
//dxlib int SRand(int seed)

// ウインドウモード関係
//dxlib int ChangeWindowMode(int flag)
//dxlib int SetMainWindowText(char *windowText)
//dxlib int SetWindowIconID(int id)
//dxlib int SetWindowSizeChangeEnableFlag(int flag, int fitScreen)
//dxlib int SetWindowSizeExtendRate(double exRateX, double exRateY)

// 通信関係
//TODO int ConnectNetWork(IPDATA ipData, int port)
//dxlib int CloseNetWork(int netHandle)
//dxlib int PreparationListenNetWork(int port)
//dxlib int StopListenNetWork()
//dxlib int GetNetWorkDataLength(int netHandle)
//dxlib int GetNetWorkSendDataLength(int netHandle)
//dxlib int GetNewAcceptNetWork()
//dxlib int GetLostNetWork()
//dxlib int GetNetWorkAcceptState(int netHandle)
//TODO int GetNetWorkIP(int netHandle, IPDATA *ipBuf)
//dxlib int MakeUDPSocket(int recvPort)
//dxlib int DeleteUDPSocket(int netUDPHandle)
//dxlib int CheckNetWorkRecvUDP(int netUDPHandle)

// ファイル読み込み関係
//dxlib int FileRead_open(char *filePath, int async)
//dxlib LONGLONG FileRead_size(char *filePath)
//dxlib int FileRead_close(int fileHandle)
//dxlib LONGLONG FileRead_tell(int fileHandle)
//dxlib int FileRead_seek(int fileHandle, LONGLONG offset, int origin)
//dxlib int FileRead_eof(int fileHandle)
//dxlib int FileRead_gets(arraychar buffer, int num, int fileHandle)
//dxlib int FileRead_getc(int fileHandle)

// ドット単位で画像にアクセスしたい関係
//dxlib int LoadSoftImage(char *fileName)
//dxlib int LoadARGB8ColorSoftImage(char *fileName)
//dxlib int LoadXRGB8ColorSoftImage(char *fileName)
//dxlib int MakeARGB8ColorSoftImage(int sizeX, int sizeY)
//dxlib int MakeXRGB8ColorSoftImage(int sizeX, int sizeY)
//dxlib int MakePAL8ColorSoftImage(int sizeX, int sizeY, int useAlpha)
//dxlib int DeleteSoftImage(int siHandle)
//dxlib int InitSoftImage()
//dxlib int GetSoftImageSize(int siHandle, int *width, int *height)
//dxlib int FillSoftImage(int siHandle, int r, int g, int b, int a)
//dxlib int SetPaletteSoftImage(int siHandle, int paletteNo, int r, int g, int b, int a)
//dxlib int GetPaletteSoftImage(int siHandle, int paletteNo, int *r, int *g, int *b, int *a)
//dxlib int DrawPixelPalCodeSoftImage(int siHandle, int x, int y, int palNo)
//dxlib int GetPixelPalCodeSoftImage(int siHandle, int x, int y)
//dxlib int DrawPixelSoftImage(int siHandle, int x, int y, int r, int g, int b, int a)
//dxlib int GetPixelSoftImage(int siHandle, int x, int y, int *r, int *g, int *b, int *a)
//dxlib int BltSoftImage(int srcX, int srcY, int srcSizeX, int srcSizeY, int srcSIHandle, int destX, int destY, int destSIHandle)
//dxlib int DrawSoftImage(int x, int y, int siHandle)
//dxlib int CreateGraphFromSoftImage(int siHandle)
//dxlib int CreateDivGraphFromSoftImage(int siHandle, int allnum, int xnum, int ynum, int sizeX, int sizeY, arrayint handleBuf)

// 非同期読み込み関係
//dxlib int SetUseASyncLoadFlag(int flag)
//dxlib int CheckHandleASyncLoad(int handle)
//dxlib int GetASyncLoadNum()

// 文字関係関数
//dxlib int SetUseCharCodeFormat(int charCodeFormat)

// マイナー関数
//dxlib int SetAlwaysRunFlag(int flag)
//dxlib int SetOutApplicationLogValidFlag(int flag)
//dxlib int SetUseDXArchiveFlag(int flag)
//dxlib int SetDXArchiveExtension(char *extension)
//dxlib int SetDXArchiveKeyString(char *keyString)
//dxlib int SetEmulation320x240(int flag)
//dxlib int SetUse3DFlag(int flag)
//dxlib int SetWaitVSyncFlag(int flag)
//dxlib int SetUseDivGraphFlag(int flag)
//dxlib int LoadPauseGraph(char *fileName)
//dxlib int ScreenCopy()
//dxlib int GetColorBitDepth()
//dxlib int SaveDrawScreen(int x1, int y1, int x2, int y2, char *fileName, int saveType, int jpegQuality, int jpegSample2x1, int pngCompressionLevel)
//dxlib int EnumFontName(arraychar nameBuffer, int nameBufferNum, int japanOnlyFlag)
//dxlib int DrawVString(int x, int y, char *str, unsigned int color, unsigned int edgeColor)
//dxlib int DrawVStringToHandle(int x, int y, char *str, unsigned int color, int fontHandle, unsigned int edgeColor)
//dxlib int ReloadFileGraphAll()
//dxlib int SetCreateSoundDataType(int soundDataType)
//dxlib int SelectMidiMode(int mode)

// Extra functions
// 直接DXライブラリの関数を呼ばずにGo言語側で処理することで呼べる関数
//ext_dxlib DrawFormatString
//ext_dxlib DrawFormatStringToHandle

// Unsupported functions
// 公式のreadme.txtより
// > 基本的に「可変引数の関数」「void 型のポインタを引数で取る関数」「コールバック関数を引数で取る関数」が使えません。
// > また、ポインタを返す関数やポインタが含まれている構造体を扱う関数も使えません
//unsupported int GraphFilter(int grHandle, int filterType, ...)
//unsupported int GraphFilterBlt(int srcGrHandle, int destGrHandle, int filterType, ...)
//unsupported int GraphFilterRectBlt(int srcGrHandle, int destGrHandle, int srcX1, int srcY1, int srcX2, int srcY2, int destX, int destY, int filterType, ...)
//unsupported int GraphBlend(int grHandle, int blendGrHandle, int blendRatio, int blendType, ...)
//unsupported int GraphBlendBlt(int srcGrHandle, int blendGrHandle, int destGrHandle, int blendRatio, int blendType, ...)
//unsupported int GraphBlendRectBlt( int srcGrHandle, int blendGrHandle, int destGrHandle, int srcX1, int srcY1, int srcX2, int srcY2, int blendX, int blendY, int destX, int destY, int blendRatio, int blendType, ...)
//unsupported int GetDrawFormatStringWidth(char *format, ...)
//unsupported int GetDrawFormatStringWidthToHandle(int fontHandle, char *format, ...)
//unsupported int printfDx(char *format, ...)
//unsupported int FileRead_scanf(int fileHandle, char *format, ...)
//unsupported int SetDataToMask(int width, int height, void *maskData, int maskHandle)
//unsupported int DrawMaskToDirectData(int x, int y, int width, int height, void *maskData, int transMode)
//unsupported int DrawFillMaskToDirectData(int x1, int y1, int x2, int y2, int width, int height, void *maskData)
//unsupported int NetWorkSend(int netHandle, void *buffer, int length)
//unsupported int NetWorkRecv(int netHandle, void *buffer, int length)
//unsupported int NetWorkRecvToPeek(int netHandle, void *buffer, int length)
//unsupported int NetWorkSendUDP(int netUDPHandle, IPDATA sendIP, int sendPort, void *buffer, int length)
//unsupported int NetWorkRecvUDP(int netUDPHandle, IPDATA *recvIP, int *recvPort, void *buffer, int length, int peek)
//unsupported int FileRead_read(void *buffer, int readSize, int fileHandle)
//unsupported int LoadSoftImageToMem(void *fileImage, int fileImageSize)
//unsupported int LoadARGB8ColorSoftImageToMem(void *fileImage, int fileImageSize)
//unsupported int LoadXRGB8ColorSoftImageToMem(void *fileImage, int fileImageSize)
//unsupported int GetCharBytes(int charCodeFormat, void *str)
//unsupported int CreateGraphFromMem(void *memImage, int memImageSize)
//unsupported int ReCreateGraphFromMem(void *memImage, int memImageSize, int grHandle)
//unsupported int SetRestoreGraphCallback(void (* callback)())
//unsupported int LoadSoundMemByMemImage(void *fileImageBuffer, int imageSize)
