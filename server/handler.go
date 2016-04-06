package server

import (
	"errors"
	"log"

	"github.com/alyoshka/calc-server/gen-go/calculate"
)

type calculateHandler struct {
}

func newCalculateHandler() *calculateHandler {
	return &calculateHandler{}
}

// Calculate - handles the rpc requests
func (ch *calculateHandler) Calculate(operation calculate.Operation, arg1 float64, arg2 float64) (float64, error) {
	log.Println("New request: ", operation, arg1, arg2)
	var result float64
	switch operation {
	case calculate.Operation_ADD:
		result = arg1 + arg2
	case calculate.Operation_SUB:
		result = arg1 - arg2
	case calculate.Operation_DIVIDE:
		if arg2 == 0 {
			return result, errors.New("Division by zero")
		}
		result = arg1 / arg2
	case calculate.Operation_MULTIPLY:
		result = arg1 * arg2
	default:
		return result, errors.New("Unknown operation")
	}
	return result, nil
}
