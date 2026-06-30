package main

import (
	"github.com/maxicapodacqua/chip8/chip8"
)

func main() {
	vm := chip8.VirtualMachine{}
	vm.Pc = 0x200
	vm.I = 0
	vm.Sp = 0
	vm.Memory[vm.Pc] = 0xA2
	vm.Memory[vm.Pc+1] = 0xF0

	vm.FetchOpCode()

}
