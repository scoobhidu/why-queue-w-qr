package bloomfilter_service

import (
	"database/sql"
	"fmt"
	"sync"
	bloomfilter_database "why-queue-w-qr/bloomfilter/database"
)

func SetCacheBits(input sql.Rows, cache *bloomfilter_database.BloomFilterCache) bool {
	// by default initialized variables have zero values stored in them hence no need to set each bucket as 0

	//fmt.Println(service.MAPHASH64("wassaup", utils.MapHashSeed) % 10000)
	//fmt.Println(service.CRC32("wassaup") % 10000)
	//fmt.Println(service.ADLER32("wassaup") % 10000)

	var mapHashLocation *uint64 = new(uint64)
	var crcLocation *uint32 = new(uint32)
	var adlerLocation *uint32 = new(uint32)

	var wg sync.WaitGroup

	for input.Next() {
		if float32(cache.OnBits) >= float32(len(cache.BitsArray))*0.7 {
			fmt.Println("Cache is full by 70%, need to reset Bloom Filter Cache")
			// TODO readjust cache
		}

		GetValueHash(&wg, v, mapHashLocation, crcLocation, adlerLocation, uint32(len(cache.BitsArray)))

		if cache.BitsArray[*mapHashLocation] != 1 {
			cache.BitsArray[*mapHashLocation] = 1
			cache.OnBits++
		}
		if cache.BitsArray[*crcLocation] != 1 {
			cache.BitsArray[*crcLocation] = 1
			cache.OnBits++
		}
		if cache.BitsArray[*adlerLocation] != 1 {
			cache.BitsArray[*adlerLocation] = 1
			cache.OnBits++
		}
	}
}
