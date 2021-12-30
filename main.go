package main

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	taskData := parseTask()
	switch taskData.Mode {
	default:
		log.Fatalln("unknown mode given. usage example: ./b64 encode from.txt to.dat")
	case "encode":
		err := encode(taskData)
		if err != nil {
			log.Fatalln(err)
		}
	case "decode":
		err := decode(taskData)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

type task struct {
	Mode     string
	FromFile string
	ToFile   string
}

func parseTask() task {
	result := task{}
	// example: sys encode from.ext to.ext
	if len(os.Args) < 4 {
		result.Mode = "unknown"
		return result
	}
	result.Mode = os.Args[1]
	result.FromFile = os.Args[2]
	result.ToFile = os.Args[3]
	return result
}

func readFile(filepath string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.New("failed to read file: " + err.Error())
	}
	return fileBytes, nil
}

func encode(taskData task) error {
	// read file
	fileBytes, err := readFile(taskData.FromFile)
	if err != nil {
		return err
	}

	// encode
	stringResult := base64.StdEncoding.EncodeToString(fileBytes)

	// write file
	return writeFile(taskData.ToFile, []byte(stringResult))
}

func writeFile(filepath string, data []byte) error {
	err := ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		return errors.New("failed to write file: " + err.Error())
	}
	return nil
}

func decode(taskData task) error {
	// read file
	fileBytes, err := readFile(taskData.FromFile)
	if err != nil {
		return err
	}

	// decode
	result, err := base64.StdEncoding.DecodeString(string(fileBytes))
	if err != nil {
		return errors.New("failed to decode base64: " + err.Error())
	}

	// write file
	return writeFile(taskData.ToFile, result)
}
