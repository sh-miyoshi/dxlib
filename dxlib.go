package dxlib

//go:generate go run cmd/generate.go -i dxlib.go -o dxlib_gen.go
//go:generate gofmt -w dxlib_gen.go

// DXライブラリリファレンス
// https://dxlib.xsrv.jp/dxfunc.html

// 基本は先頭にdxlibをつけてC言語の関数名を書く
// ただし、配列を扱いたい場合はarray<型>(arrayintなど)のように書く

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
//dxlib int CreateMaskScreen()
//dxlib int DeleteMaskScreen()
//dxlib int LoadMask(char *fileName)
//dxlib int LoadDivMask(char *fileName, int allnum, int xnum, int ynum, int xsize, int ysize, int *handleBuf)
//dxlib int DrawMask(int x, int y, int maskHandle, int transMode)
//dxlib int DrawFillMask(int x1, int y1, int x2, int y2, int maskHandle)
//dxlib int DeleteMask(int maskHandle)
//dxlib int InitMask()
//dxlib int FillMaskScreen(int flag)
//dxlib int SetUseMaskScreenFlag(int validFlag)
//dxlib int MakeMask(int width, int height)
//dxlib int GetMaskSize(int *widthBuf, int *heightBuf, int maskHandle)
//TODO int SetDataToMask(int width, int height, void *maskData, int maskHandle)
//TODO int DrawMaskToDirectData(int x, int y, int width, int height, void *maskData, int transMode)
//TODO int DrawFillMaskToDirectData(int x1, int y1, int x2, int y2, int width, int height, void *maskData)

// 入力関係の関数
//dxlib int GetJoypadNum()
//dxlib int GetJoypadInputState(int inputType)
//dxlib int GetJoypadAnalogInput(int *xbuf, int *ybuf, int inputType)
//TODO int GetJoypadDirectInputState(int inputType, DINPUT_JOYSTATE *dinputState)
//TODO int GetJoypadXInputState(int inputType, XINPUT_STATE *xinputState)
//dxlib int SetJoypadDeadZone(int inputType, double zone)
//dxlib int StartJoypadVibration(int inputType, int power, int time)
//dxlib int StopJoypadVibration(int inputType)
//dxlib int SetMouseDispFlag(int dispFlag)
//dxlib int GetMousePoint(int *xbuf, int *ybuf)
//dxlib int SetMousePoint(int pointX ,int pointY)
//dxlib int GetMouseInput()
//dxlib int GetMouseInputLog2(int *button, int *clickX, int *clickY, int *logType, int logDelete)
//dxlib int GetMouseWheelRotVol()
//dxlib int GetTouchInputNum()
//dxlib int GetTouchInput(int inputNo, int *positionX, int *positionY, int *id, int *device)
//dxlib int CheckHitKeyAll(int checkType)
//dxlib int CheckHitKey(int keyCode)
//TODO int GetHitKeyStateAll(char *keyStateBuf) // char array
//dxlib char GetInputChar(int deleteFlag)
//dxlib char GetInputCharWait(int deleteFlag)
//dxlib int ClearInputCharBuf()
//dxlib int KeyInputString(int x, int y, int charMaxLength, char *strBuffer, int cancelValidFlag)
//dxlib int KeyInputSingleCharString(int x, int y, int charMaxLength, char *strBuffer, int cancelValidFlag)
//dxlib int KeyInputNumber(int x, int y, int maxNum, int minNum, int cancelValidFlag)
//dxlib int SetKeyInputStringColor(int nmlStr, int nmlCur, int imeStrBack, int imeCur, int imeLine, int imeSelectStr, int imeModeStr, int nmlStrE, int imeSelectStrE, int imeModeStrE, int imeSelectWinE, int imeSelectWinF, int selectStrBackColor, int selectStrColor, int selectStrEdgeColor, int imeStr, int imeStrE)
//dxlib int MakeKeyInput(int maxStrLength, int cancelValidFlag, int singleCharOnlyFlag, int numCharOnlyFlag)
//dxlib int DeleteKeyInput(int inputHandle)
//dxlib int InitKeyInput()
//dxlib int SetActiveKeyInput(int inputHandle)
//dxlib int CheckKeyInput(int inputHandle)
//dxlib int DrawKeyInputString(int x, int y, int inputHandle)
//dxlib int DrawKeyInputModeString(int x, int y)
//dxlib int SetKeyInputString(char *str, int inputHandle)
//dxlib int SetKeyInputNumber(int number, int inputHandle)
//dxlib int GetKeyInputNumber(int inputHandle)

// 音利用関数
//dxlib int PlaySoundFile(char *fileName, int playType)
//dxlib int CheckSoundFile()
//dxlib int StopSoundFile()
//dxlib int LoadSoundMem(char *fileName)
//dxlib int PlaySoundMem(int soundHandle, int playType, int topPositionFlag)
//dxlib int CheckSoundMem(int soundHandle)
//dxlib int StopSoundMem(int soundHandle)
//dxlib int DeleteSoundMem(int soundHandle)
//dxlib int InitSoundMem()
//dxlib int ChangePanSoundMem(int panPal, int soundHandle)
//dxlib int ChangeVolumeSoundMem(int volumePan, int soundHandle)
//dxlib int ChangeNextPlayPanSoundMem(int panPal, int soundHandle)
//dxlib int ChangeNextPlayVolumeSoundMem(int volumePal, int soundHandle)
//dxlib int SetFrequencySoundMem(int frequencyPal, int soundHandle)
//dxlib int SetLoopPosSoundMem(int loopTime, int soundHandle)
//dxlib int SetLoopSamplePosSoundMem(int loopSamplePosition, int soundHandle)
//dxlib int SetCurrentPositionSoundMem(int samplePosition, int soundHandle)
//dxlib int DuplicateSoundMem(int srcSoundHandle)
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
//dxlib int GetNowCount()
//dxlib LONGLONG GetNowHiPerformanceCount()
//TODO int GetDateTime(DATEDATA *dataBuf)

// 乱数取得関数
//dxlib int GetRand(int randMax)
//dxlib int SRand(int seed)

// ウインドウモード関係
//dxlib int ChangeWindowMode(int flag)
//dxlib int SetMainWindowText(char *windowText)
//dxlib int SetWindowIconID(int id)
//dxlib int SetWindowSizeChangeEnableFlag(int flag)
//dxlib int SetWindowSizeExtendRate(double exRate)

// 通信関係
//TODO int ConnectNetWork(IPDATA ipData, int port)
//dxlib int CloseNetWork(int netHandle)
//dxlib int PreparationListenNetWork(int port)
//dxlib int StopListenNetWork()
//TODO int NetWorkSend(int netHandle, void *buffer, int length)
//dxlib int GetNetWorkDataLength(int netHandle)
//dxlib int GetNetWorkSendDataLength(int netHandle)
//TODO int NetWorkRecv(int netHandle, void *buffer, int length)
//TODO int NetWorkRecvToPeek(int netHandle, void *buffer, int length)
//dxlib int GetNewAcceptNetWork()
//dxlib int GetLostNetWork()
//dxlib int GetNetWorkAcceptState(int netHandle)
//TODO int GetNetWorkIP(int netHandle, IPDATA *ipBuf)
//dxlib int MakeUDPSocket(int recvPort)
//dxlib int DeleteUDPSocket(int netUDPHandle)
//TODO int NetWorkSendUDP(intnetUDPHandle, IPDATA sendIP, int sendPort, void *buffer, int length)
//TODO int NetWorkRecvUDP(int netUDPHandle, IPDATA *recvIP, int *recvPort, void *buffer, int length, int peek)
//dxlib int CheckNetWorkRecvUDP(int netUDPHandle)

// ファイル読み込み関係
//dxlib int FileRead_open(char *filePath, int async)
//dxlib LONGLONG FileRead_size(char *filePath)
//dxlib int FileRead_close(int fileHandle)
//dxlib LONGLONG FileRead_tell(int fileHandle)
//dxlib int FileRead_seek(int fileHandle, LONGLONG offset, int origin)
//TODO int FileRead_read(void *buffer, int readSize, int fileHandle)
//dxlib int FileRead_eof(int fileHandle)
//TODO int FileRead_gets(char *buffer, int num, int fileHandle) // array char
//dxlib int FileRead_getc(int fileHandle)
//TODO int FileRead_scanf(int fileHandle, char *format, ...)

// ドット単位で画像にアクセスしたい関係
//dxlib int LoadSoftImage(char *fileName)
//dxlib int LoadARGB8ColorSoftImage(char *fileName)
//dxlib int LoadXRGB8ColorSoftImage(char *fileName)
//TODO int LoadSoftImageToMem(void *fileImage, int fileImageSize)
//TODO int LoadARGB8ColorSoftImageToMem(void *fileImage, int fileImageSize)
//TODO int LoadXRGB8ColorSoftImageToMem(void *fileImage, int fileImageSize)
//dxlib int MakeARGB8ColorSoftImage(int sizeX, int sizeY)
//dxlib int MakeXRGB8ColorSoftImage(int sizeX, int sizeY)
//dxlib int MakePAL8ColorSoftImage(int sizeX, int sizeY)
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
//dxlib int CreateDivGraphFromSoftImage(int siHandle, int allnum, int xnum, int ynum, int sizeX, int sizeY, int *handleBuf)

// 非同期読み込み関係
//dxlib int SetUseASyncLoadFlag(int flag)
//dxlib int CheckHandleASyncLoad(int handle)
//dxlib int GetASyncLoadNum()

// 文字関係関数
//dxlib int SetUseCharCodeFormat(int charCodeFormat)
//TODO int GetCharBytes(int charCodeFormat, void *str)

// マイナー関数
//dxlib int SetOutApplicationLogValidFlag(int flag)
//TODO
