package main

import (
	"fmt"

	"github.com/miuer/ncepu-work/architecture/Golang-AStar/term"
	"github.com/miuer/ncepu-work/architecture/Golang-AStar/utils"
)

// Scene -
type Scene struct {
	rows, cols int
	scene      [][]byte
}

func (s *Scene) initScene(rows int, cols int) {
	s.rows = rows
	s.cols = cols

	s.scene = make([][]byte, s.rows)
	for i := 0; i < s.rows; i++ {
		s.scene[i] = make([]byte, s.cols)
		for j := 0; j < s.cols; j++ {
			if i == 0 || i == s.rows-1 || j == 0 || j == s.cols-1 {
				s.scene[i][j] = '#'
			} else {
				s.scene[i][j] = ' '
			}
		}
	}
}

func (s *Scene) draw(colour string) {
	for i := 0; i < s.rows; i++ {
		for j := 0; j < s.cols; j++ {
			var color string
			switch s.scene[i][j] {
			case '#':
				color = term.FgCyan
			case 'A':
				color = term.FgRed
			case 'B':
				color = term.FgBlue
			case '*':
				color = colour
			case 'w':
				color = term.FgCyan
			}
			fmt.Printf("%s%c%s", color, s.scene[i][j], term.Reset)
		}
		fmt.Printf("\n")
	}
}

func (s *Scene) addWalls(num int) {
	for i := 0; i < num; i++ {
		ori := utils.GetRandInt(2)
		length := utils.GetRandInt(16) + 1
		row := utils.GetRandInt(s.rows)
		col := utils.GetRandInt(s.cols)
		switch ori {
		case 0:
			// Landscape
			for i := 0; i < length; i++ {
				if col+i >= s.cols {
					break
				}
				s.scene[row][col+i] = 'w'
			}

		case 1:
			// portrait
			for i := 0; i < length; i++ {
				if row+i >= s.rows {
					break
				}
				s.scene[row+i][col] = 'w'
			}
		}
	}
}