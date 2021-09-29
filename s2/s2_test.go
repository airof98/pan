package s2geometry

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/airof98/pan/com"
	"github.com/golang/geo/s1"
	"github.com/golang/geo/s2"
)

var (
	lng0, lat0 = 129.158583945, 35.157957397 //해운대
	lng1, lat1 = 126.976972789, 37.564155705 //서울역
	lng2, lat2 = 128.199970021, 41.920170977 //백두산
	lng3, lat3 = 127.422553009, 36.335057685 //대전
	lng4, lat4 = 128.560288926, 36.842109061 //소백산
)

func TestDistance(t *testing.T) {
	pt0 := s2.PointFromLatLng(s2.LatLngFromDegrees(lat0, lng0))
	pt1 := s2.PointFromLatLng(s2.LatLngFromDegrees(lat1, lng1))
	fmt.Println(pt0.Distance(pt1).Radians() * com.EarthRadiusMeters)

	pt0 = s2.PointFromLatLng(s2.LatLngFromDegrees(0, 0))
	pt1 = s2.PointFromLatLng(s2.LatLngFromDegrees(0, 180.0))
	fmt.Printf("%f %f\n", pt0.Distance(pt1).Radians()*com.EarthRadiusMeters, com.EarthRadiusMeters*math.Pi)
}

func TestSellID(t *testing.T) {
	ll := s2.LatLngFromDegrees(lat1, lng1)
	cellID := s2.CellIDFromLatLng(ll)
	fmt.Println(cellID, cellID.Face(), cellID.Level())
	cell := s2.CellFromLatLng(ll)
	fmt.Println(cell)

	ll = s2.LatLngFromDegrees(lat2, lng2)
	cellID = s2.CellIDFromLatLng(ll)
	fmt.Println(cellID, cellID.Face(), cellID.Level(), cellID.LatLng())
	for i := 29; i >= 0; i-- {
		parent := cellID.Parent(i)
		fmt.Println(parent, parent.Face(), parent.Level(), parent.LatLng())
	}
}

func TestRegionCoverer(t *testing.T) {
	ll := s2.LatLngFromDegrees(lat1, lng1)
	point := s2.PointFromLatLng(ll)
	angle := s1.Angle(200000 / com.EarthRadiusMeters)
	sphereCap := s2.CapFromCenterAngle(point, angle)
	region := s2.Region(sphereCap)
	rc := &s2.RegionCoverer{MaxLevel: 30, MaxCells: 5}
	covering := rc.Covering(region)
	ll = s2.LatLngFromDegrees(lat2, lng2)
	fmt.Println(com.Distacne(lng1, lat1, lng2, lat2))
	fmt.Println(covering.ContainsCellID(s2.CellIDFromLatLng(ll)))
}

func TestFindPolylineEdges(t *testing.T) {

	polyline := s2.Polyline{
		s2.PointFromLatLng(s2.LatLngFromDegrees(lat0, lng0)), //해운대
		s2.PointFromLatLng(s2.LatLngFromDegrees(lat1, lng1)), //서울역
		s2.PointFromLatLng(s2.LatLngFromDegrees(lat2, lng2)), //백두산
	}

	shapeIndex := s2.NewShapeIndex()
	shapeIndex.Add(&polyline)
	query := s2.NewClosestEdgeQuery(shapeIndex, s2.NewClosestEdgeQueryOptions().MaxResults(7))
	target := s2.NewMinDistanceToPointTarget(s2.PointFromLatLng(s2.LatLngFromDegrees(lat3, lng3))) //대전

	for _, result := range query.FindEdges(target) {
		polylineIndex := result.ShapeID()
		edgeIndex := result.EdgeID()
		distance := result.Distance()
		fmt.Printf("polyline %d, edge %d is %0.1f meter\n", polylineIndex, edgeIndex,
			distance.Angle().Radians()*com.EarthRadiusMeters)
	}
}

func TestPolygon(t *testing.T) {
	latlng := s2.LatLngFromDegrees(lat1, lng1) //서울역
	polygon := s2.PolygonFromCell(s2.CellFromCellID(
		s2.CellIDFromLatLng(latlng).Parent(19)))
	assert.True(t, polygon.ContainsPoint(s2.PointFromLatLng(latlng)))
	latlng = s2.LatLngFromDegrees(lat2, lng2) //백두산
	assert.False(t, polygon.ContainsPoint(s2.PointFromLatLng(latlng)))

	loop := s2.LoopFromPoints([]s2.Point{ //counter clockWise
		s2.PointFromLatLng(s2.LatLngFromDegrees(lat0, lng0)), //해운대
		s2.PointFromLatLng(s2.LatLngFromDegrees(lat2, lng2)), //백두산
		s2.PointFromLatLng(s2.LatLngFromDegrees(lat1, lng1)), //서울역
		s2.PointFromLatLng(s2.LatLngFromDegrees(lat3, lng3)), //대전
		s2.PointFromLatLng(s2.LatLngFromDegrees(lat0, lng0)), //해운대
	})
	polygon = s2.PolygonFromLoops([]*s2.Loop{loop})
	assert.False(t, polygon.ContainsPoint(s2.PointFromLatLng(
		s2.LatLngFromDegrees(33.347019434, 126.382884031)))) //제주도
}

func TestPolyline(t *testing.T) {
	polyline := s2.Polyline{
		s2.PointFromLatLng(s2.LatLngFromDegrees(lat0, lng0)), //해운대
		s2.PointFromLatLng(s2.LatLngFromDegrees(lat1, lng1)), //서울역
		s2.PointFromLatLng(s2.LatLngFromDegrees(lat2, lng2)), //백두산
	}

	tPoint := s2.PointFromLatLng(s2.LatLngFromDegrees(lat3, lng3)) //대전
	//linestring projection
	projPoint, nVertexIdx := polyline.Project(tPoint)
	fmt.Println(s2.LatLngFromPoint(projPoint), nVertexIdx)
	//interpolation
	intPoint, nVertexIdx := polyline.Interpolate(0.3)
	fmt.Println(s2.LatLngFromPoint(intPoint), nVertexIdx)

	assert.False(t, polyline.IsOnRight(tPoint))
	tPoint = s2.PointFromLatLng(s2.LatLngFromDegrees(lat4, lng4)) //소백산
	//Is right at linstring
	assert.True(t, polyline.IsOnRight(tPoint))

	polyline1 := s2.Polyline{
		s2.PointFromLatLng(s2.LatLngFromDegrees(lat3, lng3)), //대전
		s2.PointFromLatLng(s2.LatLngFromDegrees(lat4, lng4)), //소백산
	}

	polyline1 = s2.Polyline{
		s2.PointFromLatLng(s2.LatLngFromDegrees(lat3, lng3)),                  //대전
		s2.PointFromLatLng(s2.LatLngFromDegrees(33.347019434, 126.382884031)), //제주도
	}
	//intersect
	assert.False(t, polyline.Intersects(&polyline1))
	cell := s2.CellFromLatLng(s2.LatLngFromDegrees(lat1, lng1)) //서울역
	assert.True(t, polyline.IntersectsCell(cell))
}
