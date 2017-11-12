package sdb

import (
	"fmt"
	"simple/sconfiguration"
)

type Sdb struct {
	Folder string
	TableConfigFiles map[string]string
}

//////////////////////////////////////////
// Table functions
//////////////////////////////////////////

func (sdb *Sdb) UseTable(table_name string) (table Table) {
	if config_file_location, table_exists := sdb.TableConfigFiles[table_name]; table_exists {
		table.ConfigFile = config_file_location
		table.LoadConfig()
		table.LoadAllRows()
	} else {
		fmt.Println("Table", table_name, "does not exist")
	}
	return
}

func (sdb *Sdb) CreateTable(table_name string) (table Table) {
	if _, table_exists := sdb.TableConfigFiles[table_name]; table_exists {
		fmt.Println("Table", table_name, "already exists")
	} else {
		sdb.TableConfigFiles[table_name] = sdb.Folder+"/"+table_name+"_config.csv"
		sdb.SaveConfig()
		
		table.ConfigFile = sdb.Folder+"/"+table_name+"_config.csv"
		table.DataLocation = sdb.Folder+"/"+table_name+"_data.csv"
		table.NextId = 0
		table.SaveConfig()
		table.SaveAllRows()
	}
	return
}

//////////////////////////////////////////
// Basic functions
//////////////////////////////////////////

func (sdb *Sdb) SaveConfig() {
	sconfiguration.SaveConfigToFile(sdb.Folder+"/tables.csv", sdb.TableConfigFiles)
}

func UseSdb(sdb_location string) (sdb Sdb) {
	sdb.Folder = sdb_location
	sdb.TableConfigFiles = sconfiguration.LoadConfigFromFile(sdb_location+"/tables.csv")
	return
}

func CreateSdb(sdb_location string) (sdb Sdb) {
	sdb.Folder = sdb_location
	
	empty_initial_data := make(map[string]string)
	
	sconfiguration.SaveConfigToFile(sdb_location+"/tables.csv", empty_initial_data)
	sdb.TableConfigFiles = sconfiguration.LoadConfigFromFile(sdb_location+"/tables.csv")
	return
}