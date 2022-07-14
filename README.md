# JMDict

JMDict is a simple library written in Go for parsing the raw data files for the
[JMDict](http://www.edrdg.org/jmdict/j_jmdict.html) (vocabulary)
[JMnedict](http://www.edrdg.org/enamdict/enamdict_doc.html) (names), and
[KANJIDIC](http://nihongo.monash.edu/kanjidic2/index.html) (Kanji) dictionaries. As far as I know, these are the only
publicly available Japanese dictionaries and are therefore used by a variety of tools (including
[Yomichan-Import](https://foosoft.net/projects/yomichan-import) from this site).

The XML format used to store dictionary entries and entity data was deceptively annoying to work with, leading to the
creation of this library. Please see the [documentation page](https://godoc.org/foosoft.net/projects/jmdict) for a
technical overview of how to use this library.

Please import this library from `foosoft.net/projects/jmdict` and not the GitHub path.
