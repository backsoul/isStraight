package main

import "fmt"

func isStraight(cards []int) bool {
	// last number
	last := 0
	// variable len is the number of elements in the slice
	count := 0
	valor := len(cards)
	if valor >= 5 {
		valor = 5
	}
	// sort order of the cards
	sort(cards)
	//  valida number in cards
	if inArray(14, cards) {
		cards = append(cards, 1)
		sort(cards)
		valor = valor - 1
	}

	for _, card := range cards {
		// si el numero es el mismo que el anterior no es escalera
		if card == last {
			continue
			// si el numero es el siguiente al anterior
		} else if card == last+1 {
			count++
			last = card
			// si el numero es el anterior y el siguiente
		} else {
			count = 1
			last = card
		}
		// si el numero es el ultimo y el contador es igual a la longitud del array
		if count == valor {
			return true
		}
	}
	return false
}

// validar si el numero esta en el array
func inArray(val int, array []int) bool {
	for _, v := range array {
		if v == val {
			return true
		}
	}
	return false
}

// organiza el array de menor a mayor
func sort(cards []int) {
	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			if cards[i] > cards[j] {
				cards[i], cards[j] = cards[j], cards[i]
			}
		}
	}
}
func main() {
	// escalera: 14-2-3-4-5
	// escalera: 9-10-11-12-13
	// escalera: 2-7-8-5-10-9-11
	// no es escalera: 7-8-12-13-14
	fmt.Println(isStraight([]int{14, 2, 3, 4, 5}))
	fmt.Println(isStraight([]int{9, 10, 11, 12, 13}))
	fmt.Println(isStraight([]int{2, 7, 8, 5, 10, 9, 11}))
	fmt.Println(isStraight([]int{7, 8, 12, 13, 14}))

	// test case
	results1 := isStraight([]int{2, 3, 4, 5, 6})
	fmt.Println(results1)
	results2 := isStraight([]int{14, 5, 4, 2, 3})
	fmt.Println(results2)
	results3 := isStraight([]int{7, 7, 12, 11, 3, 4, 14})
	fmt.Println(results3)
	results4 := isStraight([]int{7, 3, 2})
	fmt.Println(results4)
}
