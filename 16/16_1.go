package main

import (
	"encoding/hex"
	"log"
)

func GetVersionNumberCountFromBITS(input string) int {
	data, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}

	packet, _, err := ParsePacket(data, 0)
	if err != nil {
		log.Fatal(err)
	}

	return SumVersionNumbers(packet)
}

func SumVersionNumbers(p Packet) int {
	sum := p.Version()

	switch v := p.(type) {
	case PacketLiteral:
		// Nothing to do here
	case PacketOperator:
		for _, subPacket := range v.SubPackets {
			sum += SumVersionNumbers(subPacket)
		}
	}

	return sum
}
