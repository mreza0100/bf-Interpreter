package main

import (
	"fmt"
	"io"
	"os"
)

// brain fuck interpreter ^.^

type cell struct {
	value byte
}

type brainfuck struct {
	memory          [memorySize]cell
	pointer         int
	runnerIdx       int
	rawInstructions string

	writter io.Writer
	reader  io.Reader
}

func (d *brainfuck) print() {
	fmt.Fprintf(d.writter, "pointer: %v, string_value: %c, byte_value: %v\n", d.pointer, d.memory[d.pointer].value, d.memory[d.pointer].value)
}

func (d *brainfuck) moveForward() {
	d.pointer++
	if d.pointer >= memorySize {
		d.pointer = 0
	}
}

func (d *brainfuck) moveBackward() {
	d.pointer--
	if d.pointer < 0 {
		d.pointer = memorySize - 1
	}
}

func (d *brainfuck) increment() {
	if d.memory[d.pointer].value > 255 {
		return
	}
	d.memory[d.pointer].value++
}

func (d *brainfuck) decrement() {
	if d.memory[d.pointer].value <= 0 {
		return
	}
	d.memory[d.pointer].value--
}

func (d *brainfuck) read() {
	fmt.Fscanf(d.reader, "%c", &d.memory[d.pointer].value)
}

func (d *brainfuck) loop(instructions string) {
	insideLoop := getInsideLoop(instructions, d.runnerIdx)

	for d.memory[d.pointer].value != 0 {
		d.execute(insideLoop)
	}
}

func (d *brainfuck) execute(instructions string) {
	for d.runnerIdx = 0; d.runnerIdx < len(instructions); d.runnerIdx++ {
		switch instructions[d.runnerIdx] {
		case moveForward:
			d.moveForward()
		case moveBackward:
			d.moveBackward()
		case increment:
			d.increment()
		case decrement:
			d.decrement()
		case print:
			d.print()
		case read:
			d.read()
		case loopEnter:
			d.loop(instructions)
		}
	}
}

func (d *brainfuck) entry(instructions string) {
	d.rawInstructions += instructions
	d.execute(trim(instructions))
}

func New() *brainfuck {
	return &brainfuck{
		memory:    [memorySize]cell{},
		pointer:   0,
		runnerIdx: 0,

		writter: os.Stdout,
		reader:  os.Stdin,
	}
}

const in = `
,.
`

// nested loops
// ++
// [
// 	>+++
// 	[
// 		-.
// 	]
// 	<-
// ]

func main() {
	m := New()

	m.entry(in)
	fmt.Println("\n\n---\n", m.memory)
}
