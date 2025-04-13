package function

import (
	"github.com/zhangyiming748/pretty"
	"github.com/zhangyiming748/qweather/constant"
	"log"
	"testing"
)

func init() {
	log.SetFlags(2 | 16)
}
func TestNow(t *testing.T) {
	if rep, err := NowApi("101011000", constant.HOST, constant.APIKEY); err != nil {
		t.Error(err)
	} else {
		pretty.P(rep)
	}

}
