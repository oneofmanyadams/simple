package csv

import (
	"io"
	"fmt"
	"os"
	"encoding/csv"
	"bufio"
)

func LoadFromFile(file_location string) (lines [][]string) {
	// Open file
	file, file_error := os.Open(file_location)
	if file_error != nil {
		// For now just print errror
		// Need to handle better in future
		fmt.Println("Error opening file: ", file_error)
	}
	defer file.Close()

	// Read file as CSV data
	csv_data := csv.NewReader(bufio.NewReader(file))

	// load line by line
	line_count := 0
	for ; ; line_count++ {
		line, err := csv_data.Read()
		// breaks loop at end of file
		if err == io.EOF {
			break
		}
		lines = append(lines, line)
	}
	return
}

func SaveToFile(file_location string, lines [][]string) {	
	// Create file
	file, file_error := os.Create(file_location)
	if file_error != nil {
		// For now just print errror
		// Need to handle better in future
		fmt.Println("Error opening file: ", file_error)
	}
	defer file.Close()
	
	// Write lines
	writer := csv.NewWriter(file)
	for _, wrrt_line := range lines {
		writer.Write(wrrt_line)
	}
	defer writer.Flush()
}
