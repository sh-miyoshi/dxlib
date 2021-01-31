package dxlib

//go:generate go run cmd/generate/main.go -i dxlib.go -o dxlib_gen.go
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
//comment; DrawLine; 線を描画\n\n引数\n  x1, y1: 線の起点座標\n  x2, y2: 線の終点座標\n  color: 線の色\n  thickness: 文字の太さ(デフォルト: 1)
//dxlib int DrawLine(int x1, int y1, int x2, int y2, unsigned int color, int thickness = 1)
//comment; DrawLineAA; 線を描画(アンチエイリアス効果付き)\n\n引数\n  x1, y1: 線の起点座標\n  x2, y2: 線の終点座標\n  color: 線の色\n  thickness: 文字の太さ(デフォルト: 1.0)
//dxlib int DrawLineAA(float x1, float y1, float x2, float y2, unsigned int color, float thickness = 1.0)
//comment; DrawBox; 四角形を描画\n\n引数\n  x1, y1: 四角形の左上の頂点座標\n  x2, y2: 四角形の右下＋１の頂点座標\n  color: 四角形の色\n  fillFlag: 四角の中身を塗りつぶすか(TRUEで塗りつぶし)
//dxlib int DrawBox(int x1, int y1, int x2, int y2, unsigned int color, int fillFlag)
//comment; DrawBoxAA; 四角形を描画(アンチエイリアス効果付き)\n\n引数\n  x1, y1: 四角形の左上の頂点座標\n  x2, y2: 四角形の右下＋１の頂点座標\n  color: 四角形の色\n  fillFlag: 四角の中身を塗りつぶすか(TRUEで塗りつぶし)\n  lineThickness: 文字の太さ(デフォルト: 1.0)
//dxlib int DrawBoxAA(float x1, float y1, float x2, float y2, unsigned int color, int fillFlag, float lineThickness = 1.0)
//comment; DrawCircle; 円を描画\n\n引数\n  x, y: 円の中心座標\n  r: 半径\n  color: 円の色\n  fillFlag: 塗りつぶすか(TRUEで塗りつぶし)\n  lineThickness: 文字の太さ(デフォルト: 1)
//dxlib int DrawCircle(int x, int y, int r, unsigned int color, int fillFlag, int lineThickness = 1)
//comment; DrawCircleAA; 円を描画(アンチエイリアス効果付き)\n\n引数\n  x, y: 円の中心座標\n  r: 半径\n  posnum: 円を形成する頂点の数\n  color: 円の色\n  fillFlag: 塗りつぶすか(TRUEで塗りつぶし)\n  lineThickness: 文字の太さ(デフォルト: 1.0)
//dxlib int DrawCircleAA(float x, float y, float r, int posnum, unsigned int color, int fillFlag, float lineThickness = 1.0)
//comment; DrawOval; 楕円を描画\n\n引数\n  x, y: 楕円の中心座標\n  rx, ry: 描く楕円のX軸に対する半径とY軸に対する半径\n  color: 円の色\n  fillFlag: 塗りつぶすか(TRUEで塗りつぶし)\n  lineThickness: 文字の太さ(デフォルト: 1.0)
//dxlib int DrawOval(int x, int y, int rx, int ry, unsigned int color, int fillFlag, int lineThickness = 1.0)
//comment; DrawOvalAA; 楕円を描画(アンチエイリアス効果付き)\n\n引数\n  x, y: 楕円の中心座標\n  rx, ry: 描く楕円のX軸に対する半径とY軸に対する半径\n  posnum: 円を形成する頂点の数\n  color: 円の色\n  fillFlag: 塗りつぶすか(TRUEで塗りつぶし)\n  lineThickness: 文字の太さ(デフォルト: 1.0)
//dxlib int DrawOvalAA(float x, float y, float rx, float ry, int posnum, unsigned int color, int fillFlag, float lineThickness = 1.0)
//comment; DrawTriangle; 三角形の描画\n\n引数\n  x1, y1, x2, y2, x3, y3: 三角形を描く３つの座標\n  color: 三角形の色\n  fillFlag: 塗りつぶすか(TRUEで塗りつぶし)
//dxlib int DrawTriangle(int x1, int y1, int x2, int y2, int x3, int y3, unsigned int color, int fillFlag)
//comment; DrawTriangleAA; 三角形の描画(アンチエイリアス効果付き)\n\n引数\n  x1, y1, x2, y2, x3, y3: 三角形を描く３つの座標\n  color: 三角形の色\n  fillFlag: 塗りつぶすか(TRUEで塗りつぶし)\n  lineThickness: 文字の太さ(デフォルト: 1.0)
//dxlib int DrawTriangleAA(float x1, float y1, float x2, float y2, float x3, float y3, unsigned int color, int fillFlag, float lineThickness = 1.0)
//comment; DrawPixel; 点を描画\n\n引数\n  x, y: 座標\n  color: 点の色
//dxlib int DrawPixel(int x, int y, unsigned int color)
//comment; GetPixel; 指定点の色を取得\n\n引数\n  x, y: 座標
//dxlib unsigned int GetPixel(int x, int y)

// グラフィックデータ制御関数
//comment; LoadGraphScreen; 画像ファイルを読みこんで画面に表示する\n\n引数\n  x, y: ロードした画像を描画する矩形の左上頂点の座標\n  graphName: ロードする画像ファイルパス\n  transFlag: 画像の透明度を有効にするかどうか(TRUE: 有効)
//dxlib int LoadGraphScreen(int x, int y, char *graphName, int transFlag)
//comment; LoadGraph; 画像ファイルのメモリへの読みこみ、及び動画ファイルのロード\n\n引数\n  fileName: ファイルパス\n  notUse3DFlag: 3D機能を制限するか(デフォルト: FALSE)
//dxlib int LoadGraph(char *fileName, int notUse3DFlag = FALSE)
//comment; LoadDivGraph; 画像ファイルのメモリへの分割読みこみ\n\n引数\n  fileName: 画像ファイルパス\n  allnum: 画像の分割総数\n  xnum, ynum: 画像の横向きに対する分割数と縦に対する分割数\n  xsize, ysize: 分割された画像一つの大きさ\n  handleBuf: グラフィックハンドルを保存するバッファ\n  notUse3DFlag: 3D機能を制限するか(デフォルト: FALSE)
//dxlib int LoadDivGraph(char *fileName, int allnum, int xnum, int ynum, int xsize, int ysize, arrayint handleBuf, int notUse3DFlag = FALSE)
//comment; MakeGraph; 空のグラフィックを作成する\n\n引数\n  sizeX, sizeY: 作成する空グラフィックのサイズ\n  notUse3DFlag: 3D機能を制限するか(デフォルト: FALSE)
//dxlib int MakeGraph(int sizeX, int sizeY, int notUse3DFlag = FALSE)
//comment; MakeScreen; 描画対象にできるグラフィックを作成する\n\n引数\n  sizeX, sizeY: 作成するグラフィックのサイズ\n useAlphaChannel: 作成するグラフィックにアルファチャンネルを付けるかどうか(TRUE: つける)
//dxlib int MakeScreen(int sizeX, int sizeY, int useAlphaChannel)
//comment; SetCreateDrawValidGraphMultiSample; 描画対象にできるグラフィックのマルチサンプリング設定を行う
//dxlib int SetCreateDrawValidGraphMultiSample(int samples, int quality)
//comment; SetCreateGraphColorBitDepth; 作成するグラフィックのビット深度を設定\n\n引数\n  bitDepth: ビット震度(16 or 32)
//dxlib int SetCreateGraphColorBitDepth(int bitDepth)
//comment; SetDrawValidFloatTypeGraphCreateFlag; 描画可能な浮動小数点型のグラフィックを作成するかどうかの設定(デフォルト: FALSE)
//dxlib int SetDrawValidFloatTypeGraphCreateFlag(int flag)
//comment; SetCreateDrawValidGraphChannelNum; 作成する描画可能なグラフィックのチャンネル数の設定\n\n引数\n  channelNum: 作成する描画可能なグラフィックのチャンネル数(1, 2, or 4)
//dxlib int SetCreateDrawValidGraphChannelNum(int channelNum)
//comment; SetUsePremulAlphaConvertLoad; 読み込み時に画像を乗算済みα画像に変換するかを設定(デフォルト: FALSE)
//dxlib int SetUsePremulAlphaConvertLoad(int useFlag)
//comment; DrawGraph; メモリに読みこんだグラフィックの描画\n\n引数\n  x, y: グラフィックを描画する領域の左上頂点の座標\n  grHandle: グラフィックハンドル\n  transFlag: 画像の透明度を有効にするかどうか(TRUE: 有効にする)
//dxlib int DrawGraph(int x, int y, int grHandle, int transFlag)
//comment; DrawTurnGraph; メモリに読みこんだグラフィックのＬＲ反転描画\n\n引数\n  x, y: グラフィックを描画する領域の左上頂点の座標\n  grHandle: グラフィックハンドル\n  transFlag: 画像の透明度を有効にするかどうか(TRUE: 有効にする)
//dxlib int DrawTurnGraph(int x, int y, int grHandle, int transFlag)
//comment; DrawExtendGraph; メモリに読みこんだグラフィックの拡大縮小描画
//dxlib int DrawExtendGraph(int x1, int y1, int x2, int y2, int grHandle, int transFlag)
//comment; DrawRotaGraph; メモリに読みこんだグラフィックの回転描画\n\n引数\n  x, y: グラフィックを描画する領域の中心座標\n  extRate: 拡大率(1.0で等倍)\n  angle: 描画角度(ラジアン指定)\n  grHandle: グラフィックハンドル\n  transFlag: 画像の透明度を有効にするかどうか(TRUE: 有効にする)\n  reverseXFlag: 画像の左右反転を行うか\n  reverseYFlag: 画像の上下反転を行うか
//dxlib int DrawRotaGraph(int x, int y, double extRate, double angle, int grHandle, int transFlag, int reverseXFlag = FALSE, int reverseYFlag = FALSE)
//comment; DrawRotaGraph2; メモリに読みこんだグラフィックの回転描画(回転中心指定あり)\n\n引数\n  x, y: グラフィックを描画する領域の中心座標\n  cx, cy: 画像を回転描画する画像上の中心座標\n  extRate: 拡大率(1.0で等倍)\n  angle: 描画角度(ラジアン指定)\n  grHandle: グラフィックハンドル\n  transFlag: 画像の透明度を有効にするかどうか(TRUE: 有効にする)\n  reverseXFlag: 画像の左右反転を行うか\n  reverseYFlag: 画像の上下反転を行うか
//dxlib int DrawRotaGraph2(int x, int y, int cx, int cy, double extRate, double angle, int grHandle, int transFlag, int reverseXFlag = FALSE, int reverseYFlag = FALSE)
//comment; DrawRotaGraph3; メモリに読みこんだグラフィックの回転描画(回転中心指定あり)\n\n引数\n  x, y: グラフィックを描画する領域の中心座標\n  cx, cy: 画像を回転描画する画像上の中心座標\n  extRateX, exRateY: 拡大率(1.0で等倍)\n  angle: 描画角度(ラジアン指定)\n  grHandle: グラフィックハンドル\n  transFlag: 画像の透明度を有効にするかどうか(TRUE: 有効にする)\n  reverseXFlag: 画像の左右反転を行うか\n  reverseYFlag: 画像の上下反転を行うか
//dxlib int DrawRotaGraph3(int x, int y, int cx, int cy, double extRateX, double extRateY, double angle, int grHandle, int transFlag, int reverseXFlag = FALSE, int reverseYFlag = FALSE)
//dxlib int DrawRotaGraphFast(int x, int y, float extRate, float angle, int grHandle, int transFlag, int reverseXFlag = FALSE, int reverseYFlag = FALSE)
//dxlib int DrawRotaGraphFast2(int x, int y, int cx, int cy, float extRate, float angle, int grHandle, int transFlag, int reverseXFlag = FALSE, int reverseYFlag = FALSE)
//dxlib int DrawRotaGraphFast3(int x, int y, int cx, int cy, float extRateX, float extRateY, float angle, int grHandle, int transFlag, int reverseXFlag = FALSE, int reverseYFlag = FALSE)
//comment; DrawModiGraph; メモリに読みこんだグラフィックの自由変形描画\n\n引数\n  x1, y1, x2, y2, x3, y3, x4, y4: x1から順に描画する画像の左上、右上、右下、左下の頂点の座標\n  grHandle: グラフィックハンドル\n  transFlag: 画像の透明度を有効にするかどうか(TRUE: 有効にする)
//dxlib int DrawModiGraph(int x1, int y1, int x2, int y2, int x3, int y3, int x4, int y4, int grHandle, int transFlag)
//comment; DrawRectGraph; グラフィックの指定矩形部分のみを描画\n\n引数\n  destX, destY: グラフィックを描画する座標\n  srcX, srcY: 描画するグラフィック上の描画したい矩形の左上座標\n  width, height: 描画するグラフィックのサイズ\n  grHandle: グラフィックハンドル\n  transFlag: 画像の透明度を有効にするかどうか(TRUE: 有効にする)\n  reverseXFlag: 画像の左右反転を行うか\n  reverseYFlag: 画像の上下反転を行うか
//dxlib int DrawRectGraph(int destX, int destY, int srcX, int srcY, int width, int height, int grHandle, int transFlag, int reverseXFlag = FALSE, int reverseYFlag = FALSE)
//comment; DerivationGraph; 指定のグラフィックの指定部分だけを抜き出して新たなグラフィックを作成する\n\n引数\n  srcX, secY: グラフィック中の抜き出したい矩形の左上座標\n  width, height: 抜き出すグラフィックのサイズ\n  srcGraphHandle: グラフィックハンドル
//dxlib int DerivationGraph(int srcX, int srcY, int width, int height, int srcGraphHandle)
//comment; GetDrawScreenGraph; 描画先に設定されているグラフィック領域から指定領域のグラフィックを読みこむ\n\n引数\n  x1, y1: 取得するグラフィック領域（矩形）の左上頂点の座標\n  x2, y2: 取得するグラフィック領域の右下頂点＋１の座標\n  grHandle: 取り込んだグラフィックを保存出来るサイズのグラフィックを持つハンドル\n  useClientFlag: デフォルト TRUE
//dxlib int GetDrawScreenGraph(int x1, int y1, int x2, int y2, int grHandle, int useClientFlag = TRUE)
//comment; GetGraphSize; グラフィックのサイズを得る
//dxlib int GetGraphSize(int grHandle, int *sizeXBuf, int *sizeYBuf)
//comment; InitGraph; 読みこんだグラフィックデータをすべて削除する\n\n引数\n  logOutFlag: デフォルト FALSE
//dxlib int InitGraph(int logOutFlag = FALSE)
//comment; DeleteGraph; 指定のグラフィックをメモリ上から削除する
//dxlib int DeleteGraph(int grHandle)
//comment; SetDrawMode; 描画モードをセットする\n\n引数\n  drawMode: 描画モード(DX_DRAWMODE_NEAREST: 標準 or DX_DRAWMODE_BILINEAR)
//dxlib int SetDrawMode(int drawMode)
//comment; SetDrawBlendMode; 描画の際のブレンドモードをセット\n\n引数\n  blendMode: 描画ブレンドモード\n  pal: 描画ブレンドモードのパラメータ(0~255)
//dxlib int SetDrawBlendMode(int blendMode, int pal)
//comment; SetDrawBright; 描画輝度をセット
//dxlib int SetDrawBright(int redBright, int greenBright, int blueBright)
//comment; SetTransColor; グラフィックに設定する透過色をセットする
//dxlib int SetTransColor(int red, int green, int blue)
//comment; LoadBlendGraph; 画像ファイルからブレンド画像を読み込む
//dxlib int LoadBlendGraph(char *fileName)
//comment; DrawBlendGraph; ブレンド画像と通常画像を合成して描画する\n\n引数\n  x, y: 画像を描画する領域の左上端座標\n  grHandle: グラフィックハンドル\n  transFlag: 画像の透明度を有効にするかどうか(TRUE: 有効にする)\n  blendGraph: ブレンド画像ハンドル\n  borderParam: 境界位置(0~255)\n  borderRange: 境界幅(指定できる値は1, 64, 128, 255の４つ)
//dxlib int DrawBlendGraph(int x, int y, int grHandle, int transFlag, int blendGraph, int borderParam, int borderRange)

// 文字描画関係関数
//comment; DrawString; 文字列を描画\n\n引数\n  x, y: 文字列を描画する領域の左上の座標\n  str: 文字列\n  color: 文字列の色\n  edgeColor: デフォルト 0
//dxlib int DrawString(int x, int y, char *str, unsigned int color, unsigned int edgeColor = 0)
//comment; GetDrawStringWidth; 描画した時の文字列の幅(ドット単位)を得る\n\n引数\n str: 文字列\n  strLen: 文字列長\n  vericalFlag: デフォルト FALSE
//dxlib int GetDrawStringWidth(char *str, int strLen, int vericalFlag = FALSE)
//comment; SetFontSize; フォントのサイズをセットする
//dxlib int SetFontSize(int fontSize)
//comment; SetFontThickness; 描画する文字列の文字の太さをセットする\n\n引数\n  thickPal: 文字の太さ(0~9, デフォルト 6)
//dxlib int SetFontThickness(int thickPal)
//comment; ChangeFont; 文字列描画に使用するフォントを変更する\n\n引数\n  fontName: フォント名\n  charSet: デフォルト -1
//dxlib int ChangeFont(char *fontName, int charSet = -1)
//comment; ChangeFontType; 文字列描画に使用するフォントのタイプを変更する
//dxlib int ChangeFontType(int fontType)
//comment; CreateFontToHandle; 新しいフォントデータを作成\n\n引数\n  fontName: 作成するフォント名(NULLにするとデフォルトのフォント)\n  size: サイズ(デフォルト -1)\n  thick:  太さ(デフォルト -1)\n  fontType: フォントタイプ(デフォルト -1)\n  charSet: デフォルト -1\n  edgeSize: デフォルト -1\n  italic: デフォルト FALSE\n  handle: デフォルト -1
//dxlib int CreateFontToHandle(char *fontName = "", int size = -1, int thick = -1, int fontType = -1, int charSet = -1, int edgeSize = -1, int italic = FALSE, int handle = -1)
//comment; LoadFontDataToHandle; DXフォントデータファイルを読み込む\n\n引数\n  fileName: ファイル名\n  edgeSize: 成するフォントの縁の太さ(0を指定すると縁無し)
//dxlib int LoadFontDataToHandle(char *fileName, int edgeSize = 0)
//comment; DeleteFontToHandle; フォントデータを削除する
//dxlib int DeleteFontToHandle(int fontHandle)
//comment; SetFontCacheUsePremulAlphaFlag; 作成するフォントデータを『乗算済みα』用にするかどうかを設定する
//dxlib int SetFontCacheUsePremulAlphaFlag(int flag)
//comment; DrawStringToHandle; 指定のフォントデータで文字列を描画する\n\n引数\n  x, y: 文字列を描画する起点座標\n  str: 文字列\n  color: 文字の色\n  fontHandle: 描画に使用するフォントハンドル\n  edgeColor: 縁の色(デフォルト 0)\n  verticalFlag: デフォルト FALSE
//dxlib int DrawStringToHandle(int x, int y, char *str, unsigned int color, int fontHandle, int edgeColor = 0, int verticalFlag = FALSE)
//comment; GetDrawStringWidthToHandle; 指定のフォントデータで描画する文字列の幅(ドット単位)を得る\n\n引数\n  str: 文字列\n  strLen: 文字列長\n  fontHandle: フォントハンドル\n  verticalFlag: デフォルト FALSE
//dxlib int GetDrawStringWidthToHandle(char *str, int strLen, int fontHandle, int verticalFlag = FALSE)
//comment; GetFontStateToHandle; 指定のフォントデータの情報を得る
//dxlib int GetFontStateToHandle(char *fontName, int *size, int *thick, int fontHandle, int *fontType, int *charSet, int *edgeSize, int *italic)
//comment; InitFontToHandle; フォントデータを全て初期化する
//dxlib int InitFontToHandle()

// その他画面操作系関数
//comment; SetGraphMode; 画面モードの変更\n\n引数\n  sizeX, sizeY: 画面の解像度(デフォルト 640x480)\n  colorButNum: カラービット数(DXライブラリの標準色ビット数: 16)\n  refreshRate: デフォルト 60
//dxlib int SetGraphMode(int sizeX, int sizeY, int colorBitNum, int refreshRate)
//dxlib int SetFullScreenResolutionMode(int resolutionMode)
//dxlib int SetFullScreenScalingMode(int scalingMode, int fitScaling)
//dxlib int GetScreenState(int *sizeX, int *sizeY, int *colorBitDepth)
//dxlib int SetDrawArea(int x1, int y1, int x2, int y2)
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
//comment; LoadSoundMem; 音ファイルをメモリに読みこむ\n\n引数\n  fileName: ファイル名\n  bufferNum: デフォルト 3\n  unionHandle: デフォルト -1
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
//comment; GetNowCount; ミリ秒単位の精度を持つカウンタの現在値を得る\n\n引数\n  useRDTSCFlag: デフォルト FALSE
//dxlib int GetNowCount(int useRDTSCFlag)
//comment; GetNowHiPerformanceCount; GetNowCountの高精度バージョン\n\n引数\n  useRDTSCFlag: デフォルト FALSE
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
//ext_dxlib ClearDrawScreen

// Unsupported functions
// 公式のreadme.txtより
// > 基本的に「可変\n引数の関数」「void 型のポインタを\n引数で取る関数」「コールバック関数を\n引数で取る関数」が使えません。
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
//unsupported int FileRead_read(void *buffer, int readSize, int fileHandle)
//unsupported int LoadSoftImageToMem(void *fileImage, int fileImageSize)
//unsupported int LoadARGB8ColorSoftImageToMem(void *fileImage, int fileImageSize)
//unsupported int LoadXRGB8ColorSoftImageToMem(void *fileImage, int fileImageSize)
//unsupported int GetCharBytes(int charCodeFormat, void *str)
//unsupported int CreateGraphFromMem(void *memImage, int memImageSize)
//unsupported int ReCreateGraphFromMem(void *memImage, int memImageSize, int grHandle)
//unsupported int SetRestoreGraphCallback(void (* callback)())
//unsupported int LoadSoundMemByMemImage(void *fileImageBuffer, int imageSize)
