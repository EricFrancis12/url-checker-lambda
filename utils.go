package main

func sliceIncludes[C comparable](items []C, item C) bool {
	for _, c := range items {
		if c == item {
			return true
		}
	}
	return false
}
