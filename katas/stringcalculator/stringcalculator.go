package stringcalculator

import (
    "fmt"
    "strconv"
)

func Add(stringNumbers string) int {
	i, _ := strconv.Atoi(stringNumbers)
	fmt.Println(i)    // 42
}
