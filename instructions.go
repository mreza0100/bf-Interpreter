package brainfuck

import (
	"fmt"
)

const (
	print        = '.'
	moveForward  = '>'
	moveBackward = '<'
	increment    = '+'
	decrement    = '-'
	read         = ','
	loopEnter    = '['
	loopExit     = ']'
)

func isADefaultInstruction(instruction byte) bool {
	switch instruction {
	case print, moveForward, moveBackward, increment, decrement, read, loopEnter, loopExit:
		return true
	}

	return false
}

// abstraction to give to the executive
// why it's not a interface? because in that way I had to expose the internal (like Brainfuck.Print)
type InstructionCtl struct {
	GetMemory          func() []byte
	GetPointerPosition func() int

	Print        func()
	MoveForward  func()
	MoveBackward func()
	Increment    func()
	Decrement    func()
	Read         func()
}

type InstructionExecutive func(InstructionCtl)

type CustomInstruction struct {
	commands map[byte]InstructionExecutive
	InstructionCtl
}

func newCustomInstruction(bf *Brainfuck) *CustomInstruction {
	return &CustomInstruction{
		commands: make(map[byte]InstructionExecutive),
		InstructionCtl: InstructionCtl{
			GetMemory:          func() []byte { return bf.memory.values },
			GetPointerPosition: func() int { return bf.memPointer },

			Read:         bf.read,
			MoveForward:  bf.moveForward,
			MoveBackward: bf.moveBackward,
			Increment:    bf.increment,
			Decrement:    bf.decrement,
			Print:        bf.print,
		},
	}
}

func (ci *CustomInstruction) get(instruction byte) (InstructionExecutive, bool) {
	executive, exist := ci.commands[instruction]
	return executive, exist
}

func (ci *CustomInstruction) add(instruction byte, executive InstructionExecutive) error {
	if _, exists := ci.get(instruction); exists {
		return fmt.Errorf("instruction %c already exists", instruction)
	}

	if isADefaultInstruction(instruction) {
		return fmt.Errorf("instruction %c is a default instruction", instruction)
	}

	ci.commands[instruction] = executive
	return nil
}

func (ci *CustomInstruction) remove(instruction byte) error {
	if _, exists := ci.get(instruction); !exists {
		return fmt.Errorf("instruction %c does not exist", instruction)
	}

	delete(ci.commands, instruction)
	return nil
}
