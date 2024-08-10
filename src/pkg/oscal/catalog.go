package oscal

import oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-2"

func CatalogToTemplate() (csv []string, err error) {

	csv = []string{"ID", "Group", "Class", "Title", "Params", "Props", "Links", "Parts", "Controls"}

	return csv, nil
}

// This function should receive an []byte, convert to Catalog, and then create rows for a catalog CSV file
func CatalogToCSV(data []byte) (records [][]string, err error) {

	return records, nil
}

// This function should read a CSV file, and create groups and controls for each row
// will need to check that the group exists if the column is populated
func CSVToCatalog(records [][]string) (catalog oscalTypes.Catalog, err error) {

	return catalog, nil
}
