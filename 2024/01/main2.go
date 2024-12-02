package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {

    var similarity int = 0
    var leftList []int
    var rightList []int

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Fields(line)
        num1, _ := strconv.Atoi(parts[0])
        num2, _ := strconv.Atoi(parts[1])

        leftList = append(leftList, num1)
        rightList = append(rightList, num2)
    }

    occurrences := make(map[int]int)
    for _, num := range rightList {
        occurrences[num]++
    }

    for _, num := range leftList {
        similarity += num * occurrences[num]
    }

    fmt.Println(similarity)
}
