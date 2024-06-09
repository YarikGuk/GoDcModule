package GoDcModule

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
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

	sort.Slice(patients, func(i, j int) bool {
		return patients[i].Age < patients[j].Age
	})

	data, err := json.MarshalIndent(patients, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputPath, data, 0666)
}
