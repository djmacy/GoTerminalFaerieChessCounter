package main

import (
	"fmt"
	"strings"
)

func checkForRankII(piece int) {
	if piece > 1 || piece < 0 {
		fmt.Println("Please select 0 to 1", piece)
	}
}

func main() {
	//:= allows me to declare and assign the variable
	run := true

	for run {
		chooseDiff := true
		totalPoints := 0
		maxPoints := 0
		rank1 := 0
		rank2 := 0
		rank3 := 0

		fmt.Println("Hello! Welcome to the Faerie Chess Counter\n\nBegin by choosing the difficulty. Insert B for beginner, I for intermediate, and A for advanced")

		for chooseDiff {
			var difficulty string
			fmt.Println("\nWhat difficulty are you playing on:")
			//allows for the input to be assigned as the difficulty variable assigned earlier
			fmt.Scan(&difficulty)
			difficulty = strings.ToUpper(difficulty)

			if difficulty == "B" {
				fmt.Println("Great you will play with 60 - 65 points")
				totalPoints = 0
				rank1 = 0
				rank2 = 0
				rank3 = 0
				maxPoints = 65
				chooseDiff = false
				break
			} else if difficulty == "I" {
				fmt.Println("Great you will play with 65 - 70 points")
				totalPoints = 0
				rank1 = 0
				rank2 = 0
				rank3 = 0
				maxPoints = 70
				chooseDiff = false
				break
			} else if difficulty == "A" {
				fmt.Println("Great you will play with 70 - 75 points")
				totalPoints = 0
				rank1 = 0
				rank2 = 0
				rank3 = 0
				maxPoints = 75
				chooseDiff = false
				break
			} else {
				fmt.Println("Please choose B, I, or A")
			}
		}

		var pawn, peasant, soldier int

		for rank1 < 8 {
			rank1 = 0
			totalPoints = 0

			pawnValue := 1
			peasantValue := 2
			soldierValue := 3

			pawnLimit := 8
			peasantLimit := 2
			soldierLimit := 2

			peasantMin := 0

			fmt.Println("How many pawns would you like: ")
			fmt.Scan(&pawn)

			if pawn < 4 || pawn > pawnLimit {
				fmt.Println("Invalid input. Please select between 4 and", pawnLimit, "pawns.")
				continue
			}

			totalPoints += pawn * pawnValue
			rank1 += pawn

			if rank1 == 8 {
				fmt.Println("You have selected the maximum number of Rank I pieces")
				break
			} else if rank1 == 4 {
				peasant = 2
				soldier = 2
				totalPoints += (peasant*peasantValue + soldier*soldierValue)
				fmt.Println("Due to picking 4 pawns you will automatically get 2 peasants and 2 soldiers")
				break
			} else if rank1 > 4 {
				peasantMin = max(0, 8-(rank1+soldierLimit))
				peasantLimit = min(2, 8-rank1)
			}

			fmt.Println("You can select from", peasantMin, "to", peasantLimit, "peasants")
			fmt.Println("How many peasants would you like: ")
			fmt.Scan(&peasant)

			if peasant < peasantMin || peasant > peasantLimit {
				fmt.Println("Invalid input. Please select between", peasantMin, "and", peasantLimit)
				continue
			}

			totalPoints += peasant * peasantValue
			rank1 += peasant

			if rank1 == 8 {
				fmt.Println("You have picked the maximum number of pieces")
				break
			} else if rank1 == 6 {
				fmt.Println("You automatically get 2 soldiers")
				soldier = 2
				totalPoints += soldier * soldierValue
				rank1 += soldier
			} else if rank1 == 7 {
				fmt.Println("You automatically get 1 soldier")
				soldier = 1
				totalPoints += soldier * soldierValue
				rank1 += soldier
			}
		}

		fmt.Println("You selected", pawn, "Pawns,", peasant, "Peasants,", soldier, "Soldiers")
		fmt.Println("Total points for Rank I pieces:", totalPoints, ". You have", maxPoints-totalPoints, "points left!\nNow lets pick your Rank II pieces. You can only select 6 Rank II pieces.")

		rankIPoints := totalPoints
		doneSelecting := false

		for rank2 < 6 && !doneSelecting {
			totalPoints = rankIPoints

			rookValue := 9
			knightValue := 4
			bishopValue := 6
			catapultValue := 3
			chamberlainValue := 6
			courtesanValue := 6
			heraldValue := 6
			inquisitorValue := 6
			lancerValue := 5
			pontiffValue := 8
			thiefValue := 5
			towerValue := 10

			classicalLimit := 2

			var rook, knight, bishop, catapult, chamberlain, courtesan, herald, inquisitor, lancer, pontiff, thief, tower int
			var userInput string

			fmt.Println("\nHow many rooks would you like: ")
			fmt.Scan(&rook)
			if rook > classicalLimit || rook < 0 {
				fmt.Println("Invalid input. Please select between 0 and", classicalLimit, "rooks.")
				continue
			}
			totalPoints += rook * rookValue
			rank2 += rook

			fmt.Println("How many knights would you like: ")
			fmt.Scan(&knight)
			if knight > classicalLimit || knight < 0 {
				fmt.Println("Invalid input. Please select between 0 and", classicalLimit, "knights.")
				continue
			}
			totalPoints += knight * knightValue
			rank2 += knight

			fmt.Println("How many bishops would you like: ")
			fmt.Scan(&bishop)
			if bishop > classicalLimit || bishop < 0 {
				fmt.Println("Invalid input. Please select between 0 and", classicalLimit, "bishop.")
				continue
			}
			totalPoints += bishop * bishopValue
			rank2 += bishop

			if rank2 == 6 {
				doneSelecting = true
				break
			}

			fmt.Println("\nFor the rest of these insert y for yes and n for no")

			fmt.Println("Would you like a catapult: ")
			fmt.Scan(&userInput)
			if userInput == "y" {
				catapult = 1
				checkForRankII(catapult)
				totalPoints += catapult * catapultValue
				rank2 += catapult
			} else if userInput != "n" && userInput != "y" {
				rank2 = 0
				println("please insert a valid value: y or n")
				continue
			}
			if rank2 == 6 {
				doneSelecting = true
				break
			}

			fmt.Println("Would you like a chamberlain: ")
			fmt.Scan(&userInput)
			if userInput == "y" {
				chamberlain = 1
				checkForRankII(chamberlain)
				totalPoints += chamberlain * chamberlainValue
				rank2 += chamberlain
			} else if userInput != "n" && userInput != "y" {
				rank2 = 0
				println("please insert a valid value: y or n")
				continue
			}
			if rank2 == 6 {
				doneSelecting = true
				break
			}

			fmt.Println("Would you like a courtesan: ")
			fmt.Scan(&userInput)
			if userInput == "y" {
				courtesan = 1
				checkForRankII(courtesan)
				totalPoints += courtesan * courtesanValue
				rank2 += courtesan
			} else if userInput != "n" && userInput != "y" {
				rank2 = 0
				println("please insert a valid value: y or n")
				continue
			}
			if rank2 == 6 {
				doneSelecting = true
				break
			}

			fmt.Println("Would you like a herald: ")
			fmt.Scan(&userInput)
			if userInput == "y" {
				herald = 1
				checkForRankII(herald)
				totalPoints += herald * heraldValue
				rank2 += herald
			} else if userInput != "n" && userInput != "y" {
				rank2 = 0
				println("please insert a valid value: y or n")
				continue
			}
			if rank2 == 6 {
				doneSelecting = true
				break
			}

			fmt.Println("Would you like a inquisitor: ")
			fmt.Scan(&userInput)
			if userInput == "y" {
				inquisitor = 1
				checkForRankII(inquisitor)
				totalPoints += inquisitor * inquisitorValue
				rank2 += inquisitor
			} else if userInput != "n" && userInput != "y" {
				rank2 = 0
				println("please insert a valid value: y or n")
				continue
			}
			if rank2 == 6 {
				doneSelecting = true
				break
			}

			fmt.Println("Would you like a lancer: ")
			fmt.Scan(&userInput)
			if userInput == "y" {
				lancer = 1
				checkForRankII(lancer)
				totalPoints += lancer * lancerValue
				rank2 += lancer
			} else if userInput != "n" && userInput != "y" {
				rank2 = 0
				println("please insert a valid value: y or n")
				continue
			}
			if rank2 == 6 {
				doneSelecting = true
				break
			}

			fmt.Println("Would you like a pontiff: ")
			fmt.Scan(&userInput)
			if userInput == "y" {
				pontiff = 1
				checkForRankII(pontiff)
				totalPoints += pontiff * pontiffValue
				rank2 += pontiff
			} else if userInput != "n" && userInput != "y" {
				rank2 = 0
				println("please insert a valid value: y or n")
				continue
			}
			if rank2 == 6 {
				doneSelecting = true
				break
			}
			fmt.Println("Would you like a thief: ")
			fmt.Scan(&userInput)
			if userInput == "y" {
				thief = 1
				checkForRankII(thief)
				totalPoints += thief * thiefValue
				rank2 += thief
			} else if userInput != "n" && userInput != "y" {
				rank2 = 0
				println("please insert a valid value: y or n")
				continue
			}
			if rank2 == 6 {
				doneSelecting = true
				break
			}

			fmt.Println("Would you like a tower: ")
			fmt.Scan(&userInput)
			if userInput == "y" {
				catapult = 1
				checkForRankII(tower)
				totalPoints += tower * towerValue
				rank2 += tower
			} else if userInput != "n" && userInput != "y" {
				rank2 = 0
				println("please insert a valid value: y or n")
				continue
			}
			if rank2 == 6 {
				doneSelecting = true
				break
			} else {
				println("You only selected ", rank2, " pieces. Be sure to select 6 rank II pieces\nThe count will now reset")
			}
		}

		fmt.Println("You have selected all of your Rank II pieces")
		fmt.Println("Total points for Rank II & Rank I pieces:", totalPoints, ". You have", maxPoints-totalPoints, "points left!\nNow lets pick your Rank III pieces. You can only select 2 Rank III pieces.")

		rankIIIPoints := totalPoints

		for rank3 < 2 {
			rank3 = 0
			totalPoints = rankIIIPoints

			queenValue := 12
			kingValue := 0
			jesterValue := 12
			regentValue := 15

			var queen, king int
			var userInput string

			fmt.Println("Enter y for yes and n to no")
			fmt.Println("Would you like a queen: ")
			fmt.Scan(&userInput)

			if userInput == "y" {
				queen = 1
				totalPoints += queen * queenValue
			} else if userInput != "n" && userInput != "y" {
				rank3 = 0
				println("please insert a valid value: y or n")
			} else {
				fmt.Println("Since you did not select a queen you automatically get a jester")
				jester := 1
				totalPoints += jester * jesterValue
			}

			if maxPoints-totalPoints <= regentValue {
				fmt.Println("You do not have enough points for a regent so you will have to get a king.")
				king = 1
				totalPoints += king * kingValue
				rank3 += king
				break
			}

			fmt.Println("Would you like a king: ")
			fmt.Scan(&userInput)

			if userInput == "y" {
				king = 1
				totalPoints += king * kingValue
			} else if userInput != "n" && userInput != "y" {
				rank3 = 0
				println("please insert a valid value: y or n")
				continue
			} else {
				fmt.Println("Since you did not select a king you automatically get a regent")
				regent := 1
				totalPoints += regent * regentValue
				break
			}
		}

		fmt.Println("You have selected all of your Rank III pieces")
		fmt.Println("Total points for Rank III pieces:", totalPoints, ". You have", maxPoints-totalPoints, "points left!")
		fmt.Println("Would you like to change the difficulty or pick different pieces? Insert y or n: ")
		var redo string
		fmt.Scan(&redo)

		if redo != "y" {
			fmt.Println("Thank you. See you next time!")
			break
		} else if redo == "y" {
			chooseDiff = true
			rank1 = 0
			rank2 = 0
			rank3 = 0
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
