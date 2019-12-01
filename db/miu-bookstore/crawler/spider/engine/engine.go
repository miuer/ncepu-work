package engine

import "sync"

// EngineRuner -
type EngineRuner interface {
	EngineRun(string, *sync.WaitGroup)
}
