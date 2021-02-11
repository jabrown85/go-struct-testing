package buildpack

import (
	"encoding/json"
)

type Layer struct {
	// version 1
	Name string `json:"name"`
	// version 2
	V2Behavior *bool `json:"v2flag"`
}

func Parse(jsonString string, bpVersion int) Layer {
	var layerTOML Layer
	json.Unmarshal([]byte(jsonString), &layerTOML)

	// to default as true, we have to use a pointer and that pointer bleeds out to the users of this struct
	// we could obviously invert the flag behavior/naming and we tend to do that
	// we do HAVE to set a value for all versions, otherwise you are ParseWithMistake
	if bpVersion >= 2 && layerTOML.V2Behavior == nil {
		defaultValue := true
		layerTOML.V2Behavior = &defaultValue
	} else {
		defaultValue := false
		layerTOML.V2Behavior = &defaultValue
	}

	return layerTOML
}

func ParseWithMistake(jsonString string, bpVersion int) Layer {
	var layerTOML Layer
	json.Unmarshal([]byte(jsonString), &layerTOML)

	// subtle bug
	// writing this type of guard would allow
	// bpVersion 1 to set dosomething to whatever value
	// and get that behavior
	if bpVersion >= 2 {
		defaultValue := true
		layerTOML.V2Behavior = &defaultValue
	}

	return layerTOML
}
