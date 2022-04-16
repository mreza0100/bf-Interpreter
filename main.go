package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// brain fuck interpreter ^.^

// idea:
// 1. read from stdin
// 2. check if it is a valid brainfuck program

const (
	moveForward  = '>'
	moveBackward = '<'
	increment    = '+'
	decrement    = '-'
	print        = '.'
	read         = ','
	loopStart    = '['
	loopEnd      = ']'
)

const (
	memorySize = 30
)

type cell struct {
	value byte
}

type brainfuck struct {
	memory  [memorySize]cell
	pointer int
	charAt  int

	writter io.Writer
	reader  io.Reader
}

func (d *brainfuck) print() {
	fmt.Printf("At: %v, stringValue: %c, raw: %v\n", d.pointer, d.memory[d.pointer].value, d.memory[d.pointer].value)
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

func (d *brainfuck) read(commands string) {
	for d.charAt = 0; d.charAt < len(commands); d.charAt++ {
		switch commands[d.charAt] {
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
			temp := make([]byte, 1)
			fmt.Scanf("%c", &temp[0])
			d.memory[d.pointer].value = temp[0]
		case loopStart:
			insideLoop, endLoopIdx := getInsideLoop(commands, d.charAt)
			d.charAt = endLoopIdx
			fmt.Println("start loop\n", insideLoop)

			loopDependsOn := d.pointer
			for d.memory[loopDependsOn].value != 0 {
				time.Sleep(time.Second / 2)
				d.read(insideLoop)
			}
		}
	}
}

func New() *brainfuck {
	return &brainfuck{
		memory:  [memorySize]cell{},
		pointer: 0,
		charAt:  0,

		writter: os.Stdout,
		reader:  os.Stdin,
	}
}

const in = `
++
[
	>+++
	[
		-.
	]
	<----
]
`

// nested loops
// ++
// [
// 	>+++
// 		[-.]
// 	<-
// ]
// +++

func main() {
	m := New()

	m.read(in)
	fmt.Println("\n\n---\n", m.memory)
}
