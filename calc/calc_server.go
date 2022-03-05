package calc

import (
	"calculator-server/api"
	"context"
	"errors"
)

type Server struct {
	api.UnimplementedCalculatorServer
}

func (s *Server) DoCalc(ctx context.Context, in *api.CalculatorInputMessage) (*api.CalculatorResultMessage, error) {
	// Harusnya ga perlu return error karena sudah diprovide di CalculatorResultMessage
	num1 := in.Num1
	num2 := in.Num2
	opr := in.Operator
	var result int32
	var err error

	switch opr {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			err = errors.New("Zero divided")
		}
		result = num1 / num2
	}

	var resultMessage *api.CalculatorResultMessage
	if err != nil {
		resultMessage = &api.CalculatorResultMessage{ResNum: 0, Error: &api.Error{Code: "400", Message: err.Error()}}
	} else {
		resultMessage = &api.CalculatorResultMessage{ResNum: result}
	}

	return resultMessage, nil
}
