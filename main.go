package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Defines the attributes associated with each chess piece
type Pieces struct {
	Value     int
	MinPieces int
	MaxPieces int
}

// These constants are used to create the keys for the pieces hashmap
const (
	pawn        = "pawns"
	peasant     = "peasants"
	soldier     = "soldiers"
	rook        = "rooks"
	knight      = "knights"
	bishop      = "bishops"
	catapult    = "catapult"
	chamberlain = "chamberlain"
	courtesan   = "courtesan"
	herald      = "herald"
	inquisitor  = "inquisitor"
	lancer      = "lancer"
	pontiff     = "pontiff"
	thief       = "thief"
	tower       = "tower"
	queen       = "queen"
	jester      = "jester"
	king        = "king"
	regent      = "regent"
)

// The hashmap actually puts values to the attributes of each piece
var pieces = map[string]Pieces{
	pawn:        {1, 4, 8},
	peasant:     {2, 0, 2},
	soldier:     {3, 0, 2},
	rook:        {9, 0, 2},
	knight:      {4, 0, 2},
	bishop:      {6, 0, 2},
	catapult:    {3, 0, 1},
	chamberlain: {6, 0, 1},
	courtesan:   {6, 0, 1},
	herald:      {6, 0, 1},
	inquisitor:  {8, 0, 1},
	lancer:      {5, 0, 1},
	pontiff:     {8, 0, 1},
	thief:       {5, 0, 1},
	tower:       {10, 0, 1},
	queen:       {12, 0, 1},
	jester:      {12, 0, 1},
	king:        {0, 0, 1},
	regent:      {15, 0, 1},
}

func main() {
	for {
		//:= is used to declare and initialize a variable
		maxPoints := chooseDifficulty()
		totalPoints := calculateRankIPoints(maxPoints)
		//printf allows us to format the printing and use %d for integers
		fmt.Printf("Total points for Rank I pieces: %d. You have %d points left!\n", totalPoints, maxPoints-totalPoints)
		fmt.Println("Now let's pick your Rank II pieces. You can only select 6 Rank II pieces.")
		totalPoints = calculateRankII(maxPoints, totalPoints)
		fmt.Println("You have selected all of your Rank II pieces.")
		fmt.Printf("Total points for Rank II & Rank I pieces: %d. You have %d points left!\n", totalPoints, maxPoints-totalPoints)
		fmt.Println("Now let's pick your Rank III pieces. You can only select 2 Rank III pieces.")
		totalPoints = calculateRankIII(maxPoints, totalPoints)
		fmt.Println("You have selected all of your Rank III pieces.")
		fmt.Printf("Total points for Rank III pieces: %d. You have %d points left!\n", totalPoints, maxPoints-totalPoints)
		fmt.Println("Would you like to go again?")
		var redo string
		//Go requires using the & operator talking to the terminal. It will not wait for an input if not.
		fmt.Scanln(&redo)
		if strings.ToLower(redo) != "y" {
			fmt.Println("Thank you. See you next time!")
			break
		}
	}
}

func chooseDifficulty() int {
	var maxPoints int
	chooseDiff := true

	fmt.Println("Hello! Welcome to the Faerie Chess Counter")
	fmt.Println("Begin by choosing the difficulty.")
	fmt.Println("Insert B for beginner, I for intermediate, and A for advanced")

	for chooseDiff {
		var difficulty string
		fmt.Scanln(&difficulty)
		switch strings.ToUpper(difficulty) {
		case "B":
			fmt.Println("Great! You will play with 65 points.")
			maxPoints = 65
			chooseDiff = false
		case "I":
			fmt.Println("Great! You will play with 70 points.")
			maxPoints = 70
			chooseDiff = false
		case "A":
			fmt.Println("Great! You will play with 75 points.")
			maxPoints = 75
			chooseDiff = false
		default:
			fmt.Println("Please choose B, I, or A.")
		}
	}

	return maxPoints
}

func checkForValidClassicalInput(statement string, min int, max int) int {
	errorMessage := fmt.Sprintf("Invalid input. Please enter a number between %d and %d.", min, max)
	for {
		fmt.Print(statement)
		var userInput string
		fmt.Scanln(&userInput)
		//regexp returns two values, a boolean on whether the regular expression matches and an error. Similar to in python
		//we can put an underscore if we do not want to deal with that return value. This essentially makes sure we get a positive int
		if matched, _ := regexp.MatchString("\\d+", userInput); matched {
			intValue := 0
			//Sscan parses the string userInput into an integer and assign it as intValue. It knows to parse to an integer
			//because intValue is of integer type.
			//Sscan returns two values, how many items got scanned and if there was an error
			_, error := fmt.Sscan(userInput, &intValue)
			//if there is no error and the int value is between the passed parameters it will return the value
			if error == nil && intValue >= min && intValue <= max {
				return intValue
			}
		}
		fmt.Println(errorMessage)
	}
}

func checkForValidFaerieInput(statement string) int {
	for {
		fmt.Print(statement)
		var userInput string
		fmt.Scanln(&userInput)
		switch strings.ToLower(userInput) {
		case "y":
			return 1
		case "n":
			return 0
		default:
			fmt.Println("Please insert y or n")
		}
	}
}

func calculateRankIPoints(maxPoints int) int {
	rank1 := 0
	totalPoints := 0
	//see pieces data structure
	soldierLimit := pieces[soldier].MaxPieces

	for {
		pawnsSelected := checkForValidClassicalInput("How many pawns would you like: ", pieces[pawn].MinPieces, pieces[pawn].MaxPieces)
		var peasantsSelected, soldiersSelected int

		totalPoints += pawnsSelected * pieces[pawn].Value
		rank1 += pawnsSelected

		if rank1 == pieces[pawn].MaxPieces {
			fmt.Println("You have selected the maximum number of Rank I pieces")
			break
		} else if rank1 == pieces[pawn].MaxPieces/2 {
			peasantsSelected = 2
			soldiersSelected = 2
			totalPoints += peasantsSelected*pieces[peasant].Value + soldiersSelected*pieces[soldier].Value
			fmt.Println("Due to picking 4 pawns you will automatically get 2 peasants and 2 soldiers")
			break
		} else if rank1 > pieces[pawn].MaxPieces/2 {
			peasantMin := max(0, pieces[pawn].MaxPieces-(rank1+soldierLimit))
			peasantLimit := min(2, pieces[pawn].MaxPieces-rank1)
			fmt.Printf("You can select from %d to %d peasants\n", peasantMin, peasantLimit)
			peasantsSelected = checkForValidClassicalInput("How many peasants would you like: ", peasantMin, peasantLimit)
			totalPoints += peasantsSelected * pieces[peasant].Value
			rank1 += peasantsSelected
			if rank1 == pieces[pawn].MaxPieces {
				fmt.Println("You have picked the maximum number of pieces")
				break
			} else if rank1 == pieces[pawn].MaxPieces-2 {
				fmt.Println("You automatically get 2 soldiers")
				soldiersSelected = 2
				totalPoints += soldiersSelected * pieces[soldier].Value
				rank1 += soldiersSelected
			} else if rank1 == pieces[pawn].MaxPieces-1 {
				fmt.Println("You automatically get 1 soldier")
				soldiersSelected = 1
				totalPoints += soldiersSelected * pieces[soldier].Value
				rank1 += soldiersSelected
			}
		}
		fmt.Printf("You selected %d Pawns, %d Peasants, %d Soldiers\n", pawn, peasant, soldier)
	}
	return totalPoints
}

func calculateRankII(maxPoints, totalPoints int) int {
	//Creating a slice (or a list) in Go
	rank2Pieces := []string{rook, knight, bishop, catapult, chamberlain, courtesan, herald, inquisitor, lancer, pontiff, thief, tower}
	rank2 := 0
	rank2StartingPoints := totalPoints

	for {
		rank2 = 0
		rank2StartingPoints = totalPoints
		//when iterating through slice there are two return values, the index and the actual element
		for _, piece := range rank2Pieces {
			maxPieces := pieces[piece].MaxPieces
			if maxPieces == 2 {
				selectedPieces := checkForValidClassicalInput("How many "+piece+" would you like: ", pieces[piece].MinPieces, maxPieces)
				rank2StartingPoints += selectedPieces * pieces[piece].Value
				rank2 += selectedPieces
				if rank2 == 6 {
					break
				}
			} else if maxPieces == 1 {
				fmt.Println("\nFor the rest of these insert 'y' for yes and 'n' for no")
				selectedPiece := checkForValidFaerieInput("Would you like a " + piece + "?")
				rank2StartingPoints += selectedPiece * pieces[piece].Value
				rank2 += selectedPiece
				if rank2 == 6 {
					break
				}
			}
		}
		if rank2 == 6 {
			return rank2StartingPoints
		} else {
			fmt.Println("You only selected", rank2, "pieces. Be sure to select 6 rank II pieces. The count will now reset!")
		}
	}
}

func calculateRankIII(maxPoints, totalPoints int) int {
	rank3 := 0
	currentPoints := totalPoints

	for {
		fmt.Println("\nFor these insert 'y' for yes and 'n' for no")
		fmt.Println("Would you like a queen?")
		var queens string
		fmt.Scanln(&queens)

		if strings.ToLower(queens) == "y" {
			queenSelected := 1
			currentPoints += queenSelected * pieces[queen].Value
		} else if strings.ToLower(queens) == "n" {
			fmt.Println("Since you did not select a queen, you automatically get a jester")
			jesterSelected := 1
			currentPoints += jesterSelected * pieces[jester].Value
		} else {
			fmt.Println("Invalid input. Please select 'y' or 'n'")
			continue
		}

		if maxPoints-currentPoints <= pieces[regent].Value {
			fmt.Println("You do not have enough points for a regent, so you will have to get a king.")
			kingSelected := 1
			currentPoints += kingSelected * pieces[king].Value
			rank3 += kingSelected
			break
		}

		fmt.Println("Would you like a King?")
		var kings string
		fmt.Scanln(&kings)
		if strings.ToLower(kings) == "y" {
			kingSelected := 1
			currentPoints += kingSelected * pieces[king].Value
		} else if strings.ToLower(kings) == "n" {
			fmt.Println("Since you did not select a king, you automatically get a regent")
			regentSelected := 1
			currentPoints += regentSelected * pieces[regent].Value
			break
		} else {
			fmt.Println("Invalid input. Please select 'y' or 'n'")
			continue
		}
	}

	return currentPoints
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
