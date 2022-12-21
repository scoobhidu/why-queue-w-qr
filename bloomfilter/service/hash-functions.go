package bloomfilter_service

import (
	"hash/adler32"
	"hash/crc32"
	"hash/maphash"
	"sync"
	bloomfilter_constant "why-queue-w-qr/bloomfilter/utils"
)

// for the sake of simplicity we will be using hash32 and crc32 hash functions
// that return the values in uint32 format, so that bits can be set easily
// using MapHash hash function, but it returns uint64

func MAPHASH64(str string, seed maphash.Seed, mapHashBit *uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	//fmt.Println("Start Map")
	*mapHashBit = maphash.Bytes(seed, []byte(str))
	//fmt.Println("End Map")
}

func CRC32(str string, crcBit *uint32, wg *sync.WaitGroup) {
	defer wg.Done()
	//fmt.Println("Start crc")
	*crcBit = crc32.Checksum([]byte(str), bloomfilter_constant.CRC32_qTable)
	//fmt.Println("End Crc")
}

func ADLER32(str string, adlerBit *uint32, wg *sync.WaitGroup) {
	defer wg.Done()
	//fmt.Println("Start adler")
	*adlerBit = adler32.Checksum([]byte(str))
	//fmt.Println("End Adler")
}

func GetValueHash(wg *sync.WaitGroup, str string, mapHashLocation *uint64, crcLocation *uint32, adlerLocation *uint32, length uint32) {
	wg.Add(3)
	go MAPHASH64(str, bloomfilter_constant.MapHashSeed, mapHashLocation, wg)
	go CRC32(str, crcLocation, wg)
	go ADLER32(str, adlerLocation, wg)
	wg.Wait()

	//fmt.Println("Hash Values for " + str + ", MapHash, CRC, Adler respectively")
	*mapHashLocation = *mapHashLocation % uint64(length)
	*crcLocation = *crcLocation % length
	*adlerLocation = *adlerLocation % length
}
