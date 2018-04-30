package trip

import (
	"fmt"
	"strings"
)

// Print all subtrips
func Print(trips []SubTrip) {
	for _, t := range trips {
		if t.Type != "WALK" {
			fmt.Printf("\t%s %s -> %s %s (%s towards %s)\n",
				t.DepTime, t.From, t.To, t.ArrTime, myTitleCase(t.Type), t.Direction)
		}
	}
}

func myTitleCase(s string) string {
	lower := strings.ToLower(s)
	return strings.Title(lower)
}
