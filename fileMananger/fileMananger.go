package filemananger

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileMananger struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileMananger) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("error occured when opening the file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		// file.Close()
		return nil, errors.New("error occured when scanning the file")
	}
	return lines, nil
}

func (fm FileMananger) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to create file")
	}

	time.Sleep(3 * time.Second) // to simulate  a slow file writing process

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		// file.Close()
		return errors.New("failed to convert data to json")
	}

	// file.Close()
	return nil
}

func New(inputPath, outputPath string) FileMananger {
	return FileMananger{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
