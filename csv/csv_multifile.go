package csv

import (
	"io/ioutil"
	"fmt"
)

func LoadFromFolder(folder_location string) lines [][]string {
	// Open folder
	file_list, file_load_error := ioutil.ReadDir(folder_location)
	if file_load_error != nil {
		fmt.Println("Problem reading folder:", folder_location)
	}
	// loop through, make sure only read csv
	for _, data_file := range file_list {
		extension := data_file.Name()[len(data_file.Name())-4:]
		// make sure we only read data from csv files
		if extension == ".csv" {
			// Load data
			loaded_lines := LoadFromFile(folder_location+"/"+data_file.Name())
			lines = append(lines, loaded_lines...)
		}
	}
	return
}