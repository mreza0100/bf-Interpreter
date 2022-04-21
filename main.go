package brainfuck

import (
	"fmt"
	"io"
	"os"
)

// brainfuck interpreter ^.^

type Brainfuck struct {
	memory     *memory
	loopStack  *loopStack
	memPointer int
	errors     *errorCheck

	instructions    string
	rawInstructions string
	runnerAt        int
	Verbos          bool

	customCommands *CustomCommands
	Writter        io.Writer
	Reader         io.Reader
}

func (bf *Brainfuck) print() {
	if bf.Verbos {
		fmt.Fprintf(bf.Writter, "pointer: %v, byte_value: %v, value: %c\n", bf.memPointer, bf.memory.values[bf.memPointer], bf.memory.values[bf.memPointer])
		return
	}

	fmt.Fprintf(bf.Writter, "%c", bf.memory.values[bf.memPointer])
}

func (bf *Brainfuck) moveForward() {
	bf.memPointer++
	if !bf.memory.isOut(bf.memPointer) {
		return
	}

	if bf.memory.isStatic {
		bf.memPointer = 0
	} else {
		bf.memory.moreCap()
	}
}

func (bf *Brainfuck) moveBackward() {
	bf.memPointer--
	if bf.memory.isOut(bf.memPointer) {
		bf.memPointer = bf.memory.len() - 1
	}
}

func (bf *Brainfuck) increment() {
	if bf.memory.values[bf.memPointer] > 255 {
		return
	}
	bf.memory.values[bf.memPointer]++
}

func (bf *Brainfuck) decrement() {
	if bf.memory.values[bf.memPointer] <= 0 {
		return
	}
	bf.memory.values[bf.memPointer]--
}

func (bf *Brainfuck) read() {
	buf := make([]byte, 1)

	for {
		_, err := bf.Reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		// handeling dump input from stdin
		// we can use Fscanln too
		if bf.Reader == os.Stdin && buf[0] != '\n' {
			bf.memory.values[bf.memPointer] = buf[0]
			break
		}
	}
}

func (bf *Brainfuck) loopEnter() {
	bf.loopStack.push(bf.runnerAt - 1)
}

func (bf *Brainfuck) loopExit() {
	if err := bf.errors.noOpenedLoopCheck(bf); err != nil {
		panic(err)
	}

	if bf.memory.values[bf.memPointer] == 0 {
		bf.loopStack.pop()
		return
	}

	loopStart := bf.loopStack.pop()
	loopEnd := bf.runnerAt
	bf.runnerAt = loopStart
	insideLoop := bf.instructions[loopStart:loopEnd]

	if err := bf.errors.emptyLoopCheck(insideLoop, bf); err != nil {
		panic(err)
	}
	for _, i := range insideLoop {
		bf.execute(byte(i))
	}
}

func (bf *Brainfuck) AddCustomCommand(command byte, ie CommandExecutor) error {
	return bf.customCommands.add(command, ie)
}

func (bf *Brainfuck) RemoveCustomCommand(command byte) error {
	return bf.customCommands.remove(command)
}

func (bf *Brainfuck) isRunnerAtEdge() bool {
	return bf.runnerAt == len(bf.instructions)
}

func (bf *Brainfuck) addInstruction(instruction byte) {
	cleanInstruction := trim(instruction)

	// do not add repeated instructions from the loop
	if bf.isRunnerAtEdge() {
		bf.instructions += cleanInstruction
		bf.rawInstructions += string(instruction)
	}

	// do not add dump instructions
	isClean := cleanInstruction != ""
	if isClean {
		bf.runnerAt++
	}
}

func (bf *Brainfuck) execute(instruction byte) {
	bf.addInstruction(instruction)

	switch instruction {
	case moveForward:
		bf.moveForward()
	case moveBackward:
		bf.moveBackward()
	case increment:
		bf.increment()
	case decrement:
		bf.decrement()
	case print:
		bf.print()
	case read:
		bf.read()
	case loopEnter:
		bf.loopEnter()
	case loopExit:
		bf.loopExit()
	default:
		if commandExecutor, exist := bf.customCommands.get(instruction); exist {
			commandExecutor(bf.customCommands.CommandsCtl)
		}
	}
}

func (bf *Brainfuck) Run(stream io.Reader) {
	buf := make([]byte, 1)

	for {
		_, err := stream.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}
		bf.execute(buf[0])
	}
}

type NewOptions struct {
	// initial size of memory
	MemorySize int
	// is memory static. if you move forward and there will be no extra space, memory will get 2 times more
	// if it be static, memory will not grow and if you move forward and there will be no extra space, you will just jump back to memory[0]
	StaticMemory bool
	// verbos logging; example: pointer: 2, string_value: 97, byte_value: a
	Verbos bool
	// custom writer for loggs. default to stdout
	Writter io.Writer
	// custom reader for input. default to stdin
	Reader io.Reader
}

func New(options *NewOptions) *Brainfuck {
	if options.MemorySize == 0 {
		options.MemorySize = 300
	}

	if options.Writter == nil {
		options.Writter = os.Stdout
	}

	if options.Reader == nil {
		options.Reader = os.Stdin
	}

	bf := &Brainfuck{
		memory:          newMemory(options.MemorySize, options.StaticMemory),
		memPointer:      0,
		loopStack:       newLoopStack(),
		instructions:    "",
		rawInstructions: "",
		runnerAt:        0,
		Verbos:          options.Verbos,
		errors:          newErrorCheck(),

		Writter: options.Writter,
		Reader:  options.Reader,
	}
	bf.customCommands = newCustomCommand(bf)

	return bf
}
