package rundeck

import "encoding/xml"

// Events represents a collection of `Event`
type Events struct {
	XMLName xml.Name `xml:"events"`
	Count   int64    `xml:"count,attr"`
	Total   int64    `xml:"total,attr"`
	Max     int64    `xml:"max,attr"`
	Offset  int64    `xml:"offset,attr"`
	Events  []Event  `xml:"event"`
}

// Event represents an Event
type Event struct {
	XMLName     xml.Name `xml:"event"`
	StartTime   string   `xml:"starttime,attr"`
	EndTime     string   `xml:"endtime,attr"`
	Title       string   `xml:"title"`
	Status      string   `xml:"status"`
	Summary     string   `xml:"summary"`
	NodeSummary struct {
		XMLName   xml.Name
		Succeeded int64 `xml:"succeeded,attr"`
		Failed    int64 `xml:"failed,attr"`
		Total     int64 `xml:"total,attr"`
	} `xml:"node-summary"`
	User        string `xml:"user"`
	Project     string `xml:"project"`
	DateStarted string `xml:"date-started"`
	DateEnded   string `xml:"date-ended"`
	AbortedBy   string `xml:"abortedby,omitempty"`
	Job         *struct {
		XMLName xml.Name
		ID      string `xml:"id,attr"`
	} `xml:"job,omitempty"`
	Execution struct {
		XMLName xml.Name
		ID      int64 `xml:"id,attr"`
	} `xml:"execution,omitempty"`
}

// GetHistory returns the history for a project
func (c *Client) GetHistory(project string) (Events, error) {
	u := make(map[string]string)
	u["project"] = project
	var data Events
	var res []byte
	if err := c.Get(&res, "history", u); err != nil {
		return data, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}
