package csv

import (
	"fmt"
)

type CsvHandler struct {
	DataLoc string
	LocType string
	Lines [][]string
}


//////////////////////////
// Load Methods
//////////////////////////

func (csv *CsvHandler) Load() {
	if csv.LocType == "file" {
		csv.FileLoad()
	} else if csv.LocType == "folder" {
		csv.FolderLoad()
	} else {
		fmt.Println("Error loading, neither file or folder specified")	
	}
}

func (csv *CsvHandler) FileLoad() {
	csv.Lines = LoadFromFile(csv.DataLoc)
}

func (csv *CsvHandler) FolderLoad() {
	csv.Lines = LoadFromFolder(csv.DataLoc)
}

//////////////////////////
// Save Methods
//////////////////////////

func (csv *CsvHandler) Save() {
	if csv.LocType == "file" {
		csv.FileSave()
	} else if csv.LocType == "folder" {
		fmt.Println("Saving to multiple files not available yet")
		//csv.FileSave()
	} else {
		fmt.Println("Error loading, neither file or folder specified")	
	}
}

func (csv *CsvHandler) FileSave() {
	SaveToFile(csv.Lines, csv.DataLoc)
}

//////////////////////////
// Save Methods
//////////////////////////
