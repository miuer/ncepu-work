package main

import (
	"fmt"

	"github.com/miuer/ncepu-work/compier/grammar"
)

func main() {
	fmt.Println("Hollow World!")

	var str map[byte][]string
	str = make(map[byte][]string)

	str['E'] = []string{"TG"}
	str['G'] = []string{"+TG", "-TG", "@"}
	str['T'] = []string{"FS"}
	str['S'] = []string{"*FS", "/FS", "@"}
	str['F'] = []string{"(E)", "i"}

	fis := grammar.GetFirstFrom(str)

	fmt.Println(fis.String())

	fos := grammar.GetFollowFrom(str, 'E', fis)

	fmt.Println(fos.String())

	table := grammar.GetChartFrom(fis, fos, str)
	fmt.Println(table.String())

	item, _ := grammar.Analyze(table, 'E', "i*(i+i)-i/(i+i)")

	for _, val := range item {
		fmt.Printf("%s %s %s\n", val.Stack, val.Cur, val.Input)
	}
}
