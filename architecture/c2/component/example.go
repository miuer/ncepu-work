package component

import "encoding/json"

// ExampleAddFirst -
func ExampleAddFirst() {
	js := `[
		{
			"Operator": "add",
			"Arg": [
				122,
				2
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
				2.5
			]
		},
		{
			"Operator": "pow",
			"Arg": [
				2
			]
		}
	]`

	var c2 C2
	c2.Result = 0

	err := json.Unmarshal([]byte(js), &c2.Calculation)
	if err != nil {
		panic(err)
	}

	ch := make(chan float64)
	defer close(ch)

	c2.Run(ch)
}

// ExampleSubFirst -
func ExampleSubFirst() {
	js := `[
		{
			"Operator": "sub",
			"Arg": [
				122,
				2
			]
		},
		{
			"Operator": "add",
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
				2.5
			]
		},
		{
			"Operator": "pow",
			"Arg": [
				2
			]
		}
	]`

	var c2 C2
	c2.Result = 0

	err := json.Unmarshal([]byte(js), &c2.Calculation)
	if err != nil {
		panic(err)
	}

	ch := make(chan float64)
	defer close(ch)

	c2.Run(ch)
}

// ExampleSinFirst -
func ExampleSinFirst() {
	js := `[
		{
			"Operator": "sin",
			"Arg": [
				122
			]
		},
		{
			"Operator": "add",
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
			"Operator": "add",
			"Arg": [
				2.5
			]
		},
		{
			"Operator": "pow",
			"Arg": [
				2
			]
		}
	]`

	var c2 C2
	c2.Result = 0

	err := json.Unmarshal([]byte(js), &c2.Calculation)
	if err != nil {
		panic(err)
	}

	ch := make(chan float64)
	defer close(ch)

	c2.Run(ch)
}

// ExamplePowFirst -
func ExamplePowFirst() {
	js := `[
		{
			"Operator": "pow",
			"Arg": [
				2,
				3
			]
		},
		{
			"Operator": "add",
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
			"Operator": "add",
			"Arg": [
				2.5
			]
		},
		{
			"Operator": "add",
			"Arg": [
				2
			]
		}
	]`

	var c2 C2
	c2.Result = 0

	err := json.Unmarshal([]byte(js), &c2.Calculation)
	if err != nil {
		panic(err)
	}

	ch := make(chan float64)
	defer close(ch)

	c2.Run(ch)
}
