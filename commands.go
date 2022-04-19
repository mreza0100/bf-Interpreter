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

func isADefaultCommand(command byte) bool {
	switch command {
	case print, moveForward, moveBackward, increment, decrement, read, loopEnter, loopExit:
		return true
	}

	return false
}

// abstraction to give to the Executor
// why it's not a interface? because in that way I had to expose the internal (like Brainfuck.Print)
type CommandsCtl struct {
	GetMemory          func() []byte
	GetPointerPosition func() int

	Print        func()
	MoveForward  func()
	MoveBackward func()
	Increment    func()
	Decrement    func()
	Read         func()
}

type CommandExecutor func(CommandsCtl)

type CustomCommands struct {
	commands map[byte]CommandExecutor
	CommandsCtl
}

func newCustomCommand(bf *Brainfuck) *CustomCommands {
	return &CustomCommands{
		commands: make(map[byte]CommandExecutor),
		CommandsCtl: CommandsCtl{
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

func (ci *CustomCommands) get(command byte) (CommandExecutor, bool) {
	executive, exist := ci.commands[command]
	return executive, exist
}

func (ci *CustomCommands) add(command byte, executive CommandExecutor) error {
	if _, exists := ci.get(command); exists {
		return fmt.Errorf("command %c already exists", command)
	}

	if isADefaultCommand(command) {
		return fmt.Errorf("command %c is a default command", command)
	}

	ci.commands[command] = executive
	return nil
}

func (ci *CustomCommands) remove(command byte) error {
	if _, exists := ci.get(command); !exists {
		return fmt.Errorf("command %c does not exist", command)
	}

	delete(ci.commands, command)
	return nil
}
