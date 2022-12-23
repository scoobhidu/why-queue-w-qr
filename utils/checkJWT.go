package utils

import (
	"sync"
	_ "time"
	"why-queue-w-qr/attendance/models/MongoJWTmodels"
)

func CheckJWT(jwt string, timeStamp int64, wg *sync.WaitGroup) bool {
	defer wg.Done()
	creationTime, ok := MongoJWTmodels.JwtCollection[jwt]
	return ok && creationTime <= timeStamp && timeStamp <= creationTime+10000
}
