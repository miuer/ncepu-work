package lexical

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// As -
var As analysis

func (as *analysis) Analysis(filename, tokenName, symbleName string) {
	srcCode, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("file %s could not be found.", filename))
	}
	as.src = append(srcCode, '#')
	as.currentRow = 1
	initMachineCode()

	as.startLexicalAnalysis()
	as.writeTokensToFile(tokenName)
	as.writeSymblesToFile(symbleName)
	/*
		// --- You can view the results of the lexical analysis by uncommenting
		for i := 0; i < len(as.tokens); i++ {
			fmt.Println(as.tokens[i].ID, string(as.tokens[i].Name), as.tokens[i].MachineCode, as.tokens[i].Addr)
		}

		for i := 0; i < len(as.symbles); i++ {
			fmt.Println(as.symbles[i].ID, string(as.symbles[i].Name), as.symbles[i].Type)
		}
	*/

}

func (as *analysis) startLexicalAnalysis() {
	for true {
		if !isValid(as.src[as.forward]) {
			fmt.Printf("syntax error: unexpected literal %s at %d\n", string(as.src[as.forward]), as.currentRow)
			as.src = append(as.src[:as.forward], as.src[as.forward+1:]...)
			continue
		} else if as.src[as.forward] == ' ' {
			as.forward++
			as.lexemeBegin = as.forward
			continue
		} else if as.src[as.forward] == '	' {
			as.forward++
			as.lexemeBegin = as.forward
			continue
		} else if as.src[as.forward] == '\n' {
			as.currentRow++
			as.forward++
			as.lexemeBegin = as.forward
			continue
		} else if isLetter(as.src[as.forward]) {
			as.scanToLetter()
			continue
		} else if isDigit(as.src[as.forward]) {
			as.scanToDigit()
			continue
		} else if as.src[as.forward] == '#' {
			break
		} else {
			as.scanToSymbol()
		}
	}
}

func (as *analysis) scanToLetter() {
	as.forward++

	if !isValid(as.src[as.forward]) {
		fmt.Printf("syntax error: unexpected literal %s at %d\n", string(as.src[as.forward]), as.currentRow)
		as.src = append(as.src[:as.forward], as.src[as.forward+1:]...)
		as.scanToLetter()
	} else if isLetter(as.src[as.forward]) || isDigit(as.src[as.forward]) {
		as.scanToLetter()
	} else {
		name := as.src[as.lexemeBegin:as.forward]
		as.setToken(name)
		as.lexemeBegin = as.forward
	}
}

func (as *analysis) scanToDigit() {
	as.forward++

	if !isValid(as.src[as.forward]) {
		fmt.Printf("syntax error: unexpected literal %s at %d\n", string(as.src[as.forward]), as.currentRow)
		as.src = append(as.src[:as.forward], as.src[as.forward+1:]...)
		as.scanToDigit()
	} else if isLetter(as.src[as.forward]) {
		fmt.Printf("syntax error: unexpected literal %s at %d\n", string(as.src[as.forward]), as.currentRow)
		as.src = append(as.src[:as.forward], as.src[as.forward+1:]...)
		as.scanToDigit()
	} else if isDigit(as.src[as.forward]) {
		as.scanToDigit()
	} else if as.src[as.forward] == '.' {
		as.scanToFloat()
	} else {
		name := as.src[as.lexemeBegin:as.forward]
		as.setToken(name)
		as.lexemeBegin = as.forward
	}
}

func (as *analysis) scanToFloat() {
	as.forward++
	if !isValid(as.src[as.forward]) {
		fmt.Printf("syntax error: unexpected literal %s at %d\n", string(as.src[as.forward]), as.currentRow)
		as.src = append(as.src[:as.forward], as.src[as.forward+1:]...)
		as.scanToFloat()
	} else if isLetter(as.src[as.forward]) {
		fmt.Printf("syntax error: unexpected literal %s at %d\n", string(as.src[as.forward]), as.currentRow)
		as.src = append(as.src[:as.forward], as.src[as.forward+1:]...)
		as.scanToFloat()
	} else if isDigit(as.src[as.forward]) {
		as.scanToFloat()
	} else if as.src[as.forward] == '.' {
		fmt.Printf("syntax error: unexpected literal %s at %d\n", string(as.src[as.forward]), as.currentRow)
		as.src = append(as.src[:as.forward], as.src[as.forward+1:]...)
		as.scanToFloat()
	} else {
		name := as.src[as.lexemeBegin:as.forward]
		as.setToken(name)
		as.lexemeBegin = as.forward
	}
}

func (as *analysis) scanToSymbol() {
	if !isValid(as.src[as.forward]) {
		fmt.Printf("syntax error: unexpected literal %s at %d\n", string(as.src[as.forward]), as.currentRow)
		as.src = append(as.src[:as.forward], as.src[as.forward+1:]...)
		as.scanToSymbol()
	} else if as.src[as.forward] == ':' && as.src[as.forward+1] == '=' {
		name := as.src[as.lexemeBegin : as.forward+2]
		as.setToken(name)
		as.forward += 2
		as.lexemeBegin = as.forward
		return
	} else if as.src[as.forward] == '>' && as.src[as.forward+1] == '=' {
		name := as.src[as.lexemeBegin : as.forward+2]
		as.setToken(name)
		as.forward += 2
		as.lexemeBegin = as.forward
		return
	} else if as.src[as.forward] == '<' && as.src[as.forward+1] == '=' {
		name := as.src[as.lexemeBegin : as.forward+2]
		as.setToken(name)
		as.forward += 2
		as.lexemeBegin = as.forward
		return
	} else if as.src[as.forward] == '<' && as.src[as.forward+1] == '>' {
		name := as.src[as.lexemeBegin : as.forward+2]
		as.setToken(name)
		as.forward += 2
		as.lexemeBegin = as.forward
		return
	} else {
		name := as.src[as.lexemeBegin : as.forward+1]
		as.setToken(name)
		as.forward++
		as.lexemeBegin = as.forward
	}
}

func (as *analysis) setToken(name []byte) {
	as.label++
	mc := getMachineCode(name)

	tk := &token{
		ID:          as.label,
		Name:        name,
		MachineCode: mc,
		Addr:        as.setSymbol(name, mc),
	}

	as.tokens = append(as.tokens, tk)
}

func (as *analysis) setSymbol(name []byte, mc int) int {
	if mc < 18 || mc > 20 {
		return -1
	}

	// Handling characters that have already appeared
	for i := 0; i < len(as.symbles); i++ {
		if string(name) == string(as.symbles[i].Name) {
			return as.symbles[i].ID
		}
	}

	symble := &symbleTbl{
		ID:   len(as.symbles),
		Type: mc,
		Name: name,
	}

	as.symbles = append(as.symbles, symble)
	return len(as.symbles)
}

func (as *analysis) writeTokensToFile(filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	for i := 0; i < len(as.tokens); i++ {
		content := fmt.Sprintf("%d %s %d %d\n", as.tokens[i].ID, string(as.tokens[i].Name), as.tokens[i].MachineCode, as.tokens[i].Addr)
		_, err := io.WriteString(f, content)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (as *analysis) writeSymblesToFile(filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	for i := 0; i < len(as.symbles); i++ {
		content := fmt.Sprintf("%d %s %d\n", as.symbles[i].ID, string(as.symbles[i].Name), as.symbles[i].Type)
		_, err := io.WriteString(f, content)
		if err != nil {
			fmt.Println(err)
		}
	}
}
