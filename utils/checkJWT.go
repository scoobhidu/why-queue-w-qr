package utils

import (
	"sync"
	_ "time"
	"why-queue-w-qr/attendance/attendanceHandlers"
	"why-queue-w-qr/attendance/models/MongoJWTmodels"
)

func CheckJWT(jwt string, timeStamp int64, wg *sync.WaitGroup) {
	defer wg.Done()
	creationTime, ok := MongoJWTmodels.JwtCollection[jwt]
	if ok && creationTime <= timeStamp && timeStamp <= creationTime+10000 {
		attendanceHandlers.JwtCheck <- true
	} else {
		attendanceHandlers.JwtCheck <- false
	}
}
