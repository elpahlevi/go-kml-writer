# Go KML Writer

Go KML Writer is a powerful Go (Golang) package designed to simplify the creation of KML (Keyhole Markup Language) files. KML is a popular XML-based format used to represent geographical data and is widely supported by mapping and Geographic Information System (GIS) software.

## Features
Create KML documents such as:
1. Placemarks

Notes: Other KML documents will be added in the next release

## Getting Started
Install packages
```bash
go get github.com/elpahlevi/go-kml-writer
```
Follow the examples in the "Usage Example" section or go to `kml_test.go` file for more.

## Usage Example
```go
package main

import (
    "github.com/elpahlevi/go-kml-writer"
)

func main() {
    kml := kml.New()
	pm := placemark.New()
	pm.AddCoordinates("113.97395", "-2.99994", "0")
	pm.AddExtendedData("acqdatetimetz", "2023-10-18 09:09:00")
	pm.AddExtendedData("confidence", "80")

	kml.AddToDocument(pm)

	marshaled, err := kml.MarshalIndent("", "  ")
	if err != nil {
		panic(err)
	}
    
	kml.WriteOutput(marshaled, "./output/output.kml")
}
```
## License
Copyright (c) 2023-present <a href="https://github.com/elpahlevi">Reza Pahlevi</a>. `Go KML Writer` is open-source and licensed under the MIT License.