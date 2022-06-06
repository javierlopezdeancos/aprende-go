package leapyear

func divisibleBy(number, divisor int) bool {
    return number%divisor == 0
}

type Year int

func (y Year) IsLeap() bool {
    if divisibleBy(int(y), 400) {
        return true
    }

    if divisibleBy(int(y), 100) {
        return false
    }
    return divisibleBy(int(y), 4)
}



