package errs

import (
	"fmt"
)

type errorCheck struct{}

type ErrorCheck interface {
	OpenLoopCheck(opns *LoopCheckReq) error
	ExitLoopCheck(lc *LoopCheckReq, insideLoop string) error
}

func New() ErrorCheck {
	return new(errorCheck)
}

type LoopCheckReq struct {
	IsStackEmpty bool
	RunnerAt     int
	Instructions string
}

func (e *errorCheck) makeError(opns *LoopCheckReq, msg string) error {
	var (
		start = opns.RunnerAt - 5
		end   = opns.RunnerAt + 5
	)

	if start < 0 {
		start = 0
	}
	if end > len(opns.Instructions) {
		end = len(opns.Instructions)
	}

	return fmt.Errorf("\n----\n%v. Error at: %v.\n position: %v\n----\n", msg, opns.RunnerAt, opns.Instructions[start:end])
}

func (e *errorCheck) OpenLoopCheck(opns *LoopCheckReq) error {
	if !opns.IsStackEmpty {
		return nil
	}

	return e.makeError(opns, "no opened loop")
}

func (e *errorCheck) ExitLoopCheck(lc *LoopCheckReq, insideLoop string) error {
	if insideLoop != "[]" {
		return nil
	}

	return e.makeError(lc, "empty loop")
}
