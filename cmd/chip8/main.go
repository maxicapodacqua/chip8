package main

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/imdraw"
)

func run() {
	cfg := opengl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := opengl.NewWindow(cfg)
	imd := imdraw.New(nil)

	// 64x32
	pixelW := float64(1024 / 64)
	pixelH := float64(768 / 32)
	checkersFlag := true
	for y := range 64 {
		for x := range 32 {
			if checkersFlag {
				// draw pixel
				imd.Push(pixel.V(float64(y)*pixelW, float64(x)*pixelH))
				imd.Push(pixel.V((float64(y)*pixelW)+pixelW, (float64(x)*pixelH)+pixelH))
				imd.Rectangle(0)
			}
			checkersFlag = !checkersFlag
		}
		checkersFlag = !checkersFlag
	}
	// imd.Push(pixel.V(200, 100))
	// imd.Push(pixel.V(400, 300))
	// imd.Rectangle(0)

	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	opengl.Run(run)
}
