package grammar

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

type item struct {
	Stack string
	Cur   string
	Input string
}

func analyze(ch chart, start byte, input string) ([]*item, error) {
	if len(input) <= 0 {
		return nil, errors.New("empty")
	}
	if input[len(input)-1] != '#' {
		input += "#"
	}
	stack := []byte{'#', start}
	var step []*item
	current := input[0]
	input = input[1:]

	for {
		step = append(step, &item{
			Stack: string(stack),
			Cur:   string(current),
			Input: input,
		})

		tos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 当栈顶元素是终结符时，输入层前移
		if isTerminal(tos) {
			if tos == current {
				if current != '#' {
					current = input[0]
					input = input[1:]
				}
			} else {
				return nil, fmt.Errorf("%c != %c", tos, current)
			}

			// 栈顶元素是非终结符，根据当前输入，使用预测分析表中进行规约
		} else {
			if t, ok := ch[tos][current]; ok {
				index := bytes.LastIndexByte([]byte(t), '>')
				if t[len(t)-1] != '@' {
					// 将产生式后面的串倒叙入栈
					for i := len(t) - 1; i > index; i-- {
						stack = append(stack, t[i])
					}
				}
			} else {
				return nil, errors.New("the shell is empty: " + fmt.Sprintf("%c %c", tos, current))
			}
		}
		if tos == '#' {
			break
		}
	}
	return step, nil
}

func writeStackToFile(items []*item, filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for i := 0; i < len(items); i++ {
		content := fmt.Sprintf("%s %s %s\n", items[i].Stack, items[i].Cur, items[i].Input)
		_, err := io.WriteString(file, content)
		if err != nil {
			fmt.Println(err)
		}
	}
}
