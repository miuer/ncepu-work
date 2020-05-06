package main

import (
	"fmt"
	"math"
	"os"

	"github.com/miuer/ncepu-work/architecture/Golang-AStar/utils"
)

var (
	origin, dest              utils.Point
	openList, closeList, path []utils.Point
)

// Set the origin point
func setOrig(s *Scene) {
	origin = utils.Point{
		X:      utils.GetRandInt(s.rows-2) + 1,
		Y:      utils.GetRandInt(s.cols-2) + 1,
		H:      0,
		G:      0,
		F:      0,
		Parent: nil,
	}

	if s.scene[origin.X][origin.Y] == ' ' {
		s.scene[origin.X][origin.Y] = 'A'
	} else {
		setOrig(s)
	}
}

// Set the destination point
func setDest(s *Scene) {
	dest = utils.Point{
		X:      utils.GetRandInt(s.rows-2) + 1,
		Y:      utils.GetRandInt(s.cols-2) + 1,
		H:      0,
		G:      0,
		F:      0,
		Parent: nil,
	}

	if s.scene[dest.X][dest.Y] == ' ' {
		s.scene[dest.X][dest.Y] = 'B'
	} else {
		setDest(s)
	}
}

// Init origin, destination. Put the origin point into the openlist by the way
func initAstar(s *Scene) {
	setOrig(s)
	setDest(s)
	openList = append(openList, origin)
}

func findPath(s *Scene) bool {
	current, ok := getFMin()
	if !ok {
		return false
	}

	ok = addToCloseList(current, s)
	if !ok {
		return false
	}

	walkable := getWalkable(current, s)
	for _, p := range walkable {
		addToOpenList(p)
	}

	return true
}

// Get the point with the lowest valuation
func getFMin() (utils.Point, bool) {
	if len(openList) == 0 {
		fmt.Println("No way!!!")
		var p utils.Point
		return p, false
	}

	index := 0
	for i, p := range openList {

		// When multiple F are the same, select the last added point
		if (i > 0) && (p.F <= openList[index].F) {
			index = i
		}
	}
	return openList[index], true
}

// Get the next possible location of the current point
func getWalkable(p utils.Point, s *Scene) []utils.Point {
	var around []utils.Point
	row, col := p.X, p.Y
	left := s.scene[row][col-1]
	up := s.scene[row-1][col]
	right := s.scene[row][col+1]
	down := s.scene[row+1][col]
	leftup := s.scene[row-1][col-1]
	rightup := s.scene[row-1][col+1]
	leftdown := s.scene[row+1][col-1]
	rightdown := s.scene[row+1][col+1]
	if (left == ' ') || (left == 'B') {
		around = append(around, utils.Point{X: row, Y: col - 1, H: 0, G: 0, F: 0, Parent: &p})
	}
	if (leftup == ' ') || (leftup == 'B') {
		around = append(around, utils.Point{X: row - 1, Y: col - 1, H: 0, G: 0, F: 0, Parent: &p})
	}
	if (up == ' ') || (up == 'B') {
		around = append(around, utils.Point{X: row - 1, Y: col, H: 0, G: 0, F: 0, Parent: &p})
	}
	if (rightup == ' ') || (rightup == 'B') {
		around = append(around, utils.Point{X: row - 1, Y: col + 1, H: 0, G: 0, F: 0, Parent: &p})
	}
	if (right == ' ') || (right == 'B') {
		around = append(around, utils.Point{X: row, Y: col + 1, H: 0, G: 0, F: 0, Parent: &p})
	}
	if (rightdown == ' ') || (rightdown == 'B') {
		around = append(around, utils.Point{X: row + 1, Y: col + 1, H: 0, G: 0, F: 0, Parent: &p})
	}
	if (down == ' ') || (down == 'B') {
		around = append(around, utils.Point{X: row + 1, Y: col, H: 0, G: 0, F: 0, Parent: &p})
	}
	if (leftdown == ' ') || (leftdown == 'B') {
		around = append(around, utils.Point{X: row + 1, Y: col - 1, H: 0, G: 0, F: 0, Parent: &p})
	}
	return around
}

// Update the current point status
// if the current point is in closeList then return
// otherwise add it to openList
func addToOpenList(p utils.Point) {
	updateWeight(&p)
	if checkExist(p, closeList) {
		return
	}
	if !checkExist(p, openList) {
		openList = append(openList, p)
	} else {

		// --------------------
		if openList[findPoint(p, openList)].F > p.F {
			//New path found
			openList[findPoint(p, openList)].Parent = p.Parent
		}
	}
}

// Update G, H, F of the point
func updateWeight(p *utils.Point) {
	if checkRelativePos(*p) == 1 {
		p.G = p.Parent.G + 10
	} else {
		p.G = p.Parent.G + 14
	}
	absx := (int)(math.Abs((float64)(dest.X - p.X)))
	absy := (int)(math.Abs((float64)(dest.Y - p.Y)))
	p.H = (absx + absy) * 10
	p.F = p.G + p.H
}

// removeFromOpenList -
func removeFromOpenList(p utils.Point) {
	index := findPoint(p, openList)
	if index == -1 {
		os.Exit(0)
	}
	openList = append(openList[:index], openList[index+1:]...)
}

// Remove the current point from OpenList
// Check whether the current point is the target node
// Change the current point status and add it to CloseList
func addToCloseList(p utils.Point, s *Scene) bool {
	removeFromOpenList(p)
	if (p.X == dest.X) && (p.Y == dest.Y) {
		generatePath(p, s)
		// s.draw()
		return false
	}

	if s.scene[p.X][p.Y] != 'A' {
		s.scene[p.X][p.Y] = '·'
	}
	closeList = append(closeList, p)
	return true
}

// Check whether point is in openList or closeList
func checkExist(p utils.Point, arr []utils.Point) bool {
	for _, point := range arr {
		if p.X == point.X && p.Y == point.Y {
			return true
		}
	}
	return false
}

// findPoint - return index of the point in the list
func findPoint(p utils.Point, arr []utils.Point) int {
	for index, point := range arr {
		if p.X == point.X && p.Y == point.Y {
			return index
		}
	}

	return -1
}

// If point and parent are on the diagonal then return 2, otherwise return 1
func checkRelativePos(p utils.Point) int {
	parent := p.Parent
	hor := (int)(math.Abs((float64)(p.X - parent.X)))
	ver := (int)(math.Abs((float64)(p.Y - parent.Y)))
	return hor + ver
}

// generatePath -
func generatePath(p utils.Point, s *Scene) {
	if (s.scene[p.X][p.Y] != 'A') && (s.scene[p.X][p.Y] != 'B') {
		s.scene[p.X][p.Y] = '*'
	}
	if p.Parent != nil {
		generatePath(*(p.Parent), s)
	}
}

// ClearOpenList - reset openList
func ResetOpenList() {
	var clearList []utils.Point
	openList = clearList
	openList = append(openList, origin)
}
