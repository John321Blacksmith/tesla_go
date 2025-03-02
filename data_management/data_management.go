package data_management

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetDataset(filename string) []Category {
	var dataset []Category
	// form contents of the file to bytes
	file_bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error occurred processiong the file %v\n", filename)
	} else {
		// transform the data from bytes to the datatype
		err = json.Unmarshal(file_bytes, &dataset)
		if err != nil {
			fmt.Printf("Error while reading the file %v, Error: %v\n", filename, err)
		} else {
			return dataset
		}
	}
	return nil
}
