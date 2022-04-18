package brainfuck_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	. "github.com/mreza0100/brainfuck"
)

type E2ETest struct {
	name        string
	input       string
	expected    []byte
	shouldPanic bool
	options     *NewOptions
}

func (e E2ETest) getI() io.Reader {
	return strings.NewReader(e.input)
}

func TestE2E(t *testing.T) {
	tests := []E2ETest{
		{
			name:        "Hello World",
			input:       "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.",
			expected:    []byte("Hello World!\n"),
			options:     nil,
			shouldPanic: false,
		},
		{
			name:  "stackover flow",
			input: "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.", expected: nil,
			shouldPanic: true,
			options: &NewOptions{
				MemorySize:     1,
				IsMemoryStatic: true,
				Verbos:         false,
				Reader:         nil,
			},
		},
		{
			name:  "Simple loop",
			input: "++++[-.]", expected: []byte{3, 2, 1, 0},
			shouldPanic: false,
			options:     nil,
		},
		{
			name: "nested loop",
			input: `
			++
			[
				>+++
				[
					-.
				]
				<-
			]
			`, expected: []byte{2, 1, 0, 2, 1, 0},
			shouldPanic: false,
			options:     nil,
		},
		{
			name:  "must panic",
			input: `+++[]---`, expected: nil,
			shouldPanic: true,
			options:     nil,
		},
		{
			name:  "must panic",
			input: `+++]---`, expected: nil,
			shouldPanic: true,
			options:     nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if test.shouldPanic {
						// test passed
						return
					}

					panic(r)
				}
			}()

			writter := new(bytes.Buffer)

			var options *NewOptions
			if test.options == nil {
				options = &NewOptions{
					MemorySize:     10,
					IsMemoryStatic: true,
					Verbos:         false,
					Writter:        writter,
					Reader:         nil,
				}
			}
			bf := New(options)

			bf.Run(test.getI())

			if !bytes.Equal(writter.Bytes(), test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, writter.Bytes())
			}
		})
	}
}
