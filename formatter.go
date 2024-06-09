package GoDcModule

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
	"sort"
)

type Patient struct {
	Name  string `json:"name" xml:"Name"`
	Age   int    `json:"age" xml:"Age"`
	Email string `json:"email" xml:"Email"`
}

type Patients struct {
	XMLName  xml.Name  `xml:"patients"`
	Patients []Patient `xml:"Patient"`
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

	data, err := xml.MarshalIndent(Patients{Patients: patients}, "", "    ")
	if err != nil {
		return err
	}

	xmlHeader := []byte(xml.Header)
	data = append(xmlHeader, data...)

	return ioutil.WriteFile(outputPath, data, 0666)
}
