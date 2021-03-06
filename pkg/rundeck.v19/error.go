package rundeck

import "encoding/xml"

// Error represents a rundeck xml error
type Error struct {
	XMLName    xml.Name `xml:"result"`
	Error      bool     `xml:"error,attr"`
	APIVersion string   `xml:"apiversion,attr"`
	Message    string   `xml:"error>message"`
}
