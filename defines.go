package dxlib

// Boolean
const (
	TRUE  = 1
	FALSE = 0
)

// 描画先画面指定用定義
const (
	DX_SCREEN_FRONT     = 0xfffffffc
	DX_SCREEN_BACK      = 0xfffffffe
	DX_SCREEN_WORK      = 0xfffffffd
	DX_SCREEN_TEMPFRONT = 0xfffffff0
	DX_SCREEN_OTHER     = 0xfffffffa

	DX_NONE_GRAPH = 0xfffffffb // グラフィックなしハンドル
)
