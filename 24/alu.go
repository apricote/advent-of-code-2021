package main

import (
	"log"
	"strconv"
	"strings"
)

type ALU struct {
	Registers [4]int
	Inputs    []int
}

type Instruction struct {
	Operation string
	ValueA    string
	ValueB    string
}

type registerOrValue struct {
	Register int
	Value    int
	Type     RegisterOrValueType
}

type RegisterOrValueType bool

const (
	ROVRegister RegisterOrValueType = true
	ROVValue    RegisterOrValueType = false
)

func (a *ALU) Execute(instruction Instruction) {

	inputA := GetRegisterOrValue(instruction.ValueA)
	inputB := GetRegisterOrValue(instruction.ValueB)

	if inputA.Type == ROVValue {
		log.Printf("Invalid Instruction! A must be a Register!")
		return
	}

	switch instruction.Operation {
	case "inp":
		var input int
		input, a.Inputs = a.Inputs[0], a.Inputs[1:]
		a.Registers[inputA.Register] = input
	case "add":
		a.Registers[inputA.Register] += a.GetValue(inputB)

	case "mul":
		a.Registers[inputA.Register] *= a.GetValue(inputB)

	case "div":
		a.Registers[inputA.Register] /= a.GetValue(inputB)

	case "mod":
		a.Registers[inputA.Register] %= a.GetValue(inputB)

	case "eql":
		if a.GetValue(inputA) == a.GetValue(inputB) {
			a.Registers[inputA.Register] = 1
		} else {
			a.Registers[inputA.Register] = 0
		}
	}
}

func (a ALU) GetValue(rov registerOrValue) int {
	if rov.Type == ROVRegister {
		return a.Registers[rov.Register]
	} else {
		return rov.Value
	}
}

func GetRegisterOrValue(input string) registerOrValue {
	if len(input) == 0 {
		return registerOrValue{}
	}

	value, err := strconv.Atoi(input)

	if err == nil {
		// input was a value!
		return registerOrValue{
			Value: value,
			Type:  ROVValue,
		}
	}

	register := 0
	switch input {
	case "w":
		register = 0
	case "x":
		register = 1
	case "y":
		register = 2
	case "z":
		register = 3
	}

	return registerOrValue{Register: register, Type: ROVRegister}
}

func RunALU(instructions *[]Instruction, inputs []int) [4]int {
	alu := ALU{
		Inputs: inputs,
	}

	for _, instruction := range *instructions {
		alu.Execute(instruction)
	}

	return alu.Registers
}

func ParseProgram(program string) []Instruction {
	instructions := []Instruction{}

	for _, line := range strings.Split(program, "\n") {

		parts := strings.Split(line, " ")

		instruction := Instruction{
			Operation: parts[0],
			ValueA:    parts[1],
		}

		if len(parts) == 3 {
			// Some instructions only have one input
			instruction.ValueB = parts[2]
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}
