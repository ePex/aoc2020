package main

import (
	"sort"
)

type BoardingPass struct {
	code   string
	row    int
	column int
	seatId int
}

func FindYourSeatId(batch []string) int {
	seatIds := getSeatIds(batch)
	for idx, seatId := range seatIds {
		if idx+1 != len(seatIds) && seatId+1 != seatIds[idx+1] {
			return seatId + 1
		}
	}

	return -1
}

func GetHighestSeatId(batch []string) int {
	seatIds := getSeatIds(batch)

	return seatIds[len(seatIds)-1]
}

func getSeatIds(batch []string) (seatIds []int) {
	for _, record := range batch {
		boardingPass := getBoardingPass(record)
		seatIds = append(seatIds, boardingPass.seatId)
	}
	sort.Ints(seatIds)

	return seatIds
}

func getBoardingPass(record string) (boardingPass BoardingPass) {
	rows := make([]int, 128)
	for i := 0; i < 128; i++ {
		rows[i] = i
	}
	for i := 0; i < 7; i++ {
		rows = getRow(rows, string(record[i]))
	}

	columns := make([]int, 8)
	for i := 0; i < 8; i++ {
		columns[i] = i
	}
	for i := 7; i < 10; i++ {
		columns = getColumn(columns, string(record[i]))
	}

	boardingPass = BoardingPass{
		code:   record,
		row:    rows[0],
		column: columns[0],
		seatId: rows[0]*8 + columns[0],
	}
	/*fmt.Printf("%v: row %v, column %v, seat ID %v\n",
	boardingPass.code, boardingPass.row, boardingPass.column, boardingPass.seatId)*/

	return boardingPass
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
