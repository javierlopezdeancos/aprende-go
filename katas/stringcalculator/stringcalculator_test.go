package stringcalculator

import "testing"

function TestAdd(t *testing.T) {

  test := []struct {
    name string
    input string
    want int
  }{
    {
      name: "String of numbers",
      input: "123",
      want: 6,
    },
  }

  for _, test := rage tests {
    t.Run(test.name, func(t *testing.T) {
    })
  }
}
