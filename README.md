# JMDict #

JMDict is a simple library written in Go for parsing the raw data files for the
[JMDict](http://www.edrdg.org/enamdict/enamdict_doc.html) (vocabulary)
[JMnedict](http://www.edrdg.org/enamdict/enamdict_doc.html) (names), and
[KANJIDIC](http://nihongo.monash.edu/kanjidic2/index.html) (Kanji) dictionaries. As far as I know, these are the only
publicly available Japanese dictionaries and are therefore used by a variety of tools (including
[Yomichan-Import](https://foosoft.net/projects/yomichan-import) from this site).

The XML format used to store dictionary entries and entity data was deceptively annoying to work with, leading to the
creation of this library. Please see the [documentation page](https://godoc.org/github.com/FooSoft/jmdict) for a
technical overview of how to use this library.

## License ##

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
