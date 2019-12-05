package lexical

// MachineMap -
var MachineMap map[string]int

func initMachineCode() {
	MachineMap = make(map[string]int, 38)

	MachineMap["and"] = 1
	MachineMap["begin"] = 2
	MachineMap["bool"] = 3
	MachineMap["do"] = 4
	MachineMap["else"] = 5
	MachineMap["end"] = 6
	MachineMap["false"] = 7
	MachineMap["if"] = 8
	MachineMap["integer"] = 9
	MachineMap["not"] = 10
	MachineMap["or"] = 11
	MachineMap["program"] = 12
	MachineMap["real"] = 13
	MachineMap["then"] = 14
	MachineMap["true"] = 15
	MachineMap["var"] = 16
	MachineMap["while"] = 17
	MachineMap["标识符"] = 18
	MachineMap["整数"] = 19
	MachineMap["实数"] = 20
	MachineMap["("] = 21
	MachineMap[")"] = 22
	MachineMap["+"] = 23
	MachineMap["-"] = 24
	MachineMap["*"] = 25
	MachineMap["/"] = 26
	MachineMap["."] = 27
	MachineMap[","] = 28
	MachineMap[":"] = 29
	MachineMap[";"] = 30
	MachineMap[":="] = 31
	MachineMap["="] = 32
	MachineMap["<="] = 33
	MachineMap["<"] = 34
	MachineMap["<>"] = 35
	MachineMap[">"] = 36
	MachineMap[">="] = 37
	MachineMap["#"] = 38

}

// a symbol table is recorded for each token generated
type token struct {
	ID          int
	Name        []byte
	MachineCode int
	Addr        int
}

type symbleTbl struct {
	ID   int
	Type int
	Name []byte
}

type analysis struct {
	src        []byte
	currentRow int
	errCount   int

	tokens []*token
	label  int

	symbles []*symbleTbl

	lexemeBegin int
	forward     int
}
