package server

import (
	"testing"

	"github.com/alyoshka/calc-server/gen-go/calculate"
)

type testCase struct {
	operation calculate.Operation
	arg1      float64
	arg2      float64
	result    float64
}

func TestDivideByZero(t *testing.T) {
	handler := newCalculateHandler()
	_, err := handler.Calculate(calculate.Operation_DIVIDE, 2.0, 0.0)
	if err == nil {
		t.Error("Division by zero")
	}
}

func TestUnknownOperation(t *testing.T) {
	handler := newCalculateHandler()
	_, err := handler.Calculate(42, 2.0, 2.0)
	if err == nil {
		t.Error("Unknown operation")
	}
}

func TestHandler(t *testing.T) {
	cases := []testCase{
		{
			operation: calculate.Operation_ADD,
			arg1:      1.0,
			arg2:      2.0,
			result:    3.0,
		},
		{
			operation: calculate.Operation_SUB,
			arg1:      1.0,
			arg2:      2.0,
			result:    -1.0,
		},
		{
			operation: calculate.Operation_MULTIPLY,
			arg1:      1.0,
			arg2:      2.0,
			result:    2.0,
		},
		{
			operation: calculate.Operation_DIVIDE,
			arg1:      1.0,
			arg2:      2.0,
			result:    0.5,
		},
	}
	handler := newCalculateHandler()
	for i := range cases {
		result, err := handler.Calculate(cases[i].operation, cases[i].arg1, cases[i].arg2)
		if err != nil {
			t.Errorf("Handler error: %s, operation: %s, arg1: %f, arg2: %f", err, cases[i].operation, cases[i].arg1, cases[i].arg2)
		}
		if result != cases[i].result {
			t.Errorf("Unexpected result of %s %f and %f: get %f, want %f", cases[i].operation, cases[i].arg1, cases[i].arg2, result, cases[i].result)
		}
	}
}
