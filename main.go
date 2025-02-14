package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

func readInput(inputFile string) ([]string, error) {
	var lines []string

	if inputFile != "" {
		f, err := os.Open(inputFile)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	return lines, nil
}

func parseSubdomain(subdomain string) []string {
	var result []string
	parts := strings.Split(subdomain, ".")
	for _, part := range parts {
		result = append(result, part)

		if strings.Contains(part, "-") {
			subParts := strings.Split(part, "-")
			result = append(result, subParts...)
		}
	}
	return result
}

type wordCount struct {
	word  string
	count int
}

func processWords(words []string, limit int) []wordCount {
	counts := make(map[string]int)
	for _, w := range words {
		if w != "" {
			counts[w]++
		}
	}

	var result []wordCount
	for word, count := range counts {
		result = append(result, wordCount{word: word, count: count})
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].count == result[j].count {
			return result[i].word < result[j].word
		}
		return result[i].count > result[j].count
	})

	if limit > 0 && limit < len(result) {
		result = result[:limit]
	}

	return result
}

func main() {
	inputFile := flag.String("i", "", "input file with subdomains")
	limit := flag.Int("limit", 0, "limit output to top N most frequent words (0 for all)")
	flag.Parse()

	lines, err := readInput(*inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	var allWords []string
	for _, l := range lines {
		allWords = append(allWords, parseSubdomain(l)...)
	}

	wordCounts := processWords(allWords, *limit)

	for _, wc := range wordCounts {
		fmt.Println(wc.word)
	}
}
