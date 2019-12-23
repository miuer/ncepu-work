package grammar

import (
	"fmt"
	"strings"

	"github.com/Delveshal/compiler-LL1/util"
)

type rules map[byte][]string

func newRules() rules {
	return make(rules)
}

// 添加文法规则
func (r rules) addRules(str string) error {
	stepOne := strings.Split(str, "->")

	// 去除 ->
	if len(stepOne) != 2 {
		return fmt.Errorf("the format of input is invalid,expect X->Y but actually %s", str)
	}

	// 去除多余空格
	if len(stepOne[0]) != 1 {
		return fmt.Errorf("input is invalid,expect X on the left but actually %s", stepOne[0])
	}

	stepTwo := strings.Split(strings.Replace(stepOne[1], " ", "", -1), "|")
	for i := range stepTwo {
		r[stepOne[0][0]] = append(r[stepOne[0][0]], stepTwo[i])
	}

	return nil
}

func (r rules) getTheFirstItem(first, item byte) string {
	for _, v := range r[first] {
		if v[0] == item {
			return v
		}
	}
	return ""
}

func (r rules) Dfs(first, terminal byte) bool {
	for i := range r[first] {
		if util.IsTerminal(r[first][i][0]) && r[first][i][0] == terminal {
			return true
		} else if !util.IsTerminal(r[first][i][0]) {
			if ok := r.Dfs(r[first][i][0], terminal); ok {
				return true
			}
		}
	}
	return false
}
