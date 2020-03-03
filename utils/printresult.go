package utils

import "fmt"

// Print movements of ants
func PrintResult(a []Ant, end string) {
	for !a[len(a)-1].finish {
		for i, l := 0, len(a); i < l; i++ {
			if !a[i].finish {
				// * If the room name is equal to end room name
				// * then print result and mark ant as finished
				if a[i].route.route[a[i].roomNum].Name == end {
					fmt.Printf("L%d-%s ", a[i].ID, a[i].route.route[a[i].roomNum].Name)
					if a[i].roomNum > 0 {
						a[i].route.route[a[i].roomNum-1].busy = false
					} else {
						a[i].route.route[a[i].roomNum].busy = false
					}
					a[i].finish = true
					// * if room is not busy, print result, make room busy and increment room's number
				} else if !a[i].route.route[a[i].roomNum].busy {
					fmt.Printf("L%d-%s ", a[i].ID, a[i].route.route[a[i].roomNum].Name)
					a[i].route.route[a[i].roomNum].busy = true
					if a[i].roomNum > 0 {
						a[i].route.route[a[i].roomNum-1].busy = false
					}
					a[i].roomNum++
				}
			}
		}
		fmt.Println()
	}
}
