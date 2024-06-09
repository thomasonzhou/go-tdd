package clockface_test

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	clockface "hello/math"
)

type Svg struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    Line     `xml:"line"`
}
type Circle struct {
	Cx string `xml:"cx,attr"`
	Cy string `xml:"cy,attr"`
	R  string `xml:"r,attr"`
}

type Line []struct {
	X1 string `xml:"x1,attr"`
	Y1 string `xml:"y1,attr"`
	X2 string `xml:"x2,attr"`
	Y2 string `xml:"y2,attr"`
}

func TestSVGWriterAtMidnight(t *testing.T) {
	tm := time.Date(2024, time.October, 14, 0, 0, 0, 0, time.UTC)

	b := bytes.Buffer{}
	clockface.SVGWriter(&b, tm)

	svg := Svg{}
	xml.Unmarshal(b.Bytes(), &svg)

	x2 := "150.000"
	y2 := "60.000"

	for _, line := range svg.Line {
		if line.X2 == x2 && line.Y2 == y2 {
			return
		}
	}

	t.Errorf("Expected to find x2 %+v and y2 %+v in SVG %v", x2, y2, b.String())
}
