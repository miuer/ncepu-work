package grammar

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type firstSet map[byte]map[byte]struct{}

// GetFirstFrom -
func getFirstFrom(r rules) firstSet {
	fis := make(firstSet)
	var changed bool

	for {
		changed = false
		for k, v := range r {
			if fis[k] == nil {
				fis[k] = make(map[byte]struct{})
			}

			for _, val := range v {
				// 第一个字符是终结符
				if val[0] < 'A' || val[0] > 'Z' {
					if mergeFirstSet(fis[k], map[byte]struct{}{val[0]: {}}) != 0 {
						changed = true
					}
					continue
				}

				// 第一个字符是非终结符
				if val[0] >= 'A' && val[0] <= 'Z' {
					if removeEmptyAndMergeFirstSet(fis[k], fis[val[0]]) != 0 {
						changed = true
					}
				}
			}

		}
		if !changed {
			break
		}
	}
	return fis
}

func mergeFirstSet(a map[byte]struct{}, b map[byte]struct{}) int {
	count := 0
	for key, value := range b {
		if _, ok := a[key]; !ok {
			count++
		}
		a[key] = value
	}
	return count
}

func removeEmptyAndMergeFirstSet(a map[byte]struct{}, b map[byte]struct{}) int {
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

func (f firstSet) String() string {
	var build strings.Builder
	for key, value := range f {
		build.WriteString(fmt.Sprintf("FIRST(%c)={ ", key))
		for item := range value {
			build.WriteString(fmt.Sprintf("%c ", item))
		}
		build.WriteString("}\n")
	}
	return build.String()
}

func (f firstSet) haveEmpty(first byte) bool {
	_, ok := f[first]['@']
	return ok
}

func (f firstSet) writeFirstSetToFile(filename string) {
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
