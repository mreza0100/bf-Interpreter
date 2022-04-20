# Brainfuck Interpreter


``` bash
go get -u github.com/mreza0100/brainfuck
```

## Usage
``` go
myCustomWritter := new(bytes.Buffer)
bf := brainfuck.New(&brainfuck.NewOptions{
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
## Custom Commands
``` go
bf := brainfuck.New(&brainfuck.NewOptions{})
err := bf.AddCustomCommand('@', func(icw brainfuck.InstructionCtl) {
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
      // error in case command is already defined or it's a default command
      panic(err)
}
instructions := "@"
bf.Run(strings.NewReader(instructions))

err = bf.RemoveCustomCommand('@')
if err != nil {
      // error in case of custom command not found
      panic(err)
}
```