package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Configer struct{}

type Configuration struct {
	NrOfRecords int
	FilePath    string
}

func (c *Configer) RecordsConfig() (*Configuration, error) {
	recordsNr, err := getRecordsNr()
	if err != nil {
		return nil, fmt.Errorf("error configuring records number: %v", err)
	}
	filePath, err := getFilePath()
	if err != nil {
		return nil, fmt.Errorf("error configuring file path: %v", err)
	}

	return &Configuration{NrOfRecords: recordsNr, FilePath: filePath}, nil
}

func getRecordsNr() (int, error) {
	fmt.Println("Set the nr of records you want to read: ")
	var recordsNr int
	_, err := fmt.Scanln(&recordsNr)

	return recordsNr, err
}
func getFilePath() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("(Leave empty for current dir) Set address of file: ")
	filePath, err := reader.ReadString('\n')
	if filePath == "" || filePath == "\n" || filePath == "\r\n" {
		filePath = "records.json"
	} else {
		filePath = strings.Replace(filePath, "\r\n", "", -1)
	}
	return filePath, err

}
