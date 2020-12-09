package main

import "fmt"

func GetHighestSeatId(batch []string) (highestSeatId int) {
	for _, boardingPass := range batch {
		seatId := GetSeatId(boardingPass)
		if seatId > highestSeatId {
			highestSeatId = seatId
		}
	}

	return highestSeatId
}

func GetSeatId(boardingPass string) int {
	rows := make([]int, 128)
	for i := 0; i < 128; i++ {
		rows[i] = i
	}
	for i := 0; i < 7; i++ {
		rows = getRow(rows, string(boardingPass[i]))
	}

	columns := make([]int, 8)
	for i := 0; i < 8; i++ {
		columns[i] = i
	}
	for i := 7; i < 10; i++ {
		columns = getColumn(columns, string(boardingPass[i]))
	}

	seatId := rows[0]*8 + columns[0]

	fmt.Printf("%v: row %v, column %v, seat ID %v\n", boardingPass, rows[0], columns[0], seatId)

	return seatId
}

func getRow(rows []int, character string) []int {
	switch character {
	default:
		return nil
	case "F":
		return rows[:(len(rows) / 2)]
	case "B":
		return rows[(len(rows) / 2):]
	}
}

func getColumn(columns []int, character string) []int {
	switch character {
	default:
		return nil
	case "L":
		return columns[:(len(columns) / 2)]
	case "R":
		return columns[(len(columns) / 2):]
	}
}
