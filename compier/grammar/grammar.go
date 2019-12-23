package grammar

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Analysis -
func Analysis() {
	raw, err := ioutil.ReadFile("grammar.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	grammar := strings.Split(string(raw), "\n")
	rules := newRules()
	for i := range grammar {
		rules.addRules(grammar[i])
	}

	fis := getFirstFrom(rules)
	fis.writeFirstSetToFile("first-set.txt")

	start := grammar[0][0]
	fos := getFollowFrom(rules, start, fis)
	fos.writeFollowSetToFile("follow-set.txt")

	ch := getChartFrom(fis, fos, rules)
	ch.writeTableToFile("prediction.txt")

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	step, err := analyze(ch, start, string(input))
	if err != nil {
		fmt.Println(err)
		return
	}

	writeStackToFile(step, "stack.txt")
}
