package utils

import (
	"sync"
	_ "time"
)

var JwtCheck = make(chan bool)

var JwtCollection map[string]int64

func init() {
	JwtCollection = make(map[string]int64)
	JwtCollection["jwt7"] = 167177092
	JwtCollection["jwt1"] = 167177092
	JwtCollection["jwt2"] = 167177093
	JwtCollection["jwt3"] = 167177951
	JwtCollection["jwt4"] = 167172952
	JwtCollection["jwt5"] = 167172952
	JwtCollection["jwt6"] = 167172952
}

func CheckJWT(jwt string, timeStamp int64, wg *sync.WaitGroup) {
	defer wg.Done()
	creationTime, ok := JwtCollection[jwt]
	if ok && creationTime <= timeStamp && timeStamp <= creationTime+10000 {
		JwtCheck <- true
	} else {
		JwtCheck <- false
	}
}
