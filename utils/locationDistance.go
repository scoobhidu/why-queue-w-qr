package utils

import (
	"math"
	"sync"
)

var DistanceCheck = make(chan bool)

func GetDistanceFromLatLonInKm(lat1 float64, lon1 float64, lat2 float64, lon2 float64, wg *sync.WaitGroup) {
	defer wg.Done()

	var R float64 = 6371.345 // Radius of the earth in km
	// lat2 will always be user location
	// lat1 is college location
	var dLat = deg2rad(lat2 - lat1) // deg2rad below
	var dLon = deg2rad(lon2 - lon1)
	var a = math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(deg2rad(lat1))*math.Cos(deg2rad(lat2))*math.Sin(dLon/2)*math.Sin(dLon/2)

	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var distance = R * c // Distance in km
	if distance < 0.1 {
		DistanceCheck <- true
	} else {
		DistanceCheck <- false
	}
}

func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
