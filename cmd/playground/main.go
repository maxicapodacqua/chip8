package main

import (
	"fmt"
	"os"

	"github.com/maxicapodacqua/chip8/chip8"
)

func main() {
	vm := chip8.VirtualMachine{}

	// Loading rom
	f, err := os.ReadFile("../../roms/IBMLogo.ch8")
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

	vm.FetchOpCode()
	vm.DecodeOpCode()

	vm.FetchOpCode()
	vm.DecodeOpCode()

	vm.FetchOpCode()
	vm.DecodeOpCode()
}
