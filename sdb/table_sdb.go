package sdb

import (
	"simple/scsv"
	"simple/sconfiguration"
	"strconv"
)

type Table struct {
	ConfigFile string
	DataLocation string
	NextId int
	Rows map[int][]string
}

//////////////////////////////////////////
// Config functions
//////////////////////////////////////////

func (table *Table) LoadConfig() {
	config := sconfiguration.LoadConfigFromFile(table.ConfigFile)
	table.DataLocation = config["DataLocation"]
	table.NextId, _ = strconv.Atoi(config["NextId"])
}

func (table *Table) SaveConfig() {
	config := make(map[string]string)
	
	config["DataLocation"] = table.DataLocation
	config["NextId"] = strconv.Itoa(table.NextId)

	sconfiguration.SaveConfigToFile(table.ConfigFile, config)	
}

//////////////////////////////////////////
// All Rows function
//////////////////////////////////////////

func (table *Table) LoadAllRows() {
	idd_rows := make(map[int][]string)
	raw_rows := scsv.LoadFromFile(table.DataLocation)
	
	for _, row := range raw_rows {
		row_id, _ := strconv.Atoi(row[0])
		idd_rows[row_id] = row
	}
	
	table.Rows = idd_rows
}

func (table *Table) SaveAllRows() {
	var raw_rows [][]string
	for _, row := range table.Rows {
		raw_rows = append(raw_rows, row)
	}
	
	scsv.SaveToFile(table.DataLocation, raw_rows)
}

//////////////////////////////////////////
// Rows function
//////////////////////////////////////////

func (table *Table) AddRow(added_row []string) (row_id int) {
	var added_row_with_id []string
	
	added_row_with_id = append(added_row_with_id, strconv.Itoa(table.NextId))
	added_row_with_id = append(added_row_with_id, added_row...)
	
	table.Rows[table.NextId] = added_row_with_id
	row_id = table.NextId
	table.NextId = table.NextId + 1
	
	table.SaveConfig()
	table.SaveAllRows()
	
	return
}

func (table *Table) ChangeRow(row_id int, change_row []string) {
	var change_row_with_id []string
	
	change_row_with_id = append(change_row_with_id, strconv.Itoa(row_id))
	change_row_with_id = append(change_row_with_id, change_row...)
	
	table.Rows[row_id] = change_row_with_id
	
	table.SaveConfig()
	table.SaveAllRows()
}

