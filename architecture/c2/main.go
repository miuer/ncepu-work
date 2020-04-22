package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/miuer/ncepu-work/architecture/c2/component"
)

func main() {

	start := time.Now()

	/*
		var ca1 = &component.Calculate{
			Operator: "add",
			Arg:      []float64{122, 2},
		}

		var ca2 = &component.Calculate{
			Operator: "sub",
			Arg:      []float64{100},
		}

		var ca3 = &component.Calculate{
			Operator: "mul",
			Arg:      []float64{4},
		}

		var ca4 = &component.Calculate{
			Operator: "div",
			Arg:      []float64{2.5},
		}

	*/

	var c2 component.C2
	c2.Result = 0

	js := `[
		{
			"Operator": "add",
			"Arg": [
				122,
				3
			]
		},
		{
			"Operator": "sub",
			"Arg": [
				100
			]
		},
		{
			"Operator": "mul",
			"Arg": [
				4
			]
		},
		{
			"Operator": "div",
			"Arg": [
				2.5
			]
		},
		{
			"Operator": "sin",
			"Arg": [
			]
		},
		{
			"Operator": "pow",
			"Arg": [
				2
			]
		}
	]`

	err := json.Unmarshal([]byte(js), &c2.Calculation)
	if err != nil {
		panic(err)
	}

	ch := make(chan float64)
	defer close(ch)

	c2.Run(ch)

	elapsed := time.Since(start)

	log.Println("The total time spent is: ", elapsed)

	time.Sleep(10 * time.Second)
}
