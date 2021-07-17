package frontier

import (
	"crawler/internal/components"
	"fmt"
	"testing"
)

func Test_frontier_Push(t *testing.T) {

	frontier := NewFrontier()
	items := []*components.Item{
		{
			Value:    "red",
			Priority: 6,
		},
		{
			Value:    22,
			Priority: 1,
		},
		{
			Value:    "blue",
			Priority: 2,
		},
		{
			Value:    "yellow",
			Priority: 3,
		},
	}

	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "happy case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, i := range items {
				frontier.Push(i)
			}
			for frontier.Len() > 0 {
				fmt.Println(frontier.Pop())
			}
		})
	}
}
