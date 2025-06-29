package main

import (
"fmt"
)

func getRating(maxPos, score int) (int, int) {
	var level, rating int
	var scores = [76]int {
		793, 822, 837, 850, 860, 875, 891, 909, 958, 979, 1002, 1019, 1031, 1043, 1054, 1063, 1068, 1076, 1082, 1088, 1097, 1104, 1110, 1118, 
		1125, 1129, 1132, 1136, 1142, 1147, 1151, 1155, 1157, 1159, 1163, 1167, 1173, 1177, 1182, 1185, 1188, 1191, 1196, 1200, 1204, 1209, 1212, 1216, 1219, 
		1224, 1227, 1231, 1235, 1240, 1245, 1249, 1253, 1257, 1263, 1267, 1270, 1276, 1283, 1290, 1296, 1303, 1309, 1316, 1328, 1338, 1346, 1361, 1378, 1388, 
		1436,
	}
	scoreF := (float64(score) * float64(maxPos)) / 2286
	score = int(scoreF)
	
	if score >= 1436 {
		rating = 6050 + ((score-1436)*25)
		level = ((rating - 500) / 75) + 2
	} else {
		for i := range scores {
			if score < scores[i] {
				level = i+1
				rating = 425+((level-1)*75)
				break
			}
		}
	}
	return level, rating
}

func main() {
	var maxPos, score, level, rating int
	fmt.Print("\nMax punten mogelijk: ")
	fmt.Scanln(&maxPos)	
	for {
		fmt.Print("\nScore: ")
		fmt.Scanln(&score)
		level,rating = getRating(maxPos, score)
		fmt.Println("Level:", level, " Rating:", rating)
		fmt.Println()
	}
}
