package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var builder strings.Builder
    for scanner.Scan() {
        builder.WriteString(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Error reading input:", err)
    }
    line := builder.String()

    fmt.Println(brrbrr(line, false))
    fmt.Println(brrbrr(line, true))
}

func brrbrr(line string, obeyDont bool) int {
    mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
    var total int = 0

    if obeyDont {
        line = removeDontParts(line)
    }
    matches := mulRegex.FindAllStringSubmatch(line, -1)
    for _, match := range matches {
        x, _ := strconv.Atoi(match[1])
        y, _ := strconv.Atoi(match[2])
        total += x * y
    }
    return total
}

func removeDontParts(input string) string {
    dontRegex := regexp.MustCompile(`don't\(\).*?do\(\)`) // non-greedy
    cleanedInput := dontRegex.ReplaceAllString(input, "")

    // there is no more closing do(), so trim after the 1st occurence of don't()
    index := strings.Index(cleanedInput, "don't()")
    if index == -1 {
        return cleanedInput
    }
    return cleanedInput[:index]
}
