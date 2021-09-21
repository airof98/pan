package com

import (
	"math"
)

const EarthRadiusMeters = 6371010.0

func Distacne(slng, slat, elng, elat float64) float64 {
	//https://en.wikipedia.org/wiki/Haversine_formula
	la1 := slat * math.Pi / 180
	lo1 := slng * math.Pi / 180
	la2 := elat * math.Pi / 180
	lo2 := elng * math.Pi / 180

	hsin := func(theta float64) float64 {
		return math.Pow(math.Sin(theta/2), 2)
	}

	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)
	return 2 * EarthRadiusMeters * math.Asin(math.Sqrt(h))
}

func EuclidDistance(sx, sy, ex, ey float64) float64 {
	dx, dy := ex-sx, ey-sy
	return math.Sqrt(dx*dx + dy*dy)
}
