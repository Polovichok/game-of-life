package main

import (
	"math/rand"
)

func createArray() [width][length]int {
	var array = [width][length]int{}
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			array[i][j] = 0
		}
	}

	return array
}

func fillingArray(array *[width][length]int) {
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			array[i][j] = rand.Intn(2)
		}
	}
}

func fillingNewArray(array *[width][length]int) [width][length]int {

	for i := 1; i < length-1; i++ {
		for j := 1; j < width-1; j++ {
			sum := checkNeighbors(array, i, j)
			if sum == 3 && array[i][j] == 0 {
				array[i][j] = 1
			} else if sum == 2 || sum == 3 && array[i][j] == 1 {
				array[i][j] = 1
			} else {
				array[i][j] = 0
			}
		}
	}
	return *array
}

func checkNeighbors(array *[width][length]int, i int, j int) int {
	count := 0
	for k := i - 1; k <= i+1; k++ {
		for p := j - 1; p <= j+1; p++ {
			if k >= 0 && k < width && p >= 0 && p < length {
				if array[k][p] == 1 {
					count += 1
				}
			}
		}
	}
	count -= array[i][j]
	return count
}
