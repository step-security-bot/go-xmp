// seehuhn.de/go/xmp - Extensible Metadata Platform in Go
// Copyright (C) 2024  Jochen Voss <voss@seehuhn.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package dc

import (
	"encoding/xml"

	"seehuhn.de/go/xmp"
)

const (
	NameSpace = "http://purl.org/dc/elements/1.1/"
)

// DublinCore represents the properties in the Dublin Core namespace.
type DublinCore struct {
	// Contributor is a list of contributors to the resource.
	// This should not include names listed in the Creator field.
	Contributor xmp.UnorderedArray[xmp.ProperName]

	// Coverage is the extent or scope of the resource.
	Coverage xmp.Text

	// Creator is a list of the creators of the resource.  Entities should be
	// listed in order of decreasing precedence, if such order is significant.
	Creator xmp.OrderedArray[xmp.ProperName]

	// Date is a point or period of time associated with an event in the life
	// cycle of the resource.
	Date xmp.OrderedArray[xmp.Date]

	// Description string
	// Format      string
	// Identifier  string
	// Language    string
	// Publisher   string
	// Relation    string
	// Rights      string
	// Source      string
	// Subject     string
	// Title       string
	// Type        string
}

// EncodeXMP implements the [xmp.Model] interface.
func (dc *DublinCore) EncodeXMP(e *xmp.Encoder, pfx string) error {
	err := e.EncodeProperty(NameSpace, "contributor", dc.Contributor)
	if err != nil {
		return err
	}

	err = e.EncodeProperty(NameSpace, "coverage", dc.Coverage)
	if err != nil {
		return err
	}

	return nil
}

// NameSpaces implements the [xmp.Model] interface.
func (dc *DublinCore) NameSpaces(m map[string]struct{}) map[string]struct{} {
	if m == nil {
		m = make(map[string]struct{})
	}

	m[NameSpace] = struct{}{}
	m[xmp.RDFNameSpace] = struct{}{}

	return m
}

func updateDublinCore(m xmp.Model, name string, tokens []xml.Token) (xmp.Model, error) {
	var dc *DublinCore
	if m, ok := m.(*DublinCore); ok {
		dc = m
	} else {
		dc = &DublinCore{}
	}

	switch name {
	case "contributor":
		v, err := dc.Contributor.DecodeAnother(tokens)
		if err != nil {
			return nil, err
		}
		dc.Contributor = v.(xmp.UnorderedArray[xmp.ProperName])
	case "coverage":
		v, err := dc.Coverage.DecodeAnother(tokens)
		if err != nil {
			return nil, err
		}
		dc.Coverage = v.(xmp.Text)
	}

	return dc, nil
}

func init() {
	xmp.RegisterModel(NameSpace, "dc", updateDublinCore)
}
