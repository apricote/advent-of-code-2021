package main

import (
	"reflect"
	"testing"
)

func TestParsePacket(t *testing.T) {
	type test struct {
		name  string
		input []byte
		want  Packet
	}

	tests := []test{
		{
			name:  "LiteralValue",
			input: []byte{0xD2, 0xFE, 0x28},
			want: PacketLiteral{
				packet: packet{
					_version: 6,
					_type:    4,
				},
				Value: 2021,
			},
		},
		{
			name:  "OperatorWithLiteralSubPackets",
			input: []byte{0x38, 0x00, 0x6F, 0x45, 0x29, 0x12, 0x00},
			want: PacketOperator{
				packet: packet{
					_version: 1,
					_type:    6,
				},
				LengthType: LengthTypeBits,
				Length:     27,
				SubPackets: []Packet{
					PacketLiteral{
						packet: packet{
							_version: 6,
							_type:    4,
						},
						Value: 10,
					},
					PacketLiteral{
						packet: packet{
							_version: 2,
							_type:    4,
						},
						Value: 20,
					},
				},
			},
		},
		{
			name:  "OperatorWithOperatorSubPackets",
			input: []byte{0x8A, 0x00, 0x4A, 0x80, 0x1A, 0x80, 0x02, 0xF4, 0x78},
			want: PacketOperator{
				packet: packet{
					_version: 4,
					_type:    2,
				},
				LengthType: LengthTypePackets,
				Length:     1,
				SubPackets: []Packet{
					PacketOperator{
						packet: packet{
							_version: 1,
							_type:    2,
						},
						LengthType: LengthTypePackets,
						Length:     1,
						SubPackets: []Packet{
							PacketOperator{
								packet: packet{
									_version: 5,
									_type:    2,
								},
								LengthType: LengthTypeBits,
								Length:     11,
								SubPackets: []Packet{
									PacketLiteral{
										packet: packet{
											_version: 6,
											_type:    4,
										},
										Value: 15,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, _, err := ParsePacket(tc.input, 0)

			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("Expected %v but got %v", tc.want, got)
			}
		})
	}
}
