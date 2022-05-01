// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package translations

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en_GB": &dictionary{index: en_GBIndex, data: en_GBData},
		"en_ID": &dictionary{index: en_IDIndex, data: en_IDData},
		"en_MY": &dictionary{index: en_MYIndex, data: en_MYData},
	}
	fallback := language.MustParse("en-US")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"There are %d books\n": 1,
	"hello %s\n":           0,
}

var en_GBIndex = []uint32{ // 3 elements
	0x00000000, 0x00000019, 0x00000058,
} // Size: 36 bytes

const en_GBData string = "" + // Size: 88 bytes
	"\x04\x00\x01\x0a\x14\x02(en-gb) hello %[1]s\x04\x00\x01\x0a:\x14\x01\x81" +
	"\x01\x00=\x01\x10\x02There is 1 book\x00 \x02There are %[1]d books avail" +
	"able"

var en_IDIndex = []uint32{ // 3 elements
	0x00000000, 0x00000019, 0x00000068,
} // Size: 36 bytes

const en_IDData string = "" + // Size: 104 bytes
	"\x04\x00\x01\x0a\x14\x02(en-ID) hello %[1]s\x04\x00\x01\x0aJ\x14\x01\x81" +
	"\x01\x00=\x01\x18\x02(en-GB) There is 1 book\x00(\x02(en-GB) There are %" +
	"[1]d books available"

var en_MYIndex = []uint32{ // 3 elements
	0x00000000, 0x00000019, 0x00000068,
} // Size: 36 bytes

const en_MYData string = "" + // Size: 104 bytes
	"\x04\x00\x01\x0a\x14\x02(en-MY) hello %[1]s\x04\x00\x01\x0aJ\x14\x01\x81" +
	"\x01\x00=\x01\x18\x02(en-MY) There is 1 book\x00(\x02(en-MY) There are %" +
	"[1]d books available"

	// Total table size 404 bytes (0KiB); checksum: 97A0714C