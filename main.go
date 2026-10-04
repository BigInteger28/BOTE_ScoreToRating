package main

import (
"fmt"
)

func getRating(maxScore, lowRating, HighRating, score int) (int, int) {
	var level, rating, midRating, midScore, SixSevenScore, fourthyScore int
	midRating = (lowRating + HighRating) / 2
	midScore = maxScore / 2
	fourthyScore = (maxScore / 10) * 4
	SixSevenScore = int((float64(maxScore) / float64(100)) * float64(67))
	if score < midScore {
		if score < fourthyScore {
			ratingF := (float64(lowRating) / float64(fourthyScore)) * float64(score)
			rating = int(ratingF)
		} else {
			ratingF := ((float64(midRating) - float64(lowRating)) / (float64(midScore) - float64(fourthyScore))) * ((float64(score) - float64(fourthyScore)))
			rating = lowRating + int(ratingF)
		}
	} else if score < SixSevenScore {
		ratingF := ((float64(HighRating) - float64(midRating)) / (float64(SixSevenScore) - float64(midScore))) * ((float64(score) - float64(midScore)))
		rating = lowRating + int(ratingF)
	} else {
		rating = HighRating + ((score - SixSevenScore) * 35)
	}
	if rating < 500 {
		return 1, rating
	}
	level = ((rating - 500) / 75) + 2
	return level, rating
}

func main() {
	var maxScore, score, lowRating, HighRating int
	fmt.Println("Defaults bij engines van level 1 tot en met level 78 en 1150 engines.")
	fmt.Print("Default lowest rating: 500")
	fmt.Print("\nLowest rating engine: ")
	fmt.Scanln(&lowRating)
	fmt.Println("Default highest rating for 67 percent winrate: 15602")
	fmt.Print("Highest rating engine for 67 percent winrate: ")
	fmt.Scanln(&HighRating)
	fmt.Print("Max punten mogelijk: ")
	fmt.Scanln(&maxScore)
	for {
		fmt.Print("\nScore: ")
		fmt.Scanln(&score)
		level,rating := getRating(maxScore, lowRating, HighRating, score)
		fmt.Println("Level:", level, " Rating:", rating)
		fmt.Println()
	}
}
