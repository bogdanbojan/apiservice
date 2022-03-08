package config

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"strings"
)

// NewConfig constructs a new Config instance.
func NewConfig() *Config {
	return &Config{}
}

// Config is the current setup for the service. It includes
// the nr of records to be collectrec and the file path that the users wants to writerec to.
type Config struct {
	RecordsNr int
	FilePath  string
	SourceURL string
}

// ConfigRecords gets the nr of records and the file path that the user
// inputs at the beginning of the program.
func (c *Config) Init() error {
	rn, err := getRecordsNr()
	if err != nil {
		return fmt.Errorf("cannot configure records number: %w", err)
	}
	fp, err := getFilePath()
	if err != nil {
		return fmt.Errorf("cannot configure file path: %w", err)
	}
	su, err := getSourceURL()
	if err != nil {
		return fmt.Errorf("cannot configure envfile: %w", err)
	}

	c.RecordsNr = rn
	c.FilePath = fp
	c.SourceURL = su

	return nil
}

// getRecordsNr is the getter method for how many records from the API call the user
// wants to collect and process.
func getRecordsNr() (int, error) {
	fmt.Println("Set the nr of records you want to collect: ")
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

//go:embed service.env
var sourceURL embed.FS

// getEnvURL fetches the URL from the .env file.
func getSourceURL() (string, error) {
	key := "SOURCE_URL"
	u, err := sourceURL.ReadFile("service.env")
	if err != nil {
		return "", err
	}
	trimmedPrefixURL := strings.TrimPrefix(string(u), key+"=")
	return trimmedPrefixURL, nil
}
