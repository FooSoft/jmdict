/*
 * Copyright (c) 2016 Alex Yatskov <alex@foosoft.net>
 * Author: Alex Yatskov <alex@foosoft.net>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package jmdict

import (
	"encoding/xml"
	"io"
	"regexp"
)

type Parser func(decoder *xml.Decoder, element *xml.StartElement) error

func parseEntries(reader io.Reader, transform bool, callback Parser) (map[string]string, error) {
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
			if err := callback(decoder, &startElement); err != nil {
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
