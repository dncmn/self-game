package utils

import (
	"bytes"
	"encoding/xml"
	"io"
	"self-wechat/utils"
	"strings"
)

// xml to map
func XmlToMap(xmlData []byte) (map[string]string, error) {
	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	m := make(map[string]string)
	var token xml.Token
	var err error
	var k string
	for token, err = decoder.Token(); err == nil; token, err = decoder.Token() {
		if v, ok := token.(xml.StartElement); ok {
			k = v.Name.Local
			continue
		}
		if v, ok := token.(xml.CharData); ok {
			data := string(v.Copy())
			if strings.TrimSpace(data) == "" {
				continue
			}
			m[k] = data
		}
	}

	if err != nil && err != io.EOF {
		return nil, err
	}
	return m, nil
}

// xml []byte to struct
func XmlByteToStruct(bye []byte, rec interface{}) (err error) {
	convertMap, err := XmlToMap(bye)
	if err != nil {
		return
	}
	// map to struct
	err = utils.StructToMap(convertMap, &rec)
	if err != nil {
		return
	}
	return
}
