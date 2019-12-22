package grammar

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type followSet map[byte]map[byte]struct{}

// GetFollowFrom -
func getFollowFrom(r rules, start byte, fis firstSet) followSet {
	fos := make(followSet)

	fos[start] = make(map[byte]struct{})
	fos[start]['#'] = struct{}{}

	var changed bool
	for {
		changed = false
		for k, v := range r {
			for _, val := range v {
				for i := 0; i < len(val)-1; i++ {

					if fos[val[i]] == nil {
						fos[val[i]] = make(map[byte]struct{})
					}

					// A -> Bb | BC  终结符会直接跳过，一直到非终结符
					if !isTerminal(val[i]) {
						// 如果非终结符后面是终结符，将终结符放进非终结符的 follow
						if isTerminal(val[i+1]) {
							if mergeFollowSet(fos[val[i]], map[byte]struct{}{val[i+1]: {}}) != 0 {
								changed = true
							}

							// 如果非终结符后面是非终结符，将第二个非终结符的 first 放进第一个非终结符的 follow
						} else {
							if removeEmptyAndmergeFollowSet(fos[val[i]], fis[val[i+1]]) != 0 {
								changed = true
							}
						}
						// 如果非终结符后面是非终结符，第二个非终结符的 first 中含有空串，则将第前置的 follow 加入非终结符
						if fis.haveEmpty(val[i+1]) {
							if mergeFollowSet(fos[val[i]], fos[k]) != 0 {
								changed = true
							}
						}
					}
				}

				// A -> a 单字符（终结符）处理情况
				if fos[val[len(val)-1]] == nil {
					fos[val[len(val)-1]] = make(map[byte]struct{})
				}

				// A -> B 单字符（终结符）处理情况
				if !isTerminal(val[len(val)-1]) {
					if mergeFollowSet(fos[val[len(val)-1]], fos[k]) != 0 {
						changed = true
					}
				}
			}
		}

		if !changed {
			break
		}
	}
	return fos
}

func mergeFollowSet(a map[byte]struct{}, b map[byte]struct{}) int {
	count := 0
	for key, value := range b {
		if _, ok := a[key]; !ok {
			count++
		}
		a[key] = value
	}
	return count
}

func removeEmptyAndmergeFollowSet(a map[byte]struct{}, b map[byte]struct{}) int {
	flag := false
	if _, flag = b['@']; flag {
		flag = true
		delete(b, '@')
	}
	count := 0
	for key, value := range b {
		if _, ok := a[key]; !ok {
			count++
		}
		a[key] = value
	}
	if flag {
		b['@'] = struct{}{}
	}
	return count
}

func (f followSet) String() string {
	var build strings.Builder
	for key, value := range f {
		build.WriteString(fmt.Sprintf("FOLLOW(%c)={ ", key))
		for item := range value {
			build.WriteString(fmt.Sprintf("%c ", item))
		}
		build.WriteString("}\n")
	}
	return build.String()
}

func isTerminal(a byte) bool {
	if a < 'A' || a > 'Z' {
		return true
	}
	return false
}

func (f followSet) writeFollowSetToFile(filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = io.WriteString(file, f.String())
	if err != nil {
		fmt.Println(err)
	}
}
