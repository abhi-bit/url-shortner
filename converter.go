package urlshortner

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	BASE             = 62
	UPPERCASE_OFFSET = 55
	LOWERCASE_OFFSET = 61
	DIGIT_OFFSET     = 48
)

var (
	number int64
	input  string
)

func TrueOrder(char rune) int64 {

	var output int64

	switch {
	case char >= '0' && char < '9':
		output = int64(char) - DIGIT_OFFSET
	case char >= 'A' && char <= 'Z':
		output = int64(char) - UPPERCASE_OFFSET
	case char >= 'a' && char <= 'z':
		output = int64(char) - LOWERCASE_OFFSET
	}
	return output
}

func TrueChr(number int64) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcedfghijklmnopqrstuvwxyz"
	return string(alphanum[number])
}

func Dehydrate(number int64) string {
	var output string
	var remainder int64

	if number == 0 {
		return "0"
	}

	for number > 0 {
		remainder = number % BASE
		output = TrueChr(remainder) + output
		number /= BASE
	}

	return output
}

func Saturate(input string) int64 {
	var sum int64
	reversedInput := Reverse(input)
	for index, char := range reversedInput {
		sum += TrueOrder(char) * int64(math.Pow(BASE, float64(index)))
	}
	return sum
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	switch {
	case os.Args[1] == "-s":
		input = os.Args[2]
		fmt.Println(Saturate(input))
	case os.Args[1] == "-d":
		number, _ = strconv.ParseInt(os.Args[2], 10, 64)
		fmt.Println(Dehydrate(number))
	}
}
