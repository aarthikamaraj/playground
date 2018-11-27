package main

import "fmt"

/* Find largest island in a matrix */

var (
	neighbours = [][]int{
		//{-1,-1},
		{-1, 0},
		//{-1,1},
		{0, -1},
		{0, 1},
		//{1,-1},
		{1, 0},
		//{1,1},
	}
)

func main() {
	var (
		array [][]bool // true - shaded, false - unshaded (using bool saves space, we don't need 32 bits to store 2 values)
	)
	array = readInput()

	maxRegionSize := findLargestAdjacentCells(array)

	fmt.Println("Max Count:", maxRegionSize)

}

func findLargestAdjacentCells(array [][]bool) int {
	var (
		maxRegionSize, curRegionSize int
	)
	visited := make([][]bool, len(array))
	for i := range visited {
		visited[i] = make([]bool, len(array[0]))
	}

	for i, row := range array {
		for j, value := range row {
			if value && !visited[i][j] {
				curRegionSize = 1
				curRegionSize = curRegionSize + DFS(array, i, j, visited)
				if maxRegionSize < curRegionSize {
					maxRegionSize = curRegionSize
				}
				//fmt.Println(curRegionSize)
			}
		}
	}
	//fmt.Println(array, visited, maxRegionSize)
	return maxRegionSize
}

func DFS(array [][]bool, i, j int, visited [][]bool) int {
	var (
		curRegionSize int
	)
	visited[i][j] = true //Since array is always passed by reference, we can safely mark it here
	for k := range neighbours {
		nrow := i + neighbours[k][0]
		ncol := j + neighbours[k][1]
		if nrow < 0 || nrow >= len(array) {
			continue
		}
		if ncol < 0 || ncol >= len(array[i]) {
			continue
		}
		//fmt.Println(nrow, ncol)
		if array[nrow][ncol] && !visited[nrow][ncol] {
			curRegionSize++
			curRegionSize = curRegionSize + DFS(array, nrow, ncol, visited)
		}
	}
	return curRegionSize
}

func readInput() [][]bool {
	var (
		input    string
		array    [][]bool // true - shaded, false - unshaded (using bool saves space, we don't need 32 bits to store 2 values)
		row      *[]bool
		newSlice []bool
	)
	fmt.Scanln(&input)
	for _, c := range input {
		switch c {
		case '{':
			//fmt.Println(i, string(c))
			row = new([]bool)
		case '}':
			array = append(array, *row)
		case '1':
			newSlice = append(*row, true)
			row = &newSlice
		case '0':
			newSlice = append(*row, false)
			row = &newSlice
		}
	}
	return array
}
