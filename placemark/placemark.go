package placemark

import (
	"encoding/xml"
	"fmt"
)

type Placemark struct {
	XMLName      xml.Name     `xml:"Placemark"`
	Point        Point        `xml:"Point"`
	ExtendedData ExtendedData `xml:"ExtendedData"`
}

type Point struct {
	Coordinates string `xml:"coordinates"`
}

type ExtendedData struct {
	Data []Data `xml:"Data"`
}

type Data struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value"`
}

// Initialize a new Placemark element
func New() *Placemark {
	return &Placemark{}
}

// Add coordinates element inside Point element
func (placemark *Placemark) AddCoordinates(longitude string, latitude string, altitude string) {
	placemark.Point = Point{Coordinates: fmt.Sprintf("%s,%s,%s", longitude, latitude, altitude)}
}

// Add ExtendedData element
func (placemark *Placemark) AddExtendedData(attributeName string, value string) {
	placemark.ExtendedData.Data = append(placemark.ExtendedData.Data, Data{Name: attributeName, Value: value})
}
