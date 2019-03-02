package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

/// A struct defining the data of a JSON used
type JSONData struct {
	Pk    string
	Score string
}

/// Returns an array of JSONData for a given File with filename
func JSONGet(filename string) []JSONData {
	file, err := os.Open(filename)

	if err != nil {
		log.Printf("[!] Couldn't open the file for reading. Cause %s",
			err.Error())
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var rows []JSONData = make([]JSONData, 0)

	for scanner.Scan() {

		var row JSONData
		var data []byte = scanner.Bytes()

		if len(data) != 0 {
			err = json.Unmarshal(data, &row)

			rows = append(rows, row)

			if err != nil {
				log.Printf("[!] Couldn't unmarshal data of file. Got %s Cause %s",
					scanner.Text(), err.Error())
			}
		}
	}

	return rows
}
