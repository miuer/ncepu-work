/*
 * Revision History:
 *    Initial:                 2019/11/05    Miuer
 *    Last modify:             2019/11/05    Miuer
 *    Test environment:        go1.13.3
 */

package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var parent []float64

	fmt.Println("Hollow World!")
	parents := initial(parent)

	fmt.Print("The first generation: ")
	fmt.Println(parents)

	assessment := fitness(parents)
	optimal := selection(assessment)

	if len(optimal) == 0 {
		fmt.Println("The 1 generation is completely eliminated in the first natural selection.")
		return
	} else {
		fmt.Print("The first optimal:")
		fmt.Println(optimal)
	}

	var count = 2

	for {
		var parent []float64
		if len(optimal) < 20 {

			sel := selection(fitness(initial(parent)))
			fmt.Println("selection : ", sel)

			mut := mutation(sel)
			fmt.Println("mutation ", mut)

			cro := crossover(mut)
			fmt.Println("crossover", cro)

			optimal = append(optimal, selection(cro)...)
			if len(selection(cro)) == 0 {
				fmt.Printf("All species eliminated in the %d natural selection\n.", count)
				return
			}
			fmt.Printf("The %d generation: ", count)
			fmt.Println(optimal)
			count++
		} else {
			return
		}
	}
}

func round(f float64) float64 {
	n10 := math.Pow10(5)

	return math.Trunc((f+0.4/n10)*n10) / n10
}

func seed() {
	rand.Seed(time.Now().UnixNano())
}

func initial(parent []float64) []float64 {
	done := make(chan struct{})

	temp := produce(done)

	for i := 0; i < 100; i++ {
		parent = append(parent, round(<-temp))
	}

	close(done)

	return parent
}

func produce(done <-chan struct{}) <-chan float64 {
	data := make(chan float64)

	go func() {
		defer close(data)

		for {
			seed()

			select {
			case <-done:
				return
			case data <- rand.Float64():
			}
		}
	}()

	return data
}

func fitness(data []float64) []float64 {
	var result []float64

	for _, v := range data {
		result = append(result, round(100*v))
	}

	return result
}

func selection(data []float64) []float64 {
	var result []float64

	sort.Float64s(data)

	for _, v := range data {
		if v > 98 {
			result = append(result, v)
		}
	}

	return result
}

func mutation(data []float64) []float64 {
	mut := rand.Intn(3)

	if mut < len(data) {
		if mut > 0 {
			for i := 0; i < mut; i++ {
				seed()
				num := rand.Intn(len(data))
				data[num] = round(rand.Float64() * 100)
			}
		}
	}
	return data
}

func crossover(data []float64) []float64 {
	cro := rand.Intn(3)
	if cro < len(data) {
		if cro > 0 {

			for i := 0; i < cro; i++ {
				seed()
				p := rand.Float64()

				m := rand.Intn(len(data))
				n := rand.Intn(len(data))

				x := data[m]
				y := data[n]

				data[m] = round(x*p + y*(1-p))
				data[n] = round(x*(1-p) + y*p)
			}
		}
	}
	return data
}