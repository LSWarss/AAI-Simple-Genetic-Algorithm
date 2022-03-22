package algo

import (
	"math/rand"
	"time"

	distances "github.com/lswarss/AAI/files"
)

type CharactersMatrix struct {
	CharactersCount int
	Characters      [][]int
}

func NewCharactersMatrix(distanceMatrix distances.DistanceMatrix) CharactersMatrix {
	rand.Seed(time.Now().UnixNano())
	var characters [][]int

	for i := 0; i < distanceMatrix.Rows; i++ {
		var tempSlice []int
		for j := 0; j < distanceMatrix.Rows; j++ {
			tempSlice = append(tempSlice, j)
		}

		rand.Shuffle(len(tempSlice), func(i, j int) {
			tempSlice[i], tempSlice[j] = tempSlice[j], tempSlice[i]
		})

		characters = append(characters, tempSlice)
		tempSlice = nil
	}

	return CharactersMatrix{
		CharactersCount: distanceMatrix.Rows,
		Characters:      characters,
	}
}

func GetScore(distanceMatrix distances.DistanceMatrix, characterMatrix CharactersMatrix) []int {
	var scores []int

	for i := 0; i < characterMatrix.CharactersCount; i++ {
		var tempSum int
		for j := 0; j < characterMatrix.CharactersCount; j++ {
			characterIndex := characterMatrix.Characters[i][j]
			tempSum += distanceMatrix.Matrix[i][characterIndex]
		}

		scores = append(scores, tempSum)
		tempSum = 0
	}

	return scores
}

func GetBestCharacter(characterMatrix CharactersMatrix, scores []int) [][]int {
	bestScore := scores[0]
	var bestScoreIndex int

	for i := 0; i < len(scores)-1; i++ {
		if scores[i] < bestScore {
			bestScoreIndex = i
			bestScore = scores[i]
		}
	}

	return [][]int{
		characterMatrix.Characters[bestScoreIndex],
		{bestScore},
	}
}

func randIndex(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func GetTournamentSelection(characterMatrix CharactersMatrix) {
	
}

func tournamentSelection(characterMatrix CharactersMatrix, selectivePressure int) int {
	var bestCharacter int
	for i := 0; i < selectivePressure; i++ {
		character := characterMatrix.Characters[randIndex(0, characterMatrix.CharactersCount)][randIndex(0, characterMatrix.CharactersCount)]
		if bestCharacter == 0 || character < bestCharacter {
			bestCharacter = character
		}
	}

	return bestCharacter
}
