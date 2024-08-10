package oscal

import (
	"bytes"
	"encoding/json"
	"path/filepath"

	"github.com/defenseunicorns/go-oscal/src/pkg/files"
	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-2"
	yamlV3 "gopkg.in/yaml.v3"
)

func WriteOSCALModel(filePath string, model *oscalTypes.OscalModels) error {

	var b bytes.Buffer

	if filepath.Ext(filePath) == ".json" {
		jsonEncoder := json.NewEncoder(&b)
		jsonEncoder.SetIndent("", "  ")
		jsonEncoder.Encode(model)
	} else {
		yamlEncoder := yamlV3.NewEncoder(&b)
		yamlEncoder.SetIndent(2)
		yamlEncoder.Encode(model)
	}

	err := files.WriteOutput(b.Bytes(), filePath)
	if err != nil {
		return err
	}

	return nil
}
