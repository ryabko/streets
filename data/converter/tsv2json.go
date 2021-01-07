package converter

import (
	"encoding/json"
	"io/ioutil"
	"ru.kalcho/streets/src/model"
)

type db struct {
	streets           []map[string]string
	suburbs           []map[string]string
	streetTypes       []map[string]string
	districts         []map[string]string
	streetsDistricts  []map[string]string
	streetsOtherNames []map[string]string
	stations          []map[string]string
}

// Call from main, for example: converter.ConvertTsv2Json("data/tsv", "data/json")
func ConvertTsv2Json(tsvDir string, jsonDir string) error {
	db, err := readDb(tsvDir)
	if err != nil {
		return err
	}

	streets, err := assembleStreets(db)
	if err != nil {
		return err
	}
	streetsJson, err := json.MarshalIndent(streets, "", "    ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(jsonDir+"/streets.json", streetsJson, 0644)
	if err != nil {
		return err
	}

	stops, err := assembleStops(db)
	if err != nil {
		return err
	}
	stopsJson, err := json.MarshalIndent(stops, "", "    ")
	err = ioutil.WriteFile(jsonDir+"/stops.json", stopsJson, 0644)
	if err != nil {
		return err
	}

	return nil
}

func assembleStreets(db db) ([]model.Street, error) {
	streets := make([]model.Street, len(db.streets))
	for i, streetRecord := range db.streets {
		street := model.Street{
			Name:        streetRecord["name"],
			Description: streetRecord["description"],
		}

		suburbId, suburbIsSet := streetRecord["suburb_id"]
		if suburbIsSet && suburbId != "" {
			suburbRecord := findRecordByFieldValue(db.suburbs, "id", suburbId)
			street.Suburb = suburbRecord["name"]
			street.SuburbDescription = suburbRecord["description"]
		}

		streetTypeId, streetTypeIdIsSet := streetRecord["street_type_id"]
		if streetTypeIdIsSet && streetTypeId != "" {
			streetTypeRecord := findRecordByFieldValue(db.streetTypes, "id", streetTypeId)
			street.Type = streetTypeRecord["name"]
			street.TypeShort = streetTypeRecord["abbreviation"]
		}

		districts := findRecordsByFieldValue(db.streetsDistricts, "street_id", streetRecord["id"])
		for _, streetDistrictRecord := range districts {
			districtRecord := findRecordByFieldValue(db.districts, "id", streetDistrictRecord["district_id"])
			street.Districts = append(street.Districts, districtRecord["name"])
		}

		street.AllNames = append(street.AllNames, model.StreetName{
			Name: streetRecord["name"],
		})
		otherNames := findRecordsByFieldValue(db.streetsOtherNames, "street_id", streetRecord["id"])
		for _, streetNameRecord := range otherNames {
			streetName := model.StreetName{
				Name:        streetNameRecord["name"],
				Description: streetNameRecord["description"],
				Obsolete:    streetNameRecord["obsolete"] == "1",
			}
			street.AllNames = append(street.AllNames, streetName)
		}

		streets[i] = street
	}

	return streets, nil
}

func readDb(dir string) (db, error) {
	streets, err := readTsv(dir + "/streets.tsv")
	if err != nil {
		return db{}, err
	}

	suburbs, err := readTsv(dir + "/suburbs.tsv")
	if err != nil {
		return db{}, err
	}

	streetTypes, err := readTsv(dir + "/streets_types.tsv")
	if err != nil {
		return db{}, err
	}

	districts, err := readTsv(dir + "/districts.tsv")
	if err != nil {
		return db{}, err
	}

	streetsDistricts, err := readTsv(dir + "/streets_districts.tsv")
	if err != nil {
		return db{}, err
	}

	streetsOtherNames, err := readTsv(dir + "/streets_other_names.tsv")
	if err != nil {
		return db{}, err
	}

	stations, err := readTsv(dir + "/stations.tsv")
	if err != nil {
		return db{}, err
	}

	return db{
		streets:           streets,
		suburbs:           suburbs,
		streetTypes:       streetTypes,
		districts:         districts,
		streetsDistricts:  streetsDistricts,
		streetsOtherNames: streetsOtherNames,
		stations:          stations,
	}, nil
}

func findRecordByFieldValue(records []map[string]string, field string, value string) map[string]string {
	for _, record := range records {
		if record[field] == value {
			return record
		}
	}
	return nil
}

func findRecordsByFieldValue(records []map[string]string, field string, value string) []map[string]string {
	foundRecords := make([]map[string]string, 0)
	for _, record := range records {
		if record[field] == value {
			foundRecords = append(foundRecords, record)
		}
	}
	return foundRecords
}

func assembleStops(db db) ([]model.Stop, error) {
	stops := make([]model.Stop, len(db.stations))
	for i, stationRecord := range db.stations {
		stop := model.Stop{
			Name:    stationRecord["name"],
			Comment: stationRecord["comment"],
		}
		stops[i] = stop
	}
	return stops, nil
}
