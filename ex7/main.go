package main

import (
	"fmt"
	"strings"
)

func main() {
	diferences := hamming("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT")
	fmt.Println(diferences)
}

func hamming(dna1, dna2 string) int {
	if len(dna1) != 17 || len(dna2) != 17 {
		return 0
	}

	dna1 = strings.ToUpper(dna1)
	dna2 = strings.ToUpper(dna2)

	dnaAscii1 := returnAscii(dna1)
	dnaAscii2 := returnAscii(dna2)

	var result int
	for i := 0; i < len(dna1); i++ {
		if dnaAscii1[i] != dnaAscii2[i] {
			result += 1
		}
	}

	return result
}

func returnAscii(dna string) []rune {
	var asciiDna []rune
	for _, numbersAs := range dna {
		asciiDna = append(asciiDna, numbersAs)
	}

	return asciiDna
}
