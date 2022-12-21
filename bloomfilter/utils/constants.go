package bloomfilter_constant

import (
	"hash/crc32"
	"hash/maphash"
)

var CRC32_qTable *crc32.Table = crc32.MakeTable(0xD5828281)
var MapHashSeed = maphash.MakeSeed()
