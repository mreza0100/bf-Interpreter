package brainfuck

import (
	"fmt"
	"io"
	"os"

	"github.com/mreza0100/brainfuck/internals/errs"
	"github.com/mreza0100/brainfuck/internals/loopstack"
	"github.com/mreza0100/brainfuck/internals/memory"
)

// brainfuck interpreter ^.^

type Brainfuck struct {
	memory     memory.Memory
	loopStack  loopstack.Loopstack
	memPointer int
	errors     errs.ErrorCheck

	instructions string
	runnerAt     int
	Verbos       bool

	customCommands *customCommands
	Writter        io.Writer
	Reader         io.Reader
}

func (bf *Brainfuck) print() {
	if bf.Verbos {
		fmt.Fprintf(bf.Writter, "pointer: %v, byte_value: %v, value: %c\n", bf.memPointer, bf.memory.GetMemory()[bf.memPointer], bf.memory.GetMemory()[bf.memPointer])
		return
	}

	fmt.Fprintf(bf.Writter, "%c", bf.memory.GetMemory()[bf.memPointer])
}

func (bf *Brainfuck) moveForward() {
	bf.memPointer++
	if !bf.memory.IsOut(bf.memPointer) {
		return
	}

	if bf.memory.IsStatic() {
		bf.memPointer = 0
	} else {
		bf.memory.IncreaseCap()
	}
}

func (bf *Brainfuck) moveBackward() {
	bf.memPointer--
	if bf.memory.IsOut(bf.memPointer) {
		bf.memPointer = bf.memory.Len() - 1
	}
}

func (bf *Brainfuck) increment() {
	if bf.memory.GetMemory()[bf.memPointer] > 255 {
		return
	}
	bf.memory.GetMemory()[bf.memPointer]++
}

func (bf *Brainfuck) decrement() {
	if bf.memory.GetMemory()[bf.memPointer] <= 0 {
		return
	}
	bf.memory.GetMemory()[bf.memPointer]--
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
			bf.memory.GetMemory()[bf.memPointer] = buf[0]
			break
		}
	}
}

func (bf *Brainfuck) loopEnter() {
	bf.loopStack.Push(bf.runnerAt - 1)
}

func (bf *Brainfuck) loopExit() {
	err := bf.errors.OpenLoopCheck(&errs.LoopCheckReq{
		IsStackEmpty: bf.loopStack.IsEmpty(),
		RunnerAt:     bf.runnerAt,
		Instructions: bf.instructions,
	})
	if err != nil {
		panic(err)
	}

	if bf.memory.GetMemory()[bf.memPointer] == 0 {
		bf.loopStack.Pop()
		return
	}

	loopStart := bf.loopStack.Pop()
	loopEnd := bf.runnerAt
	bf.runnerAt = loopStart
	insideLoop := bf.instructions[loopStart:loopEnd]

	err = bf.errors.ExitLoopCheck(&errs.LoopCheckReq{
		IsStackEmpty: bf.loopStack.IsEmpty(),
		RunnerAt:     bf.runnerAt,
		Instructions: bf.instructions,
	}, insideLoop)
	if err != nil {
		panic(err)
	}
	for _, i := range insideLoop {
		bf.execute(byte(i))
	}
}

func (bf *Brainfuck) AddCustomCommand(command byte, ie CustomCMDExecutor) error {
	return bf.customCommands.add(command, ie)
}

func (bf *Brainfuck) RemoveCustomCommand(command byte) error {
	return bf.customCommands.remove(command)
}

func (bf *Brainfuck) isRunnerAtEdge() bool {
	return bf.runnerAt == len(bf.instructions)
}

func (bf *Brainfuck) addInstruction(instruction byte) {
	// do not add repeated instructions from the loop
	if bf.isRunnerAtEdge() {
		bf.instructions += string(instruction)
	}

	// do not add dump instructions
	bf.runnerAt++
}

func (bf *Brainfuck) execute(instruction byte) {
	bf.addInstruction(instruction)

	switch cmd(instruction) {
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
			commandExecutor(bf.customCommands.CommandDriver)
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
		memory:       nil,
		memPointer:   0,
		loopStack:    nil,
		instructions: "",
		runnerAt:     0,
		errors:       nil,

		Verbos:  options.Verbos,
		Writter: options.Writter,
		Reader:  options.Reader,
	}

	bf.memory = memory.New(options.MemorySize, options.StaticMemory)
	bf.loopStack = loopstack.New()
	bf.errors = errs.New()
	bf.customCommands = newCustomCommand(bf)

	return bf
}
