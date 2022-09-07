# [Brainfuck Interpreter](https://en.wikipedia.org/wiki/Brainfuck)


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
bf := brainfuck.New(&brainfuck.NewOptions{Verbos: true, StaticMemory: false, MemorySize: 10})

err := bf.AddCustomCommand('@', func(ctl brainfuck.CommandsCtl) {
      // you have tools to do almost anything
      ctl.GetMemory()
      ctl.GetPointerPosition()

      ctl.Print()
      ctl.MoveForward()
      ctl.MoveBackward()
      ctl.Increment()
      ctl.Decrement()
      ctl.Read()
})
if err != nil {
      panic(err)
}

err = bf.RemoveCustomCommand('@')
if err != nil {
      // error in case of custom instruction not found
      panic(err)
}

bf.Run(strings.NewReader("@"))
```
