package com

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {
	lng0, lat0 := 126.976972789, 37.564155705 //서울역
	lng1, lat1 := 129.158583945, 35.157957397 //해운대
	d := Distacne(lng0, lat0, lng1, lat1)
	fmt.Println(d)
}
