package lexical

import (
	"fmt"
	"io/ioutil"
)

var as analysis

func (as *analysis) Analysis(filename string) {
	srcCode, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("file %s could not be found.", filename))
	}
	as.src = append(srcCode, '#')
	as.currentRow = 1
	as.StartLexicalAnalysis()
}

func (as *analysis) StartLexicalAnalysis() {
	for true {

		if !isValid(as.src[as.forward]) {
			fmt.Printf("syntax error: unexpected literal %s, expecting name", string(as.src[as.forward]))
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
	as.lexemeBegin = as.forward
}

func (as *analysis) scanToDigit() {
	as.lexemeBegin = as.forward
}

func (as *analysis) scanToSymbol() {
	as.lexemeBegin = as.forward
	as.forward++
	if as.src[as.lexemeBegin] == ':' && as.src[as.forward] == '=' {
		name := as.src[as.lexemeBegin : as.forward+1]
		as.setToken(name)
		as.forward++
		return
	} else if as.src[as.lexemeBegin] == '>' && as.src[as.forward] == '=' {
		name := as.src[as.lexemeBegin : as.forward+1]
		as.setToken(name)
		as.forward++
		return
	} else if as.src[as.lexemeBegin] == '<' && as.src[as.forward] == '=' {
		name := as.src[as.lexemeBegin : as.forward+1]
		as.setToken(name)
		as.forward++
		return
	} else if as.src[as.forward] == '>' {
		name := as.src[as.lexemeBegin : as.forward+1]
		as.setToken(name)
		as.forward++
		return
	} else {
		name := as.src[as.lexemeBegin:as.forward]
		as.setToken(name)
		as.forward++
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
	if mc < 18 && mc > 20 {
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
	return len(as.symbles) - 1
}
