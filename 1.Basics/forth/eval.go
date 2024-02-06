//go:build !solution

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Evaluator struct {
	dict map[string][]string
}

// NewEvaluator creates evaluator.
func NewEvaluator() *Evaluator {
	return &Evaluator{
		dict: map[string][]string{
			"+":    {"+"},
			"-":    {"-"},
			"*":    {"*"},
			"/":    {"/"},
			"dup":  {"dup"},
			"swap": {"swap"},
			"drop": {"drop"},
			"over": {"over"},
		},
	}
}

var (
	DigitError    = fmt.Errorf("Digit error")
	AtoiError     = fmt.Errorf("Atoi error")
	ArgsError1    = fmt.Errorf("Not enough arguments(2)")
	DivByZero     = fmt.Errorf("Division by Zero attemption")
	ArgsError2    = fmt.Errorf("Not enough one arguments(1)")
	NoSecondValue = fmt.Errorf("SecondValue error")
	CommandError  = fmt.Errorf("Unknown command")
)

func Calculate(command string, stack []int) ([]int, error) {
	var a, b int
	if len(stack) >= 2 {
		a = stack[len(stack)-1]
		b = stack[len(stack)-2]
	}
	switch command {
	case "+":
		if len(stack) < 2 {
			return nil, ArgsError1
		}
		stack = append(stack[:len(stack)-2], a+b)
	case "-":
		if len(stack) < 2 {
			return nil, ArgsError1
		}
		stack = append(stack[:len(stack)-2], b-a)
	case "*":
		if len(stack) < 2 {
			return nil, ArgsError1
		}
		stack = append(stack[:len(stack)-2], a*b)
	case "dup":
		if len(stack) < 1 {
			return nil, ArgsError2
		}
		stack = append(stack, stack[len(stack)-1])
	case "/":
		if len(stack) < 2 {
			return nil, ArgsError1
		}
		if a == 0 {
			return nil, DivByZero
		}
		stack = append(stack[:len(stack)-2], b/a)
	case "over":
		if len(stack) < 2 {
			return nil, NoSecondValue
		}
		stack = append(stack, stack[len(stack)-2])
	case "swap":
		if len(stack) < 2 {
			return nil, ArgsError1
		}
		stack[len(stack)-1], stack[len(stack)-2] = stack[len(stack)-2], stack[len(stack)-1]
	case "drop":
		if len(stack) < 1 {
			return nil, ArgsError2
		}
		stack = stack[:len(stack)-1]
	default:
		nums, err := strconv.Atoi(command)
		if err != nil {
			return nil, AtoiError
		}
		stack = append(stack, nums)

	}
	return stack, nil
}

// Process evaluates sequence of words or definition.
//
// Returns resulting stack state and an error.
func (e *Evaluator) Process(row string) ([]int, error) {
	var stack []int
	commands := strings.Split(row, " ")
	if commands[0] == ":" {
		newCmd := strings.ToLower(commands[1])
		_, err := strconv.Atoi(newCmd)
		if err == nil {
			return nil, DigitError
		}
		temp := make([]string, 0, len(commands[2:]))
		for _, definition := range commands[2 : len(commands)-1] {
			d := strings.ToLower(definition)
			if _, ok := e.dict[d]; ok {
				temp = append(temp, e.dict[d]...)
			} else {
				_, err := strconv.Atoi(d)
				if err != nil {
					return nil, CommandError
				}
				temp = append(temp, d)
			}
		}
		e.dict[newCmd] = temp

	} else {
		for _, word := range commands {
			lower := strings.ToLower(word)
			var err error
			if _, ok := e.dict[lower]; ok {
				for _, f := range e.dict[lower] {
					stack, err = Calculate(f, stack)
					if err != nil {
						return nil, err
					}
				}
			} else {
				stack, err = Calculate(lower, stack)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return stack, nil
}
