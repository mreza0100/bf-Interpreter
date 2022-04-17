package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
)

// brain fuck interpreter ^.^

type cell struct {
	value byte
}

type brainfuck struct {
	memory     [memorySize]cell
	loopStack  *LoopStack
	memPointer int

	instructions string
	runnerAt     int

	writter io.Writer
	reader  io.Reader
}

func (bf *brainfuck) print() {
	fmt.Fprintf(bf.writter, "pointer: %v, string_value: %c, byte_value: %v\n", bf.memPointer, bf.memory[bf.memPointer].value, bf.memory[bf.memPointer].value)
}

func (bf *brainfuck) moveForward() {
	bf.memPointer++
	if bf.memPointer >= memorySize {
		bf.memPointer = 0
	}
}

func (bf *brainfuck) moveBackward() {
	bf.memPointer--
	if bf.memPointer < 0 {
		bf.memPointer = memorySize - 1
	}
}

func (bf *brainfuck) increment() {
	if bf.memory[bf.memPointer].value > 255 {
		return
	}
	bf.memory[bf.memPointer].value++
}

func (bf *brainfuck) decrement() {
	if bf.memory[bf.memPointer].value <= 0 {
		return
	}
	bf.memory[bf.memPointer].value--
}

func (bf *brainfuck) read() {
	buf := make([]byte, 1)

	for {
		_, err := bf.reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		if bf.reader == os.Stdin && buf[0] != '\n' {
			bf.memory[bf.memPointer].value = buf[0]
			break
		}
	}
}

func (bf *brainfuck) isRunnerAtEdge() bool {
	return bf.runnerAt == len(bf.instructions)
}

func (bf *brainfuck) addInstruction(instruction byte) {
	if bf.isRunnerAtEdge() {
		bf.instructions += string(instruction)
	}
	bf.runnerAt++
}

func (bf *brainfuck) execute(instruction byte) {
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
		bf.loopStack.Push(bf.runnerAt - 1)
	case loopExit:
		if bf.loopStack.IsEmpty() {
			panic("loop stack is empty")
		}

		if bf.memory[bf.memPointer].value == 0 {
			bf.loopStack.Pop()
			break
		}

		loopStart := bf.loopStack.Pop()
		loopEnd := bf.runnerAt
		bf.runnerAt = loopStart

		for _, i := range bf.instructions[loopStart:loopEnd] {
			bf.execute(byte(i))
		}
	}
}

func (bf *brainfuck) entry(stream io.Reader) {
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

func New() *brainfuck {
	return &brainfuck{
		memory:       [memorySize]cell{},
		memPointer:   0,
		loopStack:    NewLoopStack(),
		instructions: "",
		runnerAt:     0,

		writter: os.Stdout,
		reader:  os.Stdin,
	}
}

const in = `+++++[-.]`

// nested loops
const nested = `
	++
	[
		>+++
		[
			-.
		]
		<-
	]
`

func main() {
	m := New()
	runtime.KeepAlive(in)
	runtime.KeepAlive(nested)

	m.entry(strings.NewReader(nested))
	fmt.Println("\n\n---\n", m.memory)
}
