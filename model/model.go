package model

const (
	StateSuccess  = 1
	StateTransfer = 2
	StateFail     = 3
)

type UpdateFields struct {
	Field string
	Value any
	Data  map[string]any
}
