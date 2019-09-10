package main

import (
	pb "calc/CalcTest"
	"errors"
)

// Calculate multiplies 2 values and returns the outcome
func Calculate(m *pb.CalcRequest) (int32, error) {
	switch m.Type {
	case "Sum":
		return m.Num1 + m.Num2, nil
	case "Diff":
		return m.Num1 - m.Num2, nil
	case "Time":
		return m.Num1 * m.Num2, nil
	case "Div":
		return m.Num1 / m.Num2, nil
	default:
		return 0, errors.New("Invalid identifier, the correct identifiers are: Sum, Diff, Time, Div")
	}
}
