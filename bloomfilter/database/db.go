package bloomfilter_database

// whenever the number of 1's become 70% of size in BloomFilterCache
// we have to re-allocate all the storage by dynamically increasing the size of
// the object used

type BloomFilterCache struct {
	OnBits    int
	BitsArray []byte
}

func (cache *BloomFilterCache) ReAdjustCache() {
	// TODO: Get the column of the data from the database and store it in an array

	cache.OnBits = 0
	cacheSize := len(cache.BitsArray)
	cache.BitsArray = make([]byte, 2*cacheSize)

	// TODO: Rehash the column data array in the the newly made Cache.BitsArray
}
