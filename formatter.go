package GoDcModule

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Patient struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func Do(inputPath, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	var patients []Patient
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var p Patient
		if err := decoder.Decode(&p); err != nil {
			return err
		}
		patients = append(patients, p)
	}

	data, err := json.MarshalIndent(patients, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputPath, data, 0666)
}
