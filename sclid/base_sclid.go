package sclid

import (
	"fmt"
	"strings"
)


func Display(data [][]string) {
	column_max_lengths := determineMaxLengths(data)
	
	for _, row := range data {
		printFormattedRow(row, column_max_lengths)
	}

}

func printFormattedRow(row []string, column_max_lengths map[int]int) {
	var print_string string

	for column, data := range row {
		print_string = print_string + data
		print_string = print_string + strings.Repeat(" ", (column_max_lengths[column]-len(data)))
		print_string = print_string + "	"
	}
	
	fmt.Println(print_string)
}

func determineMaxLengths(data [][]string) (max_lengths map[int]int) {
	max_lengths = make(map[int]int)
	
	for _, row := range data {
		for column_number, column_length := range getColumnLengths(row) {
			if current_max_length, column_recorded := max_lengths[column_number]; column_recorded {
				if column_length > current_max_length {
					max_lengths[column_number] = column_length
				}
			} else {
				max_lengths[column_number] = column_length
			}
		}
	}
	return
}

func getColumnLengths(row []string) (lengths []int) {
	for _, column := range row {
		lengths = append(lengths, len(column))
	}
	return
}