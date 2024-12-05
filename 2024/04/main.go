package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    var charMatrix [][]rune
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        characters := []rune(line)
        charMatrix = append(charMatrix, characters)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    fmt.Println(xmas(charMatrix))
}

func xmas(charMatrix [][]rune) int {
    count := 0

    numRows := len(charMatrix)
    numCols := len(charMatrix[0])

    // horizontally
    for _, row := range charMatrix {
        count += count_xmas(row)
    }
    // vertically
    for col :=0; col < numCols; col++ {
        var column []rune
        for row := 0; row < numRows; row++ {
            column = append(column, charMatrix[row][col])
        }
        count += count_xmas(column)
    }

    // ↗️
    for i := 0; i < numRows + numCols; i++ {
        var diagonal []rune
        for j := i; j >=0; j-- {
            if j < numCols && (i-j) < numCols {
                diagonal = append(diagonal, charMatrix[j][i-j])
            }
        }
        count += count_xmas(diagonal)
    }

    // ↘️
    for i := numRows-1; i >= -numCols ; i-- {
        var diagonal []rune
        for j := 0; j <(numRows + numCols); j++ {
            if (i+j) >= 0 && (i+j) < numRows && j >= 0 && j < numCols {
                diagonal = append(diagonal, charMatrix[i+j][j])
            }
        }
        count += count_xmas(diagonal)
    }
    return count
}
func count_xmas(row []rune) int {
    count := 0
    count += strings.Count(string(row), "XMAS")
    count += strings.Count(string(row), "SAMX")
    return count
}
