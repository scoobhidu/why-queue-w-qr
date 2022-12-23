package utils

import (
	"sync"
	_ "time"
	"why-queue-w-qr/models"
)

var JwtCheck = make(chan bool)

func CheckJWT(jwt string, timeStamp int64, wg *sync.WaitGroup) {
	defer wg.Done()
	creationTime, ok := models.JwtCollection[jwt]
	if ok && creationTime <= timeStamp && timeStamp <= creationTime+10000 {
		JwtCheck <- true
	} else {
		JwtCheck <- false
	}
}
