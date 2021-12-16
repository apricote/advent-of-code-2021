package main

import (
	"fmt"
	"strconv"
)

type Packet interface {
	Version() int
	Type() PacketType
}

type packet struct {
	_version int
	_type    PacketType
}

type PacketType int

const (
	PacketTypeSum         PacketType = iota
	PacketTypeProduct     PacketType = iota
	PacketTypeMinimum     PacketType = iota
	PacketTypeMaximum     PacketType = iota
	PacketTypeLiteral     PacketType = iota
	PacketTypeGreaterThan PacketType = iota
	PacketTypeLessThan    PacketType = iota
	PacketTypeEqualTo     PacketType = iota
)

func ParsePacketType(typeInt int) (PacketType, error) {
	switch typeInt {
	case int(PacketTypeSum):
		return PacketTypeSum, nil
	case int(PacketTypeProduct):
		return PacketTypeProduct, nil
	case int(PacketTypeMinimum):
		return PacketTypeMinimum, nil
	case int(PacketTypeMaximum):
		return PacketTypeMaximum, nil
	case int(PacketTypeLiteral):
		return PacketTypeLiteral, nil
	case int(PacketTypeGreaterThan):
		return PacketTypeGreaterThan, nil
	case int(PacketTypeLessThan):
		return PacketTypeLessThan, nil
	case int(PacketTypeEqualTo):
		return PacketTypeEqualTo, nil
	}

	return 0, fmt.Errorf("unable to parse PacketType: %v", typeInt)
}

func (p packet) Version() int {
	return p._version
}

func (p packet) Type() PacketType {
	return p._type
}

type PacketLiteral struct {
	packet
	Value int
}

type PacketOperator struct {
	packet
	LengthType LengthType
	Length     int

	SubPackets []Packet
}

type LengthType int

const (
	LengthTypeBits    LengthType = 0
	LengthTypePackets LengthType = 1
)

func ParsePacket(data []byte, offset int) (Packet, int, error) {
	_version, err := GetBitsBetween(data, offset+0, offset+2)
	if err != nil {
		return packet{}, 0, err
	}

	_type, err := GetBitsBetween(data, offset+3, offset+5)
	if err != nil {
		return packet{}, 0, err
	}

	packetType, err := ParsePacketType(_type)
	if err != nil {
		return packet{}, 0, err
	}

	packet := packet{
		_version: _version,
		_type:    packetType,
	}

	switch packet._type {
	case PacketTypeLiteral:
		// Literal
		return ParsePacketLiteral(data, offset, packet)
	default:
		// Operator
		return ParsePacketOperator(data, offset, packet)
	}
}

func ParsePacketOperator(data []byte, offset int, packet packet) (PacketOperator, int, error) {

	lengthTypeBit, err := GetBitAt(data, offset+6)
	if err != nil {
		return PacketOperator{}, 0, err
	}

	var lengthType LengthType
	var length int

	subPacketOffset := 0

	switch lengthTypeBit {
	case false:
		lengthType = LengthTypeBits
		subPacketOffset += 22
		length, err = GetBitsBetween(data, offset+7, offset+21)
	case true:
		lengthType = LengthTypePackets
		subPacketOffset += 18
		length, err = GetBitsBetween(data, offset+7, offset+17)
	}
	if err != nil {
		return PacketOperator{}, 0, err
	}

	subPackets := []Packet{}
	totalSubPacketLength := 0

subPacketParsing:
	for {
		switch lengthType {
		case LengthTypeBits:
			if totalSubPacketLength >= length {
				break subPacketParsing
			}
		case LengthTypePackets:
			if len(subPackets) >= length {
				break subPacketParsing
			}
		}

		subPacket, subPacketLength, err := ParsePacket(data, offset+subPacketOffset+totalSubPacketLength)
		if err != nil {
			return PacketOperator{}, 0, err
		}

		subPackets = append(subPackets, subPacket)
		totalSubPacketLength += subPacketLength
	}

	return PacketOperator{
		packet:     packet,
		LengthType: lengthType,
		Length:     length,
		SubPackets: subPackets,
	}, subPacketOffset + totalSubPacketLength, nil
}

func ParsePacketLiteral(data []byte, offset int, packet packet) (PacketLiteral, int, error) {
	bits := ""

	var err error

	group := 0
	hasMoreGroups := true

	for hasMoreGroups {
		groupOffset := offset + 5 + group*5

		hasMoreGroups, err = GetBitAt(data, groupOffset+1)
		if err != nil {
			return PacketLiteral{}, 0, err
		}

		for i := groupOffset + 2; i <= groupOffset+5; i++ {
			bit, err := GetBitAt(data, i)
			if err != nil {
				return PacketLiteral{}, 0, err
			}

			if bit {
				bits += "1"
			} else {
				bits += "0"
			}
		}

		group += 1
	}

	value, err := strconv.ParseInt(bits, 2, 0)
	if err != nil {
		return PacketLiteral{}, 0, err
	}

	return PacketLiteral{
		packet: packet,
		Value:  int(value),
	}, 6 + group*5, nil
}

func GetBitsBetween(data []byte, start, end int) (int, error) {
	bits := ""

	for i := start; i <= end; i++ {
		bit, err := GetBitAt(data, i)
		if err != nil {
			return 0, err
		}

		if bit {
			bits += "1"
		} else {
			bits += "0"
		}
	}

	output, err := strconv.ParseInt(bits, 2, 0)
	if err != nil {
		return 0, err
	}
	return int(output), nil
}

func GetBitAt(data []byte, pos int) (bool, error) {
	byteId := pos / 8
	bitInByte := pos % 8

	if byteId > len(data) {
		return false, fmt.Errorf("pos %v not in data with length %v", byteId, len(data))
	}

	bitShifted := data[byteId] >> (7 - bitInByte)
	bit := bitShifted & 1

	return bit == 1, nil
}
