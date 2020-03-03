package utils

// Save the path
func SavePath(r *Room, path []*Room) []*Room {
	// * Recursively, through parents, add rooms to the path
	if r.start == true {
		return path
	}
	if r.start != true && r.end != true {
		Vis = append(Vis, r.Name)
	}
	path = SavePath(r.Parent, path)
	path = append(path, r)
	return path
}
