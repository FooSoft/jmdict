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
)

// Entries consist of kanji elements, reading elements
// name translation elements. Each entry must have at
// least one reading element and one sense element. Others are optional.
type EnamdictEntry struct {
	// A unique numeric sequence number for each entry
	Sequence int `xml:"ent_seq"`

	// The kanji element, or in its absence, the reading element, is
	// the defining component of each entry.
	// The overwhelming majority of entries will have a single kanji
	// element associated with an entity name in Japanese. Where there are
	// multiple kanji elements within an entry, they will be orthographical
	// variants of the same word, either using variations in okurigana, or
	// alternative and equivalent kanji. Common "mis-spellings" may be
	// included, provided they are associated with appropriate information
	// fields. Synonyms are not included; they may be indicated in the
	// cross-reference field associated with the sense element.
	Kanji []EnamdictKanji `xml:"k_ele"`

	// The reading element typically contains the valid readings
	// of the word(s) in the kanji element using modern kanadzukai.
	// Where there are multiple reading elements, they will typically be
	// alternative readings of the kanji element. In the absence of a
	// kanji element, i.e. in the case of a word or phrase written
	// entirely in kana, these elements will define the entry.
	Reading []EnamdictReading `xml:"r_ele"`

	// The trans element will record the translational equivalent
	// of the Japanese name, plus other related information.
	Translation []EnamdictTranslation `xml:"trans"`
}

type EnamdictKanji struct {
	// This element will contain an entity name in Japanese
	// which is written using at least one non-kana character (usually
	// kanji, but can be other characters). The valid
	// characters are kanji, kana, related characters such as chouon and
	// kurikaeshi, and in exceptional cases, letters from other alphabets.
	Expression string `xml:"keb"`

	// This is a coded information field related specifically to the
	// orthography of the keb, and will typically indicate some unusual
	// aspect, such as okurigana irregularity.
	Information []string `xml:"ke_inf"`

	// This and the equivalent re_pri field are provided to record
	// information about the relative priority of the entry, and are for
	// use either by applications which want to concentrate on entries of
	// a particular priority, or to generate subset files. The reason
	// both the kanji and reading elements are tagged is because on
	// occasions a priority is only associated with a particular
	// kanji/reading pair.
	Priority []string `xml:"ke_pri"`
}

type EnamdictReading struct {
	// This element content is restricted to kana and related
	// characters such as chouon and kurikaeshi. Kana usage will be
	// consistent between the keb and reb elements; e.g. if the keb
	// contains katakana, so too will the reb.
	Reading string `xml:"reb"`

	// This element is used to indicate when the reading only applies
	// to a subset of the keb elements in the entry. In its absence, all
	// readings apply to all kanji elements. The contents of this element
	// must exactly match those of one of the keb elements.
	Restrictions []string `xml:"re_restr"`

	// General coded information pertaining to the specific reading.
	// Typically it will be used to indicate some unusual aspect of
	// the reading.
	Information []string `xml:"re_inf"`

	// See the comment on ke_pri above.
	Priority []string `xml:"re_pri"`
}

type EnamdictTranslation struct {
	// The type of name, recorded in the appropriate entity codes.
	NameType []string `xml:"name_type"`

	// This element is used to indicate a cross-reference to another
	// entry with a similar or related meaning or sense. The content of
	// this element is typically a keb or reb element in another entry. In some
	// cases a keb will be followed by a reb and/or a sense number to provide
	// a precise target for the cross-reference. Where this happens, a JIS
	// "centre-dot" (0x2126) is placed between the components of the
	// cross-reference.
	References []string `xml:"xref"`

	// The actual translations of the name, usually as a transcription
	// into the target language.
	Translations []string `xml:"trans_det"`

	// The xml:lang attribute defines the target language of the
	// translated name. It will be coded using the three-letter language
	// code from the ISO 639-2 standard. When absent, the value "eng"
	// (i.e. English) is the default value. The bibliographic (B) codes
	// are used.
	Language string `xml:"lang,attr"`
}

func LoadEnamdict(reader io.Reader, transform bool) ([]EnamdictEntry, map[string]string, error) {
	var entries []EnamdictEntry

	entities, err := parseEntries(reader, transform, func(decoder *xml.Decoder, element *xml.StartElement) error {
		if element.Name.Local != "entry" {
			return nil
		}

		var entry EnamdictEntry
		if err := decoder.DecodeElement(&entry, element); err != nil {
			return err
		}

		entries = append(entries, entry)
		return nil
	})

	return entries, entities, err
}
