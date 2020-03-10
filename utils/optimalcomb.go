package utils

import "fmt"

func OptimalComb(c [][][]*Room, num int) [][]*Room {
	m := make(map[int][][]*Room)
	var tmp int
	for _, comb := range c {
		min := len(comb[0])
		max := len(comb[len(comb)-1])
		numPath := len(comb)
		// area := numPath * max
		areaEmpty := 0
		for _, p := range comb {
			if len(p) == 1 {
				fmt.Println(p)
			}
			areaEmpty += (max - len(p))
		}
		antsLeft := num - areaEmpty
		// fmt.Println("hop hey", min, max, numPath, area, areaEmpty, antsLeft)

		min += (antsLeft / numPath) + (antsLeft % numPath)
		m[min] = comb
		tmp = min
		// fmt.Println(tmp, min)
	}
	var ret [][]*Room
	for k := range m {
		if k <= tmp {
			ret = m[k]
			tmp = k
		}
	}
	return ret
}
