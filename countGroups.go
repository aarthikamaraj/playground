package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)


/*
 * Complete the 'countGroups' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING_ARRAY related as parameter.
 */

func countGroups(related []string) int32 {
    fmt.Println(related)
    var (
        visited = make([][]bool, len(related))
        count int32 = 0
    )
    for i := 0; i < len(visited); i++ {
        visited[i] = make([]bool, len(related[0]))
    }
    for i, _ := range related {
            if !visited[i][i] {
                DFS(related, i, i, visited)
                count++
                fmt.Println(i, i, count)
            }
        
    }
    return count
}

func DFS(array []string, i, j int, visited[][] bool) {
    visited[i][j] = true
    visited[j][j] = true    
    
    for k := 0; k <= len(array); k++ {
        nrow := k
        ncol := j

        if nrow < 0 || nrow >= len(array) {
            continue
        }
        if ncol < 0 || ncol >= len(array[i]) {
            continue
        }
        if array[nrow][ncol] == '1' && !visited[nrow][ncol] {
            DFS(array, nrow, ncol, visited)
        }
    }
    for k := 0; k <= len(array); k++ {
        nrow := i
        ncol := k
        if nrow < 0 || nrow >= len(array) {
            continue
        }
        if ncol < 0 || ncol >= len(array[i]) {
            continue
        }
        if array[nrow][ncol] == '1' && !visited[nrow][ncol] {
            DFS(array, nrow, ncol, visited)
        }
    }
}


func main() {
