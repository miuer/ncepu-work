package main

import (
	"sync"
	"time"

	"github.com/miuer/ncepu-work/architecture/Golang-AStar/term"
)

func main() {
	var (
		scene Scene
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	scene.initScene(23, 70)
	scene.addWalls(10)
	initAstar(&scene)

	find := func(colour string) {
		mutex.Lock()

		defer wg.Done()

		for findPath(&scene) {
			scene.draw(colour)
			time.Sleep(50 * time.Millisecond)
		}

		scene.draw(colour)
		ResetOpenList()
		mutex.Unlock()
	}

	wg.Add(3)
	go find(term.FgYellow)
	go find(term.FgBlack)
	go find(term.FgMagenta)

	wg.Wait()
}
