package jmdict

import "io"

type Kanjidic struct {
	// The single header element will contain identification information
	// about the version of the file
	Header KanjidicHeader `xml:"header"`

	Characters []KanjidicCharacter `xml:"character"`
}

type KanjidicHeader struct {
	// This field denotes the version of kanjidic2 structure, as more
	// than one version may exist.
	FileVersion string `xml:"file_version"`

	// The version of the file, in the format YYYY-NN, where NN will be
	// a number starting with 01 for the first version released in a
	// calendar year, then increasing for each version in that year.
	DatabaseVersion string `xml:"database_version"`

	// The date the file was created in international format (YYYY-MM-DD).
	DateOfCreation string `xml:"date_of_creation"`
}

type KanjidicCharacter struct {
	// The character itself in UTF8 coding.
	Literal string `xml:"literal"`

	// The codepoint element states the code of the character in the various
	// character set standards.
	Codepoint []KanjidicCodepoint `xml:"codepoint>cp_value"`

	// The radical number, in the range 1 to 214. The particular
	// classification type is stated in the rad_type attribute.
	Radical []KanjidicRadical `xml:"radical>rad_value"`

	Misc KanjidicMisc `xml:"misc"`

	// This element contains the index numbers and similar unstructured
	// information such as page numbers in a number of published dictionaries,
	// and instructional books on kanji.
	DictionaryNumbers []KanjidicDicNumber `xml:"dic_number>dic_ref"`

	// These codes contain information relating to the glyph, and can be used
	// for finding a required kanji. The type of code is defined by the
	// qc_type attribute.
	QueryCode []KanjidicQueryCode `xml:"query_code>q_code"`

	// The readings for the kanji in several languages, and the meanings, also
	// in several languages. The readings and meanings are grouped to enable
	// the handling of the situation where the meaning is differentiated by
	// reading. [T1]
	ReadingMeaning *KanjidicReadingMeaning `xml:"reading_meaning"`
}

type KanjidicCodepoint struct {
	// The cp_value contains the codepoint of the character in a particular
	// standard. The standard will be identified in the cp_type attribute.
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
	Grade *string `xml:"grade"`

	// The stroke count of the kanji, including the radical. If more than
	// one, the first is considered the accepted count, while subsequent ones
	// are common miscounts. (See Appendix E. of the KANJIDIC documentation
	// for some of the rules applied when counting strokes in some of the
	// radicals.) [S]
	StrokeCounts []string `xml:"stroke_count"`

	// Either a cross-reference code to another kanji, usually regarded as a
	// variant, or an alternative indexing code for the current kanji.
	// The type of variant is given in the var_type attribute.
	Variants []KanjidicVariant `xml:"variant"`

	// A frequency-of-use ranking. The 2,500 most-used characters have a
	// ranking; those characters that lack this field are not ranked. The
	// frequency is a number from 1 to 2,500 that expresses the relative
	// frequency of occurrence of a character in modern Japanese. This is
	// based on a survey in newspapers, so it is biassed towards kanji
	// used in newspaper articles. The discrimination between the less
	// frequently used kanji is not strong. (Actually there are 2,501
	// kanji ranked as there was a tie.)
	Frequency *string `xml:"freq"`

	// When the kanji is itself a radical and has a name, this element
	// contains the name (in hiragana.) [T2]
	RadicalName []string `xml:"rad_name"`

	// The (former) Japanese Language Proficiency test level for this kanji.
	// Values range from 1 (most advanced) to 4 (most elementary). This field
	// does not appear for kanji that were not required for any JLPT level.
	// Note that the JLPT test levels changed in 2010, with a new 5-level
	// system (N1 to N5) being introduced. No official kanji lists are
	// available for the new levels. The new levels are regarded as
	// being similar to the old levels except that the old level 2 is
	// now divided between N2 and N3.
	JlptLevel *string `xml:"jlpt"`
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
	Value string `xml:",chardata"`

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

type KanjidicQueryCode struct {
	Value string `xml:",chardata"`

	// The qc_type attribute defines the type of query code. The current values
	// are:
	//   skip -  Halpern's SKIP (System  of  Kanji  Indexing  by  Patterns)
	//   	code. The  format is n-nn-nn.  See the KANJIDIC  documentation
	//   	for  a description of the code and restrictions on  the
	//   	commercial  use  of this data. [P]  There are also
	// 	a number of misclassification codes, indicated by the
	// 	"skip_misclass" attribute.
	//   sh_desc - the descriptor codes for The Kanji Dictionary (Tuttle
	//   	1996) by Spahn and Hadamitzky. They are in the form nxnn.n,
	//   	e.g.  3k11.2, where the  kanji has 3 strokes in the
	//   	identifying radical, it is radical "k" in the SH
	//   	classification system, there are 11 other strokes, and it is
	//   	the 2nd kanji in the 3k11 sequence. (I am very grateful to
	//   	Mark Spahn for providing the list of these descriptor codes
	//   	for the kanji in this file.) [I]
	//   four_corner - the "Four Corner" code for the kanji. This is a code
	//   	invented by Wang Chen in 1928. See the KANJIDIC documentation
	//   	for  an overview of  the Four Corner System. [Q]

	//   deroo - the codes developed by the late Father Joseph De Roo, and
	//   	published in  his book "2001 Kanji" (Bonjinsha). Fr De Roo
	//   	gave his permission for these codes to be included. [DR]
	//   misclass - a possible misclassification of the kanji according
	// 	to one of the code types. (See the "Z" codes in the KANJIDIC
	// 	documentation for more details.)
	Type string `xml:"qc_type,attr"`

	// The values of this attribute indicate the type if
	// misclassification:
	// - posn - a mistake in the division of the kanji
	// - stroke_count - a mistake in the number of strokes
	// - stroke_and_posn - mistakes in both division and strokes
	// - stroke_diff - ambiguous stroke counts depending on glyph
	Misclassification string `xml:"skip_misclass,attr"`
}

type KanjidicReadingMeaning struct {
	// The reading element contains the reading or pronunciation
	// of the kanji.
	Readings []KanjidicReading `xml:"rmgroup>reading"`

	// The meaning associated with the kanji.
	Meanings []KanjidicMeaning `xml:"rmgroup>meaning"`

	// Japanese readings that are now only associated with names.
	Nanori []string `xml:"nanori"`
}

type KanjidicReading struct {
	Value string `xml:",chardata"`

	// The r_type attribute defines the type of reading in the reading
	// element. The current values are:
	//   pinyin - the modern PinYin romanization of the Chinese reading
	//   	of the kanji. The tones are represented by a concluding
	//   	digit. [Y]
	//   korean_r - the romanized form of the Korean reading(s) of the
	//   	kanji.  The readings are in the (Republic of Korea) Ministry
	//   	of Education style of romanization. [W]
	//   korean_h - the Korean reading(s) of the kanji in hangul.
	//   ja_on - the "on" Japanese reading of the kanji, in katakana.
	//   	Another attribute r_status, if present, will indicate with
	//   	a value of "jy" whether the reading is approved for a
	//   	"Jouyou kanji".
	// 	A further attribute on_type, if present,  will indicate with
	// 	a value of kan, go, tou or kan'you the type of on-reading.
	//   ja_kun - the "kun" Japanese reading of the kanji, usually in
	// 	hiragana.
	//   	Where relevant the okurigana is also included separated by a
	//   	".". Readings associated with prefixes and suffixes are
	//   	marked with a "-". A second attribute r_status, if present,
	//   	will indicate with a value of "jy" whether the reading is
	//   	approved for a "Jouyou kanji".
	Type string `xml:"r_type,attr"`

	// See under ja_on above.
	OnType *string `xml:"on_type"`

	// See under ja_on and ja_kun above.
	JouyouStatus *string `xml:"r_status"`
}

type KanjidicMeaning struct {
	// The meaning associated with the kanji.
	Meaning string `xml:",chardata"`

	// The m_lang attribute defines the target language of the meaning. It
	// will be coded using the two-letter language code from the ISO 639-1
	// standard. When absent, the value "en" (i.e. English) is implied. [{}]
	Language *string `xml:"m_lang,attr"`
}

func LoadKanjidic(reader io.Reader) (Kanjidic, error) {
	var dic Kanjidic
	_, err := parseDict(reader, &dic, true)
	return dic, err
}
