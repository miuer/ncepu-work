package component

import (
	"log"
	"math"
)

// Calculate -
type Calculate struct {
	Operator string
	Arg      []float64
}

// add -  len(Arg) is required to be  1 or 2
// if len(cal.Arg) == 2, return cal.Arg[0] + cal.Arg[1] as result
// if len(cal.Arg) == 1, return value + cal.Arg[0] as result
func (cal *Calculate) add(value float64) (result float64) {
	if len(cal.Arg) == 2 {
		return (cal.Arg[0] + cal.Arg[1])
	} else if len(cal.Arg) == 1 {
		return (value + cal.Arg[0])
	}

	log.Fatalln("Please check the data format of the calculation unit")
	return
}

// sub -  len(Arg) is required to be  1 or 2
// if len(cal.Arg) == 2, return cal.Arg[0] - cal.Arg[1] as result
// if len(cal.Arg) == 1, return value - cal.Arg[0] as result
func (cal *Calculate) sub(value float64) (result float64) {
	if len(cal.Arg) == 2 {
		return (cal.Arg[0] - cal.Arg[1])
	} else if len(cal.Arg) == 1 {
		return (value - cal.Arg[0])
	}

	log.Fatalln("Please check the data format of the calculation unit")
	return
}

// mul -  len(Arg) is required to be  1 or 2
// if len(cal.Arg) == 2, return cal.Arg[0] * cal.Arg[1] as result
// if len(cal.Arg) == 1, return value * cal.Arg[0] as result
func (cal *Calculate) mul(value float64) (result float64) {
	if len(cal.Arg) == 2 {
		return (cal.Arg[0] * cal.Arg[1])
	} else if len(cal.Arg) == 1 {
		return (value * cal.Arg[0])
	}

	log.Fatalln("Please check the data format of the calculation unit")
	return
}

// div -  len(Arg) is required to be  1 or 2
// if len(cal.Arg) == 2, return cal.Arg[0] / cal.Arg[1] as result
// if len(cal.Arg) == 1, return value / cal.Arg[0] as result
func (cal *Calculate) div(value float64) (result float64) {
	if len(cal.Arg) == 2 {
		return (cal.Arg[0] / cal.Arg[1])
	} else if len(cal.Arg) == 1 {

		return (value / cal.Arg[0])
	}

	log.Fatalln("Please check the data format of the calculation unit")
	return
}

// sin -  len(Arg) is required to be 0 or 1
// if len(cal.Arg) == 1, return sin(cal.Arg[0]) as result
// if len(cal.Arg) == 0, return sin(value) as result
func (cal *Calculate) sin(value float64) (result float64) {
	if len(cal.Arg) == 1 {
		return math.Sin(cal.Arg[0])
	} else if len(cal.Arg) == 0 {
		return math.Sin(value)
	}

	log.Fatalln("Please check the data format of the calculation unit")
	return
}

// pow len(Arg) is required to be  1 or 2
// if len(cal.Arg) == 2, return cal.Arg[0] ^ cal.Arg[1] as result
// if len(cal.Arg) == 1, return value ^ cal.Arg[0] as result
func (cal *Calculate) pow(value float64) (result float64) {
	if len(cal.Arg) == 2 {
		return math.Pow(cal.Arg[0], cal.Arg[1])
	} else if len(cal.Arg) == 1 {
		return math.Pow(value, cal.Arg[0])
	}

	log.Fatalln("Please check the data format of the calculation unit")
	return
}

// Run - all calculation methods are hidden in Run
// you can get the result by calling cal.Run()
func (cal *Calculate) Run(value float64) (result float64) {
	switch cal.Operator {
	case "add":
		result = cal.add(value)

	case "sub":
		result = cal.sub(value)

	case "mul":
		result = cal.mul(value)

	case "div":
		result = cal.div(value)

	case "sin":
		result = cal.sin(value)

	case "pow":
		return cal.pow(value)
	}

	return result
}
