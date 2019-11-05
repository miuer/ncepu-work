/*
 * Revision History:
 *     Initial: 2019/11/05       Miuer
 *     Test environment:         go1.13.3
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
		if len(optimal) < 20 {
			next := selection(fitness(initial(parent)))
			optimal = append(optimal, next...)
			fmt.Printf("The %d generation: ", count)
			fmt.Println(optimal)
			count++
		} else {
			break
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

	seed()
	mut := rand.Intn(3)

	if mut > 0 {
		seed()
		for i := 0; i < mut; i++ {
			num := rand.Intn(5)
			data[num] = round(rand.Float64())
		}
	}

	return data
}
