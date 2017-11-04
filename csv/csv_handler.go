package csv

import (
	"fmt"
)

type CsvHandler struct {
	DataLocation string
	StorageType string
	Lines [][]string
}

//////////////////////////
// Load Methods
//////////////////////////

func (csv *CsvHandler) Load() {
	if csv.StorageType == "file" {
		csv.FileLoad()
	} else if csv.StorageType == "folder" {
		csv.FolderLoad()
	} else {
		fmt.Println("Error loading, neither file or folder specified")	
	}
}

func (csv *CsvHandler) FileLoad() {
	csv.Lines = LoadFromFile(csv.DataLocation)
}

func (csv *CsvHandler) FolderLoad() {
	csv.Lines = LoadFromFolder(csv.DataLocation)
}

//////////////////////////
// Save Methods
//////////////////////////

func (csv *CsvHandler) Save() {
	if csv.StorageType == "file" {
		csv.FileSave()
	} else if csv.StorageType == "folder" {
		fmt.Println("Saving to multiple files not available yet")
		//csv.FileSave()
	} else {
		fmt.Println("Error loading, neither file or folder specified")	
	}
}

func (csv *CsvHandler) FileSave() {
	SaveToFile(csv.DataLocation, csv.Lines)
}
