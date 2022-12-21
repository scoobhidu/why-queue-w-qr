package bloomfilter_service

import (
	"fmt"
	"sync"
	bloomfilter_database "why-queue-w-qr/bloomfilter/database"
)

func CheckInCache(checkArray []string, cache *bloomfilter_database.BloomFilterCache, database *bloomfilter_database.DB) {
	var wg sync.WaitGroup

	var mapHashLocation *uint64 = new(uint64)
	var crcLocation *uint32 = new(uint32)
	var adlerLocation *uint32 = new(uint32)

	for _, v := range checkArray {
		GetValueHash(&wg, v, mapHashLocation, crcLocation, adlerLocation, uint32(len(cache.BitsArray)))

		if cache.BitsArray[*mapHashLocation] == 1 && cache.BitsArray[*adlerLocation] == 1 && cache.BitsArray[*crcLocation] == 1 {
			fmt.Println(v + " is present acc to bloom filter, might be false positive")
			var found bool = false
			for _, val := range database.Values {
				if val == v {
					fmt.Println("Found " + v + " in database -> not a false positive")
					found = true
				}
			}
			if found == false {
				fmt.Println("Not Found " + v + " in database -> false positive")
			}
		} else {
			fmt.Println(v + " is not present in database 100% sure")
		}
	}
}
