package main

import (
	"fmt"
	"os"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/imdraw"
	"github.com/maxicapodacqua/chip8/chip8"
	"golang.org/x/image/colornames"
)

var vm = chip8.VirtualMachine{}

func run() {
	cfg := opengl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := opengl.NewWindow(cfg)
	imd := imdraw.New(nil)

	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Clear(colornames.Black)
		vm.FetchOpCode()
		vm.DecodeOpCode()
		render(imd)
		imd.Draw(win)
		win.Update()
	}
}

func render(imd *imdraw.IMDraw) {
	// 64x32

	width := float64(1024 / 64)
	height := float64(768 / 32)
	for i := range 64 {
		for j := range 32 {
			pos := (31-j)*64 + i
			if vm.Graphics[pos] == 0 {
				continue
			}
			// draw pixel
			imd.Push(pixel.V(width*float64(i), height*float64(j)))
			imd.Push(pixel.V(width*float64(i)+width, height*float64(j)+height))
			imd.Rectangle(0)
		}
	}
}

func main() {

	// Loading rom
	f, err := os.ReadFile("../../roms/chip8_logo.ch8")
	// f, err := os.ReadFile("../../roms/IBMLogo.ch8")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", f)

	for i := range len(f) {
		vm.Memory[0x200+i] = f[i]
	}
	// end of loading rom

	vm.Pc = 0x200
	vm.I = 0
	vm.Sp = 0
	// vm.Memory[vm.Pc] = 0xA2
	// vm.Memory[vm.Pc+1] = 0xF0
	// vm.Memory[vm.Pc] = 0x12
	// vm.Memory[vm.Pc+1] = 0xF0

	vm.LoadFontset()
	// for {
	// 	vm.FetchOpCode()
	// 	vm.DecodeOpCode()

	// }

	opengl.Run(run)
	// vm.FetchOpCode()
	// vm.DecodeOpCode()

	// vm.FetchOpCode()
	// vm.DecodeOpCode()
}
