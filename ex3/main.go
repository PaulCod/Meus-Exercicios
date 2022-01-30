package main

import "fmt"

func main() {
	card, ok := getItem([]int{1, 2, 4, 1}, 2)
	if ok != true {
		fmt.Println(card)
		fmt.Println(ok)
		return
	}
	fmt.Println(card)
	fmt.Println(ok)

	index := 2
	newCard := 6
	res := setItem([]int{1, 2, 4, 1}, index, newCard)
	fmt.Println(res)

	slice := prefilledSlice(8, 3)
	fmt.Println(slice)

	stack := removeItem([]int{3, 2, 6, 4, 8}, 11)
	fmt.Println(stack)
}

func getItem(cards []int, index int) (int, bool) {
	length := len(cards)
	if length < index || length < 0 {
		return 0, false
	}

	return cards[index], true
}

func setItem(cards []int, index, newCard int) []int {
	length := len(cards)
	if index < 0 || index > length {
		cards[length+1] = newCard
		return cards
	}
	cards[index] = newCard

	return cards
}

func prefilledSlice(value, length int) []int {
	if length <= 0 {
		return []int{}
	}

	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = value
	}

	return slice
}

func removeItem(stack []int, index int) []int {
	if index < 0 || index >= len(stack) {
		return stack
	}

	return append(stack[:index], stack[index+1:]...)
}
