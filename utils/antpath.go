package utils

// Assign path to the ant
func AntPath(p []*Path, aN int, arr []string) {
	if len(p) < 1 {
		ErrHandler()
	}
	PrintInstructions(arr)
	var ants = []Ant{}
	for i := 0; i < aN; i++ {
		ants = append(ants, Ant{ID: i + 1})
	}
	ants[0].route = p[0]
	p[0].antCounter++
	currPath := 0
	for i, l := 1, len(ants); i < l; i++ {
		if currPath == len(p)-1 {
			currPath = 0
		}
		if currPath < len(p)-1 {
			if (len(p[currPath].route) + p[currPath].antCounter) >
				(len(p[currPath+1].route) + p[currPath+1].antCounter) {
				currPath++
				ants[i].route = p[currPath]
				p[currPath].antCounter++
			} else {
				ants[i].route = p[currPath]
				p[currPath].antCounter++
			}
		} else {
			ants[i].route = p[currPath]
		}
	}

	z := len(p[0].route) - 1
	end := p[0].route[z].Name
	PrintResult(ants, end)
}
