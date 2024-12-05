package main

import (
    "os"
    "bufio"
    "fmt"
    "slices"
    "strconv"
    "strings"
)

func main() {
    orderingRule := true
    rules := make(map[int][]int)
    correct := 0
    ordered := 0

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) == 0 {
            orderingRule = false
            continue
        }
        if orderingRule {
            parts := strings.Split(line, "|")
            before, _ := strconv.Atoi(parts[0])
            after, _ := strconv.Atoi(parts[1])
            rules[before] = append(rules[before], after)
        } else {
            compareFunc := func (a, b int) int {
                if a == b {
                    return 0
                }
                for _, v := range rules[b] {
                    if a == v {
                        return 1
                    }
                }
                return -1
            }
            parts := strings.Split(line, ",")
            var updates []int
            for _, part := range parts {
                update, _ := strconv.Atoi(part)
                updates = append(updates, update)
            }
            if slices.IsSortedFunc(updates, compareFunc) {
                correct += updates[len(updates)/2]
            } else {
                slices.SortFunc(updates, compareFunc)
                ordered += updates[len(updates)/2]
            }
        }
    }
    fmt.Println(correct)
    fmt.Println(ordered)
}
