package main

import (
	"fmt"
)

func main() {
	const (
		width  = 50
		height = 10

		empty = ' '
		ball  = '⚾'
	)

	// initialize board with empty cells
	board := make([][]rune, height)
	for y := 0; y < height; y++ {
		board[y] = make([]rune, width)
		for x := 0; x < width; x++ {
			board[y][x] = empty
		}
	}

	// draw a smiley (eyes, nose, mouth)
	points := [][2]int{
		{2, 12}, {2, 16}, // eyes
		{4, 14},          // nose
		{6, 10}, {6, 18}, // mouth corners
		{7, 12}, {7, 14}, {7, 16}, // mouth bottom
	}

	for _, p := range points {
		y, x := p[0], p[1]
		board[y][x] = ball
	}

	// print the board
	for _, row := range board {
		for _, c := range row {
			fmt.Print(string(c), " ")
		}
		fmt.Println()
	}
}
