package qweather

import (
	"github.com/zhangyiming748/pretty"
	"github.com/zhangyiming748/qweather/constant"
	"log"
	"testing"
)

func init() {
	log.SetFlags(2 | 16)
}
func TestIndices(t *testing.T) {
	if rep, err := IndicesApi("101011000", constant.HOST, constant.APIKEY,ONEDAY); err != nil {
		t.Error(err)
	} else {
		pretty.P(rep)
	}
}
