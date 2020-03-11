package utils

import (
	"strings"
)

// Checks if an argument is a start room
func IsStart(s string) bool {
	return s == "##start"
}

// Checks if an argument is an end room
func IsEnd(s string) bool {
	return s == "##end"
}

// Check if an argument is a valid room
func IsRoom(s string) string {
	tmp := strings.Split(s, " ")
	return tmp[0]
}

// Check if an argument is a valid connection between the rooms & returns the array of rooms
func IsConn(s string) []string {
	tmp := strings.Split(s, "-")
	return tmp
}

func Check(a []string) ([]string, []string) {
	var (
		aRoom []string
		aConn []string
	)

	for i, v := range a {
		if len(v) == 0 {
			continue
		}
		if !IsStart(v) && !IsEnd(v) && len(strings.Split(v, " ")) == 3 {
			if v[0] == '#' || v[0] == 'L' {
				ErrHandler()
			}
		}
		if IsStart(v) && i < len(a)-1 && len(strings.Split(a[i+1], " ")) != 3 ||
			IsEnd(v) && i < len(a)-1 && len(strings.Split(a[i+1], " ")) != 3 {
			ErrHandler()
		}
		if IsStart(v) || IsEnd(v) || len(strings.Split(v, " ")) == 3 {
			aRoom = append(aRoom, v)
			continue
		}
		if len(strings.Split(v, "-")) == 2 {
			aConn = append(aConn, v)
			continue
		}
		if len(strings.Split(v, " ")) > 3 {
			ErrHandler()
		}
	}
	return aRoom, aConn
}
