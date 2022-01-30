package main

import (
	"fmt"
	"strconv"
	"strings"
)

var tempoDeForno = 40

func main() {
	res := valid("6488846590874657")
	fmt.Println(res)
}

// Valid returns bool.
func valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	length := len(s)
	if length < 2 {
		return false
	}
	var digit, doubledUp, sum int
	var err error
	for i := 0; i < length; i++ {
		digit, err = strconv.Atoi(string(s[length-1-i]))
		if err != nil {
			return false
		}
		if i%2 == 0 {
			sum += digit
			continue
		}
		doubledUp = 2 * digit
		if doubledUp > 9 {
			doubledUp -= 9
		}
		sum += doubledUp
	}
	return sum%10 == 0
}
