package main

import (
"fmt"
)

func getRating(maxPos int, score int) (int, int) {
    var level int
    var rating int
    var ps30 int = (maxPos * 300) / 1000
    var ps50 int = (maxPos * 500) / 1000
    var ps60 int = (maxPos * 600) / 1000
	if (score * 100) / maxPos < 30 {
		return 1, 0
	}
    if (score * 100) / maxPos < 50 {
        rating = ((score - ps30) * 7)
    } else if (score * 100) / maxPos < 60 {
        rating = (((score - ps50) * 95))/10 + 2520
    } else {
        rating = ((score - ps60) * 20) + 4230
    }
    if rating < 500 {
        return 1, rating
    }
    level = ((rating - 500) / 75) + 2
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
