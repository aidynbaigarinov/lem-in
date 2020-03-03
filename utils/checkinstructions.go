package utils

// Checks if an argument is a start room
func IsStart(s string) bool {
	if s[:2] == "##" && s[2:] == "start" {
		return true
	}
	return false
}

// Checks if an argument is an end room
func IsEnd(s string) bool {
	if s[:2] == "##" && s[2:] == "end" {
		return true
	}
	return false
}

// Check if an argument is a valid room
func IsRoom(s string) (string, bool) {
	if s[0] == '#' || s[0] == 'L' {
		return "", false
	}
	for i, v := range s {
		if v == ' ' && i > 0 {
			return s[:i], true
		}
	}
	return "", false
}

// Check if an argument is a valid connection between the rooms & returns the array of rooms
func IsConn(s string) ([]string, bool) {
	var ret []string
	if s[0] == '#' || s[0] == 'L' {
		return nil, false
	}
	for i, v := range s {
		if v == '-' {
			ret = append(ret, s[:i], s[i+1:len(s)])
		}
	}
	return ret, true
}
