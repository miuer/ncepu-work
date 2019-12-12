package main

import (
	"fmt"

	"github.com/miuer/ncepu-work/compier/grammar"

	"github.com/miuer/ncepu-work/compier/lexical"
)

func main() {
	fmt.Println("Hollow World!")

	lexical.As.Analysis("src.txt", "tokens.txt", "symbles.txt")

	grammar.Analysis()
}
