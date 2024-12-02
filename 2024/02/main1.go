package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {

    scanner := bufio.NewScanner(os.Stdin)

    var safe_reports int = 0
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Fields(line)
        levels := make([]int, len(parts))
        for i, part := range parts {
            level, _ := strconv.Atoi(part)
            levels[i] = level
        }
        if is_report_safe(levels) {
            safe_reports++
        }
    }
    fmt.Println(safe_reports)
}

func is_report_safe(levels []int) bool {
    if len(levels) <= 1 {
        return true
    }
    var increasing int = 1
    if levels[1] - levels[0] < 0 {
        increasing = -1
    }
    for i := range len(levels) - 1 {
        var diff int = (levels[i+1] - levels[i]) * increasing
        if diff > 3 || diff < 1 {
            return false
        }
    }
    return true
}
