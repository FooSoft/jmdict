package jmdict

import (
	"encoding/xml"
	"io"
	"regexp"
)

func parseDict(reader io.Reader, container interface{}, transform bool) (map[string]string, error) {
	decoder := xml.NewDecoder(reader)

	var entities map[string]string
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}

		switch startElement := token.(type) {
		case xml.Directive:
			directive := token.(xml.Directive)
			entities = parseEntities(&directive)
			if transform {
				decoder.Entity = entities
			} else {
				decoder.Entity = make(map[string]string)
				for k, _ := range entities {
					decoder.Entity[k] = k
				}
			}
		case xml.StartElement:
			if err := decoder.DecodeElement(container, &startElement); err != nil {
				return nil, err
			}
		}
	}

	return entities, nil
}

func parseEntities(d *xml.Directive) map[string]string {
	re := regexp.MustCompile("<!ENTITY\\s([0-9\\-A-z]+)\\s\"(.+)\">")
	matches := re.FindAllStringSubmatch(string(*d), -1)

	entities := make(map[string]string)
	for _, match := range matches {
		entities[match[1]] = match[2]
	}

	return entities
}
