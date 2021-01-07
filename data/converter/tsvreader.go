package converter

import (
	"errors"
	"io/ioutil"
	"strings"
)

func readTsv(filename string) ([]map[string]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")
	if len(lines) == 0 {
		return nil, errors.New("empty file")
	}

	keys := strings.Split(lines[0], "\t")

	records := make([]map[string]string, len(lines)-1)

	for i, line := range lines[1:] {
		fields := strings.Split(line, "\t")
		record := make(map[string]string)
		for j, field := range fields {
			if field != "" && field != "NULL" {
				record[keys[j]] = field
			}
		}
		records[i] = record
	}

	return records, nil
}
