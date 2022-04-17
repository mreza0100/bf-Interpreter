package brainfuck

import (
	"fmt"
	"io"
	"os"
)

// brain fuck interpreter ^.^

type Brainfuck struct {
	memory     [memorySize]byte
	loopStack  *loopStack
	memPointer int

	instructions    string
	rawInstructions string
	runnerAt        int

	Writter io.Writer
	Reader  io.Reader
}

func (bf *Brainfuck) print() {
	fmt.Fprintf(bf.Writter, "pointer: %v, string_value: %v, byte_value: %c\n", bf.memPointer, bf.memory[bf.memPointer], bf.memory[bf.memPointer])
}

func (bf *Brainfuck) moveForward() {
	bf.memPointer++
	if bf.memPointer >= memorySize {
		bf.memPointer = 0
	}
}

func (bf *Brainfuck) moveBackward() {
	bf.memPointer--
	if bf.memPointer < 0 {
		bf.memPointer = memorySize - 1
	}
}

func (bf *Brainfuck) increment() {
	if bf.memory[bf.memPointer] > 255 {
		return
	}
	bf.memory[bf.memPointer]++
}

func (bf *Brainfuck) decrement() {
	if bf.memory[bf.memPointer] <= 0 {
		return
	}
	bf.memory[bf.memPointer]--
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
		if bf.Reader == os.Stdin && buf[0] != '\n' {
			bf.memory[bf.memPointer] = buf[0]
			break
		}
	}
}

func (bf *Brainfuck) loopEnter() {
	bf.loopStack.push(bf.runnerAt - 1)
}

func (bf *Brainfuck) loopExit() {
	if bf.loopStack.isEmpty() {
		panic("No loop to exit")
	}
	if bf.memory[bf.memPointer] == 0 {
		bf.loopStack.pop()
		return
	}

	loopStart := bf.loopStack.pop()
	loopEnd := bf.runnerAt
	bf.runnerAt = loopStart

	for _, i := range bf.instructions[loopStart:loopEnd] {
		bf.execute(byte(i))
	}
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
	}
}

func (bf *Brainfuck) Entry(stream io.Reader) {
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

func New() *Brainfuck {
	return &Brainfuck{
		memory:          [memorySize]byte{},
		memPointer:      0,
		loopStack:       newLoopStack(),
		instructions:    "",
		rawInstructions: "",
		runnerAt:        0,

		Writter: os.Stdout,
		Reader:  os.Stdin,
	}
}
