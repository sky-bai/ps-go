package tools

import (
	"encoding/xml"
	"io"
)

type XmlResult map[string]any

type xmlMapEntry struct {
	XMLName xml.Name
	Value   any `xml:",chardata"`
}

func (m XmlResult) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}

	return e.EncodeToken(start.End())
}

func (m *XmlResult) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = XmlResult{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}
