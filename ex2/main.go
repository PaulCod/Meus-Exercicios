package main

import "fmt"

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}

	result1 := totalBirdsCount(birdsPerDay)
	fmt.Println(result1)

	birdWeek := birdsInWeek(birdsPerDay, 2)
	fmt.Println(birdWeek)

	birdsPerDay2 := []int{2, 5, 0, 7, 4, 1}
	birdCorrectCount := fixBirdCountLog(birdsPerDay2)
	fmt.Println(birdCorrectCount)
}

func totalBirdsCount(birdsPerDay []int) int {
	var totalBirds int
	for _, birds := range birdsPerDay {
		totalBirds += birds
	}

	return totalBirds
}

func birdsInWeek(birdsPerDay []int, week int) int {
	var days int
	if week == 1 {
		days = week * 6
	} else {
		days = (week * 7) - 1
	}

	var calculateBirds int
	for i, birds := range birdsPerDay {
		if week == 1 {
			if i <= days {
				calculateBirds += birds
			}
		}

		if week != 1 {
			var reg int
			reg = days - 7
			if i <= reg {
				continue
			} else {
				calculateBirds += birds
			}
		}

		if i > days {
			break
		}

	}

	return calculateBirds
}

func fixBirdCountLog(birdsPerDay []int) []int {
	var i = 0
	for in := range birdsPerDay {
		if i == in {
			birdsPerDay[i]++
		}
	}

	return birdsPerDay
}
