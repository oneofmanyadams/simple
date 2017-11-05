package sconfiguration

import (
 	"fmt"
 	"simple/scsv"
)

func LoadConfigFromFileWithFlags(file_location string) (settings map[string]string, flags map[string]string) {
	csv_data := scsv.LoadFromFile(file_location)
	settings = make(map[string]string)
	flags = make(map[string]string)

	for _, config_record := range csv_data {
		if len(config_record) == 3 {
			settings[config_record[0]] = config_record[1]
			flags[config_record[0]] = config_record[2]
		} else {
			fmt.Println("configuration record does not have flag")
		}
	}
	return
}

func SaveConfigToFileWithFlags(file_location string, settings map[string]string, flags map[string]string) {
	
	if TestSettingsAndFlagsMatch(settings, flags) {
	
		var configuration_records [][]string
	
		for record_name, record_value := range settings {
			record_line := []string {record_name, record_value, flags[record_name]}
			configuration_records = append(configuration_records, record_line)
		}
		
		scsv.SaveToFile(file_location, configuration_records)
	}
}

func TestSettingsAndFlagsMatch(settings map[string]string, flags map[string]string) (result bool) {
	for setting_name, _ := range settings {
		if _, flag_exists := flags[setting_name]; flag_exists {
			result = true
		} else {
			result = false
			return
		}
	}

	for flag_name, _ := range flags {
		if _, setting_exists := settings[flag_name]; setting_exists {
			result = true
		} else {
			result = false
			return
		}
	}
	
	return
}