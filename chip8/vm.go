package chip8

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

func (vm *VirtualMachine) DecodeOpCode(opcode uint16) {

}
