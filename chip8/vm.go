package chip8

import "fmt"

type VirtualMachine struct {
	// uint16 because opcode is two bytes long
	Opcode uint16
	// 4k memory in total
	Memory [4096]byte
	// CPU registers
	V [16]byte
	// index register
	I uint16
	// program counter
	Pc uint16

	// display is 64 x 32 (2048 pixels)
	Graphics   [64 * 32]byte
	DelayTimer byte
	SoundTimer byte

	Stack [16]uint16
	// stack pointer
	Sp uint16

	KeyState [16]byte
}

func (vm *VirtualMachine) FetchOpCode() {
	// merge pc and pc + 1
	// the memory is stored as an array of bytes, and pc is two bytes long

	// Grab pc and shift if by one byte (8 bits)
	a := uint16(vm.Memory[vm.Pc]) << 8
	// Next memory part
	b := uint16(vm.Memory[vm.Pc+1])
	// Join them
	vm.Opcode = a | b
}

func (vm *VirtualMachine) DecodeOpCode() {
	// Get the first part of the opcode
	opcodeHighNibble := vm.Opcode & 0xF000
	fmt.Printf("instruction %x pc=%x\n", vm.Opcode, vm.Pc)
	switch opcodeHighNibble {
	// Clear the display.
	case 0x00E0:
		vm.Pc += 2
	// Jump 0x1NNN
	// The interpreter sets the program counter to nnn.
	case 0x1000:
		newPc := vm.Opcode & 0x0FFF
		vm.Pc = newPc
	// 6XNN (set register VX)
	// The interpreter puts the value kk into register Vx.
	case 0x6000:
		kk := vm.Opcode & 0x00FF
		x := vm.Opcode & 0x0F00
		vm.V[x] = byte(kk)
		vm.Pc += 2
	// 7XNN (add value to register VX)
	// Adds the value kk to the value of register Vx, then stores the result in Vx.
	case 0x7000:
		kk := vm.Opcode & 0x00FF
		x := vm.Opcode & 0x0F00
		a := vm.V[x]
		vm.V[x] = byte(kk) + a
		vm.Pc += 2
	// ANNN (set index register I)
	case 0xA000:
		vm.Pc += 2
	// DXYN (display/draw)
	case 0xD000:
		fmt.Printf("Not implemented!")

	}
	fmt.Printf("next pc=%x", vm.Pc)
}
