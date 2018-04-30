package trip

import (
	"fmt"
	"strings"
)

// Print all subtrips
func Print(alternatives [][]SubTrip) {
	for i, alt := range alternatives {
		fmt.Printf("%d.", i+1)

		for _, t := range alt {
			if t.Type != "WALK" {
				fmt.Printf("\t%s %s -> %s %s (%s towards %s)\n",
					t.DepTime, t.From, t.To, t.ArrTime, myTitleCase(t.Type), t.Direction)
			}
		}
	}
}

// PrintCompact outputs the alternatives in a compact way
func PrintCompact(alternatives [][]SubTrip) {
	for i, alt := range alternatives {
		fmt.Printf("%d.", i+1)

		lastSub := len(alt) - 1
		for j, t := range alt {
			if t.Type == "WALK" {
				continue
			}

			fmt.Printf("\t%s %s ->", t.DepTime, t.From)
			if j == lastSub {
				fmt.Printf(" %s %s", t.To, t.ArrTime)
			}
			fmt.Println()
		}
	}
}

func myTitleCase(s string) string {
	lower := strings.ToLower(s)
	return strings.Title(lower)
}
