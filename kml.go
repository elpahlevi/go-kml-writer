package kml

import (
	"encoding/xml"
	"os"
)

type KML struct {
	XMLName  xml.Name `xml:"kml"`
	Xmlns    string   `xml:"xmlns,attr"`
	Document []any    `xml:"Document"`
}

// Initialize a new KML instance
func New() *KML {
	return &KML{
		Xmlns: "http://www.opengis.net/kml/2.2",
	}
}

// Append all Document element child
func (kml *KML) AddToDocument(child any) {
	kml.Document = append(kml.Document, child)
}

// Marshal input to KML
func (kml *KML) MarshalIndent(prefix string, indent string) ([]byte, error) {
	return xml.MarshalIndent(kml, prefix, indent)
}

// Save the output as a KML file
func (kml *KML) WriteOutput(marshaled []byte, filename string) (string, error) {
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(marshaled)
	if err != nil {
		return "", err
	}
	return file.Name(), err
}
