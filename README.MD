# Brainfuck interpreter


``` bash
go get -u github.com/mreza0100/brainfuck
```

## usage
``` go
myCustomWritter := new(bytes.Buffer)
bf := New(&NewOptions{
      // inline documention for each of them is accessible in source code
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
