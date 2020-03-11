package utils

func intersect(m, n []*Room) bool {
	for i := 0; i < len(m); i++ {
		for j := i + 1; j < len(n); j++ {
			// fmt.Println("inter:", m[i].Name, n[j].Name)
			if m[i].Name == n[j].Name {
				return true
			}
		}
	}
	return false
}

func intersectInside(m []*Room, comb [][]*Room) bool {
	for j := 0; j < len(comb); j++ {
		for i := 0; i < len(m); i++ {
			for k := 0; k < len(comb[j]); k++ {
				if m[i] == comb[j][k] {
					return true
				}
			}
		}
	}
	return false
}

// Find all possible combinations not intersecting paths
func FindComb(p [][]*Room) [][][]*Room {
	for i := 0; i < len(p); i++ {
		var comb [][]*Room
		comb = append(comb, p[i])
		for j := i + 1; j < len(p); j++ {
			if !intersect(p[i][:len(p[i])-1], p[j][:len(p[j])-1]) &&
				!intersectInside(p[j][:len(p[j])-1], comb) {
				comb = append(comb, p[j])
			}
		}
		Ret = append(Ret, comb)
	}
	return Ret
}
