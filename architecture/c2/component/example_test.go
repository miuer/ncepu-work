package component_test

import (
	"testing"

	"github.com/miuer/ncepu-work/architecture/c2/component"
)

func TestC2(t *testing.T) {
	t.Run("add", testAddFirst)
	t.Run("sub", testSubFirst)
	t.Run("sin", testSinFirst)
	t.Run("pow", testPowFirst)

}

// example test
func testAddFirst(t *testing.T) {
	component.ExampleAddFirst()
}

func testSubFirst(t *testing.T) {
	component.ExampleSubFirst()
}

func testSinFirst(t *testing.T) {
	component.ExampleSinFirst()
}

func testPowFirst(t *testing.T) {
	component.ExamplePowFirst()
}
