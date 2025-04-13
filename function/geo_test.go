package function

import (
	"log"
	"testing"

	"github.com/zhangyiming748/qweather/constant"
)

func init() {
	log.SetFlags(2 | 16)
}
func TestGeo(t *testing.T) {
	GeoApi("北京", constant.HOST, constant.APIKEY)
}
