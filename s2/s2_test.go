package s2geometry

import (
	"fmt"
	"math"
	"testing"

	"github.com/airof98/pan/com"
	"github.com/golang/geo/s2"
)

func TestDistance(t *testing.T) {
	lng0, lat0 := 126.976972789, 37.564155705 //서울역
	lng1, lat1 := 129.158583945, 35.157957397 //해운대
	pt0 := s2.PointFromLatLng(s2.LatLngFromDegrees(lat0, lng0))
	pt1 := s2.PointFromLatLng(s2.LatLngFromDegrees(lat1, lng1))
	fmt.Println(pt0.Distance(pt1).Radians() * com.EarthRadiusMeters)

	pt0 = s2.PointFromLatLng(s2.LatLngFromDegrees(0, 0))
	pt1 = s2.PointFromLatLng(s2.LatLngFromDegrees(0, 180.0))
	fmt.Printf("%f %f\n", pt0.Distance(pt1).Radians()*com.EarthRadiusMeters, com.EarthRadiusMeters*math.Pi)
}
