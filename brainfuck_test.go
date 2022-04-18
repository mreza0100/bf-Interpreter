package brainfuck_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	. "github.com/mreza0100/brainfuck"
)

type E2ETest struct {
	name     string
	input    string
	expected []byte
}

func (e E2ETest) getI() io.Reader {
	return strings.NewReader(e.input)
}

func TestE2E(t *testing.T) {
	tests := []E2ETest{
		{
			name:     "Hello World",
			input:    "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.",
			expected: []byte("Hello World!\n"),
		},
		{
			name:     "Simple loop",
			input:    "++++[-.]",
			expected: []byte{3, 2, 1, 0},
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
			`,
			expected: []byte{2, 1, 0, 2, 1, 0},
		},
		{
			name:     "must panic",
			input:    `+++[]---`,
			expected: nil,
		},
		{
			name:     "must panic",
			input:    `+++]---`,
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bf := New()
			writter := new(bytes.Buffer)

			bf.Writter = writter
			bf.Verbos = false

			defer func() {
				if r := recover(); r != nil {
					if test.expected == nil {
						return
					}

					t.Errorf("panic: %v", r)
				}
			}()

			bf.Entry(test.getI())

			if !bytes.Equal(writter.Bytes(), test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, writter.Bytes())
			}
		})
	}
}
