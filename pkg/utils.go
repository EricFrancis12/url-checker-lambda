package pkg

import (
	"math/rand"
)

func Dedupe[C comparable](items []C) []C {
	uniqueItems := make(map[C]struct{})
	var result []C

	for _, item := range items {
		if _, exists := uniqueItems[item]; !exists {
			uniqueItems[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

func filter[C comparable](items []C, predicate func(C) bool) []C {
	var result []C
	for _, c := range items {
		if predicate(c) {
			result = append(result, c)
		}
	}
	return result
}

func includes[C comparable](items []C, item C) bool {
	for _, c := range items {
		if c == item {
			return true
		}
	}
	return false
}

func mustGetRand[T any](items []T) T {
	index := rand.Intn(len(items))
	return items[index]
}
