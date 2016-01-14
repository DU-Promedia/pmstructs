package pmstructs

import (
	"encoding/xml"
	//	"log"
	//	"strconv"
)

type CSResponse struct {
	Root  xml.Name `xml:"reportitems"`
	Count int      `xml:",count"`
	Items []struct {
		Title   string `xml:"title"`
		Columns []struct {
			Title string `xml:"ctitle"`
			Type  string `xml:"type"`
		} `xml:"columns>column"`
		Rows []struct {
			Column []string `xml:"c"`
		} `xml:"rows>r"`
	} `xml:"reportitem"`
}

type CSResponseMostRead struct {
	ID        string
	Pageviews int
	Browsers  int
	Virality  float64
}

func (c *CSResponse) GetMostRead(a_site string) []string {
	output := make([]string, 0)

	for _, item := range c.Items {
		for _, rows := range item.Rows {
			if rows.Column[0] == a_site {
				output = append(output, rows.Column[1])
			}
		}
	}

	return output
}

/* NOT DONE */
func (c *CSResponse) CastColumnsToFields() {
	// mostread := CSResponseMostRead{}
	// mostread := new(map[string][]CSResponseMostRead)

	// for _, item := range c.Items {
	// 	for _, rows := range item.Rows {
	// 		resp := CSResponseMostRead{}

	// 		resp.ID = rows.Column[1]
	// 		resp.Pageviews, _ = strconv.Atoi(rows.Column[2])
	// 		resp.Browsers, _ = strconv.Atoi(rows.Column[3])
	// 		resp.Virality, _ = strconv.ParseFloat(rows.Column[4], 64)

	// 		mostread[rows.Column[0]] = append(mostread[rows.Column[0]], resp)
	// 	}
	// }

	// log.Println(mostread)
}
