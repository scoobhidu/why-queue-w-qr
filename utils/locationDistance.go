package utils

import (
	"math"
	"sync"
)

func GetDistanceFromLatLonInKm(lat1 float64, lon1 float64, lat2 float64, lon2 float64, wg *sync.WaitGroup) float64 {
	defer wg.Done()

	var R float64 = 6371.345 // Radius of the earth in km
	// lat2 will always be user location
	// lat1 is college location
	var dLat = deg2rad(lat2 - lat1) // deg2rad below
	var dLon = deg2rad(lon2 - lon1)
	var a = math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(deg2rad(lat1))*math.Cos(deg2rad(lat2))*math.Sin(dLon/2)*math.Sin(dLon/2)

	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var distance = R * c // Distance in km
	return distance
}

func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
