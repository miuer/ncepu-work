package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Point  -
// X - abscissa of two-dimension array, Y - ordinate of two-dimension array
// H - evaluation value from current node to target node : (absx + absy) * 10 ,
// G - evaluation value from current node to origin node : diagonal - 14, hor or ver - 10
// F - H + G
// Parent -
type Point struct {
	X, Y    int
	H, G, F int
	Parent  *Point
}

func (p Point) String() string {
	return "[" + strconv.Itoa(p.X) + ", " + strconv.Itoa(p.Y) + ", " + strconv.Itoa(p.F) + "]"
}

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Clear -
func Clear() {
	fmt.Printf("\033[100B")
	for i := 0; i < 100; i++ {
		fmt.Printf("\033[1A")
		fmt.Printf("\033[K")
	}
}

// GetRandInt -
func GetRandInt(limit int) int {
	return r.Intn(limit)
}
