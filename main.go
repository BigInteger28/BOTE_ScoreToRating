package main

import (
"fmt"
)

func getRating(maxPos int, score int) (int,int) {
	var level int
	var rating int
	var ps31 int = (maxPos*310)/1000
	var ps50 int = (maxPos*494)/1000
	var ps60 int = (maxPos*592)/1000
	if (score*100)/maxPos < 50 {
		rating = ((score-ps31)*5)+1095
	} else if (score*100)/maxPos < 60 {
		rating = ((score-ps50)*10)+2500
	} else {
		rating = ((score-ps60)*20)+4050
	}
	if rating < 500 {
		return 1, rating
	}
	level = ((rating-500)/75)+2
	return level, rating
}

func main() {
	var maxPos, score int
	fmt.Print("\nMax punten mogelijk: ")
	fmt.Scanln(&maxPos)
	for {
		fmt.Print("\nScore: ")
		fmt.Scanln(&score)
		level,rating := getRating(maxPos, score)
		fmt.Println("Level:", level, " Rating:", rating)
		fmt.Println()
	}
}
