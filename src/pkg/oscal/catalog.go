package oscal

func CatalogToTemplate() (csv []string, err error) {

	return []string{"ID", "Group", "Class", "Title", "Params", "Props", "Links", "Parts", "Controls"}, nil
}
