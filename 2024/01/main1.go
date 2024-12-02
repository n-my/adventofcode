package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "sort"
)

func main() {

    var distance int = 0
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

    sort.Ints(leftList)
    sort.Ints(rightList)

    for i := 0; i < len(leftList); i++ {
        var diff int = leftList[i] - rightList[i]
        if diff < 0 {
            diff = - diff
        }
        distance += diff
    }

    fmt.Println(distance)
}
