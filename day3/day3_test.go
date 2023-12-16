package day3

import "testing"

func TestIsTouchingSymbol(t *testing.T) {
	tests := []struct {
		name       string
		grid       [][]string
		row        int
		col        int
		isTouching bool
	}{
		{
			name: "no symbol",
			grid: [][]string{
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", "8", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
			},
			row:        2,
			col:        2,
			isTouching: false,
		},
		{
			name: "top left",
			grid: [][]string{
				{".", ".", ".", ".", "."},
				{".", "*", ".", ".", "."},
				{".", ".", "8", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
			},
			row:        2,
			col:        2,
			isTouching: true,
		},
		{
			name: "top",
			grid: [][]string{
				{".", ".", ".", ".", "."},
				{".", ".", "*", ".", "."},
				{".", ".", "8", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
			},
			row:        2,
			col:        2,
			isTouching: true,
		},
		{
			name: "top right",
			grid: [][]string{
				{".", ".", ".", ".", "."},
				{".", ".", ".", "*", "."},
				{".", ".", "8", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
			},
			row:        2,
			col:        2,
			isTouching: true,
		},
		{
			name: "left",
			grid: [][]string{
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
				{".", "*", "8", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
			},
			row:        2,
			col:        2,
			isTouching: true,
		},
		{
			name: "right",
			grid: [][]string{
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", "8", "*", "."},
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
			},
			row:        2,
			col:        2,
			isTouching: true,
		},
		{
			name: "bottom left",
			grid: [][]string{
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", "8", ".", "."},
				{".", "*", ".", ".", "."},
				{".", ".", ".", ".", "."},
			},
			row:        2,
			col:        2,
			isTouching: true,
		},
		{
			name: "bottom",
			grid: [][]string{
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", "8", ".", "."},
				{".", ".", "*", ".", "."},
				{".", ".", ".", ".", "."},
			},
			row:        2,
			col:        2,
			isTouching: true,
		},
		{
			name: "bottom right",
			grid: [][]string{
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", "8", ".", "."},
				{".", ".", ".", "*", "."},
				{".", ".", ".", ".", "."},
			},
			row:        2,
			col:        2,
			isTouching: true,
		},
		{
			name: "all",
			grid: [][]string{
				{".", ".", ".", ".", "."},
				{".", "*", "*", "*", "."},
				{".", "*", "8", "*", "."},
				{".", "*", "*", "*", "."},
				{".", ".", ".", ".", "."},
			},
			row:        2,
			col:        2,
			isTouching: true,
		},
		{
			name: "all walls",
			grid: [][]string{
				{"8"},
			},
			row:        0,
			col:        0,
			isTouching: false,
		},
		{
			name: "left wall",
			grid: [][]string{
				{".", "."},
				{"8", "*"},
				{".", "."},
			},
			row:        1,
			col:        0,
			isTouching: true,
		},
		{
			name: "right wall",
			grid: [][]string{
				{".", "."},
				{"*", "8"},
				{".", "."},
			},
			row:        1,
			col:        1,
			isTouching: true,
		},
		{
			name: "top wall",
			grid: [][]string{
				{".", "8", "."},
				{".", ".", "*"},
			},
			row:        0,
			col:        1,
			isTouching: true,
		},
		{
			name: "bottom wall",
			grid: [][]string{
				{".", ".", "."},
				{".", "8", "*"},
			},
			row:        1,
			col:        1,
			isTouching: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isTouching := isTouching(test.grid, test.row, test.col, isSymbol)
			if isTouching != test.isTouching {
				t.Errorf("Expected isTouching to be %v, got %v", test.isTouching, isTouching)
			}
		})
	}
}
