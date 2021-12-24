package main

import (
	"log"
	"math"
)

type Burrow [8]*Amphipod

type Amphipod struct {
	Type     AmphipodType
	Room     Room
	Position int
}

type AmphipodType string

const (
	A AmphipodType = "A"
	B AmphipodType = "B"
	C AmphipodType = "C"
	D AmphipodType = "D"
)

type Room string

const (
	RoomHallway = "hallway"
	RoomA       = "A"
	RoomB       = "B"
	RoomC       = "C"
	RoomD       = "D"
)

var (
	roomPositions map[Room]int = map[Room]int{
		RoomA: 2,
		RoomB: 4,
		RoomC: 6,
		RoomD: 8,
	}

	costsPerStep map[AmphipodType]int = map[AmphipodType]int{
		A: 1,
		B: 10,
		C: 100,
		D: 1000,
	}
)

type Move struct {
	AmphipodID int
	Room       Room
	Position   int
}

func OrganizeAmphipods(input string) int {
	burrow := ParseInput(input)

	moveList := [][]Move{
		{},
	}

	finishedMoveList := [][]Move{}

	for len(moveList) > 0 {
		nextMoveList := [][]Move{}

		log.Printf("Making next move, currently we have %v move lists with depth %v", len(moveList), len(moveList[0]))

		for _, moves := range moveList {
			currentBurrow := burrow.RunMoves(moves)

			if currentBurrow.IsFinished() {
				finishedMoveList = append(finishedMoveList, moves)
				continue
			}

			possibleNextMoves := currentBurrow.PossibleMoves()

			for _, nextMove := range possibleNextMoves {
				newMoves := make([]Move, len(moves), len(moves)+1)
				copy(newMoves, moves)
				newMoves = append(newMoves, nextMove)

				nextMoveList = append(nextMoveList, newMoves)
			}
		}

		moveList = nextMoveList
	}

	minCost := math.MaxInt

	for _, moves := range finishedMoveList {
		cost := CalculateCosts(burrow, moves)

		if minCost > cost {
			minCost = cost
		}
	}

	// Evaluate score for each move list
	// Pick lowest score

	return minCost
}

func ParseInput(input string) Burrow {
	burrow := Burrow{}

	initialAmphipodPositions := [8]int{
		31,
		31 + 14,
		33,
		33 + 14,
		35,
		35 + 14,
		37,
		37 + 14,
	}

	for i, byteOffset := range initialAmphipodPositions {
		burrow[i] = &Amphipod{}

		burrow[i].Type = GetAmphipodFromByte(input[byteOffset])

		if i%2 == 1 {
			burrow[i].Position = 1
		}

		switch i / 2 {
		case 0:
			burrow[i].Room = RoomA
		case 1:
			burrow[i].Room = RoomB
		case 2:
			burrow[i].Room = RoomC
		case 3:
			burrow[i].Room = RoomD
		}
	}

	return burrow
}

func GetAmphipodFromByte(input byte) AmphipodType {
	switch input {
	case byte('A'):
		return A
	case byte('B'):
		return B
	case byte('C'):
		return C
	case byte('D'):
		return D
	}

	return A
}

func (burrow *Burrow) PossibleMoves() []Move {
	possibleMoves := []Move{}

	for id, amphipod := range burrow {
		if amphipod.Room == RoomHallway {
			room0 := burrow.GetAmphipodAtPosition(amphipod.Type.GetRoom(), 0)
			room1 := burrow.GetAmphipodAtPosition(amphipod.Type.GetRoom(), 1)

			if (room0 != nil && room0.Type != amphipod.Type) || (room1 != nil && room1.Type != amphipod.Type) {
				// Room is occupied by amphipod from other type
				// No possible move
				continue
			}

			if !burrow.RoomIsReachable(amphipod.Type.GetRoom(), amphipod.Position) {
				continue
			}

			if room1 != nil && room1.Type == amphipod.Type {
				// Other amphipod is already in final position
				// Generate move to own final position

				possibleMoves = append(possibleMoves, Move{
					AmphipodID: id,
					Room:       amphipod.Type.GetRoom(),
					Position:   0,
				})
			}

			if room1 == nil && room0 == nil {
				possibleMoves = append(possibleMoves, Move{
					AmphipodID: id,
					Room:       amphipod.Type.GetRoom(),
					Position:   1,
				})
			}

			continue
		}

		if amphipod.Type.MatchesRoom(amphipod.Room) {

			// Amphipod is in lower position
			if amphipod.Position == 1 {
				// amphipod already at final position
				continue
			}

			// If other matching amphipod is already in lower position,
			// and this amphipod is in upper position, we are also good.
			var otherAmphipodWithSameType *Amphipod
			for otherId, otherAmphipod := range burrow {
				if otherId != id && otherAmphipod.Type == amphipod.Type {
					otherAmphipodWithSameType = otherAmphipod
					break
				}
			}

			if otherAmphipodWithSameType.Type.MatchesRoom(otherAmphipodWithSameType.Room) && otherAmphipodWithSameType.Position == 1 && amphipod.Position == 0 {
				// both amphipods already at final position
				continue
			}
		}

		// Needs to move into hallway
		if amphipod.Position == 1 {
			nextSpotInRoom := burrow.GetAmphipodAtPosition(amphipod.Room, 0)

			if nextSpotInRoom != nil {
				// Path is blocked
				continue
			}
		}

		availableHallwayPositions := burrow.MovesToHallway(amphipod.Room)

		for _, hallwayPosition := range availableHallwayPositions {
			possibleMoves = append(possibleMoves, Move{
				AmphipodID: id,
				Room:       RoomHallway,
				Position:   hallwayPosition,
			})
		}
	}

	return possibleMoves
}

func (aT AmphipodType) GetRoom() Room {
	switch aT {
	case A:
		return RoomA
	case B:
		return RoomB
	case C:
		return RoomC
	case D:
		return RoomD
	}

	return RoomA
}

func (aT AmphipodType) MatchesRoom(room Room) bool {
	return aT.GetRoom() == room
}

func (b Burrow) GetAmphipodAtPosition(room Room, position int) *Amphipod {
	for _, amphipod := range b {
		if amphipod.Room == room && amphipod.Position == position {
			return amphipod
		}
	}

	return nil
}

func (b *Burrow) MovesToHallway(room Room) []int {
	availableHallwayPositions := []int{}

	for position := roomPositions[room]; position >= 0; position -= 1 {
		if position >= 2 && position <= 8 && position%2 == 0 {
			// Space is forbidden
			continue
		}

		if b.GetAmphipodAtPosition(RoomHallway, position) != nil {
			break
		}

		availableHallwayPositions = append(availableHallwayPositions, position)
	}

	for position := roomPositions[room]; position <= 10; position += 1 {
		if position >= 2 && position <= 8 && position%2 == 0 {
			// Space is forbidden
			continue
		}

		if b.GetAmphipodAtPosition(RoomHallway, position) != nil {
			break
		}

		availableHallwayPositions = append(availableHallwayPositions, position)
	}

	return availableHallwayPositions
}

func (b *Burrow) RoomIsReachable(room Room, positionInHallway int) bool {
	start := math.MaxInt
	end := math.MinInt

	if positionInHallway > roomPositions[room] {
		start = roomPositions[room]
		end = positionInHallway - 1
	} else {
		end = roomPositions[room]
		start = positionInHallway + 1
	}

	for i := start; i <= end; i++ {
		if b.GetAmphipodAtPosition(RoomHallway, i) != nil {
			return false
		}
	}

	return true
}

func (b *Burrow) IsFinished() bool {
	for _, amphipod := range b {
		if !amphipod.Type.MatchesRoom(amphipod.Room) {
			return false
		}
	}

	return true
}

func (b Burrow) Copy() Burrow {
	burrow := Burrow{}

	// Copy amphipods to new burrow
	for i, amphipod := range b {
		burrow[i] = &Amphipod{
			Type:     amphipod.Type,
			Room:     amphipod.Room,
			Position: amphipod.Position,
		}
	}

	return burrow
}

func (b Burrow) RunMoves(moves []Move) Burrow {
	burrow := b.Copy()

	for _, move := range moves {
		burrow[move.AmphipodID].Room = move.Room
		burrow[move.AmphipodID].Position = move.Position
	}

	return burrow
}

func CalculateCosts(burrow Burrow, moves []Move) int {
	totalCost := 0

	burrow = burrow.Copy()

	for _, move := range moves {
		amphipod := burrow[move.AmphipodID]

		steps := 0
		if move.Room == RoomHallway {
			// room -> hallway
			steps = amphipod.Position + 1 + int(math.Abs(float64(move.Position-roomPositions[amphipod.Room])))
		} else {
			// hallway -> room
			steps = int(math.Abs(float64(amphipod.Position-roomPositions[move.Room]))) + move.Position + 1
		}

		// Update burrow
		burrow[move.AmphipodID].Room = move.Room
		burrow[move.AmphipodID].Position = move.Position

		totalCost += steps * costsPerStep[amphipod.Type]
	}

	return totalCost
}
