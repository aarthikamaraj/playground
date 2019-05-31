package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

var neighbors = [][]int32 {
    {0,1},
    {0,-1},
    {1, 0},
    {-1,0},
}

func minimumMoves(grid []string, startX int32, startY int32, goalX int32, goalY int32) int32 {
    if len(grid) == 0 {return -1}
    var visited = make([][]bool, len(grid))
    for i := 0; i < len(grid); i++ {
        visited[i] = make([]bool, len(grid[0]))
    }
    steps, found := dfsRecursion(grid, 0, visited, startX, startY,goalX,goalY)
    fmt.Println(steps, found)
    if !found { return -1}
    return steps
}

func dfsRecursion(grid []string, steps int32, visited[][]bool, i int32, j int32, goalX int32, goalY int32) (int32, bool) {
    var (
        row, col int32
        curLength int32
        found,pathExists bool
    )
    if i == goalX && j == goalY {
        return steps, true
    }
    fmt.Println(i, j, goalX, goalY)
    visited[i][j] = true
    if grid[i][j] == 'X' {
        return -1, false
    }
    var shortestPath = int32(len(grid) * len(grid[i]))
    for _, neighbor := range neighbors {
        row = i + neighbor[0]
        col = j + neighbor[1]
        if row < 0 || col < 0 { continue}
        if int(row) >= len(grid) || int(col) >= len(grid[i]) { continue}
        if !visited[row][col] {
            fmt.Println("DFS:", "i, j", i, j, row, col, "Steps:", steps, curLength)
            curLength, found = dfsRecursion(grid, 1 + steps, visited, row, col, goalX , goalY )
            if found && curLength < shortestPath {
                shortestPath = curLength
                pathExists = true
            }
        }
    }
    return shortestPath, pathExists
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    var grid []string

    for i := 0; i < int(n); i++ {
        gridItem := readLine(reader)
        grid = append(grid, gridItem)
    }

    startXStartY := strings.Split(readLine(reader), " ")

    startXTemp, err := strconv.ParseInt(startXStartY[0], 10, 64)
    checkError(err)
    startX := int32(startXTemp)

    startYTemp, err := strconv.ParseInt(startXStartY[1], 10, 64)
    checkError(err)
    startY := int32(startYTemp)

    goalXTemp, err := strconv.ParseInt(startXStartY[2], 10, 64)
    checkError(err)
    goalX := int32(goalXTemp)

    goalYTemp, err := strconv.ParseInt(startXStartY[3], 10, 64)
    checkError(err)
    goalY := int32(goalYTemp)

    result := minimumMoves(grid, startX, startY, goalX, goalY)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

