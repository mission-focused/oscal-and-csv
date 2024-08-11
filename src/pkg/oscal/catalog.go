package oscal

import (
	"gopkg.in/yaml.v3"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-2"
)

func CatalogTemplate() (csv []string, err error) {

	csv = []string{"ID", "Group", "Class", "Title", "Params", "Props", "Links", "Parts", "Controls"}

	return csv, nil
}

// This function should receive an []byte, convert to Catalog, and then create rows for a catalog CSV file
func CatalogToCSV(data []byte) (records [][]string, err error) {

	catalog, err := NewCatalog(data)
	if err != nil {
		return records, err
	}

	if catalog.Groups != nil {
		tmpRecords := searchGroups(catalog.Groups)
		records = append(records, tmpRecords...)
	}

	if catalog.Controls != nil {
		tmpRecords := searchControls(catalog.Controls, "")
		records = append(records, tmpRecords...)
	}

	return records, nil
}

// This function should read a CSV file, and create groups and controls for each row
// will need to check that the group exists if the column is populated
func CSVToCatalog(records [][]string) (catalog oscalTypes.Catalog, err error) {

	return catalog, nil
}

// NewCatalog creates a new catalog object from the given data.
func NewCatalog(data []byte) (catalog *oscalTypes.Catalog, err error) {
	var oscalModels oscalTypes.OscalModels

	// unmarshal the catalog
	// yaml.v3 unmarshal handles both json and yaml
	err = yaml.Unmarshal(data, &oscalModels)
	if err != nil {
		return catalog, err
	}

	return oscalModels.Catalog, nil
}

func WriteCatalog(catalog oscalTypes.Catalog, filePath string) error {
	var model = oscalTypes.OscalModels{
		Catalog: &catalog,
	}

	err := WriteOSCALModel(filePath, &model)
	if err != nil {
		return err
	}
	return nil
}

func searchGroups(groups *[]oscalTypes.Group) [][]string {
	records := make([][]string, 0)

	for _, group := range *groups {
		if group.Groups != nil {
			tmpRecords := searchGroups(group.Groups)
			records = append(records, tmpRecords...)
		}
		if group.Controls != nil {
			tmpRecords := searchControls(group.Controls, group.ID)
			records = append(records, tmpRecords...)
		}
	}

	return records

}

func searchControls(controls *[]oscalTypes.Control, group string) [][]string {

	records := make([][]string, 0)

	// TODO: create string representations for the remaining fields

	for _, control := range *controls {
		record := []string{control.ID, group, control.Class, control.Title, "", "", "", "", ""}
		records = append(records, record)
	}

	return records
}
