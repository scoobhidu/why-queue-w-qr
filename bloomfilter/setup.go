package bloomfilter

import bloomfilter_database "why-queue-w-qr/bloomfilter/database"

func setup(initCacheSize int) bloomfilter_database.BloomFilterCache {
	cache := bloomfilter_database.BloomFilterCache{}
	cache.OnBits = 0
	cache.BitsArray = make([]byte, initCacheSize)
	return cache
}
