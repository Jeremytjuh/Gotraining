package main

import (
	pb "calc/CalcTest"
	"errors"
)

// Calculate multiplies 2 values and returns the outcome
func Calculate(m *pb.CalcRequest) (int32, error) {
	switch m.Type {
	case "Add":
		return m.Num1 + m.Num2, nil
	case "Subtract":
		return m.Num1 - m.Num2, nil
	case "Multiply":
		return m.Num1 * m.Num2, nil
	case "Divide":
		return m.Num1 / m.Num2, nil
	default:
		return 0, errors.New("Invalid identifier, the correct identifiers are: Sum, Diff, Time, Div")
	}
}
