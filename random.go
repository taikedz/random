package main

import (
    "fmt"
    "os"
    "bufio"
    "math/rand/v2"
)

func readLines() ([]string, error) {
    // https://golangdocs.com/golang-read-file-line-by-line
    fileScanner := bufio.NewScanner(os.Stdin)

    fileScanner.Split(bufio.ScanLines)
    var lines []string

    for fileScanner.Scan() {
        lines = append(lines, fileScanner.Text())
    }

    return lines, nil
}

func loadargs() []string {
    var all_terms []string
    var seen_stdin bool = false

    for _,token := range os.Args[1:] {
	    if token == "-" {
                if seen_stdin { continue }
		seen_stdin = true
                items, err := readLines()
                if err != nil { fail(1, err) }
		all_terms = append(all_terms, items...)
	    } else {
                all_terms = append(all_terms, token)
            }
    }
    return all_terms
}

func fail(code int, message any) {
    fmt.Printf("%v\n", message)
    os.Exit(code)
}

func main() {
    tokens := loadargs()
    max := uint32(len(tokens))
    if max == 0 { fail(1, "No tokens to use") }
    idx := rand.Uint32N(max)
    fmt.Printf("%s\n", tokens[idx])
}
