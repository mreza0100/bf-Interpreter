package brainfuck

import (
	"fmt"
)

type errorCheck struct{}

func newErrorCheck() *errorCheck {
	return new(errorCheck)
}

func (e *errorCheck) makeError(bf *Brainfuck, msg string) error {
	var (
		start = bf.runnerAt - 5
		end   = bf.runnerAt + 5
	)

	if start < 0 {
		start = 0
	}
	if end > len(bf.instructions) {
		end = len(bf.instructions)
	}

	return fmt.Errorf("\n----\n%v. Error at: %v.\n position: %v\n----\n", msg, bf.runnerAt, bf.instructions[start:end])
}

func (e *errorCheck) noOpenedLoopCheck(bf *Brainfuck) error {
	if !bf.loopStack.isEmpty() {
		return nil
	}

	return e.makeError(bf, "no opened loop")
}

func (e *errorCheck) emptyLoopCheck(insideLoop string, bf *Brainfuck) error {
	if insideLoop != "[]" {
		return nil
	}

	return e.makeError(bf, "empty loop")
}
