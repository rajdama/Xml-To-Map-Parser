package xmltoyaml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Func XmlToMap takes xml decoder as input and convert it into map[string]interface{}
func XmlToMap(skey string, a []xml.Attr, dec *xml.Decoder) (map[string]interface{}, error) {

	var n, na map[string]interface{}
	if skey != "" {
		n = make(map[string]interface{})
		na = make(map[string]interface{})
		if len(a) > 0 {

			for _, v := range a {
				na[v.Name.Local] = v.Value
			}

		}
	}
	for {
		t, err := dec.Token()
		if err != nil {

			if err != io.EOF {
				return nil, errors.New(fmt.Sprintf("bad xml token: %s", err.Error()))
			}

			return nil, err
		}

		switch kind := t.(type) {

		case xml.StartElement:
			if skey == "" {
				return XmlToMap(kind.Name.Local, kind.Attr, dec)
			}
			cn, err := XmlToMap(kind.Name.Local, kind.Attr, dec)
			if err != nil {
				return nil, err
			}

			var key string
			var val interface{}
			for key, val = range cn {
				break
			}

			if v, ok := na[key]; ok {
				var sibling []interface{}
				switch v.(type) {
				case []interface{}:
					sibling = v.([]interface{})
				default:
					sibling = []interface{}{v}
				}

				sibling = append(sibling, val)
				na[key] = sibling
			} else {
				na[key] = val
			}

		case xml.EndElement:
			if len(n) == 0 {

				if len(na) > 0 {
					n[skey] = na
				} else {
					n[skey] = ""
				}
			}

			return n, nil
		case xml.CharData:
			tt := strings.TrimSpace((string(kind)))
			if len(tt) > 0 {

				if len(na) > 0 {
					na["#test"] = tt
					n[skey] = na
				} else {
					n[skey] = tt
				}

			}
		}

	}
}
