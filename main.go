package main

import (
"fmt"
)

func getRating(maxScore, lowRating, HighRating, score int) (int, int) {
	var level, rating, midRating, midScore, SixTwoScore, fourthyScore int
	midRating = (lowRating + HighRating) / 2
	midScore = maxScore / 2
	fourthyScore = (maxScore / 10) * 4
	SixTwoScore = int((float64(maxScore) / float64(100)) * float64(62))
	if score < midScore {
		if score < fourthyScore {
			ratingF := (float64(lowRating) / float64(fourthyScore)) * float64(score)
			rating = int(ratingF)
		} else {
			ratingF := ((float64(midRating) - float64(lowRating)) / (float64(midScore) - float64(fourthyScore))) * ((float64(score) - float64(fourthyScore)))
			rating = lowRating + int(ratingF)
		}
	} else if score < SixTwoScore {
		ratingF := ((float64(HighRating) - float64(midRating)) / (float64(SixTwoScore) - float64(midScore))) * ((float64(score) - float64(midScore)))
		rating = midRating + int(ratingF)
	} else {
		rating = HighRating + ((score - SixTwoScore) * 35)
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
	fmt.Println("Default lowest rating (start 1 level higher than lowest engine): 500")
	fmt.Println("Default highest rating for 62 percent winrate: 6425")
	fmt.Println("\nDefaults bij engines van level 51 tot oneindig.")
	fmt.Println("Default lowest rating (start 1 level higher than lowest engine): 4250")
	fmt.Println("Default highest rating for 62 percent winrate: 10175")
	
	fmt.Print("\nLowest rating engine: ")
	fmt.Scanln(&lowRating)
	
	fmt.Print("Highest rating engine for 62 percent winrate: ")
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
