package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Card struct {
	suit   string
	number int
}

func BubbleSortForCards(cards []Card, N int) {
	for i := 0; i < N; i++ {
		for j := N - 1; j >= i+1; j-- {
			if cards[j].number < cards[j-1].number {
				cards[j], cards[j-1] = cards[j-1], cards[j]
			}
		}
	}
}

func SelectionSortForCards(cards []Card, N int) {
	for i := 0; i < N; i++ {
		minJ := i
		for j := i; j < N; j++ {
			if cards[j].number < cards[minJ].number {
				minJ = j
			}
		}
		cards[i], cards[minJ] = cards[minJ], cards[i]
	}
}

func printCards(cards []Card) {
	for i := 0; i < len(cards); i++ {
		fmt.Printf("%s%d", cards[i].suit, cards[i].number)
		if i == len(cards)-1 {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	cards := make([]Card, n)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Scan(&tmp)
		number, _ := strconv.Atoi(tmp[1:])
		cards[i] = Card{suit: tmp[:1], number: number}
	}
	bCards := make([]Card, n)
	sCards := make([]Card, n)
	_ = copy(bCards, cards)
	_ = copy(sCards, cards)

	BubbleSortForCards(bCards, n)
	printCards(bCards)
	fmt.Println("Stable")

	SelectionSortForCards(sCards, n)
	printCards(sCards)
	if reflect.DeepEqual(bCards, sCards) {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
}
