package main

import (
"fmt"
)

func getRating(maxPos int, midScore int, score int) (int, int) {
	var level, rating, perM int
	var PRM = []int {0, 25, 925, 1275, 2500, 3950, 4700, 5300, 5750, 6275}
	perM = (score * 1000) / maxPos
	if midScore != 2500 {
		var multiplier int = (midScore * 1000) / 2500
		for i := 1; i < 10; i++ {
			PRM[i] =(PRM[i] * multiplier) / 1000
		}
	}		
	if perM < 250 {
		rating = ((PRM[0] / (25-0)) * perM) / 10 // 0 rating per promile onder 250 promile
	} else if perM < 300 {
		rating = PRM[0] + ((((PRM[1]-PRM[0]) / (30-25)) * (perM-250)) / 10) //0.5 rating per promile boven 250 promile 
	} else if perM < 400 {
		rating = PRM[1] + ((((PRM[2]-PRM[1]) / (40-30)) * (perM-300)) / 10) //9 rating per promile boven 300 promile
	} else if perM < 450 {
		rating = PRM[2] + ((((PRM[3]-PRM[2]) / (45-40)) * (perM-400)) / 10) //7 rating per promile boven 300 promile
	} else if perM < 500 {
		rating = PRM[3] + ((((PRM[4]-PRM[3]) / (50-45)) * (perM-450)) / 10) //25 rating per promile boven 400 promile
	} else if perM < 530 {
		rating = PRM[4] + ((((PRM[5]-PRM[4]) / (53-50)) * (perM-500)) / 10) //48 rating per promile boven 500 promile
	} else if perM < 550 {
		rating = PRM[5] + ((((PRM[6]-PRM[5]) / (55-53)) * (perM-530)) / 10) //37 rating per promile boven 530 promile
	} else if perM < 570 {
		rating = PRM[6] + ((((PRM[7]-PRM[6]) / (57-55)) * (perM-550)) / 10) //30 rating per promile boven 550 promile
	} else if perM < 600 {
		rating = PRM[7] + ((((PRM[8]-PRM[7]) / (60-57)) * (perM-570)) / 10) //15 rating per promile boven 570 promile
	} else if perM < 650 {
		rating = PRM[8] + ((((PRM[9]-PRM[8]) / (65-60)) * (perM-600)) / 10) //20 rating per promile ONDER 650 PROMILE
	} else {
		rating = PRM[9] + ((perM-700)*20) //20 rating per promile boven 700 promile
	}
	if rating < 500 {
		return 1, rating
	}
	level = ((rating - 500) / 75) + 2
	return level, rating
}

func main() {
	var maxPos, score, midScore int
	fmt.Print("\nRating op 50% (default 2500): ")
	fmt.Scanln(&midScore)
	fmt.Print("\nMax punten mogelijk: ")
	fmt.Scanln(&maxPos)
	for {
		fmt.Print("\nScore: ")
		fmt.Scanln(&score)
		level,rating := getRating(maxPos, midScore, score)
		fmt.Println("Level:", level, " Rating:", rating)
		fmt.Println()
	}
}
