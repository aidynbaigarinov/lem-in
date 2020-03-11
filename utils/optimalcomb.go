package utils

// Find the optimal combinations with least steps,
// depending on number of ants
func OptimalComb(c [][][]*Room, num int) [][]*Room {
	m := make(map[int][][]*Room)
	var ret [][]*Room
	var tmp int
	for _, comb := range c {
		min := len(comb[0])
		max := len(comb[len(comb)-1])
		numPath := len(comb)
		areaEmpty := 0
		for _, p := range comb {

			if len(p) == 1 {
				ret = append(ret, p)
				return ret
			}
			areaEmpty += (max - len(p))
		}
		antsLeft := num - areaEmpty

		min += (antsLeft / numPath) + (antsLeft % numPath)
		m[min] = comb
		tmp = min
	}
	for k := range m {
		if k <= tmp {
			ret = m[k]
			tmp = k
		}
	}
	return ret
}
