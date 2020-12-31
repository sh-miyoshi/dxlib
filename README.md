# go-dxlib

go言語で[DXライブラリ](https://dxlib.xsrv.jp/)を使用するためのパッケージです。

## インストール方法

dxlibのインストールは通常のgoのパッケージと同様に`go get`コマンドで取得できます。

```bat
go get -u "github.com/sh-miyoshi/dxlib"
```

## 使い方

### 1. DxLib.dllを取得

- 方法1: DXライブラリの公式サイトの[ダウンロードページ](https://dxlib.xsrv.jp/dxdload.html)からVisual C#用パッケージをダウンロードし、その中のDxLib_x64.dll(32bit環境ならDxLib.dll)を取得
- 方法2: このリポジトリの[example/DxLib.dll](https://github.com/sh-miyoshi/dxlib/raw/master/example/DxLib.dll)を取得

### 2. コードを記述

```go
package main

import (
    "github.com/sh-miyoshi/dxlib"
)

func main() {
    dxlib.Init("DxLib.dll")

    dxlib.ChangeWindowMode(dxlib.TRUE)

    dxlib.DxLib_Init()
    dxlib.SetDrawScreen(dxlib.DX_SCREEN_BACK)

    for dxlib.ScreenFlip() == 0 && dxlib.ProcessMessage() == 0 && dxlib.ClearDrawScreen() == 0 {
        dxlib.DrawString(10, 10, "Hello, world", dxlib.GetColor(255, 255, 255), 0)
    }

    dxlib.DxLib_End()
}
```

### 3. 実行

※DXライブラリはWindows向けのアプリケーションなのでWindows環境で実施してください。

```bat
go build -o project.exe
.\project.exe
```
