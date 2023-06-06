package models

import (
	"encoding/csv"
	"os"
)

// FileName declares how the File is called, that is used to get the data from
const FileName = "data.csv"

// SampleDB declares how the File is called, where the Samples are stored for another run
const SampleDB = "sampleDB.csv"

// DelimiterInDataStorage defines how the values are seperated
const DelimiterInDataStorage = ","

// Data stored in File or in Program
var filePersistence bool = false

// EnableFilePersistence enables the file persistence
func EnableFilePersistence() {
	filePersistence = true
}

// DisableFilePersistence disables the file persistence
func DisableFilePersistence() {
	filePersistence = false
}

// gets the data from the previous application run.
func getDataFromDB() ([][]string, error) {
	file, err := os.OpenFile(SampleDB, os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = 4
	records, err := csvReader.ReadAll()
	return records, err
}
func updateDB(data [][]string) error {
	file, err := os.Create(SampleDB)
	defer file.Close()
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(file)
	err = csvWriter.WriteAll(data)
	return err
}

// UpdateSampleCSV gets data from a new File if there are any new Files to be entered.
func UpdateDataNewDataFile() ([][]string, error) {
	file, err := os.OpenFile(FileName, os.O_RDWR, 0755)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = 4
	records, err := csvReader.ReadAll()
	return records, err
}

func NewEntry(dataNew []string) error {
	var data [][]string
	data = append(data, dataNew)
	file, err := os.OpenFile(FileName, os.O_RDWR, 0755)
	defer file.Close()
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(file)
	err = csvWriter.WriteAll(data)
	return err
}

func GetTable() ([][]string, error) {
	var TempSampleList [][]string
	//get stored samples
	oldSamples, err := getDataFromDB()
	if err != nil {
		return nil, err
	}
	//get new samples
	newSamples, err := UpdateDataNewDataFile()
	if err != nil {
		return nil, err
	}
	//merge old and new
	for _, v := range newSamples {
		TempSampleList = append(TempSampleList, v)
	}
	for _, v := range oldSamples {
		TempSampleList = append(TempSampleList, v)
	}
	for i := 0; i <= 10; i++ {
		TempSampleList = append(TempSampleList, []string{"open", "open", "open", "open"})
	}
	//overwrite the DB with new values
	updateDB(TempSampleList[:10])
	//take the newest 10 & return Table
	return TempSampleList[:10], err
}
