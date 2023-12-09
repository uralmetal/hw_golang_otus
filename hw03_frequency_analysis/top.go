package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

func normalizeWord(word string) string {
	/*
		1. Приводим в к одному регистру
		2. Очищаем от знаков препинания в конце и начале
		3.
	*/
	return word
}

func Top10(text string) []string {
	frequency := make(map[string]int)
	var top []Pair
	var topWords []string
	var topLen int

	words := strings.Fields(text)
	for _, word := range words {
		normalizedWord := normalizeWord(word)
		if len(normalizedWord) > 0 {
			frequency[normalizedWord]++
		}
	}
	for key, value := range frequency {
		top = append(top, Pair{Key: key, Value: value})
	}
	sort.Slice(top, func(i, j int) bool {
		if top[i].Value != top[j].Value {
			return top[i].Value > top[j].Value
		} else {
			return top[i].Key < top[j].Key
		}
	})
	if len(top) >= 10 {
		topLen = 10
	} else {
		topLen = len(top)
	}
	for _, word := range top[:topLen] {
		topWords = append(topWords, word.Key)
	}
	return topWords
}
