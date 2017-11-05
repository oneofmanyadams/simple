package sconfiguration

import (
 	"simple/scsv"
)

func LoadConfigFromFile(file_location string) (configuration map[string]string) {
	csv_data := scsv.LoadFromFile(file_location)
	configuration = make(map[string]string)
	for _, config_record := range csv_data {
		configuration[config_record[0]] = config_record[1]
	}
	return
}

func SaveConfigToFile(file_location string, configuration map[string]string) {
	var configuration_records [][]string
	
	for record_name, record_value := range configuration {
		record_line := []string {record_name, record_value}
		configuration_records = append(configuration_records, record_line)
	}
	
	scsv.SaveToFile(file_location, configuration_records)
}