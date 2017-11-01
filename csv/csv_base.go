package csv

import (
	"io"
	"io/ioutil"
	"fmt"
	"os"
	"encoding/csv"
	"bufio"
)

func Load(location string, file_or_folder string) [][]string {
	if file_or_folder == "file" {
		return LoadFromFile(location)
	}
	if file_or_folder == "folder" {
		return LoadFromFolder(location)
	}
	var null_return [][]string
	fmt.Println("Error loading, neither file or folder specified")
	return null_return
}

func LoadFromFolder(folder_location string) [][]string {
	// Open folder
	file_list, file_load_error := ioutil.ReadDir(folder_location)
	if file_load_error != nil {
		fmt.Println("Problem reading folder:", folder_location)
	}
	// loop through, make sure only read csv
	var lines [][]string
	for _, data_file := range file_list {
		extension := data_file.Name()[len(data_file.Name())-4:]
		// make sure we only read data from csv files
		if extension == ".csv" {
			// Load data
			loaded_lines := FromFile(folder_location+"/"+data_file.Name())
			lines = append(lines, loaded_lines...)
		}
	}
	return lines
}

func LoadFromFile(file_location string) [][]string {
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
	var lines [][]string
	line_count := 0
	for ; ; line_count++ {
		line, err := csv_data.Read()
		// breaks loop at end of file
		if err == io.EOF {
			break
		}
		lines = append(lines, line)
	}
	return lines
}

func Write(lines [][]string, location string) {
	
	file, _ := os.Create(location)
	defer file.Close()
	
	writer := csv.NewWriter(file)

	for _, wrrt_line := range lines {
		writer.Write(wrrt_line)
	}
	
	defer writer.Flush()
}
