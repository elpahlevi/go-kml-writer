package kml_test

import (
	"fmt"
	"testing"

	"github.com/elpahlevi/go-kml-writer"
	"github.com/elpahlevi/go-kml-writer/placemark"
)

func TestCreatePlacemark(t *testing.T) {
	pm := placemark.New()
	pm.AddCoordinates("113.97395", "-2.99994", "0")
	pm.AddExtendedData("acqdatetimetz", "2023-10-18 09:09:00")
	pm.AddExtendedData("confidence", "80")

	kml := kml.New()
	kml.AddToDocument(pm)

	marshaled, err := kml.MarshalIndent("", "  ")
	if err != nil {
		panic(err)
	}

	res, err := kml.WriteOutput(marshaled, "./output/output.kml")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestCreateMultiplePlacemark(t *testing.T) {
	// create a slice of Placemark struct to store a multiple placemark
	var pms []placemark.Placemark

	// you can use any kinds of data as long as it's written in struct form
	type hotspotData struct {
		latitude      string
		longitude     string
		altitude      string
		acqdatetimetz string
		confidence    string
	}

	hotspots := []hotspotData{
		{
			latitude:      "-2.99994",
			longitude:     "113.97395",
			altitude:      "0",
			acqdatetimetz: "2023-10-18 09:09:00",
			confidence:    "80",
		},
		{
			latitude:      "-2.98971",
			longitude:     "113.97543",
			altitude:      "0",
			acqdatetimetz: "2023-10-18 09:09:00",
			confidence:    "HIGH",
		},
	}

	for _, hotspot := range hotspots {
		pm := placemark.New()
		pm.AddCoordinates(hotspot.longitude, hotspot.latitude, hotspot.altitude)
		pm.AddExtendedData("acqdatetimetz", hotspot.acqdatetimetz)
		pm.AddExtendedData("confidence", hotspot.confidence)

		pms = append(pms, *pm)
	}

	kml := kml.New()
	kml.AddToDocument(pms)

	marshaled, err := kml.MarshalIndent("", "  ")
	if err != nil {
		panic(err)
	}

	res, err := kml.WriteOutput(marshaled, "/Users/yourusername/output.kml")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
