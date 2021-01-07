package converter

import (
	"testing"
)

func TestRead(t *testing.T) {
	filename := "../tsv/suburbs.tsv"
	records, err := readTsv(filename)
	if err != nil {
		t.Fatalf(`Cannnot read %q. Error: %v`, filename, err)
	}

	expectedLen := 32
	if len(records) != expectedLen {
		t.Fatalf(`len(records) = %d, but wanted %d`, len(records), expectedLen)
	}

	expectedRecord0 := map[string]string{
		"id":   "1",
		"name": "Подгорное",
	}
	record0 := records[0]
	if record0["id"] != expectedRecord0["id"] ||
		record0["name"] != expectedRecord0["name"] ||
		record0["description"] != expectedRecord0["description"] {

		t.Fatalf(`records[0] = %v, but wanted %v`, record0, expectedRecord0)
	}

	expectedRecord19 := map[string]string{
		"id":   "20",
		"name": "Труд",
	}
	record19 := records[19]
	if record19["id"] != expectedRecord19["id"] ||
		record19["name"] != expectedRecord19["name"] ||
		record19["description"] != expectedRecord19["description"] {

		t.Fatalf(`records[19] = %v, but wanted %v`, record19, expectedRecord19)
	}

	expectedRecord29 := map[string]string{
		"id":          "31",
		"name":        "Масловский",
		"description": "Совхоз",
	}
	record29 := records[29]
	if record29["id"] != expectedRecord29["id"] ||
		record29["name"] != expectedRecord29["name"] ||
		record29["description"] != expectedRecord29["description"] {

		t.Fatalf(`records[29] = %v, but wanted %v`, record29, expectedRecord29)
	}

}
