package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FileConfiger struct {
	Configuration
}

// NewFileConfiger constructs a new FileConfiger instance.
func NewFileConfiger() *FileConfiger {
	return &FileConfiger{} // init diff?
}

// Configuration is the current setup for the service. It includes
// the nr of records to be read and the file path that the users wants to write to.
type Configuration struct {
	NrOfRecords int
	FilePath    string
}

// ConfigRecords gets the nr of records and the file path that the user
// inputs at the beginning of the program.
func (fc *FileConfiger) ConfigRecords() (*FileConfiger, error) {
	recordsNr, err := getRecordsNr()
	if err != nil {
		return nil, fmt.Errorf("error configuring records number: %w", err)
	}
	filePath, err := getFilePath()
	if err != nil {
		return nil, fmt.Errorf("error configuring file path: %w", err)
	}

	return &FileConfiger{Configuration{NrOfRecords: recordsNr, FilePath: filePath}}, nil
}

// getRecordsNr is the getter method for how many records from the API call the user
// wants to read and process.
func getRecordsNr() (int, error) {
	fmt.Println("Set the nr of records you want to read: ")
	var recordsNr int
	_, err := fmt.Scanln(&recordsNr)

	return recordsNr, err
}

// getFilePath is the getter method for the location in which the user wants to save
// the transformed file. It defaults to the current directory.
func getFilePath() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fileName := "records.json"
	fmt.Println("(Leave empty for current dir) Set address of file: ")
	filePath, err := reader.ReadString('\n')
	if filePath == "" || filePath == "\n" || filePath == "\r\n" {
		filePath = fileName
	} else {
		filePath = strings.Replace(filePath+`\`+fileName, "\r\n", "", -1)
	}
	return filePath, err

}
