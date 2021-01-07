package converter

import (
	"reflect"
	"ru.kalcho/streets/src/model"
	"testing"
)

func TestAssembleStreets(t *testing.T) {
	db, err := readDb("../tsv")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	streets, err := assembleStreets(db)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	expectedLen := 1582
	if len(streets) != expectedLen {
		t.Fatalf(`len(streets) = %d, but wanted %d`, len(streets), expectedLen)
	}

	expected0 := model.Street{
		Name:      "1 апреля",
		Suburb:    "Отрадное",
		Type:      "улица",
		TypeShort: "ул.",
		AllNames:  []model.StreetName{{Name: "1 апреля"}},
	}
	street0 := streets[0]
	if !streetsAreEqual(expected0, street0) {
		t.Fatalf(`streets[0] = %v, but wanted %v`, street0, expected0)
	}

	expected7 := model.Street{
		Name:      "20 лет Октября",
		Type:      "улица",
		TypeShort: "ул.",
		Districts: []string{"20 лет Октября - Чапаева", "Острогожская - Цирк", "Центр"},
		AllNames:  []model.StreetName{{Name: "20 лет Октября"}},
	}
	street7 := streets[7]
	if !streetsAreEqual(expected7, street7) {
		t.Fatalf(`streets[7] = %v, but wanted %v`, street7, expected7)
	}

	expected225 := model.Street{
		Name:      "Почтовая",
		Type:      "улица",
		TypeShort: "ул.",
		Suburb:    "1 Мая",
		AllNames:  []model.StreetName{{Name: "Почтовая"}, {Name: "Гагарина", Description: "До 2011 года", Obsolete: true}},
	}
	street225 := streets[225]
	if !streetsAreEqual(expected225, street225) {
		t.Fatalf(`streets[225] = %v, but wanted %v`, street225, expected225)
	}

	expected1440 := model.Street{
		Name:              "Ленина",
		Suburb:            "Масловский",
		SuburbDescription: "Совхоз",
		Type:              "улица",
		TypeShort:         "ул.",
		AllNames:          []model.StreetName{{Name: "Ленина"}},
	}
	street1440 := streets[1440]
	if !streetsAreEqual(expected1440, street1440) {
		t.Fatalf(`streets[0] = %v, but wanted %v`, street1440, expected1440)
	}
}

func streetsAreEqual(expected model.Street, actual model.Street) bool {
	return expected.Name == actual.Name &&
		expected.Description == actual.Description &&
		expected.Type == actual.Type &&
		expected.TypeShort == actual.TypeShort &&
		expected.Suburb == actual.Suburb &&
		expected.SuburbDescription == actual.SuburbDescription &&
		reflect.DeepEqual(expected.Districts, actual.Districts) &&
		streetNamesAreEqual(expected.AllNames, actual.AllNames)
}

func streetNamesAreEqual(expected []model.StreetName, actual []model.StreetName) bool {
	if len(expected) != len(actual) {
		return false
	}
	for i, expectedItem := range expected {
		actualItem := actual[i]
		if expectedItem.Name != actualItem.Name ||
			expectedItem.Description != actualItem.Description ||
			expectedItem.Obsolete != actualItem.Obsolete {
			return false
		}
	}
	return true
}

func TestAssembleStops(t *testing.T) {
	db, err := readDb("../tsv")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	stops, err := assembleStops(db)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	expectedLen := 290
	if len(stops) != expectedLen {
		t.Fatalf(`len(stops) = %d, but wanted %d`, len(stops), expectedLen)
	}

	expected0 := model.Stop{
		Name:    "15 квартал",
		Comment: "ТЦ Московский проспект",
	}
	street0 := stops[0]
	if !stopsAreEqual(expected0, street0) {
		t.Fatalf(`stops[0] = %v, but wanted %v`, street0, expected0)
	}

	expected1 := model.Stop{
		Name: "17 квартал",
	}
	street1 := stops[1]
	if !stopsAreEqual(expected1, street1) {
		t.Fatalf(`stops[1] = %v, but wanted %v`, street1, expected1)
	}
}

func stopsAreEqual(expected model.Stop, actual model.Stop) bool {
	return expected.Name == actual.Name &&
		expected.Comment == actual.Comment
}
