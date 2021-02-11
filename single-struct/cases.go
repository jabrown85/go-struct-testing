package main

import (
	"fmt"
	"lifecycle/pkg/buildpack"
)

func main() {

	bpVersion := 1
	var layerTOML buildpack.Layer

	fmt.Println("******")
	fmt.Println("Case 1: (v1 Behavior)")
	layerTOML = buildpack.Parse(`{"name":"custom-layer"}`, bpVersion)
	doWork(layerTOML)

	fmt.Println("")
	fmt.Println("******")
	fmt.Println("Case 2: (v1 Behavior, disallows setting v2flag)")
	layerTOML = buildpack.Parse(`{"name":"custom-layer","v2flag":true}`, bpVersion)
	doWork(layerTOML)

	fmt.Println("")
	fmt.Println("******")
	fmt.Println("Case 2: with parsing mistake allows setting of value not in spec for this version")
	layerTOML = buildpack.ParseWithMistake(`{"name":"custom-layer","v2flag":true}`, bpVersion)
	doWork(layerTOML)

	fmt.Println("")
	fmt.Println("******")
	fmt.Println("Case 3: bumped bpVersion, v2 Behavior")
	bpVersion = 2
	layerTOML = buildpack.Parse(`{"name":"custom-layer"}`, bpVersion)
	doWork(layerTOML)

	fmt.Println("")
	fmt.Println("******")
	fmt.Println("Case 4: bumped bpVersion, v1 Behavior")
	bpVersion = 2
	layerTOML = buildpack.Parse(`{"name":"custom-layer","v2flag":false}`, bpVersion)
	doWork(layerTOML)
}

func doWork(layerTOML buildpack.Layer) {
	// fmt.Printf("Parsed as: %+v\n", layerTOML)
	// assume it isn't nil here o_0
	v2Behavior := *layerTOML.V2Behavior
	if v2Behavior {
		// we have allowed an old version to use this new feature
		fmt.Println("v2 Behavior")
	} else {
		fmt.Println("V1 Behavior")
	}
}
