package component

import "log"

// C2 -
// Calculation is the component used for calculation
// Result is used to store the intermediate or final value calculated by the component
type C2 struct {
	Calculation []*Calculate
	Result      float64
}

// Run -
func (c2 *C2) Run(ch chan float64) {

	for _, cal := range c2.Calculation {

		// create a calculation goroutinue
		go func() {
			log.Print("the calcution is:", *cal)
			ch <- cal.Run(c2.Result)
		}()

		// The main process captures the value of calculation component
		// then create a new calculation goroutinue
		c2.Result = <-ch
		log.Println("the result is:", c2.Result)

	}

}
