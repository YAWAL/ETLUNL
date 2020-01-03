package extract

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/YAWAL/ETLUNL/model"
)

// Extract
func Extract(filePath string) ([]model.RawData, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ProcessRawData(file), nil
}

func ProcessRawData(file io.Reader) []model.RawData {
	var rawCSVData []model.RawData
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		rawRecord := fillRawData(record)
		rawCSVData = append(rawCSVData, rawRecord)
	}
	return rawCSVData[1 : len(rawCSVData)-1]
}

func fillRawData(record []string) (rawRecord model.RawData) {
	for _, rec := range record {
		features := strings.Split(rec, ";")
		fmt.Println(len(features))
		fmt.Println(features)
		rawRecord.Game = features[0]
		rawRecord.Date = features[1]
		rawRecord.Lototron = features[2]
		rawRecord.BallSet = features[3]
		rawRecord.Ball1 = features[4]
		rawRecord.Ball2 = features[5]
		rawRecord.Ball3 = features[6]
		rawRecord.Ball4 = features[7]
		rawRecord.Ball5 = features[8]
		rawRecord.Ball6 = features[9]
		rawRecord.Ball2Winners = features[10]
		rawRecord.Ball2Price = features[11]
		rawRecord.Ball3Winners = features[12]
		rawRecord.Ball3Price = features[13]
		rawRecord.Ball4Winners = features[14]
		rawRecord.Ball4Price = features[15]
		rawRecord.Ball5Winners = features[16]
		rawRecord.Ball5Price = features[17]
		rawRecord.Ball6Winners = features[18]
		rawRecord.Ball6Price = features[19]
	}
	return rawRecord
}
