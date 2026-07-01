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

// FontSet found in http://www.multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter
var fontSet = [80]byte{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xe0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0x80, // C
	0xF0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

func (vm *VirtualMachine) LoadFontset() {
	for i := range 80 {
		vm.Memory[i] = fontSet[i]
	}

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
	fmt.Printf("instruction %04x pc=%x\n", vm.Opcode, vm.Pc)
	switch opcodeHighNibble {
	// Clear the display.
	case 0x0000: // TODO: add case for 00e0, and other codes that start with 0
		// case 0x00E0:
		vm.Pc += 2
	// Jump 0x1NNN
	// The interpreter sets the program counter to nnn.
	case 0x1000:
		newPc := vm.Opcode & 0x0FFF
		vm.Pc = newPc
		// panic("asdasd")
	// 6XNN (set register VX)
	// The interpreter puts the value kk into register Vx.
	case 0x6000:
		kk := vm.Opcode & 0x00FF
		x := (vm.Opcode & 0x0F00) >> 8
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
		nnn := vm.Opcode & 0x0FFF
		vm.I = nnn
		vm.Pc += 2
	// DXYN (display/draw)
	// Dxyn - DRW Vx, Vy, nibble
	// Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.
	case 0xD000:
		vm.draw()
		// fmt.Printf("Not implemented!\n")
		vm.Pc += 2

	}
	fmt.Printf("next pc=%x\n", vm.Pc)
}

func (vm *VirtualMachine) draw() {
	// get vx and make it a value, shift by 8, because is the third nibble
	// basically >> 8, moves the value two nibbles to the right
	// 0x0A00 >> 8 = 0x000A
	vx := (vm.Opcode & 0x0F00) >> 8
	// get vy and make it a value, shift by 4, because is the second nibble
	vy := (vm.Opcode & 0x00F0) >> 4
	x := uint16(vm.V[vx])
	y := uint16(vm.V[vy])
	height := vm.Opcode & 0x00F
	var pixel byte
	// Set collision flag to 0
	vm.V[0xF] = 0

	for yline := range height {
		pixel = vm.Memory[vm.I+yline]
		for xline := range uint16(8) {
			// 0x80 is 10000000
			// in this loop, xline will go from 0 to 8
			// this bit shift will create a bitmask with different values per loop
			// xline=0, bitmask = 10000000
			// xline=1, bitmask = 01000000
			// xline=2, bitmask = 00100000
			bitmaskForXLine := 0x80 >> xline
			// By & on the pixel byte, you get the single big for that position
			//
			// the & with the bitmask will check for a single bit
			// if pixel is 01100000 and bitmask is 01000000
			// 01100000 & 01000000 = 00100000
			//  ^          ^
			isPixelSet := pixel&(byte(bitmaskForXLine)) != 0
			if isPixelSet {
				graphIndex := x + xline + ((y + yline) * 64)
				// Set vF register
				if vm.Graphics[graphIndex] == 1 {
					vm.V[0xF] = 1
				}
				// Set value for graphic (always XOR)
				vm.Graphics[graphIndex] ^= 1
			}
		}
	}
}

func (vm *VirtualMachine) Render() {

}
