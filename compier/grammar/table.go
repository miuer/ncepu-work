package grammar

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type chart map[byte]map[byte]string

// GetChartFrom -
func getChartFrom(fis firstSet, fos followSet, rus rules) chart {
	ch := make(chart)
	for NT, rule := range rus {
		for i := range rule {
			if ch[NT] == nil {
				ch[NT] = make(map[byte]string)
			}

			for T := range fis[NT] {
				if T != '@' {
					// 首字符是终结符
					if t := rus.getTheFirstItem(NT, T); t != "" {
						ch[NT][T] = string(NT) + "->" + t

						// 如果 first 与输入串匹配，添加到预测分析表中
					} else if rus.Dfs(NT, T) {
						ch[NT][T] = string(NT) + "->" + rule[i]
					}
				} else {
					// A -> ε
					for k := range fos[NT] {
						ch[NT][k] = string(NT) + "->" + string(T)
					}
				}
			}
		}
	}

	return ch
}

func (c chart) String() string {
	var builder strings.Builder
	for row, v := range c {
		for col, formula := range v {
			builder.Write([]byte{'{', '[', row, ' ', col, ']', ' '})
			builder.WriteString(formula + "} ")
		}
		builder.WriteByte('\n')
	}
	return builder.String()
}

func (c chart) writeTableToFile(filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = io.WriteString(file, c.String())
	if err != nil {
		fmt.Println(err)
	}
}
