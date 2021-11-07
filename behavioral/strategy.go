/*
Sort function uses defferent comparing strategies
*/
package main

type fn func(a, b int) int

func sort(collection *[]int, comparator fn) {
	var tmp int
	size := len(*collection)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			l, r := &(*collection)[i], &(*collection)[j]
			if comparator(*l, *r) < 0 {
				tmp = *l
				*l = *r
				*r = tmp
			}
		}
	}
}

func naturalOrder(a, b int) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

func reverseOrder(a, b int) int {
	switch {
	case a > b:
		return -1
	case a < b:
		return 1
	default:
		return 0
	}
}
