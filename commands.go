package brainfuck

import (
	"fmt"
)

type cmd byte

const (
	print                     cmd = '.'
	read                      cmd = ','
	moveForward, moveBackward cmd = '>', '<'
	increment, decrement      cmd = '+', '-'
	loopEnter, loopExit       cmd = '[', ']'
)

func (c cmd) isADefaultCommand() bool {
	switch c {
	case print, moveForward, moveBackward, increment, decrement, read, loopEnter, loopExit:
		return true
	}

	return false
}

// abstraction to give to the Executor
// why it's not a interface? because in that way I had to expose the internal (like Brainfuck.Print)
type CommandDriver struct {
	GetMemory          func() []byte
	GetPointerPosition func() int

	Print        func()
	MoveForward  func()
	MoveBackward func()
	Increment    func()
	Decrement    func()
	Read         func()
}

type CustomCMDExecutor func(CommandDriver)

type customCommands struct {
	commands map[byte]CustomCMDExecutor
	CommandDriver
}

func newCustomCommand(bf *Brainfuck) *customCommands {
	return &customCommands{
		commands: make(map[byte]CustomCMDExecutor),
		CommandDriver: CommandDriver{
			GetMemory:          func() []byte { return bf.memory.GetMemory() },
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

func (cc *customCommands) get(command byte) (CustomCMDExecutor, bool) {
	executive, exist := cc.commands[command]
	return executive, exist
}

func (cc *customCommands) add(command byte, executive CustomCMDExecutor) error {
	if _, exists := cc.get(command); exists {
		return fmt.Errorf("command %c already exists", command)
	}

	if cmd(command).isADefaultCommand() {
		return fmt.Errorf("command %c is a default command", command)
	}

	cc.commands[command] = executive
	return nil
}

func (cc *customCommands) remove(command byte) error {
	if _, exists := cc.get(command); !exists {
		return fmt.Errorf("command %c does not exist", command)
	}

	delete(cc.commands, command)
	return nil
}
