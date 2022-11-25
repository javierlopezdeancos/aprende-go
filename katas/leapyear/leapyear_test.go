package leapyear

import "testing"

func TestLeapYear(t *testing.T) {
  tests := []struct {
    name string
    year Year
    want bool
  }{
    {
      name: "Common year",
      year: 1995,
      want: false,
    },
    {
      name: "Leap year 2020",
      year: 2020,
      want: true,
    },
    {
      name: "Leap year 1996",
      year: 1996,
      want: true,
    },
    {
      name: "Extra common year",
      year: 1900,
      want: false,
    },
    {
      name: "Extra leap year",
      year: 2000,
      want: true,
    },
  }

  for _, test := range tests {
    t.Run(test.name, func(t *testing.T) {
      got := test.year.IsLeap()

      if got != test.want {
        t.Errorf("Expected %#v, got %#v", test.want, got)
      }
    })
  }
}
