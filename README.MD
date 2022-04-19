# Brainfuck Interpreter


``` bash
go get -u github.com/mreza0100/brainfuck
```

## Usage
``` go
myCustomWritter := new(bytes.Buffer)
bf := New(&NewOptions{
      // inline documention for each of them is accessible in the source code
      MemorySize:     100,
      IsMemoryStatic: true,
      Verbos:         false,
      Writter:        myCustomWritter,
      Reader:         nil, // will fallback to os.Stdin
})
instructions := "++++[-.>+<]"
bf.Run(strings.NewReader(instructions))
// continue to run more instructions
instructions = ">."
bf.Run(strings.NewReader(instructions))
```
## Custom Instructions
``` go
bf := New(&NewOptions{})
err := bf.AddCustomInstruction('@', func(icw brainfuck.InstructionCtl) {
      // you have tools to do almost anything
		icw.GetMemory()
		icw.GetPointerPosition()

		icw.Print()
		icw.MoveForward()
		icw.MoveBackward()
		icw.Increment()
		icw.Decrement()
		icw.Read()
})
if err != nil {
      // error in case instruction is already defined or it's a default instruction
      panic(err)
}
instructions := "@"
bf.Run(strings.NewReader(instructions))

err = bf.RemoveCustomInstruction('@')
if err != nil {
      // error in case of custom instruction not found
      panic(err)
}


```