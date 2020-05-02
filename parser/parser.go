package parser

import (
	"io/ioutil"
	"strings"
)

type Letter struct {
	Title string
	Body  []string
}

func Parse(path string) ([]Letter, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return []Letter{}, nil
	}

	return letters(removeEmptyAndTrim(strings.Split(string(content), "\n"))), nil
}

func removeEmptyAndTrim(lines []string) []string {
	result := []string{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		result = append(result, strings.TrimSpace(line))
	}

	return result
}

func letters(lines []string) []Letter {
	result := []Letter{}
	for _, line := range lines {
		if isTitle(line) {
			result = append(result, Letter{})
			result[len(result)-1].Title = line
		} else {
			result[len(result)-1].Body = append(result[len(result)-1].Body, line)
		}
	}

	return result
}

func isTitle(line string) bool {
	return strings.HasPrefix(line, "Letter")
}
