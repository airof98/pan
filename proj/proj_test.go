package proj

import (
	"fmt"
	"testing"

	"github.com/pebbe/go-proj-4/proj"
	"github.com/stretchr/testify/assert"
)

const (
	//google
	epsg3857 = "+proj=merc " +
		"+a=6378137 +b=6378137 +lat_ts=0.0 +lon_0=0.0 " +
		"+x_0=0.0 +y_0=0 +k=1.0 +units=m " +
		"+nadgrids=@null +no_defs"
	//wgs84 lng/lat
	epsg4326 = "+proj=longlat +ellps=WGS84 +datum=WGS84 +no_defs"
	//naver
	epsg5179 = "+proj=tmerc +lat_0=38 +lon_0=127.5 +k=0.9996 +x_0=1000000 +y_0=2000000 +ellps=GRS80 +units=m +no_defs"
	//kakao
	epsg5181 = "+proj=tmerc +lat_0=38 +lon_0=127 +k=1 +x_0=200000 +y_0=500000 +ellps=GRS80 +units=m +no_defs"
	lng      = 127.028126674
	lat      = 37.499143848
)

var (
	p3857 *proj.Proj
	p4326 *proj.Proj
	p5179 *proj.Proj
	p5181 *proj.Proj
)

func init() {
	p3857, _ = proj.NewProj(epsg3857)
	p4326, _ = proj.NewProj(epsg4326)
	p5179, _ = proj.NewProj(epsg5179)
	p5181, _ = proj.NewProj(epsg5181)
}

func TestGoogleToLnglat(t *testing.T) {
	tlng, tlat, err := proj.Transform2(p4326, p3857, proj.DegToRad(lng), proj.DegToRad(lat))
	assert.NoError(t, err)
	fmt.Printf("%.1f %.1f\n", tlng, tlat)
	//https://mts1.google.com/vt/x=1325&y=3143&z=13&s=Galile
}

func TestLnglatToNaver(t *testing.T) {
	x, y, err := proj.Transform2(p4326, p3857, proj.DegToRad(lng), proj.DegToRad(lat))
	assert.NoError(t, err)
	fmt.Printf("https://map.naver.com/v5/entry/place/1665298766?c=%f,%f,18,0,0,0,dh\n", x, y)
}

func TestLnglatToKakao(t *testing.T) {
	x, y, err := proj.Transform2(p4326, p5181, proj.DegToRad(lng), proj.DegToRad(lat))
	assert.NoError(t, err)
	fmt.Printf("https://map.kakao.com/?map_type=TYPE_MAP&q=test&urlLevel=3&urlX=%d&urlY=%d\n", int(x*2.5), int(y*2.5))
}
