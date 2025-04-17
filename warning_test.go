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
func TestWarning(t *testing.T) {
	if rep, err := WarningApi("101011000", constant.HOST, constant.APIKEY); err != nil {
		t.Error(err)
	} else {
		for i, v := range rep.Warning{
			log.Printf("第%d个预警:%s", i, v.Text)
			pretty.P(rep)
		}
		l:=len(rep.Warning)
		log.Printf("总共%d个预警", l)
	}
}
