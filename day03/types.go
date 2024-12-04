package main

type Day3Inputs struct {
	Rows []string
}

type Operation struct {
	Operation OperationType
	op1       int
	op2       int
}
type OperationType int

const (
	MUL OperationType = iota
	DO
	DONT
)

func (o OperationType) String() string {
	switch o {
	case MUL:
		return "MUL"
	case DO:
		return "DO"
	case DONT:
		return "DONT"
	default:
		return "UNKNOWN"
	}
}
