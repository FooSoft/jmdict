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

type KanjidicCharacter struct {
	// The character itself in UTF8 coding.
	Literal string `xml:"literal"`

	// The codepoint element states the code of the character in the various
	// character set standards.
	Codepoint KanjidicCodepoint `xml:"codepoint"`

	// The radical number, in the range 1 to 214. The particular
	// classification type is stated in the rad_type attribute.
	Radical KanjidicRadical `xml:"rad_value"`

	Misc KanjidicMisc `xml:"misc"`

	// This element contains the index numbers and similar unstructured
	// information such as page numbers in a number of published dictionaries,
	// and instructional books on kanji.
	DictionaryNumbers KanjidicDicNumber `xml:"dic_number"`

	// These codes contain information relating to the glyph, and can be used
	// for finding a required kanji. The type of code is defined by the
	// qc_type attribute.
	QueryCode KanjidicQueryCode `xml:"query_code"`
}

type KanjidicCodepoint struct {
	// The cp_value contains the codepoint of the character in a particular
	// standard. The standard will be identified in the cp_type attribute.
	Values []KanjidicCodepointValue `xml:"cp_value"`
}

type KanjidicCodepointValue struct {
	Value string `xml:",chardata"`

	// The cp_type attribute states the coding standard applying to the
	// element. The values assigned so far are:
	// 	jis208 - JIS X 0208-1997 - kuten coding (nn-nn)
	// 	jis212 - JIS X 0212-1990 - kuten coding (nn-nn)
	// 	jis213 - JIS X 0213-2000 - kuten coding (p-nn-nn)
	// 	ucs - Unicode 4.0 - hex coding (4 or 5 hexadecimal digits)
	Type string `xml:"cp_type,attr"`
}

type KanjidicRadical struct {
	// The radical number, in the range 1 to 214. The particular
	// classification type is stated in the rad_type attribute.
	Values []KanjidicCodepointValue `xml:"rad_value"`
}

type KanjidicRadicalValue struct {
	Value string `xml:",chardata"`

	// The rad_type attribute states the type of radical classification.
	// classical - as recorded in the KangXi Zidian.
	// nelson_c - as used in the Nelson "Modern Japanese-English
	// Character Dictionary" (i.e. the Classic, not the New Nelson).
	// This will only be used where Nelson reclassified the kanji.
	Type string `xml:"rad_type,attr"`
}

type KanjidicMisc struct {
	// The kanji grade level. 1 through 6 indicates a Kyouiku kanji
	// and the grade in which the kanji is taught in Japanese schools.
	// 8 indicates it is one of the remaining Jouyou Kanji to be learned
	// in junior high school, and 9 or 10 indicates it is a Jinmeiyou (for use
	// in names) kanji. [G]
	Grade string `xml:"grade"`

	// The stroke count of the kanji, including the radical. If more than
	// one, the first is considered the accepted count, while subsequent ones
	// are common miscounts. (See Appendix E. of the KANJIDIC documentation
	// for some of the rules applied when counting strokes in some of the
	// radicals.) [S]
	StrokeCounts []string `xml:"stroke_count"`

	// Either a cross-reference code to another kanji, usually regarded as a
	// variant, or an alternative indexing code for the current kanji.
	// The type of variant is given in the var_type attribute.
	Variant KanjidicVariant `xml:"variant"`

	// A frequency-of-use ranking. The 2,500 most-used characters have a
	// ranking; those characters that lack this field are not ranked. The
	// frequency is a number from 1 to 2,500 that expresses the relative
	// frequency of occurrence of a character in modern Japanese. This is
	// based on a survey in newspapers, so it is biassed towards kanji
	// used in newspaper articles. The discrimination between the less
	// frequently used kanji is not strong. (Actually there are 2,501
	// kanji ranked as there was a tie.)
	Frequency string `xml:"freq"`

	// When the kanji is itself a radical and has a name, this element
	// contains the name (in hiragana.) [T2]
	RadicalName string `xml:"rad_name"`

	// The (former) Japanese Language Proficiency test level for this kanji.
	// Values range from 1 (most advanced) to 4 (most elementary). This field
	// does not appear for kanji that were not required for any JLPT level.
	// Note that the JLPT test levels changed in 2010, with a new 5-level
	// system (N1 to N5) being introduced. No official kanji lists are
	// available for the new levels. The new levels are regarded as
	// being similar to the old levels except that the old level 2 is
	// now divided between N2 and N3.
	JlptLevel string `xml:"jlpt"`
}

type KanjidicVariant struct {
	Value string `xml:",chardata"`

	// The var_type attribute indicates the type of variant code. The current
	// values are:
	// 	jis208 - in JIS X 0208 - kuten coding
	// 	jis212 - in JIS X 0212 - kuten coding
	// 	jis213 - in JIS X 0213 - kuten coding
	// 	  (most of the above relate to "shinjitai/kyuujitai"
	// 	  alternative character glyphs)
	// 	deroo - De Roo number - numeric
	// 	njecd - Halpern NJECD index number - numeric
	// 	s_h - The Kanji Dictionary (Spahn & Hadamitzky) - descriptor
	// 	nelson_c - "Classic" Nelson - numeric
	// 	oneill - Japanese Names (O'Neill) - numeric
	// 	ucs - Unicode codepoint- hex
	Type string `xml:"var_type"`
}

type KanjidicDicNumber struct {
	// Each dic_ref contains an index number. The particular dictionary,
	// etc. is defined by the dr_type attribute.
	DictionaryReferences []KanjiDicReference `xml:"dic_ref"`
}

type KanjiDicReference struct {
	// The dr_type defines the dictionary or reference book, etc. to which
	// dic_ref element applies. The initial allocation is:
	//   nelson_c - "Modern Reader's Japanese-English Character Dictionary",
	//   	edited by Andrew Nelson (now published as the "Classic"
	//   	Nelson).
	//   nelson_n - "The New Nelson Japanese-English Character Dictionary",
	//   	edited by John Haig.
	//   halpern_njecd - "New Japanese-English Character Dictionary",
	//   	edited by Jack Halpern.
	//   halpern_kkd - "Kodansha Kanji Dictionary", (2nd Ed. of the NJECD)
	//   	edited by Jack Halpern.
	//   halpern_kkld - "Kanji Learners Dictionary" (Kodansha) edited by
	//   	Jack Halpern.
	//   halpern_kkld_2ed - "Kanji Learners Dictionary" (Kodansha), 2nd edition
	//     (2013) edited by Jack Halpern.
	//   heisig - "Remembering The  Kanji"  by  James Heisig.
	//   heisig6 - "Remembering The  Kanji, Sixth Ed."  by  James Heisig.
	//   gakken - "A  New Dictionary of Kanji Usage" (Gakken)
	//   oneill_names - "Japanese Names", by P.G. O'Neill.
	//   oneill_kk - "Essential Kanji" by P.G. O'Neill.
	//   moro - "Daikanwajiten" compiled by Morohashi. For some kanji two
	//   	additional attributes are used: m_vol:  the volume of the
	//   	dictionary in which the kanji is found, and m_page: the page
	//   	number in the volume.
	//   henshall - "A Guide To Remembering Japanese Characters" by
	//   	Kenneth G.  Henshall.
	//   sh_kk - "Kanji and Kana" by Spahn and Hadamitzky.
	//   sh_kk2 - "Kanji and Kana" by Spahn and Hadamitzky (2011 edition).
	//   sakade - "A Guide To Reading and Writing Japanese" edited by
	//   	Florence Sakade.
	//   jf_cards - Japanese Kanji Flashcards, by Max Hodges and
	// 	Tomoko Okazaki. (Series 1)
	//   henshall3 - "A Guide To Reading and Writing Japanese" 3rd
	// 	edition, edited by Henshall, Seeley and De Groot.
	//   tutt_cards - Tuttle Kanji Cards, compiled by Alexander Kask.
	//   crowley - "The Kanji Way to Japanese Language Power" by
	//   	Dale Crowley.
	//   kanji_in_context - "Kanji in Context" by Nishiguchi and Kono.
	//   busy_people - "Japanese For Busy People" vols I-III, published
	// 	by the AJLT. The codes are the volume.chapter.
	//   kodansha_compact - the "Kodansha Compact Kanji Guide".
	//   maniette - codes from Yves Maniette's "Les Kanjis dans la tete" French adaptation of Heisig.
	Type string `xml:"dr_type,attr"`

	// See above under "moro".
	Volume string `xml:"m_vol,attr"`

	// See above under "moro".
	Page string `xml:"m_page,attr"`
}
