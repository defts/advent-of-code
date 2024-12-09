package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

// Processus d'entrée pour créer le disque
func processInput(input []int) []interface{} {
	var storage []interface{}
	fileId := 0

	for i, blockCount := range input {
		isFile := i%2 == 0
		for b := 0; b < blockCount; b++ {
			if isFile {
				storage = append(storage, fileId)
			} else {
				storage = append(storage, ".")
			}
		}
		if isFile {
			fileId++
		}
	}
	return storage
}

// Trouve le prochain bloc libre pour un fichier unique
func findNextFreeBlock(storage []interface{}) int {
	nextFreeBlock := 0
	for i := nextFreeBlock; i < len(storage); i++ {
		if storage[i] == "." {
			return i
		}
	}
	return -1
}

// Compression simple : déplace les fichiers vers les premières cases libres
func compress(storage []interface{}) {
	for i := len(storage) - 1; i >= 0; i-- {
		if _, isFile := storage[i].(int); isFile {
			nextFreeBlock := findNextFreeBlock(storage)
			if nextFreeBlock > i {
				return // Rien à déplacer
			}
			storage[nextFreeBlock] = storage[i]
			storage[i] = "."
		}
	}
}

// Renvoie le début du fichier se terminant à `end`
func getFileStart(storage []interface{}, end int) int {
	fileID := storage[end]
	for i := end; i > 0; i-- {
		if storage[i-1] != fileID {
			return i
		}
	}
	return 0
}

// Compression avec défragmentation
func defragmentedCompress(storage []interface{}) {
	for i := len(storage) - 1; i >= 0; i-- {
		if _, isFile := storage[i].(int); isFile {
			start := getFileStart(storage, i)
			size := (i - start) + 1
			newStart := findNextFreeBlockBySize(storage, size)
			if newStart != -1 && newStart < start {
				moveFile(storage, start, newStart, size)
			}
			i = start // Sauter au début du fichier traité
		}
	}
}

// Trouve le prochain bloc libre d'une taille donnée
func findNextFreeBlockBySize(storage []interface{}, size int) int {
	currentBlockSize := 0
	for i := 0; i < len(storage); i++ {
		if storage[i] == "." {
			currentBlockSize++
			if currentBlockSize == size {
				return i - size + 1
			}
		} else {
			currentBlockSize = 0
		}
	}
	return -1
}

// Déplace un fichier de `start` à `newStart` sur `size` blocs
func moveFile(storage []interface{}, start, newStart, size int) {
	for i := 0; i < size; i++ {
		storage[newStart+i] = storage[start+i]
		storage[start+i] = "."
	}
}

// Partie 1 : Compression simple
func part1(diskMap []int) int {
	defer utils.Timer("part1")()

	storage := processInput(diskMap)
	compress(storage)

	sum := 0
	for i, block := range storage {
		if fileID, isFile := block.(int); isFile {
			sum += i * fileID
		}
	}
	return sum
}

// Partie 2 : Compression avec défragmentation
func part2(diskMap []int) int {
	defer utils.Timer("part2")()

	storage := processInput(diskMap)
	defragmentedCompress(storage)

	sum := 0
	for i, block := range storage {
		if fileID, isFile := block.(int); isFile {
			sum += i * fileID
		}
	}
	return sum
}

func main() {
	loadFile("input.txt")
	fmt.Println("part1: ", part1(diskMap))
	fmt.Println("part2: ", part2(diskMap))
}
