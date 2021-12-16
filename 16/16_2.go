package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math"
)

func SolvePacketEquation(input string) int {
	data, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}

	packet, _, err := ParsePacket(data, 0)
	if err != nil {
		log.Fatal(err)
	}

	solution, err := EvaluatePacket(packet)
	if err != nil {
		log.Fatal(err)
	}

	return solution
}

func EvaluatePacket(p Packet) (int, error) {
	if packet, ok := p.(PacketLiteral); ok {
		return packet.Value, nil
	}

	packet, ok := p.(PacketOperator)
	if !ok {
		return 0, fmt.Errorf("unable to convert packet %v", p)
	}

	subPacketValues := []int{}

	for _, subPacket := range packet.SubPackets {
		value, err := EvaluatePacket(subPacket)
		if err != nil {
			return 0, fmt.Errorf("unable to evaluate subpacket (%v) with error: %v", subPacket, err)
		}

		subPacketValues = append(subPacketValues, value)
	}

	switch packet.Type() {
	case PacketTypeSum:
		sum := 0
		for _, value := range subPacketValues {
			sum += value
		}

		return sum, nil

	case PacketTypeProduct:
		product := 1
		for _, value := range subPacketValues {
			product *= value
		}

		return product, nil

	case PacketTypeMinimum:
		min := math.MaxInt
		for _, value := range subPacketValues {
			if value < min {
				min = value
			}
		}

		return min, nil

	case PacketTypeMaximum:
		max := math.MinInt
		for _, value := range subPacketValues {
			if value > max {
				max = value
			}
		}

		return max, nil

	case PacketTypeGreaterThan:
		if len(subPacketValues) != 2 {
			return 0, fmt.Errorf("GreaterThan packet has more or less than two sub packets")
		}

		if subPacketValues[0] > subPacketValues[1] {
			return 1, nil
		} else {
			return 0, nil
		}

	case PacketTypeLessThan:
		if len(subPacketValues) != 2 {
			return 0, fmt.Errorf("LessThan packet has more or less than two sub packets")
		}

		if subPacketValues[0] < subPacketValues[1] {
			return 1, nil
		} else {
			return 0, nil
		}

	case PacketTypeEqualTo:
		if len(subPacketValues) != 2 {
			return 0, fmt.Errorf("EqualTo packet has more or less than two sub packets")
		}

		if subPacketValues[0] == subPacketValues[1] {
			return 1, nil
		} else {
			return 0, nil
		}

	default:
		return 0, fmt.Errorf("unable to match packet type")
	}

}
