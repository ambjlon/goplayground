package app

import (
	"fmt"
	"math"
)

func TestEarthDistance() {
	lat1 := 120.358017
	lng1 := 36.102131

	lat2 := 120.359688
	lng2 := 36.102222
	fmt.Println(EarthDistance(lat1, lng1, lat2, lng2))
}

func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6371000 // 6378137
	rad := math.Pi / 180.0

	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad

	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))

	return dist * float64(radius)
}
