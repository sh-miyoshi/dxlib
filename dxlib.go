package dxlib

//go:generate go run cmd/generate.go -i dxlib.go -o dxlib_gen.go
//go:generate gofmt -w dxlib_gen.go

//dxlib int DxLib_Init()
//dxlib int DxLib_End()
//dxlib int ProcessMessage()

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
